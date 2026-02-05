package models

import "time"

type Trip struct {
	ID             string  `gorm:"primaryKey" json:"id"`
	UserID         string  `gorm:"column:user_id;not null" json:"userId"`
	BoatID         string  `gorm:"column:boat_id;not null" json:"boatId"`
	Title          string  `gorm:"not null" json:"title"`
	PracticalInfo  string  `gorm:"column:practical_info" json:"practicalInfo"`
	TripType       string  `gorm:"column:trip_type" json:"tripType"`
	RateType       string  `gorm:"column:rate_type" json:"rateType"`
	PassengerCount int     `gorm:"column:passenger_count" json:"passengerCount"`
	Price          float64 `json:"price"`

	Schedules []TripSchedule `gorm:"foreignKey:TripID" json:"schedules"`
}

type TripSchedule struct {
	TripID        string     `gorm:"column:trip_id;primaryKey" json:"-"`
	StartDate     time.Time  `gorm:"column:start_date;primaryKey" json:"startDate"`
	EndDate       *time.Time `gorm:"column:end_date" json:"endDate"`
	DepartureTime string     `gorm:"column:departure_time;primaryKey" json:"departureTime"` // Format HH:MM:SS
	EndTime       string     `gorm:"column:end_time" json:"endTime"`
}
