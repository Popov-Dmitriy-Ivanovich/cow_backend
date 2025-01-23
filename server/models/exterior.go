package models

import (
	"errors"
	"gorm.io/gorm"
)

type Exterior struct {
	ID                    uint `gorm:"primaryKey"`
	CowID                 uint `gorm:"index;"`
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

	// 100 бальные
	MilkStrength  *float64
	BodyStructure *float64
	Limbs         *float64
	Udder         *float64
	// 100 бальные

	ForeUdderWidth *float64
	HindUdderWidth *float64
	AcrumLength    *float64

	PicturePath *string
}

// Validate
// 100 строк, 100, Карл!
func (e *Exterior) Validate() error {
	if e.MilkStrength != nil && (*e.MilkStrength < 0 || *e.MilkStrength > 100) {
		return errors.New("milkStrength must be between 0 and 100")
	}
	if e.BodyStructure != nil && (*e.BodyStructure < 0 || *e.BodyStructure > 100) {
		return errors.New("bodyStructure must be between 0 and 100")
	}
	if e.Limbs != nil && (*e.Limbs < 0 || *e.Limbs > 100) {
		return errors.New("limbs must be between 0 and 100")
	}
	if e.Udder != nil && (*e.Udder < 0 || *e.Udder > 100) {
		return errors.New("udder must be between 0 and 100")
	}

	if e.ChestWidth != nil && (*e.ChestWidth < 0 || *e.ChestWidth > 10) {
		return errors.New("chestWidth must be between 0 and 10")
	}
	if e.PelvicWidth != nil && (*e.PelvicWidth < 0 || *e.PelvicWidth > 10) {
		return errors.New("pelvicWidth must be between 0 and 10")
	}
	if e.SacrumHeight != nil && (*e.SacrumHeight < 0 || *e.SacrumHeight > 10) {
		return errors.New("sacrumHeight must be between 0 and 10")
	}
	if e.BodyDepth != nil && (*e.BodyDepth < 0 || *e.BodyDepth > 10) {
		return errors.New("bodyDepth must be between 0 and 10")
	}
	if e.ExteriorType != nil && (*e.ExteriorType < 0 || *e.ExteriorType > 10) {
		return errors.New("exteriorType must be between 0 and 10")
	}
	if e.BoneQHockJointRear != nil && (*e.BoneQHockJointRear < 0 || *e.BoneQHockJointRear > 10) {
		return errors.New("boneQHockJointRear must be between 0 and 10")
	}
	if e.SacrumAngle != nil && (*e.SacrumAngle < 0 || *e.SacrumAngle > 10) {
		return errors.New("sacrumAngle must be between 0 and 10")
	}
	if e.TopLine != nil && (*e.TopLine < 0 || *e.TopLine > 10) {
		return errors.New("topLine must be between 0 and 10")
	}
	if e.HoofAngle != nil && (*e.HoofAngle < 0 || *e.HoofAngle > 10) {
		return errors.New("hoofAngle must be between 0 and 10")
	}
	if e.HindLegPosSide != nil && (*e.HindLegPosSide < 0 || *e.HindLegPosSide > 10) {
		return errors.New("hindLegPosSide must be between 0 and 10")
	}
	if e.HindLegPosRead != nil && (*e.HindLegPosRead < 0 || *e.HindLegPosRead > 10) {
		return errors.New("hindLegPosRead must be between 0 and 10")
	}
	if e.ForeLegPosFront != nil && (*e.ForeLegPosFront < 0 || *e.ForeLegPosFront > 10) {
		return errors.New("foreLegPosFront must be between 0 and 10")
	}
	if e.UdderDepth != nil && (*e.UdderDepth < 0 || *e.UdderDepth > 10) {
		return errors.New("udderDepth must be between 0 and 10")
	}
	if e.UdderBalance != nil && (*e.UdderBalance < 0 || *e.UdderBalance > 10) {
		return errors.New("udderBalance must be between 0 and 10")
	}
	if e.HeightOfUdderAttach != nil && (*e.HeightOfUdderAttach < 0 || *e.HeightOfUdderAttach > 10) {
		return errors.New("heightOfUdderAttach must be between 0 and 10")
	}
	if e.ForeUdderAttach != nil && (*e.ForeUdderAttach < 0 || *e.ForeUdderAttach > 10) {
		return errors.New("foreUdderAttach must be between 0 and 10")
	}
	if e.ForeUdderPlcRear != nil && (*e.ForeUdderPlcRear < 0 || *e.ForeUdderPlcRear > 10) {
		return errors.New("foreUdderPlcRear must be between 0 and 10")
	}
	if e.HindTeatPlc != nil && (*e.HindTeatPlc < 0 || *e.HindTeatPlc > 10) {
		return errors.New("hindTeatPlc must be between 0 and 10")
	}
	if e.ForeTeatLendth != nil && (*e.ForeTeatLendth < 0 || *e.ForeTeatLendth > 10) {
		return errors.New("foreTeatLength must be between 0 and 10")
	}
	if e.RearTeatLength != nil && (*e.RearTeatLength < 0 || *e.RearTeatLength > 10) {
		return errors.New("rearTeatLength must be between 0 and 10")
	}
	if e.ForeTeatDiameter != nil && (*e.ForeTeatDiameter < 0 || *e.ForeTeatDiameter > 10) {
		return errors.New("foreTeatDiameter must be between 0 and 10")
	}
	if e.RearTeatDiameter != nil && (*e.RearTeatDiameter < 0 || *e.RearTeatDiameter > 10) {
		return errors.New("rearTeatDiameter must be between 0 and 10")
	}
	if e.CenterLigamentDepth != nil && (*e.CenterLigamentDepth < 0 || *e.CenterLigamentDepth > 10) {
		return errors.New("centerLigamentDepth must be between 0 and 10")
	}
	if e.HarmonyOfMovement != nil && (*e.HarmonyOfMovement < 0 || *e.HarmonyOfMovement > 10) {
		return errors.New("harmonyOfMovement must be between 0 and 10")
	}
	if e.Conditioning != nil && (*e.Conditioning < 0 || *e.Conditioning > 10) {
		return errors.New("conditioning must be between 0 and 10")
	}
	if e.ProminenceOfMilkVeins != nil && (*e.ProminenceOfMilkVeins < 0 || *e.ProminenceOfMilkVeins > 10) {
		return errors.New("prominenceOfMilkVeins must be between 0 and 10")
	}
	if e.ForeUdderWidth != nil && (*e.ForeUdderWidth < 0 || *e.ForeUdderWidth > 10) {
		return errors.New("foreUdderWidth must be between 0 and 10")
	}
	if e.HindUdderWidth != nil && (*e.HindUdderWidth < 0 || *e.HindUdderWidth > 10) {
		return errors.New("hindUdderWidth must be between 0 and 10")
	}
	if e.AcrumLength != nil && (*e.AcrumLength < 0 || *e.AcrumLength > 10) {
		return errors.New("acrumLength must be between 0 and 10")
	}
	return nil
}

func (e *Exterior) BeforeCreate(tx *gorm.DB) error {
	if e.MilkStrength == nil || e.BodyStructure == nil || e.Limbs == nil || e.Udder == nil {
		return nil //errors.New("не возможно рассчитать рейтинг, нет одного из признаков со стобальной оценкой")
	}
	e.Rating = (*e.MilkStrength + *e.BodyStructure + *e.Limbs + *e.Udder) / 4.0
	return e.Validate()
}

func (e *Exterior) BeforeUpdate(tx *gorm.DB) error {
	if e.MilkStrength == nil || e.BodyStructure == nil || e.Limbs == nil || e.Udder == nil {
		return nil //errors.New("не возможно рассчитать рейтинг, нет одного из признаков со стобальной оценкой")
	}
	e.Rating = (*e.MilkStrength + *e.BodyStructure + *e.Limbs + *e.Udder) / 4.0
	return e.Validate()
}
