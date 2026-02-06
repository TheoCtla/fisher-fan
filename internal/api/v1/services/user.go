package services

import (
	"fisherfan/internal/api/v1/models"
	"fisherfan/internal/api/v1/repository"

	"github.com/google/uuid"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

// GetAllUsers récupère tous les utilisateurs en appliquant les filtres optionnels
func (s *UserService) GetAllUsers(lastName, firstName, email, status string) ([]models.User, error) {
	// Ici, on pourrait ajouter une logique métier,
	// par exemple vérifier que le status est valide avant d'appeler le repo.
	return s.repo.FindAll(lastName, firstName, email, status)
}

// CreateUser gère la création d'un utilisateur
func (s *UserService) CreateUser(user *models.User) error {
	user.ID = uuid.New().String()
	return s.repo.Create(user)
}

// GetUser récupère un utilisateur par son ID
func (s *UserService) GetUser(id string) (models.User, error) {
	return s.repo.FindByID(id)
}

// UpdateUser gère la mise à jour des informations
func (s *UserService) UpdateUser(user *models.User) (*models.User, error) {

	// 1. On vérifie d'abord si l'utilisateur existe
	_, err := s.repo.FindByID(user.ID)
	if err != nil {
		return nil, err
	}

	// 2. On demande au repo de sauvegarder les modifications
	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser supprime un utilisateur
func (s *UserService) DeleteUser(id string) error {
	// On pourrait ajouter une vérification : est-ce que l'utilisateur a des réservations en cours ?
	// Si oui, on pourrait empêcher la suppression.
	return s.repo.Delete(id)
}
