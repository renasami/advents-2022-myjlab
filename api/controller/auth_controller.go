package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/renasami/advents-2022-myjlab/api/service/application/auth_service"
)

func AuthController(r *gin.RouterGroup){
	r.POST("/login",func(c *gin.Context) {
		c.JSON(200, gin.H{
				"message": "Hello World",
		})
	})
	r.POST("/register",auth_service.Register)
}