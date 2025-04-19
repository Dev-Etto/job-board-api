package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine){
	v1 := router.Group("/api/v1")

	{
		v1.GET("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message":"GET Opening"})
		})
		v1.GET("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message":"GET Openings"})
		})
		v1.POST("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message":"POST Opening"})
		})
		v1.DELETE("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message":"DELETE Opening"})
		})
		v1.PUT("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message":"PUT Opening"})
		})
	}
}
