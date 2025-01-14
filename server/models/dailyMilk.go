package models

import (
	"gorm.io/gorm"
)

type DailyMilk struct {
	ID             uint     `gorm:"primaryKey"`
	LactationId    uint     `gorm:"index"` // ID лактации во время котороый была дойка `example:"1"`
	Date           DateOnly `gorm:"index"` // Дата дойки
	Milk           int      // Суммарный надой `example:"12"`
	MilkMorning    *int     // Надой утром `example:"12"`
	MilkNoon       *int     // Надой днем `example:"12"`
	MilkEvening    *int     // Надой вечером `example:"12"`
	Fat            int      // Суммарный жир `example:"12"`
	FatMorning     *int     // Жир утром`example:"12"`
	FatNoon        *int     // Жир днем `example:"12"`
	FatEvening     *int     // Жир вечером `example:"12"`
	Protein        int      // Суммарный белок `example:"12"`
	ProteinMorning *int     // Белок утром `example:"12"`
	ProteinNoon    *int     // Белок днем `example:"12"`
	ProteinEvening *int     // Белок вечером `example:"12"`
}

func (dm *DailyMilk) Validate() error {
	return nil
}

func (dm *DailyMilk) BeforeCreate(tx *gorm.DB) error {
	return dm.Validate()
}

func (dm *DailyMilk) BeforeUpdate(tx *gorm.DB) error {
	return dm.Validate()
}
