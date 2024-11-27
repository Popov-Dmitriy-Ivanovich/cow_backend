package models

import "time"

type Cow struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	Farm      Farm
	FarmID    uint

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

	Exterior                float64
	InbrindingCoeffByFamily float64

	Approved    int // int to load database dump
	BirthDate   time.Time
	DepartDate  time.Time
	DeathDate   time.Time
	BirkingDate time.Time
}
