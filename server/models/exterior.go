package models

import (
	"errors"

	"gorm.io/gorm"
)

type Exterior struct {
	ID                    uint
	CowID                 uint
	Rating                float64
	ChestWidth            *float64
	PelvicWidth           *float64
	SacrumHeight          *float64
	BodyDepth             *float64
	ExteriorType          *float64
	BoneQHockJointRear    *float64
	SacrumAngle           *float64
	TopLine               *float64
	HoofAngle             *float64
	HindLegPosSide        *float64
	HindLegPosRead        *float64
	ForeLegPosFront       *float64
	UdderDepth            *float64
	UdderBalance          *float64
	HeightOfUdderAttach   *float64
	ForeUdderAttach       *float64
	ForeUdderPlcRear      *float64
	HindTeatPlc           *float64
	ForeTeatLendth        *float64
	RearTeatLength        *float64
	ForeTeatDiameter      *float64
	RearTeatDiameter      *float64
	CenterLigamentDepth   *float64
	HarmonyOfMovement     *float64
	Conditioning          *float64
	ProminenceOfMilkVeins *float64
	MilkStrength          *float64
	BodyStructure         *float64
	Limbs                 *float64
	Udder                 *float64

	ForeUdderWidth *float64
	HindUdderWidth *float64
	AcrumLength    *float64

	PicturePath *string
}

func (e *Exterior) BeforeCreate(tx *gorm.DB) error {
	if e.MilkStrength == nil || e.BodyStructure == nil || e.Limbs == nil || e.Udder == nil {
		return errors.New("не возможно рассчитать рейтинг, нет одного из признаков со стобальной оценкой")
	}
	e.Rating = (*e.MilkStrength + *e.BodyStructure + *e.Limbs + *e.Udder) / 4.0
	return nil
}
