package services

import (
	"fisherman/internal/api/v1/models"
	"fisherman/internal/api/v1/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.repo.Create(user)
}

func (s *UserService) GetUser(id uint) (models.User, error) {
	return s.repo.FindByID(id)
}
