package models

import (
	"errors"
	"regexp"
)

type Partner struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  // Название партнера
	Address     *string // Адрес
	Phone       *string // телефон
	Email       *string // Эл. почта
	Description string  // Описание партнера
	LogoPath    *string // Путь к логотипу партнера относительно genmlik.ru/api/static/partners
}

func (p *Partner) Validate() error {
	if p.Email != nil {
		matched, err := regexp.MatchString(emailRegexp, *p.Email)
		if err != nil {
			return err
		}
		if !matched {
			return errors.New("invalid email address")
		}
	}
	if p.Phone != nil {
		matched, err := regexp.MatchString(phoneRegexp, *p.Phone)
		if err != nil {
			return err
		}
		if !matched {
			return errors.New("invalid phone number")
		}
	}
	return nil
}
