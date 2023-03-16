package models

import "time"

type Pendaftaran struct {
	ID                 string `gorm:"primary_key" json:"id"`
	Nik                int
	Rincian_info       string `gorm:"type:varchar(255)" json:"rincian_info"`
	Tujuan_penggunaan  string `gorm:"type:varchar(255)" json:"tujuan_penggunaan"`
	No_pendaftaran     string `gorm:"type:varchar(255)" json:"no_pendaftaran"`
	Tgl_pendaftaran    time.Time
	Status             string `gorm:"type:varchar(255)" json:"status"`
	Kategori           string `gorm:"type:varchar(255)" json:"Kategori"`
	Salinan_informasi  string `gorm:"type:varchar(255)" json:"salinan_informasi"`
	Pengajuan_komplain string `gorm:"type:varchar(255)" json:"pengajuan_komplain"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type Pendaftarans []Pendaftaran
