package models

import (
	"errors"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Cow struct {
	ID        uint      `gorm:"primaryKey" example:"1"` // ID коровы
	CreatedAt time.Time `example:"2007-01-01"`          // Время создания коровы в базе данных

	Farm   *Farm `json:"-"`
	FarmID *uint `gorm:"index" example:"1"` // ID фермы, которой корова принадлежит

	FarmGroup   Farm `json:"-"`
	FarmGroupId uint `gorm:"index" example:"1"` // ID хозяйства, которому корова принадлежит

	Holding   *Farm
	HoldingID *uint `gorm:"index"` // ID холдинга, которому принадлежит корова

	Breed   Breed `json:"-"`
	BreedId uint  `gorm:"index" example:"1"` // ID породы коровы

	Sex   Sex  `json:"-"`
	SexId uint `gorm:"index" example:"1"` // ID пола коровы

	Events []Event `json:"-"`

	GradeRegion   *Grade `json:"-"`
	GradeRegionId *uint  `example:"1"` // оценка по региону

	GradeHoz   *Grade `json:"-"`
	GradeHozId *uint  `example:"1"` // оценка по хозяйству

	FatherSelecs *uint64 // ID коровы отца коровы

	MotherSelecs *uint64 // ID коровы матери коровы

	// CreatedBy   *User `json:"-"` // пользователь, создавший корову
	// CreatedByID *uint `example:"1"`

	Genetic   *Genetic
	Exterior  *Exterior
	Lactation []Lactation `json:"-"`

	IdentificationNumber *string `gorm:"index"`                      // он все-таки есть! это какой-то не российский номер коровы
	InventoryNumber      *string `gorm:"index" example:"1213321"`    // Инвентарный номер коровы
	SelecsNumber         *uint64 `gorm:"index" example:"98989"`      // Селекс номер коровы
	RSHNNumber           *string `gorm:"index" example:"1323323232"` // РСХН номер коровы
	Name                 string  `gorm:"index" example:"Дима"`       // Кличка коровы

	// Exterior                float64  `example:"3.14"` // Оценка экстерьера коровы, будет переделано в ID экстерьера коровы
	InbrindingCoeffByFamily *float64 `gorm:"index" example:"3.14"` // Коэф. инбриндинга по роду

	Approved    int       `gorm:"index" example:"1"` // Целое число, 0 - корова не подтверждена, 1 - корова подтверждена, -1 - корова отклонена
	BirthDate   DateOnly  `gorm:"index"`             // День рождения
	DepartDate  *DateOnly `gorm:"index"`             // День отбытия из коровника
	DeathDate   *DateOnly `gorm:"index"`             // Дата смерти
	BirkingDate *DateOnly `gorm:"index"`             // Дата перебирковки

	// Новые поля
	PreviousHoz   *Farm   `json:"-"`
	PreviousHozId *uint   // ID предыдущего хозяйства, когда корову продают, она переходит к новому владельцу и становится "новой коровой"
	BirthHoz      *Farm   `json:"-"`
	BirthHozId    *uint   // ID хозяйства рождения
	BirthMethod   *string // способ зачатия: клон, эмбрион, искусственное осеменени, естественное осеменение

	PreviousInventoryNumber *string `json:"-"` // Одна и та же реальная корова имеет разные инвент. номера, это предыдущий селекс коровы

	Documents []Document `json:"-"` // документы коровы
}

type Document struct {
	ID    uint   // ID
	CowID uint   `gorm:"index"` // ID коровы, для которой хранитя документ
	Path  string // путь к документу относительно genmilk.ru/api/static/documents
}

func (c *Cow) BeforeCreate(tx *gorm.DB) error {
	if c.RSHNNumber == nil {
		c.RSHNNumber = new(string)
		if c.SelecsNumber == nil {
			return errors.New("нет ни селекса ни РСХН")
		}
		*c.RSHNNumber = "!" + strconv.FormatUint(uint64(*c.SelecsNumber), 10)
	}
	return nil
}
