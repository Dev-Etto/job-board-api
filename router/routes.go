package router

import (
	docs "github.com/Dev-Etto/job-board-api/docs"
	"github.com/Dev-Etto/job-board-api/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(router *gin.Engine){
	handler.InitializeHandler()

	basePath := "/api/v1"

	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)

	{
		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.GET("/opening", handler.ListOpeningHandler)

		v1.GET("/opening/:id", handler.FindOpeningHandler)

		v1.DELETE("/opening/:id", handler.DeleteOpeningHandler)

		v1.PUT("/opening/:id", handler.UpdateOpeningHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler) )
}
