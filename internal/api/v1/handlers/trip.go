package handlers

import (
	"fisherfan/internal/api/v1/models"
	"fisherfan/internal/api/v1/services"

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
	c.JSON(201, trip)
}

func (h *TripHandler) GetTrips(c *gin.Context) {
	filters := map[string]string{
		"userId":    c.Query("userId"),
		"title":     c.Query("title"),
		"tripType":  c.Query("tripType"),
		"startDate": c.Query("startDate"),
		"endDate":   c.Query("endDate"),
	}

	trips, err := h.service.GetTrips(filters)
	if err != nil {
		c.JSON(500, gin.H{"code": "SERVER_ERROR", "message": err.Error()})
		return
	}
	c.JSON(200, trips)
}

func (h *TripHandler) GetTripByID(c *gin.Context) {
	id := c.Param("id")
	trip, err := h.service.GetTripByID(id)
	if err != nil {
		c.JSON(404, gin.H{"code": "NOT_FOUND", "message": "Trip not found"})
		return
	}
	c.JSON(200, trip)
}

func (h *TripHandler) UpdateTrip(c *gin.Context) {
	id := c.Param("id")
	var trip models.Trip
	if err := c.ShouldBindJSON(&trip); err != nil {
		c.JSON(400, gin.H{"message": "Invalid data"})
		return
	}
	if err := h.service.UpdateTrip(id, &trip); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, trip)
}

func (h *TripHandler) DeleteTrip(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteTrip(id); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(204, nil)
}
