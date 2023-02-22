package main

import (
	"fmt"
	"net/http"
	"os"
	"shopeecode/jwt-api/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var hmacSampleSecret []byte

type RegisterBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Birthday string `json:"birthday" binding:"required"`
	Tel      string `json:"tel" binding:"required"`
}
type Users struct {
	gorm.Model
	Username string
	Password string
	Fullname string `json:"Fullname"`
	Address  string `json:"Address"`
	Birthday string
	Tel      string
	Avatar   string `json:"Avatar"`
}
type LoginBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}
	dsn := os.Getenv("MYPOSTGRES_DNS")
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migtrate the schema
	Db.AutoMigrate(&Users{})
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization")
	r.Use(cors.New(config))
	///Register user
	r.POST("/register", func(c *gin.Context) {
		var json RegisterBody
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//   Check User Exists
		var userExist Users
		Db.Where("username =  ?", json.Username).First(&userExist)
		if userExist.ID > 0 {
			c.JSON(http.StatusOK, gin.H{"Status": "Error", "Message": "User Exists"})
			return
		}
		//Create User
		encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(json.Password), 10)
		user := Users{Username: json.Username, Password: string(encryptedPassword), Fullname: json.Fullname, Address: json.Address, Birthday: json.Birthday, Tel: json.Tel}
		Db.Create(&user)
		if user.ID > 0 {
			c.JSON(http.StatusOK, gin.H{"Status": "OK", "Message": "User Registed", "User ID": user.ID})
			fmt.Printf("Read body: %+v\n", user)
		} else {
			c.JSON(http.StatusOK, gin.H{"Status": "Error", "Message": "User Not Registed"})
		}
	})
	/// User login
	r.POST("/login", func(c *gin.Context) {
		var json LoginBody
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//   Check User Exists
		var userExist Users
		Db.Where("username =  ?", json.Username).First(&userExist)
		if userExist.ID == 0 {
			c.JSON(http.StatusOK, gin.H{"Status": "Error", "Message": "User Does not Exists"})
			return
		}
		err := bcrypt.CompareHashAndPassword([]byte(userExist.Password), []byte(json.Password))
		if err == nil {
			hmacSampleSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"userId": userExist.ID,
				"exp":    time.Now().Add(time.Minute * 1).Unix(), //token have 1 min for expire
			})
			tokenString, err := token.SignedString(hmacSampleSecret)
			fmt.Println(tokenString, err)
			c.JSON(http.StatusOK, gin.H{"Status": "OK", "Message": "Login Success", "token": tokenString, "Username": json.Username})

		} else {
			c.JSON(http.StatusOK, gin.H{"Status": "Error", "Message": "Login Failed"})
		}

	})

	/// User Information
	authorized := r.Group("/users", middleware.JWTAuthen())
	authorized.GET("/readall", func(c *gin.Context) {
		var users []Users
		Db.Find(&users)
		c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "User information read success", "users": users})
	})
	// User Profile
	authorized.GET("/profile", func(c *gin.Context) {
		userId := c.MustGet("userId").(float64)
		var user []Users
		Db.First(&user, userId)
		c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "User information read success", "user": user})
	})
	authorized.PUT("/update", func(c *gin.Context) {
		var json Users
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		userId := c.MustGet("userId").(float64)

		var user []Users
		newUser := Users{Fullname: json.Fullname, Address: json.Address, Avatar: json.Avatar}
		Db.First(&user, userId).Updates(&newUser)

		if err == nil {
			Db.Save(&user)
			c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "User Updated Already Done!", "user": userId, `Updated Info`: user})
		} else {
			c.JSON(http.StatusOK, gin.H{"Status": "Error", "Message": "User Updated Failed"})
		}
	})
	authorized.DELETE("/delete", func(c *gin.Context) {
		userId := c.MustGet("userId").(float64)
		var user []Users
		Db.Delete(&user, userId)

		if err == nil {
			c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "User Deleted Already Done!", "user": userId, `Updated Info`: user})
		} else {
			c.JSON(http.StatusOK, gin.H{"Status": "Error", "Message": "User Updated Failed"})
		}
	})
	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
