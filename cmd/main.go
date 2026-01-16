package main

import (
	"fisherman/internal/database"
	"fisherman/internal/server"
	"log"
)

func main() {
	log.Println("🐟 Démarrage de FisherFan API...")

	// Connexion à la base de données
	if err := database.Connect(); err != nil {
		log.Fatalf("❌ Impossible de se connecter à la base de données: %v", err)
	}
	defer database.Close()

	// Démarrage du serveur
	server.InitServer()
}
