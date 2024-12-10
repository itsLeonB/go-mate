package route

import (
	"github.com/gin-gonic/gin"
	"github.com/itsLeonB/go-mate/internal/delivery/http/middleware"
	"github.com/itsLeonB/go-mate/internal/provider"
)

func SetupRoutes(router *gin.Engine, handlers *provider.Handlers, utils *provider.Utils) {
	router.HandleMethodNotAllowed = true
	router.ContextWithFallback = true

	authRoutes := router.Group("/auth")
	authRoutes.POST("/register", handlers.Auth.HandleRegister())
	authRoutes.POST("/login", handlers.Auth.HandleLogin())

	recommendationRoutes := router.Group("/recommendations", middleware.Authorize(utils.JWT))
	recommendationRoutes.GET("", handlers.Recommendation.HandleGetUserRecommendations())
}
