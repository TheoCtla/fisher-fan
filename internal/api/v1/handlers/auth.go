package handlers

import (
	"errors"
	"fisherman/internal/api/v1/models"
	"fisherman/internal/api/v1/repository"
	"fisherman/internal/api/v1/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthHandler struct {
	authService *services.AuthService
	userRepo    *repository.UserRepository
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{
		authService: services.NewAuthService(),
		userRepo:    repository.NewUserRepository(db),
	}
}

// RegisterRequest représente les données d'inscription
type RegisterRequest struct {
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=8"`
	FirstName   string `json:"firstName" binding:"required"`
	LastName    string `json:"lastName" binding:"required"`
	Status      string `json:"status" binding:"required,oneof=individual professional"`
	BirthDate   string `json:"birthDate"`
	BoatLicense string `json:"boatLicense"`
	// Champs professionnels optionnels
	CompanyName  string `json:"companyName"`
	ActivityType string `json:"activityType"`
	SiretNumber  string `json:"siretNumber"`
	RcNumber     string `json:"rcNumber"`
}

// LoginRequest représente les données de connexion
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// RefreshRequest représente la demande de refresh token
type RefreshRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

// AuthResponse représente la réponse d'authentification
type AuthResponse struct {
	AccessToken  string       `json:"accessToken"`
	RefreshToken string       `json:"refreshToken"`
	User         *models.User `json:"user"`
}

// Register gère POST /api/v1/auth/register
func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides: " + err.Error()})
		return
	}

	// Vérifier si l'email existe déjà
	existingUser, _ := h.userRepo.FindByID(req.Email)
	if existingUser.Email != "" {
		c.JSON(http.StatusConflict, gin.H{"error": "Un utilisateur avec cet email existe déjà"})
		return
	}

	// Hasher le mot de passe
	hashedPassword, err := h.authService.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors du hashage du mot de passe"})
		return
	}

	// Générer un refresh token
	refreshToken, err := h.authService.GenerateRefreshToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la génération du refresh token"})
		return
	}

	// Créer l'utilisateur
	user := &models.User{
		ID:           uuid.New().String(),
		Email:        req.Email,
		Password:     hashedPassword,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Status:       req.Status,
		BirthDate:    req.BirthDate,
		BoatLicense:  req.BoatLicense,
		CompanyName:  req.CompanyName,
		ActivityType: req.ActivityType,
		SiretNumber:  req.SiretNumber,
		RcNumber:     req.RcNumber,
		RefreshToken: refreshToken,
	}

	if err := h.userRepo.Create(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la création de l'utilisateur"})
		return
	}

	// Générer l'access token
	accessToken, err := h.authService.GenerateAccessToken(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la génération du token"})
		return
	}

	c.JSON(http.StatusCreated, AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         user,
	})
}

// Login gère POST /api/v1/auth/login
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides: " + err.Error()})
		return
	}

	// Rechercher l'utilisateur par email
	users, err := h.userRepo.FindAll("", "", req.Email, "")
	if err != nil || len(users) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email ou mot de passe incorrect"})
		return
	}

	user := users[0]

	// Vérifier le mot de passe
	if err := h.authService.VerifyPassword(user.Password, req.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email ou mot de passe incorrect"})
		return
	}

	// Générer un nouveau refresh token
	refreshToken, err := h.authService.GenerateRefreshToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la génération du refresh token"})
		return
	}

	// Mettre à jour le refresh token dans la base
	user.RefreshToken = refreshToken
	if err := h.userRepo.Update(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la mise à jour du token"})
		return
	}

	// Générer l'access token
	accessToken, err := h.authService.GenerateAccessToken(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la génération du token"})
		return
	}

	c.JSON(http.StatusOK, AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         &user,
	})
}

// RefreshToken gère POST /api/v1/auth/refresh
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Refresh token manquant"})
		return
	}

	// Rechercher l'utilisateur avec ce refresh token
	users, err := h.userRepo.FindAll("", "", "", "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur serveur"})
		return
	}

	var user *models.User
	for i := range users {
		if users[i].RefreshToken == req.RefreshToken {
			user = &users[i]
			break
		}
	}

	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token invalide"})
		return
	}

	// Générer un nouvel access token
	accessToken, err := h.authService.GenerateAccessToken(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la génération du token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken": accessToken,
	})
}

// GetCurrentUser gère GET /api/v1/auth/me
func (h *AuthHandler) GetCurrentUser(c *gin.Context) {
	// Récupérer l'ID utilisateur depuis le contexte (injecté par le middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Non authentifié"})
		return
	}

	// Récupérer l'utilisateur
	user, err := h.userRepo.FindByID(userID.(string))
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
