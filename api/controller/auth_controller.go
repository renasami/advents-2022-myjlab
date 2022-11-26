package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/renasami/advents-2022-myjlab/api/service/application/auth_service"
)

func AuthController(r *gin.RouterGroup) {
	authMiddleware := auth_service.Jwt()
	r.POST("/lll", authMiddleware.LoginHandler)
	r.POST("/login", auth_service.Login)
	r.POST("/register", auth_service.Register)
	r.GET("/refresh_token", authMiddleware.RefreshHandler)
	r.Use(authMiddleware.MiddlewareFunc())
	r.GET("/logout", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	r.GET("/hello", auth_service.HelloHandler)
}
