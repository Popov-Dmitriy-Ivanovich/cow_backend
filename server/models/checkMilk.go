package models

import "time"

type CheckMilk struct {
	ID uint `gorm:"primaryKey"`

	Lactation   Lactation
	LactationId uint

	CheckDate time.Time
	Milk      int
	Fat       int
	Protein   int
}
