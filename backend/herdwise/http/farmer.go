package http

import (
	"herdwise/middleware"
	api "herdwise/service/farmer"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func farmer_group(rg *gin.RouterGroup, f *api.Service) {
	farmer_group := rg.Group("/")

	farmer_group.GET("/farmerId/:id", func(c *gin.Context) {
		id_param := c.Param("id")
		id, err := strconv.Atoi(id_param)
		if err != nil {
			log.Println("Selected farmer, invalid id: ")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid ID"})
			return
		}

		farmer, farmer_error := f.SelectById(id)
		if farmer_error != api.FarmerErrorNil {
			log.Println("Selected farmer id failed:", farmer_error, "data[", farmer, "]")
			switch farmer_error {
			case api.FarmerErrorFarmerNotFound:
				c.JSON(http.StatusNotFound, gin.H{"error": string(api.FarmerErrorFarmerNotFound)})
				return
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": string(farmer_error)})
				return
			}
		}
		c.JSON(http.StatusOK, farmer)
	})

	//http request to find user by email
	farmer_group.GET("/email/:email", func(c *gin.Context) {
		email_param := c.Param("email")

		//switch case error handling if user is not registered
		farmer, farmer_error := f.SelectByEmail(email_param)
		if farmer_error != api.FarmerErrorNil {
			log.Println("Select farmer email failed:", farmer_error, "data[", farmer, "]")
			switch farmer_error {
			case api.FarmerErrorFarmerNotFound:
				c.JSON(http.StatusNotFound, gin.H{"error": string(api.FarmerErrorFarmerNotFound)})
				return
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": string(farmer_error)})
				return
			}
		}

		c.JSON(http.StatusOK, farmer)
	})

	//http request to create farmer
	farmer_group.POST("/create", func(c *gin.Context) {
		farmer := api.Farmer{}
		if err := c.ShouldBindJSON(&farmer); err != nil {
			log.Println("Create farmer failed to bind to JSON")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data sent"})
			return
		}

		log.Println(farmer)

		//switch case error handling for farmer who already registered
		farmer_resp, farmer_error := f.Create(&farmer)
		if farmer_error != api.FarmerErrorNil {
			log.Println("Create user failed:", farmer_error, "data[", farmer, "]")
			switch farmer_error {
			case api.FarmerErrorFarmerEmailExists:
				c.JSON(http.StatusConflict, gin.H{"error": string(api.FarmerErrorFarmerEmailExists)})
				return
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": string(farmer_error)})
				return
			}
		}

		c.JSON(http.StatusCreated, farmer_resp)
	})

	//http request to update farmer by id
	farmer_group.PUT("/update/:id", func(c *gin.Context) {
		id_param := c.Param("id")
		id, err := strconv.Atoi(id_param)

		//error handling if farmer doesn't exist
		if err != nil {
			log.Println("Select farmer, invalid id")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid ID"})
			return
		}
		//error handling if farmer failed to be updated
		farmer := api.Farmer{}
		if err := c.ShouldBindJSON(&farmer); err != nil {
			log.Println("Update farmer failed to bind to JSON")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data sent"})
			return
		}

		farmer.ID = id

		farmer_resp, farmer_error := f.Update(&farmer)
		if farmer_error != api.FarmerErrorNil {
			c.JSON(http.StatusConflict, gin.H{"error": string(farmer_error)})
			return
		}

		c.JSON(http.StatusOK, farmer_resp)
	})

	farmer_group.POST("/login", func(c *gin.Context) {
		var user *api.Farmer
		var input struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		//empid_param := c.Param("empid")
		if err := c.ShouldBindJSON(&input); err != nil {
			log.Println("Update user failed to bind to JSON")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data sent"})
			return
		}

		// Fetch user from database
		if err := f.Db.Where("Email = ?", input.Email).First(&user).Error; err != nil {
			log.Println("User not found with Email:", input.Email)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}

		user, user_error := f.Login(input.Email, input.Password)
		if user_error != api.FarmerErrorNil {
			log.Println("Login failed:", user_error)
			c.JSON(http.StatusUnauthorized, gin.H{"error": string(user_error)})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		token, _ := middleware.GenerateToken(uint(user.ID), c)
		log.Println("Generated token:", token)
		c.Writer.Header().Set("Authorization", "Bearer "+token)

		c.JSON(http.StatusOK, gin.H{
			"user":  user,
			"token": token,
		})
	})

}
