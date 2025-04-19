package handler

import (
	"fmt"
	"net/http"

	"github.com/Dev-Etto/job-board-api/schema"
	"github.com/gin-gonic/gin"
)

// @basePath /api/v1

// @Summary Update opening
// @Description Update an existing job opening
// @Tags openings
// @Accept json
// @Produce json
// @Param id path string true "Opening ID"
// @Param opening body UpdateOpeningRequest true "Update opening request"
// @Success 200 {object} OpeningResponseModel
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening/{id} [put]
func UpdateOpeningHandler(context *gin.Context) {
	request := UpdateOpeningRequest{}

	if err := context.BindJSON(&request); err != nil {
			logger.Errorf("error binding request: %v", err.Error())
			sendError(context, http.StatusBadRequest, "invalid request body")
			return
	}

	if err := request.Validate(); err != nil {
			logger.Errorf("validation error: %v", err.Error())
			sendError(context, http.StatusBadRequest, err.Error())
			return
	}

	id := context.Param("id")

	if id == "" {
			sendError(context, http.StatusBadRequest, errParamIsRequered("id", "queryParams").Error())
			return
	}

	opening := schema.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
			sendError(context, http.StatusNotFound, fmt.Sprintf("opening with id: %v not found", id))
			return
	}

	if request.Role != "" {
			opening.Role = request.Role
	}

	if request.Company != "" {
			opening.Company = request.Company
	}

	if request.Location != "" {
			opening.Location = request.Location
	}

	if request.Link != "" {
			opening.Link = request.Link
	}

	if request.Remote != nil {
			opening.Remote = *request.Remote
	}

	if request.Salary > 0 {
			opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
			logger.Errorf("error updating opening: %v", err.Error())
			sendError(context, http.StatusInternalServerError, "error updating opening")
			return
	}

	sendSuccess(context, "update-opening", schema.OpeningResponse{
		ID:       opening.ID,
		Role:     opening.Role,
		Company:  opening.Company,
		Location: opening.Location,
		Link:     opening.Link,
		Remote:   opening.Remote,
		Salary:   opening.Salary,
	})
}
