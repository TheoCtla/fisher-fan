package main

import (
	"fisherman/internal/api/v1/models"
	"fisherman/internal/database"
	"fisherman/internal/server"
	"log"
)

func main() {
	log.Println("🐟 Démarrage de FisherFan API...")

	// 1. Connexion
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("❌ Impossible de se connecter: %v", err)
	}

	// 2. Auto-Migration (Optionnel mais recommandé ici)
	db.AutoMigrate(&models.User{})

	// 3. Lancement du serveur avec la DB injectée
	server.InitServer(db)
}
