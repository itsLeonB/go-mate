package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsLeonB/go-mate/internal/model"
	"github.com/itsLeonB/go-mate/internal/service"
	"github.com/rotisserie/eris"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (ah *AuthHandler) HandleRegister() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := new(model.RegisterRequest)
		err := ctx.ShouldBindJSON(request)
		if err != nil {
			_ = ctx.Error(eris.Wrap(err, "error binding request"))
			return
		}

		err = ah.authService.Register(ctx, request)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, model.NewSuccessResponse(model.NewRegisterResponse()))
	}
}

func (ah *AuthHandler) HandleLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := new(model.LoginRequest)
		err := ctx.ShouldBindJSON(request)
		if err != nil {
			_ = ctx.Error(eris.Wrap(err, "error binding request"))
			return
		}

		response, err := ah.authService.Login(ctx, request)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, model.NewSuccessResponse(response))
	}
}
