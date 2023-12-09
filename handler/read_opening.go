package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasassuncao/gopportunities/schemas"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "string").Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
	}

	sendSuccess(ctx, "show-opening", opening)
}
