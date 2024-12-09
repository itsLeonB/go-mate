package route

import (
	"github.com/gin-gonic/gin"
	"github.com/itsLeonB/go-mate/internal/provider"
)

func SetupRoutes(router *gin.Engine, handlers *provider.Handlers) {
	router.HandleMethodNotAllowed = true
	router.ContextWithFallback = true

	authRoutes := router.Group("/auth")
	authRoutes.POST("/register", handlers.Auth.HandleRegister())
	authRoutes.POST("/login", handlers.Auth.HandleLogin())

	recommendationRoutes := router.Group("/recommendations")
	recommendationRoutes.GET("", handlers.Recommendation.HandleGetUserRecommendations())
}
