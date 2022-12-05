package application

import (
	"github.com/gin-gonic/gin"
	"github.com/renasami/advents-2022-myjlab/api/service/database"
)

func DeleteItem(c *gin.Context) {
	id := c.Query("id")
	result := database.DeleteItemWithId(id)
	c.JSON(200,
		gin.H{
			"msg": result,
		})
	return
}
