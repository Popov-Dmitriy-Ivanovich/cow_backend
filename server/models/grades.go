package models

type Grade struct {
	ID    uint `gorm:"primaryKey" example:"1"` // ID оценки
	CowID uint `gorm:"index;"`

	GeneralValueReliability *float64 // Достоверность расчета общей оценки
	GeneralValue            *float64 // Общая оценка по EBV

	EbvMilkReliability *float64 // Достоверность расчета оценки удоя
	EbvMilk            *float64 // Оценка удоя по EBV

	EbvFatReliability *float64 // Достоверность расчета оценки жира
	EbvFat            *float64 // Оценка жира по EBV

	EbvProteinReliability *float64 // Достоверность расчета оценки белка
	EbvProtein            *float64 // Оценка белка по EBV

	EbvInsemenationReliability *float64 // Достоверность расчета оценки кратности осеменения
	EbvInsemenation            *float64 // Оценка кратности осеменения по EBV

	EbvServiceReliability *float64 // Достоверность расчета оценки сервисного периода
	EbvService            *float64 // Оценка длительности сервисного периода по EBV

	EbvMastit            *float64 // оценка мастита
	EbvMastitReliability *float64 // достоверность мастита

	EbvFatPercents            *float64 // оценка жира в %
	EbvFatPercentsReliability *float64 // достоверность оценки жира в процентах

	EbvProteinPercents            *float64 // оценка белка в процентах
	EbvProteinPercentsReliability *float64 // достоверность оценки белка в процентах

	EbvSomaticNucs            *float64 // оценка числа соматических клеток
	EbvSomaticNucsReliability *float64 // достоверность оценки числа соматических клеток

	EbvProductiveLongevity            *float64 // оценка продуктивного долголетия
	EbvProductiveLongevityReliability *float64 // достоверность оценки продуктивного долголетия
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
