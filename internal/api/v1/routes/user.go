package routes

import (
	"fisherfan/internal/api/v1/handlers"
	"fisherfan/internal/api/v1/middleware"
	"fisherfan/internal/api/v1/repository"
	"fisherfan/internal/api/v1/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(router *gin.RouterGroup, database *gorm.DB) {
	// 1. On crée le repository avec la DB
	repo := repository.NewUserRepository(database)

	// 2. On crée le service avec le repo
	service := services.NewUserService(repo)

	// 3. On crée le handler avec le service
	handler := handlers.NewUserHandler(service)

	// 4. Routes protégées par authentification
	protected := router.Group("/users")
	protected.Use(middleware.AuthRequired())
	{
		protected.GET("", handler.GetUsers)
		protected.GET("/:id", handler.GetUser)
		protected.PUT("/:id", handler.UpdateUser)
		protected.DELETE("/:id", handler.DeleteUser)
	}

	// Route publique pour la création (backward compatibility)
	// Note: Il est recommandé d'utiliser /auth/register à la place
	router.POST("/users", handler.CreateUser)
}
