package application

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/renasami/advents-2022-myjlab/api/service/database"
)

func GetAllItem(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	fmt.Println(claims)
	email := claims["email"].(string)
	fmt.Println("items")
	result := database.GetAllItems(email)
	c.JSON(200, result)
	return
}
