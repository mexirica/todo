package models

import "time"

type Task struct {
	ID          uint      `gorm:"primary_key;auto_increment" json:"-"`
	Description string    `gorm:"size:255" json:"description"`
	UserID      float64   `gorm:"not null" json:"-"`
	Completed   bool      `gorm:"default:false" json:"completed"`
	CompletedAt time.Time `gorm:"default:null" json:"completed_at"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP;not null" json:"created_at"`
}
