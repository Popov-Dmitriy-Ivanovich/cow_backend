package models

type Event struct {
	ID uint `gorm:"primaryKey"`

	CowId uint `gorm:"index"` // ID коровы

	EventType   EventType
	EventTypeId uint `gorm:"index"` // Стандартизированная группа события

	EventType1   EventType
	EventType1Id uint `gorm:"index"` // Стандартизированная название события

	EventType2   *EventType
	EventType2Id *uint `gorm:"index"` // Стандартизированное разновидность события

	DataResourse      *string // Источник данных
	DaysFromLactation uint    // Дни от начала лактации

	Date     DateOnly `gorm:"index"` // Дата ветеринарного события
	Comment1 *string  // Комментарий 1 (по всей видимости сюда что-то пришит врач)
	Comment2 *string  // Комментарий 2
}

type EventType struct { // бывший EventList
	ID uint `gorm:"primaryKey"`

	Parent   *EventType `json:"-"`
	ParentId *uint      // ID старшего в иерархии типов события типа (для разновидности события ID группы событий, которой эта разновидность принадлежит)

	Name string // Название группы/названия/разновидности события

	Code uint // код группы или разновидности или названия события
	Type uint `gorm:"index"` // 1 - группа события, 2 - разновидность события, 3 - название события
}
