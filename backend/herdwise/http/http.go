package http

import (
	"fmt"
	//"time"

	"herdwise/service/device"
	"herdwise/service/farmer"

	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Host          string
	Port          string
	FarmerService *farmer.Service
	DeviceService *device.Service
}

func (c *Config) Addr() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

func Start(c *Config) {
	router := gin.Default()

	f := router.Group("/farmer")
	farmer_group(f, c.FarmerService)

	d := router.Group("/device")
	device_group(d, c.DeviceService)
	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:4200"},                   // Angular's dev server address
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},            // Allowed HTTP methods
	// 	AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"}, // Allowed headers
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour, // Preflight cache duration
	// }))
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization,Origin")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	router.Run(c.Addr())
}
