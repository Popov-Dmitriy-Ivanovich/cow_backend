package models

type Farm struct {
	ID uint `gorm:"primaryKey"`

	Region   Region
	RegionId uint

	District   *District
	DistrictId *uint
	Parrent    *Farm
	ParrentId  *uint

	Type      uint
	Name      string
	NameShort string
	Inn       string

	Address     string
	Phone       string
	Email       string
	Description string
	CowsCount   uint
}
