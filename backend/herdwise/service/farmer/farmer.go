package farmer

import (
	"herdwise/service/database"
	"log"

	"reflect"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Service struct {
	Db *gorm.DB
}

type FarmerError string

// error handling messages
const (
	FarmerErrorNil               FarmerError = ""
	FarmerErrorFarmerEmailExists FarmerError = "Farmer with the same email address already exist"
	FarmerErrorFarmerEMPIDExists FarmerError = "Farmer with the same Farmer ID number already exist"
	FarmerErrorFarmerNotFound    FarmerError = "Farmer not found"
	FarmerErrorInvalidLogin      FarmerError = "invalid login username or password"
)

type Farmer struct {
	ID               int    `gorm:"column:ID;primaryKey;autoIncrement" json:"farmerId"`
	FirstName        string `gorm:"column:FirstName;type:varchar(100);not null" json:"first_name"`
	LastName         string `gorm:"column:LastName;type:varchar(100);not null" json:"last_name"`
	Email            string `gorm:"column:Email;type:varchar(100);not null" json:"email"`
	Contact          string `gorm:"column:Contact;type:bigint;" json:"contact,omitempty"`
	Password         string `gorm:"column:Password;not null" json:"password"`
	LatitudeF        string `gorm:"column:LatitudeF;type:varchar(100);not null" json:"latFarm"`
	LongitudeF       string `gorm:"column:LongitudeF;type:varchar(100);not null" json:"longFarm"`
	FarmSize         int    `gorm:"column:FarmSize;type:varchar(100);not null" json:"size"`
	CreatedDateD     []byte `gorm:"column:CreatedDate;type:DATETIME DEFAULT CURRENT_TIMESTAMP" json:"-"`
	CreatedDateJ     string `gorm:"-" json:"created_date"`
	LastUpdatedDateD []byte `gorm:"column:LastUpdatedDate;type:DATETIME DEFAULT CURRENT_TIMESTAMP" json:"-"`
	LastUpdatedDateJ string `gorm:"-" json:"last_updated_date"`
}

// naming the Farmer table
func (u *Farmer) TableName() string {
	return "farmer"
}

func NewFarmer(db_config *database.Config) *Service {
	var err error

	s := Service{}

	// Open a new GORM database connection
	s.Db, err = gorm.Open(mysql.Open(db_config.ConnectString()), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Perform auto-migration to create or update the table based on the farmer struct
	err = s.Db.AutoMigrate(&Farmer{})
	if err != nil {
		panic("Auto-migration failed: " + err.Error())
	}
	s.Db = s.Db.Debug()
	log.Println("Auto-migration completed successfully.")

	// Remember to close the database connection when done.
	//s.Db.Close()

	return &s
}

// create a new farmer
func (s *Service) Create(farmer *Farmer) (farmer_resp *Farmer, farmer_error FarmerError) {
	_, farmer_error = s.SelectByEmail(farmer.Email)
	if farmer_error != FarmerErrorFarmerNotFound {
		return nil, FarmerErrorFarmerEmailExists
	}

	// Reset the ID at it will be omitted and created on insert
	farmer.ID = 0

	err := s.Db.Transaction(func(tx *gorm.DB) error {
		farmer.CreatedDateD = []byte(time.Now().Format("2006-01-02 15:04:05"))
		farmer.LastUpdatedDateD = []byte(time.Now().Format("2006-01-02 15:04:05"))

		//hashing the password

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(farmer.Password), bcrypt.DefaultCost)
		farmer.Password = string(hashedPassword)

		//restricting the id to be updated
		err := tx.Omit("ID").Create(farmer).Error
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, FarmerError(err.Error())
	}

	farmer, farmer_error = s.SelectById(farmer.ID)
	if farmer_error != FarmerErrorNil {
		return nil, farmer_error
	}

	return farmer, FarmerErrorNil
}

// update the farmer information
func (s *Service) Update(farmer *Farmer) (farmer_resp *Farmer, farmer_error FarmerError) {
	_, farmer_error = s.SelectById(farmer.ID)
	if farmer_error != FarmerErrorNil {
		return nil, farmer_error
	}

	update_data := make(map[string]interface{})
	farmer_struct := reflect.ValueOf(farmer).Elem()

	farmer.LastUpdatedDateD = []byte(time.Now().Format("2006-01-02 15:04:05"))
	//assigning name to the farmer structre for it to be updated
	for i := 0; i < farmer_struct.NumField(); i++ {
		name := farmer_struct.Type().Field(i).Name
		value := farmer_struct.Field(i).Interface()

		//blocking te farmer to updating the spefic fields
		if name == "ID" || name == "Password" || strings.Contains(name, "CreatedDate") || strings.Contains(name, "LastUpdatedDate") {
			continue
		}
		//updatable fields
		for j := 0; j < farmer_struct.NumField(); j++ {
			update_data[name] = value
		}

	}

	err := s.Db.Transaction(func(tx *gorm.DB) error {
		err := tx.Table(farmer.TableName()).Where("ID = ?", farmer.ID).Updates(&update_data).Error
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, FarmerError(err.Error())
	}

	farmer, farmer_error = s.SelectById(farmer.ID)
	if farmer_error != FarmerErrorNil {
		return nil, farmer_error
	}

	return farmer, FarmerErrorNil
}

// find the farmer by ID
func (s *Service) SelectById(farmer_id int) (farmer *Farmer, farmer_error FarmerError) {
	farmer = &Farmer{}
	err := s.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Find(farmer, "ID = ?", farmer_id).Error
	})

	if err != nil {
		return nil, FarmerError(err.Error())
	}

	if farmer.ID < 1 {
		return nil, FarmerErrorFarmerNotFound
	}

	farmer.CreatedDateJ = byteTimeStampToString(farmer.CreatedDateD)
	farmer.LastUpdatedDateJ = byteTimeStampToString(farmer.CreatedDateD)

	return farmer, FarmerErrorNil
}

// find the Farmer by email
func (s *Service) SelectByEmail(email string) (farmer *Farmer, farmer_error FarmerError) {
	farmer = &Farmer{}
	err := s.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Find(farmer, "Email = ?", email).Error
	})

	if err != nil {
		return nil, FarmerError(err.Error())
	}

	if farmer.ID < 1 {
		return nil, FarmerErrorFarmerNotFound
	}

	return farmer, FarmerErrorNil
}

func (s *Service) Login(email string, password string) (farmer *Farmer, user_error FarmerError) {
	farmer = &Farmer{}

	err := s.Db.Transaction(func(tx *gorm.DB) error {
		return tx.First(farmer, "email = ?", email).Error
	})

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, FarmerErrorInvalidLogin
		}
		return nil, FarmerError(err.Error())
	}

	return farmer, FarmerErrorNil
}

// converting the time stamp to a string so it can be used by the front end
func byteTimeStampToString(str []byte) string {
	time_str, _ := time.Parse("2006-01-02 15:04:05", string(str))
	t, err := time.Parse(time.RFC3339, time_str.Format(time.RFC3339))
	if err != nil {
		panic(err)
	}

	return t.Format("2006-01-02 15:04:05")
}
