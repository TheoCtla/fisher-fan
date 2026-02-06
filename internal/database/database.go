package database

import (
	"fisherfan/internal/variables"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		variables.GlobalConfig.DB.Host, variables.GlobalConfig.DB.Port, variables.GlobalConfig.DB.User,
		variables.GlobalConfig.DB.Password, variables.GlobalConfig.DB.Name, variables.GlobalConfig.DB.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	log.Println("✅ Connexion PostgreSQL établie")
	return db, nil
}
