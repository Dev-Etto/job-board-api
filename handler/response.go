package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(context *gin.Context, code int, msg string) {
	context.Header("Content-type", "application/json")
	
	context.JSON(code, gin.H{
		"message":msg,
		"errorCode":code,
	})
}


func sendSuccess(context *gin.Context, op string, data interface{}) {
	context.Header("Content-type","application/json")

	context.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":data,
	})
}
