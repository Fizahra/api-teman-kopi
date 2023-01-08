package model

type Staff struct {
	ID     int    `gorm:"primaryKey"`
	Nama   string `gorm:"not null" json:"nama" form:"nama" valid:"required~Your Nama is required"`
	Posisi string `gorm:"not null" json:"posisi" form:"posisi" valid:"required~Your Posisi is required"`
	Shift  string `gorm:"not null" json:"shift" form:"shift" valid:"required~Your Shift is required"`
}
