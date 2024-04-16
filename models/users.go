package models

import "time"

type User struct {
	ID           uint      `gorm:"primary_key;auto_increment" json:"id"`
	Name         string    `gorm:"size:20;not null" json:"name"`
	Surname      string    `gorm:"size:20;not null" json:"surname"`
	Email        string    `gorm:"size:50;not null" json:"email"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP; not null" json:"created_at"`
	Birthday     time.Time `gorm:"not null" json:"birthday"`
	Company      string    `gorm:"size:50;not null" json:"company"`
	Tasks        []Task    `gorm:"foreignkey:UserID" json:"tasks"`
	RefreshToken string    `json:"refresh_token"`
	Password     string    `json:"password" gorm:"not null"`
}

type NewUser struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Birthday string `json:"birthday"`
	Company  string `json:"company"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
