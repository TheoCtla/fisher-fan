package models

type Log struct {
	ID     string `gorm:"primaryKey" json:"id"`
	UserID string `gorm:"column:user_id;unique;not null" json:"userId"`
	Pages  []Page `gorm:"foreignKey:LogID" json:"pages"`
}
