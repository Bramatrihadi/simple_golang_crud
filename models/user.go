package models

import "time"

type User struct {
	ID             string `gorm:"primary_key" json:"id"`
	Nik            int
	Foto_ktp       string `gorm:"type:varchar(255)" json:"foto_ktp"`
	No_kk          int
	Nama_lengkap   string `gorm:"type:varchar(255);NOT NULL" json:"nama_lengkap" binding:"required"`
	Alamat_lengkap string `gorm:"type:text" json:"alamat_lengkap"`
	Email          string `gorm:"type:varchar(255)" json:"email"`
	Pekerjaan      string `gorm:"type:varchar(255)" json:"pekerjaan"`
	No_hp          string `gorm:"type:varchar(100);NOT NULL;UNIQUE;UNIQUE_INDEX" json:"no_hp" binding:"required"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Users []User
