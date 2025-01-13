package models

type CheckMilk struct {
	ID uint `gorm:"primaryKey" example:"1"` // ID контрольной дойки

	LactationId uint `gorm:"index" example:"1"` // ID лактации для которой выполнена контрольная дойка

	CheckDate       DateOnly `gorm:"index"` // Дата конрольной дойки
	Milk            float64  `example:"1"`  // Параметр контрольной дойки, как я понимаю кол-во молока
	Fat             float64  `example:"1"`  // Параметр контрольной дойки, как я понимаю кол-во жира в молоке
	Protein         float64  `example:"1"`  // Параметр контрольной дойки, как я понимаю кол-во белка в молоке
	SomaticNucCount *float64 // количество соматических клеток
	ProbeNumber     *uint    // номер пробы
	DryMatter       *float64 // сухой материал
}
