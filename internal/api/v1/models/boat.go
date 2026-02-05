package models

import (
	"time"

	"gorm.io/gorm"
)

type Boat struct {
	ID        string         `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	UserID          string     `gorm:"column:user_id;not null" json:"userId"`
	Name            string     `gorm:"not null" json:"name"`
	Description     string     `json:"description"`
	Brand           string     `json:"brand"`
	ManufactureYear *time.Time `gorm:"column:manufacture_year;type:date" json:"manufactureYear"`
	PhotoURL        string     `gorm:"column:photo_url" json:"photoUrl"`
	LicenseType     string     `gorm:"column:license_type" json:"licenseType"`
	BoatType        string     `gorm:"column:boat_type" json:"boatType"`
	DepositAmount   float64    `gorm:"column:deposit_amount" json:"depositAmount"`
	MaxCapacity     int        `gorm:"column:max_capacity" json:"maxCapacity"`
	NumberOfBeds    int        `gorm:"column:number_of_beds" json:"numberOfBeds"`
	HomePort        string     `gorm:"column:home_port" json:"homePort"`
	Latitude        float64    `json:"latitude"`
	Longitude       float64    `json:"longitude"`
	EngineType      string     `gorm:"column:engine_type" json:"engineType"`
	EnginePower     int        `gorm:"column:engine_power" json:"enginePower"`

	Equipments []BoatEquipment `gorm:"foreignKey:BoatID" json:"equipment"`
}

type BoatEquipment struct {
	BoatID string `gorm:"column:boat_id;primaryKey" json:"-"`
	Name   string `gorm:"primaryKey" json:"name"`
}
