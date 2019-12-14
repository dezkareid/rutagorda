package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		data := gin.H{
			"message": "Hello World",
		}
		c.JSON(http.StatusOK, data)
	})
	r.Run()
}
