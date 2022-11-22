package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func GetAllRouter() *gin.Engine {
	router := gin.Default()
	//先に定義しないと反映されないので注意
	router.Use(cors.New(cors.Config{
		//今回は全てのアクセスを許容
		AllowOrigins: []string{
			"*",
		},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Access-Control-Allow-Headers", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge: 6 * time.Hour,
	}))
	router.SetTrustedProxies([]string{"host.docker.internal"})
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	items := router.Group("/items")
	{
		ItemController(items)
	}
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
				"message": "Hello World",
		})
	})

	return router
}