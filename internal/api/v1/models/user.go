package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string         `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	LastName    string `gorm:"not null" json:"lastName"`
	FirstName   string `gorm:"not null" json:"firstName"`
	BirthDate   string `gorm:"type:date" json:"birthDate"`
	Email       string `gorm:"uniqueIndex;not null" json:"email"`
	BoatLicense string `gorm:"type:varchar(8)" json:"boatLicense"`
	Status      string `gorm:"type:varchar(20);check:status IN ('individual', 'professional')" json:"status"`

	CompanyName  string `json:"companyName"`
	ActivityType string `gorm:"type:varchar(20);check:activity_type IN ('rental', 'fishing guide', '')" json:"activityType"`
	SiretNumber  string `gorm:"type:varchar(14)" json:"siretNumber"`
	RcNumber     string `json:"rcNumber"`

	Boats        []Boat        `gorm:"foreignKey:UserID" json:"boats"`
	Trips        []Trip        `gorm:"foreignKey:UserID" json:"trips"`
	Reservations []Reservation `gorm:"foreignKey:UserID" json:"reservations"`
	Log          *Log          `gorm:"foreignKey:UserID" json:"log"`
}
