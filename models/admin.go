package models

import "time"

type Admin struct {
	ID        string `gorm:"primary_key" json:"id"`
	Nama      string `gorm:"type:varchar(255);NOT NULL" json:"nama" binding:"required"`
	Email     string `gorm:"type:varchar(255)" json:"email"`
	Password  string `gorm:"type:varchar(255)" json:"password"`
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Admins []Admin
