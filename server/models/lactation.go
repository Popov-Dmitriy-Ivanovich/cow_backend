package models

type Lactation struct {
	ID uint `gorm:"primaryKey"`

	CowId uint

	CheckMilks []CheckMilk `json:"-"`
	DailyMilks []DailyMilk `json:"-"`

	Number uint // номер лактации
	Date   DateOnly // дата начала лактации

	InsemenationNum  int
	InsemenationDate DateOnly

	CalvingCount int
	CalvingDate  DateOnly

	Abort      bool
	MilkAll    *int
	Milk305    *int
	FatAll     *int
	Fat305     *int
	ProteinAll *int
	Protein305 *int
	Days       *int // количество дней, когда корова дает молоко
}
