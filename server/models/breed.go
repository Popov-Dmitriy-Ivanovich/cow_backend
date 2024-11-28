package models

type Breed struct {
	ID   uint   `gorm:"primaryKey" example:"1"` // ID породы
	Name string `example:"Какая-нибудь порода"` // Название породы
}
