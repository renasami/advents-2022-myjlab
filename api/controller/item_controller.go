package controller

import "github.com/gin-gonic/gin"

func ItemController(r *gin.RouterGroup) {
	r.POST("/")
	r.DELETE("/")
	r.PUT("/edit")
	r.GET("/")
}
