package auth_service

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renasami/advents-2022-myjlab/api/models"
	"github.com/renasami/advents-2022-myjlab/api/service/database"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var loginRequest models.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := database.GetUserByEmail(loginRequest.Email)
	if err != nil {
		log.Errorf("Get ERROR %v", err)
		return
	}
	usr := *result
	pas := []byte(loginRequest.Password)
	err = bcrypt.CompareHashAndPassword([]byte(usr.Password), pas)
	if err != nil {
		log.Errorf("err %v", err)
		return
	}
	c.JSON(200, gin.H{"message": "login"})
}

func Register(c *gin.Context) {
	var authRequest models.AuthRequest
	if err := c.ShouldBindJSON(&authRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(authRequest)

	// 参考: https://zenn.dev/kou_pg_0131/articles/go-digest-and-compare-by-bcrypt
	password := []byte(authRequest.Password)
	hashed, _ := bcrypt.GenerateFromPassword(password, 10)

	err := bcrypt.CompareHashAndPassword(hashed, password)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	auth_data := models.Auth{
		Username: authRequest.Username,
		Email:    authRequest.Email,
		Password: string(hashed),
	}
	message, err := database.AddUser(auth_data)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	// 他にいい書き方ないかなあ
	if message == "this email already exists" {
		c.JSON(400, gin.H{"message": message})
		return
	}
	c.JSON(200, gin.H{"message": message})
}

func Logout(c *gin.Context) {
	c.JSON(200, gin.H{"message": "loggedout"})
}
