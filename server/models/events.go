package models

type Event struct {
	ID uint `gorm:"primaryKey"`

	CowId uint

	EventType   	EventType
	EventTypeId 	uint

	EventType1 		EventType
	EventType1Id 	uint

	EventType2 		EventType
	EventType2Id 	uint

	DataResourse      string // источник данные
	DaysFromLactation uint   // дни от начала лактации

	Date     DateOnly
	Comment1 *string
	Comment2 *string
}

type EventType struct { // бывший EventList
	ID uint `gorm:"primaryKey"`

	Parent   *EventType `json:"-"`
	ParentId *uint

	Name string

	Code uint
	Type uint // Should be enum
}
