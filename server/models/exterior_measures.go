package models

type Measures struct {
	ID         uint `gorm:"primaryKey" json:"-"`
	ExteriorID uint `json:"-" gorm:"index"`

	SacrumHeight *float64 // высота в крестце (Сантиметры или градусы)
	ChestWidth   *float64 // Ширина груди (Сантиметры или градусы)
	SacrumAngle  *float64 // Угол наклона крестца (Сантиметры или градусы)
	SacrumWidth  *float64 // ширина в крестце (Сантиметры или градусы)

	HindLegWalkSideView *float64 // Поступь задних ног сбоку (Сантиметры или градусы)
	HoofAngle           *float64 // Угол копыта (Сантиметры или градусы)

	UdderDepth                      *float64 // Глубина вымени (Сантиметры или градусы)
	FrontUdderSegmentsLocation      *float64 // Прикрепление передних долей вымени (Сантиметры или градусы)
	BackUdderSegmentsLocationHeight *float64 // Высота прикрепления задних долей вымени (Сантиметры или градусы)
	BackUdderSegmentsWidth          *float64 // Ширина задних долей вымени (Сантиметры или градусы)
	CentralLigamentDepth            *float64 // Глубина центральной связки (Сантиметры или градусы)
	FrontNippleLength               *float64 // Длинна передних сосков (Сантиметры или градусы)

	UdderBalance        *float64 // Баланс вымени (Сантиметры или градусы)
	FrontNippleDiameter *float64 // Диаметр передних сосков (Сантиметры или градусы)
	BackNippleDiameter  *float64 // Диаметр задних сосков (Сантиметры или градусы)
}
