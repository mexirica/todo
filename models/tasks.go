package models

import "time"

type Task struct {
	ID          uint      `gorm:"primary_key;auto_increment" json:"id"`
	Description string    `gorm:"size:255" json:"description"`
	UserID      uint      `gorm:"not null" json:"user_id"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP;not null" json:"created_at"`
}
