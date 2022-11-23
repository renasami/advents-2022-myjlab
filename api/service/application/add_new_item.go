package application

import (
	"github.com/gin-gonic/gin"
)

func AddNewItem(c *gin.Context){
	c.JSON(200,gin.H{
		"message": "Hello World",
})
	return
}