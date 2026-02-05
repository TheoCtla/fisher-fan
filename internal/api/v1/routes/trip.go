package routes

import (
	"fisherman/internal/api/v1/handlers"
	"fisherman/internal/api/v1/repository"
	"fisherman/internal/api/v1/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupTripRoutes(router *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewTripRepository(db)
	service := services.NewTripService(repo)
	handler := handlers.NewTripHandler(service)

	router.POST("/trips", handler.CreateTrip)

	// router.POST("/boats", handler.CreateBoat)
	// router.GET("/boats", handler.GetBoats)
	// router.GET("/boats/:id", handler.GetBoatByID)
	// router.PUT("/boats/:id", handler.UpdateBoat)
	// router.DELETE("/boats/:id", handler.DeleteBoat)
	// router.GET("/boats/bbox", handler.GetBoatsByBBox)
}
