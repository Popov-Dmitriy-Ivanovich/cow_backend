package models

type Region struct {
	ID   uint   `gorm:"primaryKey" default:"1"`
	Name string `example:"Усть-Каменский"`
}
