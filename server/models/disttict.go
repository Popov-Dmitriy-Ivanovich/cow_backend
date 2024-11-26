package models

type District struct {
	ID uint `gorm:"primaryKey"`
	Region Region
	RegionId uint
	Name string
}