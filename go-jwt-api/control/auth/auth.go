package auth

import (
	"fmt"
	"net/http"
	"shopeecode/jwt-api/orm"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Birthday string `json:"birthday" binding:"required"`
	Tel      string `json:"tel" binding:"required"`
}

func Register(c *gin.Context) {
	var json RegisterBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//   Check User Exists
	var userExist orm.Users
	orm.Db.Where("username =  ?", json.Username).First(&userExist)
	if userExist.ID > 0 {
		c.JSON(http.StatusOK, gin.H{"Status": "Error", "Message": "User Exists"})
		return
	}
	//Create User
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(json.Password), 10)
	user := orm.Users{Username: json.Username, Password: string(encryptedPassword), Fullname: json.Fullname, Address: json.Address, Birthday: json.Birthday, Tel: json.Tel}
	orm.Db.Create(&user)
	if user.ID > 0 {
		c.JSON(http.StatusOK, gin.H{"Status": "OK", "Message": "User Registed", "User ID": user.ID})
		fmt.Printf("Read body: %+v\n", user)
	} else {
		c.JSON(http.StatusOK, gin.H{"Status": "Error", "Message": "User Not Registed"})
	}
}
