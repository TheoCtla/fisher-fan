package middleware

import (
	"fisherman/internal/api/v1/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthRequired est un middleware qui vérifie la présence et la validité du JWT
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Récupérer le header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token d'authentification manquant"})
			c.Abort()
			return
		}

		// Vérifier le format "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Format du token invalide. Utilisez: Bearer <token>"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Valider le token
		authService := services.NewAuthService()
		claims, err := authService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token invalide ou expiré"})
			c.Abort()
			return
		}

		// Extraire l'ID utilisateur et l'ajouter au contexte
		userID, ok := (*claims)["user_id"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token invalide: user_id manquant"})
			c.Abort()
			return
		}

		// Stocker l'ID utilisateur dans le contexte pour les handlers suivants
		c.Set("userID", userID)

		// Continuer vers le handler suivant
		c.Next()
	}
}
