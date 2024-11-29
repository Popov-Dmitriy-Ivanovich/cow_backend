package models

type Event struct {
	ID uint `gorm:"primaryKey"`

	Cow   Cow `json:"-"`
	CowId uint

	EventType   EventType `json:"-"`
	EventTypeId uint

	Date     DateOnly
	Comment1 string
	Comment2 string
}

type EventType struct {
	ID uint `gorm:"primaryKey"`

	Parent   *EventType `json:"-"`
	ParentId uint

	Name string

	Code int
	Type int
}
