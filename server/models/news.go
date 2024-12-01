package models

// import "time"

type News struct {
	ID        uint `gorm:"primaryKey"`
	Date DateOnly

	Region   *Region `json:"-"`
	RegionId *uint
	Title    string
	Text     string
}
