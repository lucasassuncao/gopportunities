package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasassuncao/gopportunities/schemas"
)

func ListOpeningHandler(ctx *gin.Context) {
	opening := []schemas.Opening{}

	if err := db.Find(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-opening", opening)
}
