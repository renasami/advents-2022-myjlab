package application

import (
	"github.com/gin-gonic/gin"
	"github.com/renasami/advents-2022-myjlab/api/models"
	"github.com/renasami/advents-2022-myjlab/api/service/database"
	"net/http"
)

func EditItem(c *gin.Context) {
	var body models.Item
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.EditItem(body)
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
	return
}
