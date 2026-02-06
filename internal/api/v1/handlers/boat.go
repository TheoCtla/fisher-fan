package handlers

import (
	"fisherfan/internal/api/v1/models"
	"fisherfan/internal/api/v1/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BoatHandler struct {
	service *services.BoatService
}

func NewBoatHandler(service *services.BoatService) *BoatHandler {
	return &BoatHandler{service: service}
}

func (h *BoatHandler) GetBoats(c *gin.Context) {
	filters := map[string]string{
		"userId":   c.Query("userId"),
		"y1":       c.Query("y1"),
		"y2":       c.Query("y2"),
		"x1":       c.Query("x1"),
		"x2":       c.Query("x2"),
		"name":     c.Query("name"),
		"brand":    c.Query("brand"),
		"boatType": c.Query("boatType"),
		"homePort": c.Query("homePort"),
	}

	boats, err := h.service.GetAllBoats(filters)
	if err != nil {
		c.JSON(500, gin.H{"code": "SERVER_ERROR", "message": err.Error()})
		return
	}
	c.JSON(200, boats)
}

func (h *BoatHandler) GetBoatsByBBox(c *gin.Context) {
	latMin, _ := strconv.ParseFloat(c.Query("lat_min"), 64)
	latMax, _ := strconv.ParseFloat(c.Query("lat_max"), 64)
	lonMin, _ := strconv.ParseFloat(c.Query("lon_min"), 64)
	lonMax, _ := strconv.ParseFloat(c.Query("lon_max"), 64)

	boats, err := h.service.GetByBBox(latMin, latMax, lonMin, lonMax)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, boats)
}

func (h *BoatHandler) CreateBoat(c *gin.Context) {
	var boat models.Boat
	boat.ID = ""

	if err := c.ShouldBindJSON(&boat); err != nil {
		c.JSON(400, gin.H{"message": "Invalid data"})
		return
	}
	if err := h.service.CreateBoat(&boat); err != nil {
		c.JSON(422, gin.H{"message": "Could not create boat"})
		return
	}
	c.JSON(201, boat)
}

func (h *BoatHandler) GetBoatByID(c *gin.Context) {
	id := c.Param("id")
	boat, err := h.service.GetBoatByID(id)
	if err != nil {
		c.JSON(404, gin.H{"message": "Boat not found"})
		return
	}
	c.JSON(200, boat)
}

func (h *BoatHandler) UpdateBoat(c *gin.Context) {
	id := c.Param("id")
	var boat models.Boat
	if err := c.ShouldBindJSON(&boat); err != nil {
		c.JSON(400, gin.H{"message": "Invalid data"})
		return
	}
	if err := h.service.UpdateBoat(id, &boat); err != nil {
		c.JSON(500, gin.H{"message": "Update failed"})
		return
	}
	c.JSON(200, boat)
}

func (h *BoatHandler) DeleteBoat(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteBoat(id); err != nil {
		c.JSON(500, gin.H{"message": "Delete failed"})
		return
	}
	c.Status(204)
}
