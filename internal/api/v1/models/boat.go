package models

type Boat struct {
	ID     uint    `json:"id" gorm:"primaryKey"`
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Length float64 `json:"length"`
}
