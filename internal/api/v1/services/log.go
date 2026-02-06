package services

import (
	"fisherfan/internal/api/v1/models"
	"fisherfan/internal/api/v1/repository"

	"github.com/google/uuid"
)

type LogService struct {
	repo *repository.LogRepository
}

func NewLogService(r *repository.LogRepository) *LogService {
	return &LogService{repo: r}
}

func (s *LogService) GetByUserID(userID string) (*models.Log, error) {
	return s.repo.GetByUserID(userID)
}

func (s *LogService) CreateLog(log *models.Log) error {
	log.ID = uuid.New().String()
	return s.repo.Create(log)
}

func (s *LogService) GetPage(pageID string) (*models.Page, error) {
	return s.repo.GetPage(pageID)
}

func (s *LogService) UpdatePage(pageID string, page *models.Page) error {
	return s.repo.UpdatePage(pageID, page)
}

func (s *LogService) DeletePage(pageID string) error {
	return s.repo.DeletePage(pageID)
}
