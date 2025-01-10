package models

type Farm struct {
	ID uint `gorm:"primaryKey"`

	// Region   Region `json:"-"`
	// RegionId uint
	HozNumber  *string  // номер хоз-ва
	District   District `json:"-"`
	DistrictId uint     // ID района, в котором находится хозяйство
	Parrent    *Farm    `json:"-"`
	ParrentId  *uint    // ID более управляющего хоз-ва (для хозяйства - холдинг, для фермы - хозяйство)

	Type      uint    // Тип: хозяйство, ферма, холдинг
	Name      string  // Название хозяйства
	NameShort string  // Краткое название хозйства
	Inn       *string // ИНН

	Address     string  // Адрес
	Phone       *string // телефон
	Email       *string // эл. почта
	Description *string // описание
	CowsCount   *uint   // количество коров
}
