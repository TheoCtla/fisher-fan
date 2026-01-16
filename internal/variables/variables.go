package variables

import "os"

// Configuration du serveur
var (
	Adress = getEnv("SERVER_ADDRESS", "0.0.0.0")
	Port   = getEnv("SERVER_PORT", "8080")
)

// Configuration de la base de données PostgreSQL
var (
	DBHost     = getEnv("DB_HOST", "localhost")
	DBPort     = getEnv("DB_PORT", "5432")
	DBUser     = getEnv("DB_USER", "fisherfan")
	DBPassword = getEnv("DB_PASSWORD", "fisherfan")
	DBName     = getEnv("DB_NAME", "fisherfan")
	DBSSLMode  = getEnv("DB_SSLMODE", "disable")
)

// getEnv récupère une variable d'environnement ou retourne une valeur par défaut
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
