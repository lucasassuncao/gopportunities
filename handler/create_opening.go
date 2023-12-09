package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasassuncao/gopportunities/schemas"
)

// @BasePath /api/v1

// @Summary Create Opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request Body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	err := request.Validate()
	if err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Link:     request.Link,
		Location: request.Location,
		Remote:   *request.Remote,
		Salary:   request.Salary,
	}

	logger.Infof("Request received: %+v", request)
	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
	}

	sendSuccess(ctx, "create-opening", opening)
}
