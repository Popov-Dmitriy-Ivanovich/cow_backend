package models

type Farm struct {
	ID uint `gorm:"primaryKey"`

	Region   Region
	RegionId uint

	Parrent   *Farm
	ParrentId *uint

	Type        uint
	Name        string
	Inn         string
	District    string
	Address     string
	Phone       string
	Email       string
	Description string
	CowsCount   uint
}
