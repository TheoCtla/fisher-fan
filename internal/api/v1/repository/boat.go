package repository

import (
	"fisherman/internal/api/v1/models"

	"gorm.io/gorm"
)

type BoatRepository struct {
	db *gorm.DB
}

func NewBoatRepository(db *gorm.DB) *BoatRepository {
	return &BoatRepository{db: db}
}

func (r *BoatRepository) Create(boat *models.Boat) error {
	return r.db.Create(boat).Error
}

func (r *BoatRepository) GetAll(filters map[string]string) ([]models.Boat, error) {
	var boats []models.Boat
	query := r.db.Model(&models.Boat{})

	// Filtres exacts
	if v := filters["userId"]; v != "" {
		query = query.Where("user_id = ?", v)
	}
	if v := filters["brand"]; v != "" {
		query = query.Where("brand = ?", v)
	}
	if v := filters["boatType"]; v != "" {
		query = query.Where("boat_type = ?", v)
	}
	if v := filters["homePort"]; v != "" {
		query = query.Where("home_port = ?", v)
	}

	// Recherche textuelle
	if v := filters["name"]; v != "" {
		query = query.Where("name ILIKE ?", "%"+v+"%")
	}

	// Bounding Box (y=lat, x=lon)
	if filters["y1"] != "" && filters["y2"] != "" {
		query = query.Where("latitude BETWEEN ? AND ?", filters["y1"], filters["y2"])
	}
	if filters["x1"] != "" && filters["x2"] != "" {
		query = query.Where("longitude BETWEEN ? AND ?", filters["x1"], filters["x2"])
	}

	err := query.Find(&boats).Error
	return boats, err
}

func (r *BoatRepository) FindByID(id string) (*models.Boat, error) {
	var boat models.Boat
	err := r.db.First(&boat, "id = ?", id).Error
	return &boat, err
}

func (r *BoatRepository) Update(id string, boat *models.Boat) error {
	return r.db.Model(&models.Boat{}).Where("id = ?", id).Updates(boat).Error
}

func (r *BoatRepository) Delete(id string) error {
	return r.db.Delete(&models.Boat{}, "id = ?", id).Error
}

func (r *BoatRepository) FindByBBox(latMin, latMax, lonMin, lonMax float64) ([]models.Boat, error) {
	var boats []models.Boat
	err := r.db.Where("latitude BETWEEN ? AND ? AND longitude BETWEEN ? AND ?",
		latMin, latMax, lonMin, lonMax).Find(&boats).Error
	return boats, err
}
