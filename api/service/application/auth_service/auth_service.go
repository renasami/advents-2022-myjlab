package auth_service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renasami/advents-2022-myjlab/api/models"
	"github.com/renasami/advents-2022-myjlab/api/service/database"
	"golang.org/x/crypto/bcrypt"
)

var authRequest models.AuthRequest



func Login(c *gin.Context){
	c.JSON(200, gin.H{"message":"login"})
}

func Register(c *gin.Context) {
	if err := c.ShouldBindJSON(&authRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(authRequest)

	// 参考: https://zenn.dev/kou_pg_0131/articles/go-digest-and-compare-by-bcrypt
	password := []byte(authRequest.Password)
	hashed, _ := bcrypt.GenerateFromPassword(password, 10)

	err := bcrypt.CompareHashAndPassword(hashed, password)
	if err != nil {
		c.JSON(500, gin.H{"error":err.Error()})
	}
	auth_data := models.Auth {
		Username:authRequest.Username,
		Email :authRequest.Email,
		Password: string(hashed),
	}
	message, err := database.AddUser(auth_data)
	if err != nil {
		c.JSON(500, gin.H{"error":err.Error()})
	}
	// 他にいい書き方ないかなあ
	if message == "this email already exists" {
		c.JSON(400, gin.H{"message":message})
		return
	}
	c.JSON(200, gin.H{"message":message})
}

func Logout(c *gin.Context) {
	c.JSON(200, gin.H{"message":"loggedout"})
}