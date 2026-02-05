package models

import "time"

type Page struct {
	ID           string     `gorm:"primaryKey" json:"id"`
	LogID        string     `gorm:"column:log_id;not null" json:"logId"`
	UserID       string     `gorm:"column:user_id;not null" json:"userId"`
	FishName     string     `gorm:"column:fish_name" json:"fishName"`
	FishPhotoURL string     `gorm:"column:fish_photo_url" json:"fishPhotoUrl"`
	Comment      string     `json:"comment"`
	Length       float64    `json:"length"`
	Weight       float64    `json:"weight"`
	FishingSpot  string     `gorm:"column:fishing_spot" json:"fishingSpot"`
	FishingDate  *time.Time `gorm:"column:fishing_date;type:date" json:"fishingDate"`
	Release      bool       `json:"release"`
}
