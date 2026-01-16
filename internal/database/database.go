package database

import (
	"fmt"
	"log"

	"fisherman/internal/variables"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB est l'instance globale de la base de données
var DB *gorm.DB

// Connect initialise la connexion à la base de données PostgreSQL
func Connect() error {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		variables.DBHost,
		variables.DBPort,
		variables.DBUser,
		variables.DBPassword,
		variables.DBName,
		variables.DBSSLMode,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return fmt.Errorf("erreur de connexion à la base de données: %w", err)
	}

	log.Println("✅ Connexion à la base de données PostgreSQL établie")
	return nil
}

// GetDB retourne l'instance de la base de données
func GetDB() *gorm.DB {
	return DB
}

// Close ferme la connexion à la base de données
func Close() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
