package handlers

import (
	"fisherman/internal/api/v1/models"
	"fisherman/internal/api/v1/services"

	"github.com/gin-gonic/gin"
)

type ReservationHandler struct {
	service *services.ReservationService
}

func NewReservationHandler(service *services.ReservationService) *ReservationHandler {
	return &ReservationHandler{service: service}
}

func (h *ReservationHandler) GetReservations(c *gin.Context) {
	filters := map[string]string{
		"userId": c.Query("userId"),
		"tripId": c.Query("tripId"),
		"date":   c.Query("date"),
	}

	reservations, err := h.service.GetAllReservations(filters)
	if err != nil {
		c.JSON(500, gin.H{"code": "SERVER_ERROR", "message": err.Error()})
		return
	}
	c.JSON(200, reservations)
}

func (h *ReservationHandler) GetReservationByID(c *gin.Context) {
	id := c.Param("id")
	reservation, err := h.service.GetReservationByID(id)
	if err != nil {
		c.JSON(404, gin.H{"code": "NOT_FOUND", "message": "Reservation not found"})
		return
	}
	c.JSON(200, reservation)
}

func (h *ReservationHandler) CreateReservation(c *gin.Context) {
	var reservation models.Reservation
	if err := c.ShouldBindJSON(&reservation); err != nil {
		c.JSON(400, gin.H{"code": "BAD_REQUEST", "message": "Invalid data"})
		return
	}
	err := h.service.CreateReservation(&reservation)
	if err != nil {
		c.JSON(500, gin.H{"code": "SERVER_ERROR", "message": err.Error()})
		return
	}
	c.JSON(201, reservation)
}

func (h *ReservationHandler) UpdateReservation(c *gin.Context) {
	id := c.Param("id")
	var reservation models.Reservation
	if err := c.ShouldBindJSON(&reservation); err != nil {
		c.JSON(400, gin.H{"code": "BAD_REQUEST", "message": "Invalid data"})
		return
	}
	err := h.service.UpdateReservation(id, &reservation)
	if err != nil {
		c.JSON(500, gin.H{"code": "SERVER_ERROR", "message": err.Error()})
		return
	}
	c.JSON(200, reservation)
}

func (h *ReservationHandler) DeleteReservation(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeleteReservation(id)
	if err != nil {
		c.JSON(500, gin.H{"code": "SERVER_ERROR", "message": err.Error()})
		return
	}
	c.JSON(204, nil)
}
