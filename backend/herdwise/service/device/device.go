package device

import (
	"fmt"
	"herdwise/service/database"
	"log"
	"math/rand"
	"reflect"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Service struct {
	Db *gorm.DB
}

type DeviceError string

const (
	DeviceErrorNil            DeviceError = ""
	DeviceErrorDeviceNotFound DeviceError = "Device not found"
)

type Device struct {
	ID               uint   `gorm:"column:ID;primaryKey;autoIncrement" json:"deviceId"`
	DeviceRefference string `gorm:"column:DeviceRefference;type:varchar(100);not null" json:"ref"`
	Longitude        string `gorm:"column:Longitude;not null" json:"long"`
	Latitude         string `gorm:"column:Latitude;not null" json:"lati"`
	Temperature      string `gorm:"column:Temperature;type:varchar(100);not null" json:"temp"`
	AnimalType       string `gorm:"column:AnimalType;type:varchar(100);not null" json:"type"`
	HighTemperature  string `gorm:"column:HighTemperature;type:varchar(100);not null" json:"highTemp"`
	LowTemperature   string `gorm:"column:LowTemperature;type:varchar(100);not null" json:"lowTemp"`
	Model            string `gorm:"column:Model;type:varchar(100);not null" json:"model"`
	FarmerId         int    `gorm:"column:FarmerId;not null" json:"farmerId"`
	CreatedDateD     []byte `gorm:"column:CreatedDate;type:DATETIME DEFAULT CURRENT_TIMESTAMP" json:"-"`
	CreatedDateJ     string `gorm:"-" json:"created_date"`
	LastUpdatedDateD []byte `gorm:"column:LastUpdatedDate;type:DATETIME DEFAULT CURRENT_TIMESTAMP" json:"-"`
	LastUpdatedDateJ string `gorm:"-" json:"last_updated_date"`
}

func (d *Device) TableName() string {
	return "device"
}

func NewDevice(db_config *database.Config) *Service {
	var err error

	s := Service{}

	// Open a new GORM database connection
	s.Db, err = gorm.Open(mysql.Open(db_config.ConnectString()), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Perform auto-migration to create or update the table based on the Device struct
	err = s.Db.AutoMigrate(&Device{})
	if err != nil {
		panic("Auto-migration failed: " + err.Error())
	}

	s.Db = s.Db.Debug()

	log.Println("Auto-migration completed successfully.")

	// Remember to close the database connection when done.
	//s.Db.Close()

	return &s
}

func (s *Service) Create(device *Device) (device_resp *Device, device_error DeviceError) {
	device.ID = 0

	err := s.Db.Transaction(func(tx *gorm.DB) error {
		device.CreatedDateD = []byte(time.Now().Format("2006-01-02T15:04:05"))
		device.LastUpdatedDateD = []byte(time.Now().Format("2006-01-02 15:04:05"))
		device.DeviceRefference = generateRandomID()
		err := tx.Omit("ID").Create(device).Error
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, DeviceError(err.Error())
	}

	device, device_error = s.SelectById(int(device.ID))
	if device_error != DeviceErrorNil {
		return nil, device_error
	}

	return device, DeviceErrorNil
}

// find the Device by ID
func (s *Service) SelectById(device_id int) (device *Device, device_error DeviceError) {
	device = &Device{}
	err := s.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Find(device, "ID = ?", device_id).Error
	})

	if err != nil {
		return nil, DeviceError(err.Error())
	}

	if device.ID < 1 {
		return nil, DeviceErrorDeviceNotFound
	}

	device.CreatedDateJ = byteTimeStampToString(device.CreatedDateD)

	return device, DeviceErrorNil
}

func (s *Service) SelectByRef(ref string) (device *Device, farmer_error DeviceError) {
	device = &Device{}
	err := s.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Find(device, "DeviceRefference = ?", ref).Error
	})

	if err != nil {
		return nil, DeviceError(err.Error())
	}

	if device.ID < 1 {
		return nil, DeviceErrorDeviceNotFound
	}

	return device, DeviceErrorNil
}

// update the user information
func (s *Service) Update(device *Device) (device_resp *Device, device_error DeviceError) {
	_, device_error = s.SelectById(int(device.ID))
	if device_error != DeviceErrorNil {
		return nil, device_error
	}

	update_data := make(map[string]interface{})
	device_struct := reflect.ValueOf(device).Elem()

	device.LastUpdatedDateD = []byte(time.Now().Format("2006-01-02 15:04:05"))
	//assigning name to the user structre for it to be updated
	for i := 0; i < device_struct.NumField(); i++ {
		name := device_struct.Type().Field(i).Name
		value := device_struct.Field(i).Interface()

		//blocking user to updating the spefic fields
		if name == "ID" || name == "DeviceRefference" || name == "farmer_id" || strings.Contains(name, "CreatedDate") || strings.Contains(name, "LastUpdatedDate") {
			continue
		}
		//updatable fields
		for j := 0; j < device_struct.NumField(); j++ {
			update_data[name] = value
		}

	}

	err := s.Db.Transaction(func(tx *gorm.DB) error {
		err := tx.Table(device.TableName()).Where("ID = ?", device.ID).Updates(&update_data).Error
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, DeviceError(err.Error())
	}

	device, device_error = s.SelectById(int(device.ID))
	if device_error != DeviceErrorNil {
		return nil, device_error
	}

	return device, DeviceErrorNil
}

// func (s *Service) List() (devices []*Device, user_error DeviceError) {
// 	devices = make([]*Device, 0)

// 	err := s.Db.Transaction(func(tx *gorm.DB) error {
// 		return tx.Find(&devices).Error
// 	})

// 	if err != nil {
// 		return nil, DeviceError(err.Error())
// 	}

// 	return devices, DeviceErrorNil
// }

func byteTimeStampToString(str []byte) string {
	time_str, _ := time.Parse("2006-01-02 15:04:05", string(str))
	t, err := time.Parse(time.RFC3339, time_str.Format(time.RFC3339))
	if err != nil {
		panic(err)
	}

	return t.Format("2006-01-02 15:04:05")
}

func (s *Service) List(farmer_id int) (devices []*Device, device_error DeviceError) {

	err := s.Db.Where("FarmerId = ?", farmer_id).Find(&devices).Error
	if err != nil {
		return nil, DeviceError(err.Error())

	}

	return devices, DeviceErrorNil

}

func generateRandomID() string {
	rand.Seed(time.Now().UnixNano())

	randomrefId := rand.Intn(9000) + 1000

	return fmt.Sprintf("device%x", randomrefId)
}
