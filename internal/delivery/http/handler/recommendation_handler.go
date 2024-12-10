package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsLeonB/go-mate/internal/model"
	"github.com/itsLeonB/go-mate/internal/service"
	"github.com/rotisserie/eris"
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

func (rh *RecommendationHandler) HandleLogAction() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := new(model.LogActionRequest)
		err := ctx.ShouldBindJSON(&request)
		if err != nil {
			_ = ctx.Error(eris.Wrap(err, "error binding request"))
			return
		}

		response, err := rh.recommendationService.LogAction(ctx, request)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, model.NewSuccessResponse(response))
	}
}
