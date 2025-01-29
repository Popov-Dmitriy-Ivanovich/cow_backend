package models

import (
	"errors"
	"gorm.io/gorm"
)

type Event struct {
	ID uint `gorm:"primaryKey"`

	CowId uint `gorm:"index"` // ID коровы

	EventType   EventType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	EventTypeId uint      `gorm:"index"` // Стандартизированная группа события

	EventType1   EventType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	EventType1Id uint      `gorm:"index"` // Стандартизированная название события

	EventType2   *EventType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	EventType2Id *uint      `gorm:"index"` // Стандартизированное разновидность события

	DataResourse      *string // Источник данных
	DaysFromLactation uint    // Дни от начала лактации

	Date     *DateOnly `gorm:"index"` // Дата ветеринарного события
	Comment1 *string   // Комментарий 1 (по всей видимости сюда что-то пришит врач)
	Comment2 *string   // Комментарий 2
}

func (e *Event) Validate() error {
	db := dbConnection
	cow := Cow{}
	if err := db.First(&cow, e.CowId).Error; err != nil {
		return errors.New("не найдена корова, для которой добавляется вет. событие")
	}
	if e.Date != nil {
		if cow.BirthDate.After(e.Date.Time) {
			return errors.New("вет. событие не может случиться до рождения коровы")
		}
	}
	return nil
}

func (e *Event) BeforeCreate(tx *gorm.DB) error {
	return e.Validate()
}

func (e *Event) BeforeUpdate(tx *gorm.DB) error {
	return e.Validate()
}

type EventType struct { // бывший EventList
	ID uint `gorm:"primaryKey"`

	Parent   *EventType `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	ParentId *uint      // ID старшего в иерархии типов события типа (для разновидности события ID группы событий, которой эта разновидность принадлежит)

	Name string // Название группы/названия/разновидности события

	Code uint // Код группы или разновидности или названия события
	Type uint `gorm:"index"` // 1 - группа события, 2 - разновидность события, 3 - название события
}
