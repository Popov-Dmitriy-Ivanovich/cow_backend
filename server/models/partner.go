package models

type Partner struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Address     *string
	Phone       *string
	Email       *string
	Description string
	LogoPath    *string // replaced byte data
}
