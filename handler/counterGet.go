package handler

import (
	"counter/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CounterGet(repository *repository.CounterRepository) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, repository)
	}
}
