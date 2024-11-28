package models

type District struct {
	ID       uint   `gorm:"primaryKey"`
	Region   Region `json:"-"`
	RegionId uint
	Name     string
}
