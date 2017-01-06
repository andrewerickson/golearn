package main

import (
	// nice stuff, https://github.com/gin-gonic/gin
	"github.com/andrewerickson/golearn/routes"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/index", routes.GetIndexPage)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}
