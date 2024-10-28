package http

import (
	api "herdwise/service/device"

	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Grouping all routers links into one through the http.go
func device_group(rg *gin.RouterGroup, d *api.Service) {
	device_group := rg.Group("/")

	device_group.GET("/deviceID/:id", func(c *gin.Context) {
		id_param := c.Param("id")
		id, err := strconv.Atoi(id_param)
		if err != nil {
			log.Println("Selected Device, invalid id: ")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid ID"})
			return
		}

		device, device_error := d.SelectById(id)
		if device_error != api.DeviceErrorNil {
			log.Println("Selected Device id failed:", device_error, "data[", device, "]")
			switch device_error {
			case api.DeviceErrorDeviceNotFound:
				c.JSON(http.StatusNotFound, gin.H{"error": string(api.DeviceErrorDeviceNotFound)})
				return
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": string(device_error)})
				return
			}
		}
		c.JSON(http.StatusOK, device)
	})

	device_group.GET("/list/:id", func(c *gin.Context) {
		farmer_id := c.Param("id")
		id, err := strconv.Atoi(farmer_id)
		if err != nil {
			log.Println("Select form, invalid refNo")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid refNo"})
			return
		}
		devices, device_error := d.List(id)
		if device_error != api.DeviceErrorNil {
			log.Println("List devices failed", device_error, "data[", devices, "]")
			c.JSON(http.StatusInternalServerError, gin.H{"error": string(device_error)})
			return
		}
		c.JSON(http.StatusOK, devices)
	})

	//http request to create devices
	device_group.POST("/create", func(c *gin.Context) {
		device := api.Device{}
		if err := c.ShouldBindJSON(&device); err != nil {
			log.Println("Create device failed to bind to JSON")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data sent"})
			return
		}

		log.Println(device)

		device_resp, device_error := d.Create(&device)
		if device_error != api.DeviceErrorNil {
			log.Println("Create device failed:", device_error, "data[", device, "]")
			c.JSON(http.StatusInternalServerError, gin.H{"error": string(device_error)})
			return
		}
		c.JSON(http.StatusCreated, device_resp)
	})

	//http request to find user by ref
	device_group.GET("/ref/:ref", func(c *gin.Context) {
		ref_param := c.Param("ref")

		//switch case error handling if user is not registered
		device, device_error := d.SelectByRef(ref_param)
		if device_error != api.DeviceErrorNil {
			log.Println("Select device ref failed:", device_error, "data[", device, "]")
			switch device_error {
			case api.DeviceErrorDeviceNotFound:
				c.JSON(http.StatusNotFound, gin.H{"error": string(api.DeviceErrorDeviceNotFound)})
				return
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": string(device_error)})
				return
			}
		}

		c.JSON(http.StatusOK, device)
	})

	//http request to update user by id
	device_group.PUT("/update/:id", func(c *gin.Context) {
		id_param := c.Param("id")
		id, err := strconv.Atoi(id_param)

		//error handling if user doesn't exist
		if err != nil {
			log.Println("Select user, invalid id")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid ID"})
			return
		}
		//error handling if user failed to be updated
		device := api.Device{}
		if err := c.ShouldBindJSON(&device); err != nil {
			log.Println("Update user failed to bind to JSON")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data sent"})
			return
		}

		device.ID = uint(id)

		device_resp, device_error := d.Update(&device)
		if device_error != api.DeviceErrorNil {
			c.JSON(http.StatusConflict, gin.H{"error": string(device_error)})
			return
		}

		c.JSON(http.StatusOK, device_resp)
	})

}
