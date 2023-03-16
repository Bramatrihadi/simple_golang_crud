package models

import "time"

type Dokumen struct {
	ID        string `gorm:"primary_key" json:"id"`
	Judul     string `gorm:"type:varchar(255)" json:"judul"`
	File      string `gorm:"type:varchar(255)" json:"file"`
	Jenis     string `gorm:"type:varchar(255)" json:"jenis"`
	Status    string `gorm:"type:varchar(255)" json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Dokumens []Dokumen
