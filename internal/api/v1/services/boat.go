package services

import (
	"fisherfan/internal/api/v1/models"
	"fisherfan/internal/api/v1/repository"

	"github.com/google/uuid"
)

type BoatService struct {
	repo *repository.BoatRepository
}

func NewBoatService(r *repository.BoatRepository) *BoatService {
	return &BoatService{repo: r}
}

func (s *BoatService) GetAllBoats(filters map[string]string) ([]models.Boat, error) {
	return s.repo.GetAll(filters)
}

func (s *BoatService) GetBoatByID(id string) (*models.Boat, error) {
	return s.repo.FindByID(id)
}

func (s *BoatService) CreateBoat(boat *models.Boat) error {
	boat.ID = uuid.New().String()
	return s.repo.Create(boat)
}

func (s *BoatService) UpdateBoat(id string, boat *models.Boat) error {
	return s.repo.Update(id, boat)
}

func (s *BoatService) DeleteBoat(id string) error {
	return s.repo.Delete(id)
}

func (s *BoatService) GetByBBox(latMin, latMax, lonMin, lonMax float64) ([]models.Boat, error) {
	return s.repo.FindByBBox(latMin, latMax, lonMin, lonMax)
}
