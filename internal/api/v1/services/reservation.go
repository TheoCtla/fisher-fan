package services

import (
	"fisherman/internal/api/v1/models"
	"fisherman/internal/api/v1/repository"

	"github.com/google/uuid"
)

type ReservationService struct {
	repo *repository.ReservationRepository
}

func NewReservationService(r *repository.ReservationRepository) *ReservationService {
	return &ReservationService{repo: r}
}

func (s *ReservationService) GetAllReservations(filters map[string]string) ([]models.Reservation, error) {
	return s.repo.GetAll(filters)
}

func (s *ReservationService) GetReservationByID(id string) (*models.Reservation, error) {
	return s.repo.FindByID(id)
}

func (s *ReservationService) CreateReservation(reservation *models.Reservation) error {
	reservation.ID = uuid.New().String()
	return s.repo.Create(reservation)
}

func (s *ReservationService) UpdateReservation(id string, reservation *models.Reservation) error {
	return s.repo.Update(id, reservation)
}

func (s *ReservationService) DeleteReservation(id string) error {
	return s.repo.Delete(id)
}
