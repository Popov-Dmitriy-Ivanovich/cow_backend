package models

import "time"

type Lactation struct {
	ID uint `gorm:"primaryKey"`

	CowId uint

	CheckMilks []CheckMilk
	DailyMilks []DailyMilk

	Number uint
	Date   time.Time

	InsemenationNum  int
	InsemenationDate time.Time

	CalvingCount int
	CalvingDate  time.Time

	Abort      bool
	MilkAll    int
	Milk305    int
	FatAll     int
	Fat305     int
	ProteinAll int
	Protein305 int
	Days       int
}
