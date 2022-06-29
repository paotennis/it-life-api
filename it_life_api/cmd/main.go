package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/rikuhatano09/it-life-api/internal/interfaces/handler"
)

func main() {
	engine := gin.Default()

	// Set CORS config.
	engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"*",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		MaxAge: 24 * time.Hour,
	}))

	engine.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello ITLife!!",
		})
	})

	engine.GET("/users", handler.GetUsers)
	engine.POST("/users", handler.CreateUser)

	engine.GET("/users/:id", handler.GetUserByID)

	engine.POST("/events", handler.CreateEvents)

	engine.Run(":" + os.Getenv("PORT"))
}
