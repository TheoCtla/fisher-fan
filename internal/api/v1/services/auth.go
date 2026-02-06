package services

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fisherfan/internal/variables"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// AuthService gère la logique d'authentification
type AuthService struct{}

// NewAuthService crée une nouvelle instance d'AuthService
func NewAuthService() *AuthService {
	return &AuthService{}
}

// HashPassword hash un mot de passe avec bcrypt
func (s *AuthService) HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// VerifyPassword vérifie si un mot de passe correspond au hash
func (s *AuthService) VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// GenerateAccessToken génère un JWT access token
func (s *AuthService) GenerateAccessToken(userID, email string) (string, error) {
	// Parse la durée d'expiration
	duration, err := time.ParseDuration(variables.GlobalConfig.JWT.AccessTokenExpiry)
	if err != nil {
		duration = 24 * time.Hour // Fallback à 24h
	}

	// Créer les claims
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(duration).Unix(),
		"iat":     time.Now().Unix(),
	}

	// Créer le token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Signer le token avec le secret
	tokenString, err := token.SignedString([]byte(variables.GlobalConfig.JWT.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// GenerateRefreshToken génère un refresh token aléatoire sécurisé
func (s *AuthService) GenerateRefreshToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// ValidateToken valide un JWT token et retourne les claims
func (s *AuthService) ValidateToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Vérifier la méthode de signature
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("méthode de signature invalide")
		}
		return []byte(variables.GlobalConfig.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}

	return nil, errors.New("token invalide")
}

// ExtractUserIDFromToken extrait l'ID utilisateur d'un token valide
func (s *AuthService) ExtractUserIDFromToken(tokenString string) (string, error) {
	claims, err := s.ValidateToken(tokenString)
	if err != nil {
		return "", err
	}

	userID, ok := (*claims)["user_id"].(string)
	if !ok {
		return "", errors.New("user_id non trouvé dans le token")
	}

	return userID, nil
}
