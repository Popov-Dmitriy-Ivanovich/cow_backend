package models

import (
	"errors"
	"gorm.io/gorm"
	"regexp"
)

type UserRegisterRequest struct {
	ID        uint
	HozNumber string // Номер хоз-ва к которому привязвыается пользователь: либо существует, либо newHoz

	NameSurnamePatronimic string
	RoleId                uint
	Email                 string `gorm:"uniqueIndex"`
	Phone                 string
	Password              string

	RegionId uint
}

func (urr *UserRegisterRequest) Validate() error {
	matched, err := regexp.MatchString(emailRegexp, urr.Email)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("email format is invalid")
	}
	matchedPhone, err := regexp.MatchString(phoneRegexp, urr.Phone)
	if err != nil {
		return err
	}
	if !matchedPhone {
		return errors.New("phone format is invalid")
	}
	return nil
}

func (urr *UserRegisterRequest) BeforeCreate(tx *gorm.DB) error {
	return urr.Validate()
}

func (urr *UserRegisterRequest) BeforeUpdate(tx *gorm.DB) error {
	return urr.Validate()
}

type HozRegisterRequest struct {
	ID uint

	HoldNumber string

	HozNumber  string
	DistrictId uint

	Name        string
	ShortName   string
	Inn         *string
	Address     *string
	Phone       *string
	Email       *string
	Description *string
	CowsCount   *uint
}
type HoldRegisterRequest struct {
	ID          uint
	HozNumber   string
	DistrictId  string
	Name        string
	ShortName   string
	Inn         *string
	Address     *string
	Phone       *string
	Email       *string
	Description *string
	CowsCount   *uint
}
