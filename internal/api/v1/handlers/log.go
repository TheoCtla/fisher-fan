package handlers

import (
	"errors"
	"fisherman/internal/api/v1/models"
	"fisherman/internal/api/v1/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LogHandler struct {
	service *services.LogService
}

func NewLogHandler(service *services.LogService) *LogHandler {
	return &LogHandler{service: service}
}

func (h *LogHandler) GetLogByUserID(c *gin.Context) {
	userID := c.Param("id")
	log, err := h.service.GetByUserID(userID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"code": "NOT_FOUND", "message": "Aucun journal pour cet utilisateur"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"code": "SERVER_ERROR", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, log)
}

func (h *LogHandler) CreateLogByUserID(c *gin.Context) {
	var log models.Log
	log.UserID = c.Param("id")

	if err := c.ShouldBindJSON(&log); err != nil {
		c.JSON(400, gin.H{"message": "Invalid JSON data", "error": err.Error()})
		return
	}

	if err := h.service.CreateLog(&log); err != nil {
		c.JSON(500, gin.H{"code": "DB_ERROR", "message": err.Error()})
		return
	}
	c.JSON(201, log)
}

func (h *LogHandler) GetPage(c *gin.Context) {
	pageId := c.Param("page_id")
	page, err := h.service.GetPage(pageId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Page non trouvée"})
		return
	}
	c.JSON(http.StatusOK, page)
}

func (h *LogHandler) UpdatePage(c *gin.Context) {
	pageId := c.Param("page_id")
	var page models.Page
	if err := c.ShouldBindJSON(&page); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Data invalide"})
		return
	}
	if err := h.service.UpdatePage(pageId, &page); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Echec mise à jour"})
		return
	}
	c.JSON(http.StatusOK, page)
}

func (h *LogHandler) DeletePage(c *gin.Context) {
	pageId := c.Param("page_id")
	if err := h.service.DeletePage(pageId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Echec suppression"})
		return
	}
	c.Status(http.StatusNoContent)
}
