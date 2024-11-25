package models

import "time"

type Lactation struct {
	ID uint `gorm:"primaryKey"`

	Cow   Cow
	CowId uint

	Number uint
	Date   time.Time

	OsemenNum  int
	OsemenDate time.Time

	OtelNum  int
	OtelDate time.Time

	Abort      int
	MilkAll    int
	Milk305    int
	FatAll     int
	Fat305     int
	ProteinAll int
	Protein305 int
	Days       int
}
