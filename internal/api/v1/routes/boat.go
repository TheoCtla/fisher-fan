package routes

import (
	"fisherfan/internal/api/v1/handlers"
	"fisherfan/internal/api/v1/repository"
	"fisherfan/internal/api/v1/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupBoatRoutes(router *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewBoatRepository(db)
	service := services.NewBoatService(repo)
	handler := handlers.NewBoatHandler(service)

	router.POST("/boats", handler.CreateBoat)
	router.GET("/boats", handler.GetBoats)
	router.GET("/boats/:id", handler.GetBoatByID)
	router.PUT("/boats/:id", handler.UpdateBoat)
	router.DELETE("/boats/:id", handler.DeleteBoat)
	router.GET("/boats/bbox", handler.GetBoatsByBBox)
}
