package models

import (
	"errors"
	"gorm.io/gorm"
	"regexp"
)

type User struct {
	ID                    uint   `gorm:"primaryKey"`
	NameSurnamePatronimic string // ФИО
	Role                  Role
	RoleId                int    // ID роли
	Email                 string `gorm:"uniqueIndex"` // Почта
	Phone                 string // телефон
	Password              []byte `json:"-"`
	Farm                  *Farm  `json:"-"`
	FarmId                *uint  // ID хозяйства
	Region                Region `json:"-"`
	RegionId              uint   `example:"1"` // ID региона
}

func (u *User) Validate() error {
	matchedMail, err := regexp.MatchString(emailRegexp, u.Email)
	if err != nil {
		return err
	}
	if !matchedMail {
		return errors.New("invalid email address")
	}
	matchedPhone, err := regexp.MatchString(phoneRegexp, u.Phone)
	if err != nil {
		return err
	}
	if !matchedPhone {
		return errors.New("invalid phone number")
	}
	return nil
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	return u.Validate()
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	return u.Validate()
}
