package models

import (
	"time"
)

type DailyMilk struct {
	ID uint `gorm:"primaryKey"`

	Date           time.Time
	Milk           int
	MilkMorning    int
	MilkNoon       int
	MilkEvening    int
	Fat            int
	FatMorning     int
	FatNoon        int
	FatEvening     int
	Protein        int
	ProteinMorning int
	ProteinNoon    int
	ProteinEvening int
}
