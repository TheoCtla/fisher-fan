package main

import (
	"fisherfan/internal/database"
	"fisherfan/internal/server"
	"log"
)

func main() {
	log.Println("🐟 Démarrage de FisherFan API...")
	log.Println("Init db...")
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("❌ Impossible de se connecter: %v", err)
	}

	server.InitServer(db)
}
