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

func (r *TripRepository) GetTrips(filters map[string]string) ([]models.Trip, error) {
	var trips []models.Trip
	query := r.db.Model(&models.Trip{})

	if v := filters["userId"]; v != "" {
		query = query.Where("user_id = ?", v)
	}
	if v := filters["title"]; v != "" {
		query = query.Where("title ILIKE ?", "%"+v+"%")
	}
	if v := filters["tripType"]; v != "" {
		query = query.Where("trip_type = ?", v)
	}
	if v := filters["startDate"]; v != "" {
		query = query.Where("start_date >= ?", v)
	}
	if v := filters["endDate"]; v != "" {
		query = query.Where("end_date <= ?", v)
	}

	err := query.Find(&trips).Error
	return trips, err
}

func (r *TripRepository) GetTripByID(id string) (*models.Trip, error) {
	var trip models.Trip
	err := r.db.First(&trip, "id = ?", id).Error
	return &trip, err
}

func (r *TripRepository) UpdateTrip(id string, trip *models.Trip) error {
	return r.db.Model(&models.Trip{}).Where("id = ?", id).Updates(trip).Error
}

func (r *TripRepository) DeleteTrip(id string) error {
	return r.db.Delete(&models.Trip{}, "id = ?", id).Error
}
