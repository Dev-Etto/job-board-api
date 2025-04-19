package handler

import (
	"net/http"

	"github.com/Dev-Etto/job-board-api/schema"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Find opening
// @Description Find a job opening by ID
// @Tags openings
// @Accept json
// @Produce json
// @Param id path string true "Opening ID"
// @Success 200 {object} OpeningResponseModel
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening/{id} [get]
func FindOpeningHandler(context *gin.Context) {
	id := context.Param("id")

	if id == "" {
		sendError(context, http.StatusNotFound, errParamIsRequered("id", "string").Error())
		return
	}

	opening := schema.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(context, http.StatusNotFound, "opening not found")
		return
	}

	sendSuccess(context, "find-opening", opening)	
}

