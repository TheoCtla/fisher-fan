package handlers

import (
	"fisherman/internal/api/v1/models"
	"fisherman/internal/api/v1/services"

	"github.com/gin-gonic/gin"
)

type TripHandler struct {
	service *services.TripService
}

func NewTripHandler(service *services.TripService) *TripHandler {
	return &TripHandler{service: service}
}

func (h *TripHandler) CreateTrip(c *gin.Context) {
	var trip models.Trip

	if err := c.ShouldBindJSON(&trip); err != nil {
		c.JSON(400, gin.H{"message": "Invalid data"})
		return
	}

	if err := h.service.CreateTrip(&trip); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
}
