package rest_api

import (
	"fmt"
	"net/http"

	"github.com/brunoeduardodev/trellogo/internal/infra/config"
	"github.com/brunoeduardodev/trellogo/internal/infra/database/users"
	"github.com/gin-gonic/gin"
)

func Start(userRepository users.UserRepository) {
	router := gin.Default()
	v1 := router.Group("/api/v1")

	config := config.NewConfig()

	{

		v1.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "Alive"})
		})

		v1.GET("/users", func(ctx *gin.Context) {
			users, err := userRepository.List()

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "Could not connect to database",
				})
				return
			}

			ctx.JSON(http.StatusOK, users)
		})

		v1.POST("/users", func(ctx *gin.Context) {
			var input users.CreateUserInput

			err := ctx.ShouldBind(&input)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": fmt.Sprintf("Invalid request: %v", err),
				})

				return
			}

			err = userRepository.Create(input)

			fmt.Printf("%v", input.Name)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "Could not connect save user",
				})

				return
			}

			ctx.JSON(http.StatusOK, gin.H{"success": true})
		})

	}

	router.Run(":" + config.RestApi.Port)
}
