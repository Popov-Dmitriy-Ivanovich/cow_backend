package models

type Grade struct {
	ID              uint `gorm:"primaryKey" example:"1"` // ID оценки
	CowID           uint
	GeneralValue    *float64
	EbvMilk         *float64
	EbvFat          *float64
	EbvProtein      *float64
	EbvInsemenation *float64
	EvbService      *float64
}
