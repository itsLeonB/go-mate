package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsLeonB/go-mate/internal/model"
	"github.com/itsLeonB/go-mate/internal/service"
)

type RecommendationHandler struct {
	recommendationService service.RecommendationService
}

func NewRecommendationHandler(recommendationService service.RecommendationService) *RecommendationHandler {
	return &RecommendationHandler{recommendationService: recommendationService}
}

func (rh *RecommendationHandler) HandleGetUserRecommendations() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response, err := rh.recommendationService.GetUserRecommendations(ctx)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, model.NewSuccessResponse(response))
	}
}
