package models

import "time"

type CheckMilk struct {
	ID uint `gorm:"primaryKey" example:"1"` // ID контрольной дойки

	LactationId *uint `example:"1"` // ID лактации для которой выполнена контрольная дойка

	CheckDate time.Time `example:"1999-02-12"` // Время конрольной дойки
	Milk      int       `example:"1"`          // Параметр контрольной дойки, как я понимаю кол-во молока
	Fat       int       `example:"1"`          // Параметр контрольной дойки, как я понимаю кол-во жира в молоке
	Protein   int       `example:"1"`          // Параметр контрольной дойки, как я понимаю кол-во белка в молоке
}
