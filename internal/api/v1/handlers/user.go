package handlers

import (
	"errors"
	"net/http"

	"fisherfan/internal/api/v1/models"
	"fisherfan/internal/api/v1/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// GetUsers gère GET /v1/users (avec filtres query)
func (h *UserHandler) GetUsers(c *gin.Context) {
	// Extraction des filtres depuis la query string
	lastName := c.Query("lastName")
	firstName := c.Query("firstName")
	email := c.Query("email")
	status := c.Query("status")

	users, err := h.service.GetAllUsers(lastName, firstName, email, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des utilisateurs"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// CreateUser gère POST /v1/users
func (h *UserHandler) CreateUser(c *gin.Context) {

	var user models.User
	user.ID = ""

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON invalide: " + err.Error()})
		return
	}

	if err := h.service.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la création"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetUser gère GET /v1/users/:userId
func (h *UserHandler) GetUser(c *gin.Context) {
	idParam := c.Param("id") // Correspond au nom dans tes routes Gin

	user, err := h.service.GetUser(idParam)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Utilisateur non trouvé"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur serveur"})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser gère PUT /v1/users/:id
func (h *UserHandler) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")

	var userUpdate models.User
	if err := c.ShouldBindJSON(&userUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides"})
		return
	}

	// On s'assure que l'ID de l'URL est bien celui appliqué à l'objet
	userUpdate.ID = idParam

	updatedUser, err := h.service.UpdateUser(&userUpdate)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Utilisateur non trouvé"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la mise à jour"})
		}
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

// DeleteUser gère DELETE /v1/users/:id
func (h *UserHandler) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")

	if err := h.service.DeleteUser(idParam); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la suppression"})
		return
	}

	c.Status(http.StatusNoContent)
}
