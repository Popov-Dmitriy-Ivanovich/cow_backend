package models

type Breed struct {
	ID   uint   `gorm:"primaryKey" example:"1"`                    // ID породы
	Name string `gorm:"uniqueIndex" example:"Какая-нибудь порода"` // Название породы
}
