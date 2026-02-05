package repository

import (
	"fisherman/internal/api/v1/models"

	"gorm.io/gorm"
)

type LogRepository struct {
	db *gorm.DB
}

func NewLogRepository(db *gorm.DB) *LogRepository {
	return &LogRepository{db: db}
}

func (r *LogRepository) GetByUserID(userID string) (*models.Log, error) {
	var log models.Log
	err := r.db.Preload("Pages").Where("user_id = ?", userID).First(&log).Error
	return &log, err
}

func (r *LogRepository) Create(log *models.Log) error {
	return r.db.Create(log).Error
}

func (r *LogRepository) GetPage(pageID string) (*models.Page, error) {
	var page models.Page
	err := r.db.First(&page, "id = ?", pageID).Error
	return &page, err
}

func (r *LogRepository) UpdatePage(pageID string, page *models.Page) error {
	return r.db.Save(page).Error
}

func (r *LogRepository) DeletePage(pageID string) error {
	return r.db.Delete(&models.Page{}, "id = ?", pageID).Error
}
