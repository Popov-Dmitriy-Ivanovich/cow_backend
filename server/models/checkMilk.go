package models

import (
	"gorm.io/gorm"
)

type CheckMilk struct {
	ID uint `gorm:"primaryKey" example:"1"` // ID контрольной дойки

	LactationId uint `gorm:"index" example:"1"` // ID лактации для которой выполнена контрольная дойка

	CheckDate       DateOnly `gorm:"index"` // Дата контрольной дойки
	Milk            float64  `example:"1"`  // Параметр контрольной дойки, как я понимаю кол-во молока
	Fat             float64  `example:"1"`  // Параметр контрольной дойки, как я понимаю кол-во жира в молоке
	Protein         float64  `example:"1"`  // Параметр контрольной дойки, как я понимаю кол-во белка в молоке
	SomaticNucCount *float64 // Количество соматических клеток
	ProbeNumber     *uint    // Номер пробы
	DryMatter       *float64 // Сухой материал
}

func (cm *CheckMilk) Validate() error {

	return nil
}

func (cm *CheckMilk) BeforeCreate(tx *gorm.DB) error {
	return cm.Validate()
}

func (cm *CheckMilk) BeforeUpdate(tx *gorm.DB) error {
	return cm.Validate()
}
