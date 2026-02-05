package repository

import (
	"fisherman/internal/api/v1/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository injecte l'instance GORM
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create insère un nouvel utilisateur en base
func (repo *UserRepository) Create(user *models.User) error {
	user.ID = uuid.New().String()
	return repo.db.Create(user).Error
}

// FindByID récupère un utilisateur par sa clé primaire
func (repo *UserRepository) FindByID(id string) (models.User, error) {
	var user models.User

	err := repo.db.First(&user, "id = ?", id).Error
	return user, err
}

// FindAll gère la recherche avec les filtres du Swagger
func (r *UserRepository) FindAll(lastName, firstName, email, status string) ([]models.User, error) {
	var users []models.User
	query := r.db

	// Filtres dynamiques
	if lastName != "" {
		// "ILIKE" pour PostgreSQL permet une recherche insensible à la casse
		query = query.Where("last_name ILIKE ?", "%"+lastName+"%")
	}
	if firstName != "" {
		query = query.Where("first_name ILIKE ?", "%"+firstName+"%")
	}
	if email != "" {
		query = query.Where("email = ?", email)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	err := query.Find(&users).Error
	return users, err
}

// Update met à jour l'ensemble des champs de l'utilisateur
func (r *UserRepository) Update(user *models.User) error {
	// Save met à jour tous les champs, y compris ceux à zéro
	// Si tu veux mettre à jour uniquement certains champs, utilise .Updates()
	return r.db.Save(user).Error
}

// Delete effectue un "Soft Delete" (si DeletedAt est présent dans le modèle)
func (r *UserRepository) Delete(id string) error {
	return r.db.Delete(&models.User{}, "id = ?", id).Error
}
