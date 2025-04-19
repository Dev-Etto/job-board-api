package handler

import (
	"fmt"
	"net/http"

	"github.com/Dev-Etto/job-board-api/schema"
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


type ErrorResponse struct  {
	Message string `json:"message"`
	ErrorCode string `json:"error_code"`
}


type OpeningResponseModel struct {
	Message string `json:"message"`
	Data schema.OpeningResponse `json:"data"`
}

type OpeningListResponseModel struct {
	Message string `json:"message"`
	Data []schema.OpeningResponse `json:"data"`
}
