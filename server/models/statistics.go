package models

import "github.com/lib/pq"

type GaussianStatistics struct {
	ID uint

	Farm   *Farm `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	FarmID *uint `gorm:"index"` // не null, если статистика собрана по холдингу или хозяйству

	Region   *Region `constraint:"OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RegionID *uint   `gorm:"index"` // не null, если статистика собрана по региону

	MinIndex  float64       // Значение минимального индекса
	MinCount  uint          // Количество коров с минимальным индексом
	MinCowIds pq.Int64Array `gorm:"type:bigint[]" swaggertype:"array,number"` // ID коров с минимальным значением индекса

	AvgIndex  float64       // Значение среднего индекса
	AvgCount  uint          // Количество коров с минимальным индексом
	AvgCowIds pq.Int64Array `gorm:"type:bigint[]" swaggertype:"array,number"` // ID коров с средним значением индекса

	MaxIndex  float64       // Значение максимального индекса
	MaxCount  uint          // Количество коров с максимальным индексом
	MaxCowIds pq.Int64Array `gorm:"type:bigint[]" swaggertype:"array,number"` // ID коров с максимальным значением индекса
}
