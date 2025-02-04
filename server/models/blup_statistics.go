package models

import "time"

type BlupStatistics struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time

	AverageEbvGeneralValueRegion float64 // средняя Оценка общая по EBV
	AverageEbvMilkRegion         float64 // средняя Оценка удоя по EBV
	AverageEbvFatRegion          float64 // средняя Оценка жира по EBV
	AverageEbvProteinRegion      float64 // средняя Оценка белка по EBV
	AverageEbvInsemenationRegion float64 // средняя Оценка кратности осеменения по EBV
	AverageEbvServiceRegion      float64 // средняя Оценка длительности сервисного периода по EBV

	MinEbvGeneralValueRegion float64 // минимальная Оценка общая по EBV
	MinEbvMilkRegion         float64 // минимальная Оценка удоя по EBV
	MinEbvFatRegion          float64 // минимальная Оценка жира по EBV
	MinEbvProteinRegion      float64 // минимальная Оценка белка по EBV
	MinEbvInsemenationRegion float64 // минимальная Оценка кратности осеменения по EBV
	MinEbvServiceRegion      float64 // минимальная Оценка длительности сервисного периода по EBV

	MaxEbvGeneralValueRegion float64 // максимальная Оценка общая по EBV
	MaxEbvMilkRegion         float64 // максимальная Оценка удоя по EBV
	MaxEbvFatRegion          float64 // максимальная Оценка жира по EBV
	MaxEbvProteinRegion      float64 // максимальная Оценка белка по EBV
	MaxEbvInsemenationRegion float64 // максимальная Оценка кратности осеменения по EBV
	MaxEbvServiceRegion      float64 // максимальная Оценка длительности сервисного периода по EBV
}
