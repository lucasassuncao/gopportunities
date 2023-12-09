package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasassuncao/gopportunities/schemas"
)

// @BasePath /api/v1

// @Summary List Openings
// @Description List all job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningHandler(ctx *gin.Context) {
	opening := []schemas.Opening{}

	if err := db.Find(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-opening", opening)
}
