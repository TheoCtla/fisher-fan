package repository

import (
	"fisherman/internal/api/v1/models"

	"gorm.io/gorm"
)

type ReservationRepository struct {
	db *gorm.DB
}

func NewReservationRepository(db *gorm.DB) *ReservationRepository {
	return &ReservationRepository{db: db}
}

func (r *ReservationRepository) Create(reservation *models.Reservation) error {
	return r.db.Create(reservation).Error
}

func (r *ReservationRepository) GetAll(filters map[string]string) ([]models.Reservation, error) {
	var reservations []models.Reservation
	query := r.db.Model(&models.Reservation{})

	// Filtres exacts
	if v := filters["userId"]; v != "" {
		query = query.Where("user_id = ?", v)
	}
	if v := filters["tripId"]; v != "" {
		query = query.Where("trip_id = ?", v)
	}
	if v := filters["date"]; v != "" {
		query = query.Where("date = ?", v)
	}

	err := query.Find(&reservations).Error
	return reservations, err
}

func (r *ReservationRepository) FindByID(id string) (*models.Reservation, error) {
	var reservation models.Reservation
	err := r.db.First(&reservation, "id = ?", id).Error
	return &reservation, err
}

func (r *ReservationRepository) Update(id string, reservation *models.Reservation) error {
	return r.db.Model(&models.Reservation{}).Where("id = ?", id).Updates(reservation).Error
}

func (r *ReservationRepository) Delete(id string) error {
	return r.db.Delete(&models.Reservation{}, "id = ?", id).Error
}
