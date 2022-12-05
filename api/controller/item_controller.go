package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/renasami/advents-2022-myjlab/api/service/application"
	"github.com/renasami/advents-2022-myjlab/api/service/application/auth_service"
)

func ItemController(r *gin.RouterGroup) {
	authMiddleware := auth_service.Jwt()
	r.Use(authMiddleware.MiddlewareFunc())

	r.POST("/", application.AddNewItem)
	r.DELETE("/delete", application.DeleteItem)
	r.PUT("/edit", application.EditItem)
	r.GET("/", application.GetAllItem)
}
