package models

type Event struct {
	ID uint `gorm:"primaryKey"`

	CowId uint

	EventType   EventType // стандартизированная группа события
	EventTypeId uint

	EventType1   EventType // стандартизированная название события
	EventType1Id uint

	EventType2   *EventType // стандартизированное разновидность события
	EventType2Id *uint

	DataResourse      *string // источник данные
	DaysFromLactation uint    // дни от начала лактации

	Date     DateOnly
	Comment1 *string
	Comment2 *string
}

type EventType struct { // бывший EventList
	ID uint `gorm:"primaryKey"`

	Parent   *EventType `json:"-"`
	ParentId *uint

	Name string // имя этой штуки

	Code uint // код группы или разновидности или события
	Type uint // 1 - группа события, 2 - разновидность события, 3 - название события
}
