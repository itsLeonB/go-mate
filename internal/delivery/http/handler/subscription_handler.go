package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsLeonB/go-mate/internal/model"
	"github.com/itsLeonB/go-mate/internal/service"
	"github.com/rotisserie/eris"
)

type SubscriptionHandler struct {
	subscriptionService service.SubscriptionService
}

func NewSubscriptionHandler(subscriptionService service.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{
		subscriptionService: subscriptionService,
	}
}

func (sh *SubscriptionHandler) HandleAddSubscription() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := &model.NewSubscriptionRequest{}
		err := ctx.ShouldBindJSON(request)
		if err != nil {
			_ = ctx.Error(eris.Wrap(err, "error binding request"))
			return
		}

		response, err := sh.subscriptionService.AddSubscription(ctx, request)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, model.NewSuccessResponse(response))
	}
}
