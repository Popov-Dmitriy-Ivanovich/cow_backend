package models

import "time"

type CheckMilk struct {
	ID uint `gorm:"primaryKey"`

	LactationId *uint

	CheckDate time.Time
	Milk      int
	Fat       int
	Protein   int
}
