package models

import "time"

type News struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time

	Region   *Region
	RegionId *uint
	Title    string
	Text     string
}
