package models

type Region struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}
