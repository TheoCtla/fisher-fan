package services

import (
	"fisherman/internal/api/v1/models"
	"fisherman/internal/api/v1/repository"

	"github.com/google/uuid"
)

type TripService struct {
	repo *repository.TripRepository
}

func NewTripService(repo *repository.TripRepository) *TripService {
	return &TripService{repo: repo}
}

func (s *TripService) CreateTrip(trip *models.Trip) error {
	trip.ID = uuid.New().String()
	return s.repo.CreateTrip(trip)
}

func (s *TripService) GetTrips(filters map[string]string) ([]models.Trip, error) {
	return s.repo.GetTrips(filters)
}

func (s *TripService) GetTripByID(id string) (*models.Trip, error) {
	return s.repo.GetTripByID(id)
}

func (s *TripService) UpdateTrip(id string, trip *models.Trip) error {
	return s.repo.UpdateTrip(id, trip)
}

func (s *TripService) DeleteTrip(id string) error {
	return s.repo.DeleteTrip(id)
}
