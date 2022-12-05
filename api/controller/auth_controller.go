package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/renasami/advents-2022-myjlab/api/service/application/auth_service"
)

func AuthController(r *gin.RouterGroup) {
	authMiddleware := auth_service.Jwt()
	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/register", auth_service.Register)
	r.Use(authMiddleware.MiddlewareFunc())
	r.GET("/refresh_token", authMiddleware.RefreshHandler)
	r.GET("/hello", auth_service.HelloHandler)
	r.GET("/logout", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
}
