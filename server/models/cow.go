package models

import "time"

type Cow struct {
	ID        uint      `gorm:"primaryKey" example:"1"` // ID коровы
	CreatedAt time.Time `example:"2007-01-01"`          // Время создания коровы в базе данных
	Farm      *Farm     `json:"-"`
	FarmID    *uint     `example:"1"` // ID фермы, которой корова принадлежит

	FarmGroup   Farm `json:"-"`
	FarmGroupId uint `example:"1"` // ID хозяйства, которому корова принадлежит

	Breed   Breed `json:"-"`
	BreedId uint  `example:"1"` // ID породы коровы

	Sex   Sex  `json:"-"`
	SexId uint `example:"1"` // ID пола коровы

	Father   *Cow  `json:"-"`
	FatherId *uint `example:"1"` // ID коровы отца коровы

	Mother   *Cow  `json:"-"`
	MotherId *uint `example:"1"` // ID коровы матери коровы

	CreatedBy   *User `json:"-"` // пользователь, создавший корову
	CreatedByID *uint `example:"1"`

	Genetic   *Genetic
	Exterior  *Exterior
	Lactation []Lactation `json:"-"`

	IdentificationNumber *string // он все-таки есть! это какой-то не российский номер коровы
	InventoryNumber      *string `example:"1213321"`    // Инвентарный номер коровы
	SelecsNumber         *string `example:"98989"`      // Селекс номер коровы
	RSHNNumber           *string `example:"1323323232"` // РСХН номер коровы
	Name                 string  `example:"Дима"`       // Кличка коровы

	// Exterior                float64  `example:"3.14"` // Оценка экстерьера коровы, будет переделано в ID экстерьера коровы
	InbrindingCoeffByFamily *float64 `example:"3.14"` // Коэф. инбриндинга по роду

	Approved    int       `example:"1"` // Целое число, что-то для админов, чтобы подтверждать коров
	BirthDate   DateOnly  // День рождения
	DepartDate  *DateOnly // День отбытия из коровника
	DeathDate   *DateOnly // Дата смерти
	BirkingDate *DateOnly // Дата перебирковки

	// Новые поля
	PreviousHoz             *Farm   `json:"-"`
	PreviousHozId           *uint   // ID предыдущего хозяйства, когда корову продают, она переходит к новому владельцу и становится "новой коровой"
	BirthHoz                *Farm   `json:"-"`
	BirthHozId              *uint   // ID хозяйства рождения
	PreviousReincarnation   *Cow    `json:"-"` // Одна и та же реальная корова имеет разные инвент. номера, это указатель на эту же корову в другом хоз-ве с другим инв. номером
	PreviousReincarnationId *uint   // см PreviousReincarnation
	BirthMethod             *string // способ зачатия: клон, эмбрион, искусственное осеменени, естественное осеменение
}
