package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsLeonB/go-mate/internal/apperror"
	"github.com/itsLeonB/go-mate/internal/model"
	"github.com/rotisserie/eris"
)

func HandleError() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		if lastErr := ctx.Errors.Last(); lastErr != nil {
			err := lastErr.Err
			appError, ok := err.(*apperror.AppError)
			if !ok {
				originalErr := eris.Unwrap(err)
				log.Printf("unhandled error type: %T", originalErr)
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewErrorResponse(apperror.InternalServerError()))
				return
			}

			ctx.AbortWithStatusJSON(appError.HttpStatusCode, model.NewErrorResponse(appError))
		}
	}
}
