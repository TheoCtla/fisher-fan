package repository

import (
	"fisherman/internal/api/v1/models"

	"gorm.io/gorm"
)

type TripRepository struct {
	db *gorm.DB
}

func NewTripRepository(db *gorm.DB) *TripRepository {
	return &TripRepository{db: db}
}

func (r *TripRepository) CreateTrip(trip *models.Trip) error {
	return r.db.Create(trip).Error
}
