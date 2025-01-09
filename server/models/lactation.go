package models

type Lactation struct {
	ID uint `gorm:"primaryKey"`

	CowId uint // ID коровы, данные о лактации которой записаны

	CheckMilks []CheckMilk `json:"-"`
	DailyMilks []DailyMilk `json:"-"`

	Number uint // номер лактации

	InsemenationNum  int // Количество осеменений
	InsemenationDate DateOnly // Дата осеменения

	CalvingCount int // Количество рожденных телят: 0 - мертворождение, 2 - двойня
	CalvingDate  DateOnly

	ServicePeriod *uint // сервис период коровы: время от отела до осеменения
	Abort         bool // был ли аборт
	MilkAll       *float64 // Суммарный надой
	Milk305       *float64 // Суммарный надой за 305 дней
	FatAll        *float64 // Суммарный жир
	Fat305        *float64 // Суммарный жир за 305 дней
	ProteinAll    *float64 // Суммарный белок
	Protein305    *float64 // Суммарный белок за 305 дней
	Days          *int // количество дней, когда корова дает молоко
}
