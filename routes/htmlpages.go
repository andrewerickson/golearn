package routes

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

func GetIndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Main website",
	})
}
