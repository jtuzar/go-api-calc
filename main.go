package main

import (
	"go-api-calc/api"
	calc_handlers "go-api-calc/internal/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type hello struct {
	Message string
}

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, hello{Message: "Hello World"})
}

func main() {
	calcServer := calc_handlers.NewCalcServer()

	router := gin.Default()
	router.GET("/hello", Hello)

	api.RegisterHandlers(router, calcServer)

	router.Run()
}
