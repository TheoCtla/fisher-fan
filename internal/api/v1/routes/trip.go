package routes

import (
	"fisherfan/internal/api/v1/handlers"
	"fisherfan/internal/api/v1/repository"
	"fisherfan/internal/api/v1/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupTripRoutes(router *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewTripRepository(db)
	service := services.NewTripService(repo)
	handler := handlers.NewTripHandler(service)

	router.POST("/trips", handler.CreateTrip)
	router.GET("/trips", handler.GetTrips)
	router.GET("/trips/:id", handler.GetTripByID)
	router.PUT("/trips/:id", handler.UpdateTrip)
	router.DELETE("/trips/:id", handler.DeleteTrip)
}
