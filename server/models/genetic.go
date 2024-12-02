package models

type Genetic struct {
	ID                       uint `gorm:"primaryKey"`
	CowID                    uint
	ProbeNumber              string
	BloodDate                DateOnly
	ResultDate               DateOnly
	IbrindingCoeffByGenotype float64
	GeneticIllnesses         []GeneticIllness `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;many2many:genetic_genetic_illnesses;"`
}
type GeneticIllness struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Descritpion string
	OMIA        *string
}
