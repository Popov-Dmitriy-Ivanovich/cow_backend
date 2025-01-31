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

	Farm   *Farm `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	FarmID *uint `gorm:"index" example:"1"` // ID фермы, которой корова принадлежит

	FarmGroup   Farm `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	FarmGroupId uint `gorm:"index" example:"1"` // ID хозяйства, которому корова принадлежит

	Holding   *Farm `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	HoldingID *uint `gorm:"index"` // ID холдинга, которому принадлежит корова

	Breed   Breed `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	BreedId uint  `gorm:"index" example:"1"` // ID породы коровы

	Sex   Sex  `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	SexId uint `gorm:"index" example:"1"` // ID пола коровы

	Events []Event `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	GradeRegion *GradeRegion `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	GradeHoz *GradeHoz `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	GradeCountry *GradeCountry `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	FatherSelecs *uint64 // ID коровы отца коровы

	MotherSelecs *uint64 // ID коровы матери коровы

	// CreatedBy   *User `json:"-"` // пользователь, создавший корову
	// CreatedByID *uint `example:"1"`

	Genetic   *Genetic    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Exterior  *Exterior   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Lactation []Lactation `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	IdentificationNumber *string `gorm:"index"`                      // Он все-таки есть! Это какой-то не российский номер коровы
	InventoryNumber      *string `gorm:"index" example:"1213321"`    // Инвентарный номер коровы
	SelecsNumber         *uint64 `gorm:"index" example:"98989"`      // Селекс номер коровы
	RSHNNumber           *string `gorm:"index" example:"1323323232"` // РСХН номер коровы
	Name                 string  `gorm:"index" example:"Дима"`       // Кличка коровы

	// Exterior                float64  `example:"3.14"` // Оценка экстерьера коровы, будет переделано в ID экстерьера коровы
	InbrindingCoeffByFamily *float64 `gorm:"index" example:"3.14"` // Коэф. инбриндинга по роду

	Approved    int       `gorm:"index" example:"1"` // Целое число, 0 - корова не подтверждена, 1 - корова подтверждена, -1 - корова отклонена
	BirthDate   *DateOnly `gorm:"index"`             // День рождения
	DepartDate  *DateOnly `gorm:"index"`             // День отбытия из коровника
	DeathDate   *DateOnly `gorm:"index"`             // Дата смерти
	BirkingDate *DateOnly `gorm:"index"`             // Дата перебирковки

	// Новые поля
	PreviousHoz   *Farm   `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PreviousHozId *uint   // ID предыдущего хозяйства, когда корову продают, она переходит к новому владельцу и становится "новой коровой"
	BirthHoz      *Farm   `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BirthHozId    *uint   // ID хозяйства рождения
	BirthMethod   *string // Способ зачатия: клон, эмбрион, искусственное осеменени, естественное осеменение

	PreviousInventoryNumber *string `json:"-"` // Одна и та же реальная корова имеет разные инвент. номера, это предыдущий селекс коровы

	Documents []Document `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Документы коровы
}

type Document struct {
	ID    uint   // ID
	CowID uint   `gorm:"index;"` // ID коровы, для которой хранитя документ
	Path  string // Путь к документу относительно genmilk.ru/api/static/documents
}

func (c *Cow) Validate() error {
	if c.DepartDate != nil && c.DepartDate.Before(c.BirthDate.Time) {
		return errors.New("дата выбытия не может быть меньше даты рождения")
	}
	if c.DeathDate != nil && c.DeathDate.Before(c.BirthDate.Time) {
		return errors.New("дата смерти не может быть меньше даты рождения")
	}
	if c.BirkingDate != nil && c.BirkingDate.Before(c.BirthDate.Time) {
		return errors.New("дата перебирковки не может быть меньше даты рождения")
	}
	return nil
}

func (c *Cow) BeforeCreate(tx *gorm.DB) error {
	if c.RSHNNumber == nil {
		c.RSHNNumber = new(string)
		if c.SelecsNumber == nil {
			return errors.New("нет ни селекса ни РСХН")
		}
		*c.RSHNNumber = "!" + strconv.FormatUint(uint64(*c.SelecsNumber), 10)
	}
	return c.Validate()
}

func (c *Cow) BeforeUpdate(tx *gorm.DB) error {
	if c.SelecsNumber != nil && c.RSHNNumber == nil {
		c.RSHNNumber = new(string)
		*c.RSHNNumber = "!" + strconv.FormatUint(uint64(*c.SelecsNumber), 10)
	}
	return c.Validate()
}
