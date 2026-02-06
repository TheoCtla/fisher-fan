package variables

import (
	"log"
	"os"

	// On enlève le "autoload" pour charger manuellement et voir les erreurs
	"github.com/joho/godotenv"
)

var GlobalConfig *Config

type Config struct {
	ServerAddress string
	ServerPort    string
	DB            PostgresConfig
	JWT           JWTConfig
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type JWTConfig struct {
	Secret             string
	AccessTokenExpiry  string
	RefreshTokenExpiry string
}

func init() {
	// 1. On affiche le dossier où Go s'exécute pour vérifier qu'on est au bon endroit
	cwd, _ := os.Getwd()
	log.Println("📂 Dossier d'exécution (CWD) :", cwd)

	// 2. On essaie de charger le .env explicitement
	err := godotenv.Load()
	if err != nil {
		// On affiche l'erreur exacte (ex: fichier introuvable, problème de format...)
		log.Printf("⚠️  ATTENTION: Impossible de charger le fichier .env : %v", err)
		log.Println("ℹ️  Lecture des variables système uniquement...")
	} else {
		log.Println("✅ Fichier .env trouvé et chargé !")
	}

	log.Println("⚙️  Chargement de la configuration globale...")
	GlobalConfig = loadConfig()
}

func loadConfig() *Config {
	mustGetEnv := func(key string) string {
		val := os.Getenv(key)
		if val == "" {
			// On arrête tout ici si c'est vide
			log.Fatalf("❌ Erreur critique : La variable '%s' est vide ou manquante.", key)
		}
		return val
	}

	return &Config{
		ServerAddress: mustGetEnv("SERVER_ADDRESS"),
		ServerPort:    mustGetEnv("SERVER_PORT"),
		DB: PostgresConfig{
			Host:     mustGetEnv("DB_HOST"),
			Port:     mustGetEnv("DB_PORT"),
			User:     mustGetEnv("DB_USER"),
			Password: mustGetEnv("DB_PASSWORD"),
			Name:     mustGetEnv("DB_NAME"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
		},
		JWT: JWTConfig{
			Secret:             mustGetEnv("JWT_SECRET"),
			AccessTokenExpiry:  os.Getenv("ACCESS_TOKEN_EXPIRY"),
			RefreshTokenExpiry: os.Getenv("REFRESH_TOKEN_EXPIRY"),
		},
	}
}
