package routes

import (
	"gopkg.in/gin-gonic/gin.v1"
)

func GetPingResponse(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
