package models

type Lactation struct {
	ID uint `gorm:"primaryKey"`

	CowId uint

	CheckMilks []CheckMilk `json:"-"`
	DailyMilks []DailyMilk `json:"-"`

	Number uint // номер лактации

	InsemenationNum  int
	InsemenationDate DateOnly

	CalvingCount int
	CalvingDate  DateOnly

	ServicePeriod *uint // сервис период коровы: время от отела до осеменения
	Abort         bool
	MilkAll       *float64
	Milk305       *float64
	FatAll        *float64
	Fat305        *float64
	ProteinAll    *float64
	Protein305    *float64
	Days          *int // количество дней, когда корова дает молоко
}
