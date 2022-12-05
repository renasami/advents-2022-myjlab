package application

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/renasami/advents-2022-myjlab/api/models"
	"github.com/renasami/advents-2022-myjlab/api/service/database"
	"net/http"
)

func AddNewItem(c *gin.Context) {
	var body models.Item
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	claims := jwt.ExtractClaims(c)
	email := claims["email"].(string)
	status := database.AddItem(body, email)
	if status == false {
		c.JSON(500, gin.H{
			"message": "err",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})

	fmt.Println(email)

	return
}
