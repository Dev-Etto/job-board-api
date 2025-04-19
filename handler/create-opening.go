package handler

import (
	"net/http"

	"github.com/Dev-Etto/job-board-api/schema"
	"github.com/gin-gonic/gin"
)


func CreateOpeningHandler(context *gin.Context) {
	request := CreateOpeningRequest{}

	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validating error: %v", err.Error())
		sendError(context, http.StatusBadRequest, err.Error())
		return
	}


	opening := schema.Opening{
		Role: request.Role,
		Company: request.Company,
		Location: request.Location,
		Remote: *request.Remote,
		Link: request.Link,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opning: %v", err.Error())
		sendError(context, http.StatusInternalServerError, "error creating opening on database.")
		return
	}

	sendSuccess(context, "create-opening", schema.OpeningResponse{
		ID: opening.ID,
		Role: opening.Role,
		Company: opening.Company,
		Location: opening.Location,
		Remote: opening.Remote,
		Link: opening.Link,
	})
}
