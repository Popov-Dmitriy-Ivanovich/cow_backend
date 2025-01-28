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
}
