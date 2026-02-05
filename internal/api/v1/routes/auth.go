package routes

import (
	"fisherman/internal/api/v1/handlers"
	"fisherman/internal/api/v1/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAuthRoutes(router *gin.RouterGroup, database *gorm.DB) {
	// Créer le handler d'authentification
	authHandler := handlers.NewAuthHandler(database)

	// Routes publiques (sans authentification)
	router.POST("/auth/register", authHandler.Register)
	router.POST("/auth/login", authHandler.Login)
	router.POST("/auth/refresh", authHandler.RefreshToken)

	// Routes protégées (avec authentification)
	protected := router.Group("/auth")
	protected.Use(middleware.AuthRequired())
	{
		protected.GET("/me", authHandler.GetCurrentUser)
	}
}
