package models

import (
	"errors"
	"gorm.io/gorm"
	"regexp"
)

type Farm struct {
	ID uint `gorm:"primaryKey"`

	// Region   Region `json:"-"`
	// RegionId uint
	HozNumber  *string  `gorm:"index"` // Номер хоз-ва
	District   District `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	DistrictId uint     // ID района, в котором находится хозяйство
	Parrent    *Farm    `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	ParrentId  *uint    // ID более управляющего хоз-ва (для хозяйства - холдинг, для фермы - хозяйство)

	Type      uint    // Тип: хозяйство, ферма, холдинг
	Name      string  // Название хозяйства
	NameShort string  // Краткое название хозяйства
	Inn       *string // ИНН

	Address     string  // Адрес
	Phone       *string // телефон
	Email       *string // Эл. почта
	Description *string // описание
	CowsCount   *uint   // Количество коров
}

func (f *Farm) Validate() error {
	if f.Email != nil {
		matched, err := regexp.MatchString(emailRegexp, *f.Email)
		if err != nil {
			return err
		}
		if !matched {
			return errors.New("email address is invalid " + *f.Email)
		}
	}
	if f.Phone != nil {
		matched, err := regexp.MatchString(phoneRegexp, *f.Phone)
		if err != nil {
			return err
		}
		if !matched {
			return errors.New("phone number is invalid " + *f.Phone)
		}
	}
	return nil
}

func (f *Farm) BeforeCreate(tx *gorm.DB) (err error) {
	return f.Validate()
}

func (f *Farm) BeforeUpdate(tx *gorm.DB) (err error) {
	return f.Validate()
}
