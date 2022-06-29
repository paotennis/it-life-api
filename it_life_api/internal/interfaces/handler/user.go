package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/rikuhatano09/it-life-api/pkg/domain/contract"
	"github.com/rikuhatano09/it-life-api/pkg/usecase"
)

func CreateUser(context *gin.Context) {
	requestBody := contract.UserPostRequestBody{}

	err := context.ShouldBindJSON(&requestBody)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Bad request error: %s", err.Error()),
		})
		return
	}

	user, err := usecase.CreateUser(requestBody)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Internal server error: %s", err.Error()),
		})
		return
	}
	context.JSON(http.StatusOK, user)
}

func GetUsers(context *gin.Context) {
	users, err := usecase.GetUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Internal server error: %s", err.Error()),
		})
		return
	}
	context.JSON(http.StatusOK, users)
}

func GetUserByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil || context.Query("startDate") == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Bad request error: %s", err.Error()),
		})
		return
	}

	user, err := usecase.GetUserByID(id, context.Query("startDate"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Internal server error: %s", err.Error()),
		})
		return
	}
	context.JSON(http.StatusOK, user)
}
