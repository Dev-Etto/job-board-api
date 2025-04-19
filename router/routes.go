package router

import (
	"github.com/Dev-Etto/job-board-api/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine){
	handler.InitializeHandler()

	v1 := router.Group("/api/v1")

	{
		v1.GET("/opening/:id", handler.FindOpeningHandler)

		v1.GET("/opening", handler.ListOpeningHandler)

		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		v1.PUT("/opening", handler.UpdateOpeningHandler)
	}
}
