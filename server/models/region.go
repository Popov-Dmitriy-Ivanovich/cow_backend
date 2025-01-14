package models

type Region struct {
	ID     uint   `gorm:"primaryKey" default:"1"`
	Name   string `example:"Усть-Каменский"` // Название региона
	News   []News
	RegNum uint // Номер региона (Архангельская область = 29)
}
