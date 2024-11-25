package models

import "time"

type Event struct {
	ID uint `gorm:"primaryKey"`

	Cow   Cow
	CowId uint

	EventType   EventType
	EventTypeId uint

	Date     time.Time
	Comment1 string
	Comment2 string
}

type EventType struct {
	ID uint `gorm:"primaryKey"`

	Parent   *EventType
	ParentId uint

	Name string

	Code int
	Type int
}
