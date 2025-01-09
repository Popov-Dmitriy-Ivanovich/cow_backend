package models

// import "time"

type News struct {
	ID   uint `gorm:"primaryKey"`
	Date DateOnly // Дата новости

	Region   *Region `json:"-"`
	RegionId *uint // ID региона для которого записана новость
	Title    string // Заголовок новости
	Text     string // Текст новости
}
