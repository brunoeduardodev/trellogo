package rest_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start(port string) {
	router := gin.Default()

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Alive"})
	})

	router.Run(":" + port)
}
