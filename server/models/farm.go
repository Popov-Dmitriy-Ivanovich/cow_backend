package models

type Farm struct {
	ID uint `gorm:"primaryKey"`

	// Region   Region `json:"-"`
	// RegionId uint
	HozNumber  *string  // номер хоз-ва, интересно а можно ли его сдлеать интом?
	District   District `json:"-"`
	DistrictId uint
	Parrent    *Farm `json:"-"`
	ParrentId  *uint

	Type      uint
	Name      string
	NameShort string
	Inn       *string

	Address     string
	Phone       *string
	Email       *string
	Description *string
	CowsCount   *uint
}
