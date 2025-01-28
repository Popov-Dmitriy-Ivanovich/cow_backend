package models

type Grade struct {
	ID                       uint     `gorm:"primaryKey" example:"1"` // ID оценки
	CowID                    uint     `gorm:"index;"`
	ReliabilityOfCalculation *float64 // Достоверность расчета
	GeneralValue             *float64 // Общая оценка по EBV
	EbvMilk                  *float64 // Оценка удоя по EBV
	EbvFat                   *float64 // Оценка жира по EBV
	EbvProtein               *float64 // Оценка белка по EBV
	EbvInsemenation          *float64 // Оценка кратности осеменения по EBV
	EvbService               *float64 // Оценка длительности сервисного периода по EBV
}

type GradeRegion struct {
	Grade
	CowID uint `example:"1"`
}

type GradeHoz struct {
	Grade
	CowID uint `example:"1"`
}

type GradeCountry struct {
	Grade
	CowID uint `example:"1"`
}
