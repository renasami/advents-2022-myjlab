package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/renasami/advents-2022-myjlab/api/service/application"
)

func ItemController(r *gin.RouterGroup) {
	r.POST("/", application.AddNewItem)
	r.DELETE("/",application.AddNewItem)
	r.PUT("/edit",application.AddNewItem)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
				"message": "Hello World",
		})
	})
}
