package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rikuhatano09/it-life-api/pkg/domain/contract"
	"github.com/rikuhatano09/it-life-api/pkg/usecase"
)

func CreateEvent(context *gin.Context) {
	requestBody := contract.EventPostRequestBody{}

	err := context.ShouldBindJSON(&requestBody)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Bad request error: %s", err.Error()),
		})
		return
	}

	event, _, err := usecase.CreateEvent(requestBody)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Internal server error: %s", err.Error()),
		})
		return
	}
	context.JSON(http.StatusOK, event)
}
