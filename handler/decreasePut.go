package handler

import (
	"counter/models"
	"counter/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DecreasePut(repository *repository.CounterRepository) gin.HandlerFunc {
	return func(context *gin.Context) {
		body := models.DecreaseByDto{}
		context.Bind(&body)
		if body.Value == 0 {
			body.Value = 1
		}
		repository.DecreaseBy(body.Value)
		context.JSON(http.StatusOK, repository)
	}
}
