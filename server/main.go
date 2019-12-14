package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerEndpoint(router *gin.Engine, path string, handler gin.HandlerFunc) (endpoint *gin.RouterGroup) {
	endpoint = router.Group(path)
	{
		endpoint.GET("/", handler)
	}
	return
}

func rootHandler(context *gin.Context) {
	response := gin.H{
		"api":     "ruta gorda",
		"version": 1,
	}
	context.JSON(http.StatusOK, response)
}

func getStablishments(context *gin.Context) {
	response := []string{"taquitos"}
	context.JSON(http.StatusOK, response)
}

func main() {
	router := gin.Default()
	router.GET("/", rootHandler)
	registerEndpoint(router, "establishments", getStablishments)
	router.Run()
}
