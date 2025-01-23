package models

import (
	"errors"
	"gorm.io/gorm"
)

type Lactation struct {
	ID uint `gorm:"primaryKey"`

	CowId uint `gorm:"index;"` // ID коровы, данные о лактации которой записаны

	CheckMilks []CheckMilk `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DailyMilks []DailyMilk `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Number uint // номер лактации

	InsemenationNum  int      // Количество осеменений
	InsemenationDate DateOnly `gorm:"index"` // Дата осеменения

	CalvingCount int      `gorm:"index"` // Количество рожденных телят: 0 - мертворождение, 2 - двойня
	CalvingDate  DateOnly `gorm:"index"`

	ServicePeriod *uint    `gorm:"index"` // сервис период коровы: время от отела до осеменения
	Abort         bool     `gorm:"index"` // был ли аборт
	MilkAll       *float64 // Суммарный надой
	Milk305       *float64 // Суммарный надой за 305 дней
	FatAll        *float64 // Суммарный жир
	Fat305        *float64 // Суммарный жир за 305 дней
	ProteinAll    *float64 // Суммарный белок
	Protein305    *float64 // Суммарный белок за 305 дней
	Days          *int     // Количество дней, когда корова дает молоко
}

func (l *Lactation) Validate() error {
	if l.CalvingCount < 0 || l.CalvingCount > 2 {
		return errors.New("calving count must be between 0 and 2")
	}
	db := dbConnection
	cow := Cow{}
	if err := db.First(&cow, l.CowId).Error; err != nil {
		return errors.New("cow not found")
	}
	if cow.BirthDate.After(l.InsemenationDate.Time) {
		return errors.New("корова не может родиться после осеменения")
	}
	if l.InsemenationDate.Time.After(l.CalvingDate.Time) {
		return errors.New("отел не может произойти до осеменения")
	}
	return nil
}

func (l *Lactation) BeforeCreate(tx *gorm.DB) error {
	return l.Validate()
}
