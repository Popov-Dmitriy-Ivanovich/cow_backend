package models

type Sex struct {
	ID   uint   `gorm:"primaryKey"`
	Name string // Название пола
}
