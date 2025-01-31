package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Genetic struct {
	ID                        uint      `gorm:"primaryKey"` // ID записи о генотипировании
	CowID                     uint      `gorm:"index;"`     // ID коровы
	ProbeNumber               string    // Номер пробы
	BloodDate                 *DateOnly `gorm:"index"` // Дата взятия пробы крови
	ResultDate                *DateOnly `gorm:"index"` // Дата получения результата
	InbrindingCoeffByGenotype *float64  `gorm:"index"` // Коэф. инбриндинга по генотипу

	GeneticIllnessesData []GeneticIllnessData `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Список генетических заболеваний, пустой если нет

	GtcFilePath *string // Путь к gtc файлу относительно genmilk.ru/api/static/gtc
}

func (g *Genetic) Validate() error {
	db := dbConnection
	cow := Cow{}
	if err := db.First(&cow, g.CowID).Error; err != nil {
		return errors.New("не найдена корова")
	}
	if cow.BirthDate.After(g.BloodDate.Time) {
		return errors.New("корова должна родиться до сдачи крови")
	}
	if g.ResultDate.Before(g.BloodDate.Time) {
		return errors.New("результат должен быть получен после сдачи крови")
	}
	return nil
}

type GeneticIllnessData struct {
	ID        uint `gorm:"primaryKey"`
	GeneticID uint `gorm:"index"` // ID данных о генотипировании коровы, к которым относятся данные о заболеваниях коровы
	Status    *GeneticIllnessStatus
	StatusID  *uint `gorm:"index"` // ID статус заболевания
	Illness   GeneticIllness
	IllnessID uint `gorm:"index"` // ID заболевание
}

type GeneticIllness struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  // имя генетического заболевания
	Description string  // описание генетического заболевания
	OMIA        *string // Какой-то там ОМИЯ номер
}

type GeneticIllnessStatus struct {
	ID     uint   `gorm:"primaryKey"`
	Status string // Статус заболевания: FREE, CARIER, BAD ...
}

func (g *Genetic) BeforeCreate(tx *gorm.DB) error {
	now := time.Now().UTC()
	if g.ResultDate == nil {
		g.ResultDate = &DateOnly{Time: now.AddDate(-1, 0, 1)}
	}
	if g.BloodDate == nil {
		g.BloodDate = &DateOnly{Time: now}
	}
	return g.Validate()
}

func (g *Genetic) BeforeUpdate(tx *gorm.DB) error {
	now := time.Now().UTC()
	if g.ResultDate == nil {
		g.ResultDate = &DateOnly{Time: now.AddDate(-1, 0, 1)}
	}
	if g.BloodDate == nil {
		g.BloodDate = &DateOnly{Time: now}
	}
	return g.Validate()
}
