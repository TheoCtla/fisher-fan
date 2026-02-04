package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model             // Contient déjà ID (uint), CreatedAt, UpdatedAt, DeletedAt
	LastName     string    `json:"lastName"`
	FirstName    string    `json:"firstName"`
	BirthDate    time.Time `json:"birthDate"`
	Email        string    `json:"email" gorm:"unique;not null"`
	BoatLicense  string    `json:"boatLicense"`
	Status       string    `json:"status"` // individual, professional
	CompanyName  string    `json:"companyName,omitempty"`
	ActivityType string    `json:"activityType,omitempty"`
	SiretNumber  string    `json:"siretNumber,omitempty"`
	RCNumber     string    `json:"rcNumber,omitempty"`
}
