package main

import (
	"strconv"

	"github.com/FriedrichWilken/TinyGoExamples/DockerizeAnApi/handler"
	"github.com/FriedrichWilken/TinyGoExamples/DockerizeAnApi/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	port := 8080
	repository := repository.New(1, "Bonjour")
	engine := gin.Default()

	engine.GET("/greeting", handler.CounterGet(repository))

	engine.Run(":" + strconv.Itoa(port))
}
