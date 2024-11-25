package models

type Breed struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}
