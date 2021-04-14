package handler

import (
	"net/http"

	"github.com/FriedrichWilken/TinyGoExamples/DockerizeAnApi/repository"
	"github.com/gin-gonic/gin"
)

func CounterGet(repository *repository.GreetingRepository) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, repository)
	}
}
