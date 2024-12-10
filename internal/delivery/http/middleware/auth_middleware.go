package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/itsLeonB/go-mate/internal/appconstant"
	"github.com/itsLeonB/go-mate/internal/apperror"
	"github.com/itsLeonB/go-mate/internal/util"
)

func Authorize(jwt util.JWT) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			_ = ctx.Error(apperror.MissingTokenError())
			ctx.Abort()
			return
		}

		splits := strings.Split(token, " ")
		if len(splits) != 2 || splits[0] != "Bearer" {
			_ = ctx.Error(apperror.InvalidTokenError())
			ctx.Abort()
			return
		}

		claims, err := jwt.VerifyToken(splits[1])
		if err != nil {
			_ = ctx.Error(err)
			ctx.Abort()
			return
		}

		ctx.Set(appconstant.ContextUserID, claims.Data[appconstant.ContextUserID])
		ctx.Next()
	}
}
