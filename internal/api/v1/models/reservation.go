package models

import "time"

type Reservation struct {
	ID            string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	TripID        string    `gorm:"column:trip_id;not null" json:"tripId"`
	UserID        string    `gorm:"column:user_id;not null" json:"userId"`
	Date          time.Time `gorm:"not null" json:"date"`
	ReservedSeats int       `gorm:"column:reserved_seats" json:"reservedSeats"`
	TotalPrice    float64   `gorm:"column:total_price" json:"totalPrice"`
}
