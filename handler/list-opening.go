package handler

import (
	"net/http"

	"github.com/Dev-Etto/job-board-api/schema"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List openings
// @Description List all job openings
// @Tags openings
// @Accept json
// @Produce json
// @Success 200 {object} OpeningListResponseModel
// @Failure 500 {object} ErrorResponse
// @Router /opening [get]
func ListOpeningHandler(context *gin.Context) {
	openings := []schema.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(context, http.StatusInternalServerError, "error listing openings")
		return
	}

			openingResponses := make([]schema.OpeningResponse, len(openings))

			for index, opening := range openings {
					openingResponses[index] = schema.OpeningResponse{
							ID:       opening.ID,
							Role:     opening.Role,
							Company:  opening.Company,
							Location: opening.Location,
							Remote:   opening.Remote,
							Link:     opening.Link,
					}
			}

	sendSuccess(context, "list-openings", openingResponses)
}
