package routes

import (
	"fisherman/internal/api/v1/handlers"
	"fisherman/internal/api/v1/repository"
	"fisherman/internal/api/v1/services"

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

	// 4. On dit à Gin : "Quand quelqu'un appelle POST /users, utilise la méthode du handler handler"
	router.POST("/users", handler.CreateUser)
	router.GET("/users", handler.GetUsers)
	router.GET("/users/:id", handler.GetUser)
	router.PUT("/users/:id", handler.UpdateUser)
	router.DELETE("/users/:id", handler.DeleteUser)
}
