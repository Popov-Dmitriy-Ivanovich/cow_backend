package models

type Lactation struct {
	ID uint `gorm:"primaryKey"`

	CowId uint `gorm:"index"` // ID коровы, данные о лактации которой записаны

	CheckMilks []CheckMilk `json:"-"`
	DailyMilks []DailyMilk `json:"-"`

	Number uint // номер лактации

	InsemenationNum  int      // Количество осеменений
	InsemenationDate DateOnly `gorm:"index"` // Дата осеменения

	CalvingCount int      `gorm:"index"` // Количество рожденных телят: 0 - мертворождение, 2 - двойня
	CalvingDate  DateOnly `gorm:"index"`

	ServicePeriod *uint    `gorm:"index"` // сервис период коровы: время от отела до осеменения
	Abort         bool     `gorm:"index"` // был ли аборт
	MilkAll       *float64 // Суммарный надой
	Milk305       *float64 // Суммарный надой за 305 дней
	FatAll        *float64 // Суммарный жир
	Fat305        *float64 // Суммарный жир за 305 дней
	ProteinAll    *float64 // Суммарный белок
	Protein305    *float64 // Суммарный белок за 305 дней
	Days          *int     // количество дней, когда корова дает молоко
}
