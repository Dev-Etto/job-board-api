package handler

import (
	"fmt"
	"net/http"

	"github.com/Dev-Etto/job-board-api/schema"
	"github.com/gin-gonic/gin"
)


func DeleteOpeningHandler(context *gin.Context) {
	id := context.Param("id")

	if id == "" {
		sendError(context, http.StatusBadRequest, errParamIsRequered("id", "queryParameter").Error())
		return
	}

	opening := schema.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(context, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	if err := db.Delete(&opening).Error; err!= nil {
		sendError(context, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s", id))
		return
	}

	sendSuccess(context,"delete-opening", opening)
}

