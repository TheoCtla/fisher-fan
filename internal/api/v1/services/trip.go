package services

import (
	"fisherman/internal/api/v1/models"
	"fisherman/internal/api/v1/repository"
)

type TripService struct {
	repo *repository.TripRepository
}

func NewTripService(repo *repository.TripRepository) *TripService {
	return &TripService{repo: repo}
}

func (s *TripService) CreateTrip(trip *models.Trip) error {
	return s.repo.CreateTrip(trip)
}
