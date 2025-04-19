package handler

import (
	"fmt"
	"net/http"

	"github.com/Dev-Etto/job-board-api/schema"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Delete opening
// @Description Delete a new job opening
// @Tags openings
// @Accept json
// @Produce json
// @Param id path string true "Opening ID"
// @Success 200 {object} OpeningResponseModel
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening/{id} [DELETE]
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

	sendSuccess(context,"delete-opening", schema.OpeningResponse{
		ID:       opening.ID,
		Role:     opening.Role,
		Company:  opening.Company,
		Location: opening.Location,
		Link:     opening.Link,
		Remote:   opening.Remote,
		Salary:   opening.Salary,
	})
}

