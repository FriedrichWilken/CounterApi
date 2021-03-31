package main

import (
	"counter/handler"
	"counter/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	port := 8080
	repository := repository.New(0)
	engine := gin.Default()

	engine.GET("/counter", handler.CounterGet(repository))
	engine.PUT("/counter/increase", handler.IncreasePut(repository))
	engine.PUT("/counter/decrease", handler.DecreasePut(repository))

	engine.Run(":" + strconv.Itoa(port))
}
