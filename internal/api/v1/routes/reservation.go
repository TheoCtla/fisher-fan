package routes

import (
	"fisherfan/internal/api/v1/handlers"
	"fisherfan/internal/api/v1/repository"
	"fisherfan/internal/api/v1/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupReservationRoutes(router *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewReservationRepository(db)
	service := services.NewReservationService(repo)
	handler := handlers.NewReservationHandler(service)

	router.POST("/reservations", handler.CreateReservation)
	router.GET("/reservations", handler.GetReservations)
	router.GET("/reservations/:id", handler.GetReservationByID)
	router.PUT("/reservations/:id", handler.UpdateReservation)
	router.DELETE("/reservations/:id", handler.DeleteReservation)
}
