package handlers

import (
	"net/http"
	"strconv"

	"fisherman/internal/api/v1/models"
	"fisherman/internal/api/v1/services"

	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserHandler structure qui contient le service nécessaire à la logique
type UserHandler struct {
	service *services.UserService
}

// NewUserHandler est le constructeur du handler (utilisé dans le main.go)
func NewUserHandler(s *services.UserService) *UserHandler {
	return &UserHandler{service: s}
}

// CreateUser gère la route POST /users
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User

	// 1. On tente de binder le JSON reçu vers notre struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON invalide: " + err.Error()})
		return
	}

	// 2. On appelle le service pour la création
	if err := h.service.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la création de l'utilisateur"})
		return
	}

	// 3. Succès
	c.JSON(http.StatusCreated, user)
}

// GetUser gère la route GET /users/:id
func (h *UserHandler) GetUser(c *gin.Context) {
	// 1. On récupère l'ID dans l'URL et on le convertit en int
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide, doit être un nombre"})
		return
	}

	// 2. On appelle le service
	user, err := h.service.GetUser(uint(id))
	if err != nil {
		// On vérifie si c'est une erreur "pas trouvé" ou une erreur serveur
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Utilisateur non trouvé"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur serveur"})
		}
		return
	}

	// 3. On renvoie l'utilisateur trouvé
	c.JSON(http.StatusOK, user)
}
