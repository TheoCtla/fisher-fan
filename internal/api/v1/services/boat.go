package services

import (
	"errors"
	"fisherfan/internal/api/v1/models"
	"fisherfan/internal/api/v1/repository"

	"github.com/google/uuid"
)

type BoatService struct {
	repo     *repository.BoatRepository
	userRepo *repository.UserRepository
}

func NewBoatService(r *repository.BoatRepository, u *repository.UserRepository) *BoatService {
	return &BoatService{repo: r, userRepo: u}
}

func (s *BoatService) GetAllBoats(filters map[string]string) ([]models.Boat, error) {
	return s.repo.GetAll(filters)
}

func (s *BoatService) GetBoatByID(id string) (*models.Boat, error) {
	return s.repo.FindByID(id)
}

func (s *BoatService) CreateBoat(boat *models.Boat) error {
	if boat.Latitude < -90 || boat.Latitude > 90 {
		return errors.New("latitude invalide : doit être comprise entre -90 et +90")
	}
	if boat.Longitude < -180 || boat.Longitude > 180 {
		return errors.New("longitude invalide : doit être comprise entre -180 et +180")
	}

	user, err := s.userRepo.FindByID(boat.UserID)
	if err != nil {
		return errors.New("utilisateur introuvable : impossible de vérifier le permis")
	}

	if user.BoatLicense == "" {
		return errors.New("action interdite : vous devez renseigner un numéro de permis dans votre profil utilisateur avant d'ajouter un bateau")
	}
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
