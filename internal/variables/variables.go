package variables

import (
	"log"
	"os"
)

// Configuration du serveur
var (
	Adress = getEnv("SERVER_ADDRESS")
	Port   = getEnv("SERVER_PORT")
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
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("❌ Variable d'environnement requise non définie: %s", key)
	}
	return value
}
