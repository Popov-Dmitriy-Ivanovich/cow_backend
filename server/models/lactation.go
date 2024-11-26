package models

import "time"

type Lactation struct {
	ID uint `gorm:"primaryKey"`

	CowId uint

	CheckMilks []CheckMilk
	DailyMilks []DailyMilk

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
