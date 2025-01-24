package models

// import "time"

type News struct {
	ID   uint     `gorm:"primaryKey"`
	Date DateOnly // Дата новости

	Region   *Region `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RegionId *uint   `gorm:"index"` // ID региона для которого записана новость
	Title    string  // Заголовок новости
	Text     string  // Текст новости
}
