package models

type Event struct {
	ID uint `gorm:"primaryKey"`

	CowId uint // ID коровы

	EventType   EventType
	EventTypeId uint // стандартизированная группа события

	EventType1   EventType
	EventType1Id uint // стандартизированная название события

	EventType2   *EventType
	EventType2Id *uint // стандартизированное разновидность события

	DataResourse      *string // источник данных
	DaysFromLactation uint    // дни от начала лактации

	Date     DateOnly // Дата ветеринарного события
	Comment1 *string  // Коментарий 1 (по всей видиости сюда что-то пришет врач)
	Comment2 *string  // Коментарий 2
}

type EventType struct { // бывший EventList
	ID uint `gorm:"primaryKey"`

	Parent   *EventType `json:"-"`
	ParentId *uint      // ID старшего в иерархии типов события типа (для разновидности события ID группы событий, которой эта разновидность принадлежит)

	Name string // Название группы/названия/разновидности события

	Code uint // код группы или разновидности или названия события
	Type uint // 1 - группа события, 2 - разновидность события, 3 - название события
}
