package main

import (
	//go:generate go tool oapi-codegen -config ./api/openapi/oapi-codegen-config.yaml ./api/openapi/api-spec.yaml
	"net/http"

	"github.com/gin-gonic/gin"
)

type hello struct {
	message string
}

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, hello{message: "Hello World!"})
}

func main() {
	router := gin.Default()
	router.GET("/hello", Hello)

	router.Run()
}
