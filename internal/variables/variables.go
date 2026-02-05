package variables

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Configuration du serveur
var (
	Address = getEnv("SERVER_ADDRESS")
	Port    = getEnv("SERVER_PORT")
)

// Configuration JWT
var (
	JWTSecret          = getEnv("JWT_SECRET")
	AccessTokenExpiry  = getEnvWithDefault("ACCESS_TOKEN_EXPIRY", "24h")
	RefreshTokenExpiry = getEnvWithDefault("REFRESH_TOKEN_EXPIRY", "168h")
)

// Configuration de la base de données PostgreSQL
var (
	DBHost     = getEnv("DB_HOST")
	DBPort     = getEnv("DB_PORT")
	DBUser     = getEnv("DB_USER")
	DBPassword = getEnv("DB_PASSWORD")
	DBName     = getEnv("DB_NAME")
	DBSSLMode  = getEnv("DB_SSLMODE")
)

// getEnv récupère une variable d'environnement ou crash si absente
func getEnv(key string) string {
	godotenv.Load()
	value, exists := os.LookupEnv(key)
	log.Println("Loading env variable:", key, "Value:", value)
	if !exists {
		log.Fatalf("❌ Variable d'environnement requise non définie: %s", key)
	}
	return value
}

// getEnvWithDefault récupère une variable d'environnement ou retourne une valeur par défaut
func getEnvWithDefault(key, defaultValue string) string {
	godotenv.Load()
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		log.Printf("⚠️  Variable d'environnement '%s' non définie, utilisation de la valeur par défaut: %s", key, defaultValue)
		return defaultValue
	}
	log.Println("Loading env variable:", key, "Value:", value)
	return value
}
