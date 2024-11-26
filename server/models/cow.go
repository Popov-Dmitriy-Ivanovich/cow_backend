package models

import "time"

type Cow struct {
	ID uint `gorm:"primaryKey"`

	Farm   Farm
	FarmID uint

	FarmGroup   Farm
	FarmGroupId uint

	Breed   Breed
	BreedId uint

	Sex   Sex
	SexId uint

	Father   *Cow
	FatherId *uint

	Mother   *Cow
	MotherId *uint

	Lactation []Lactation

	InventoryNumber      string
	SelecsNumber         string
	RSHNNumber           string
	Name                 string
	IdentificationNumber string

	IsDead bool

	BirthDate  time.Time
	DepartDate time.Time
	DeathDate  time.Time
}
