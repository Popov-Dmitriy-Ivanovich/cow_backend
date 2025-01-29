package models

// DownSides
// Поля количество сделаны числами с плавающей запятой из предположения, что когда-нибудь кто-нибудь подумает, что
// недостаток "рак матки" сильно критичнее, чем недостаток "некрасивый окрас"
type DownSides struct {
	ID         uint `gorm:"primary_key" json:"-"`
	ExteriorID uint `json:"-" gorm:"index"`

	MilkTypeCount float64 // Молочный тип
	BodyCount     float64 // Туловище
	LimbsCount    float64 // Конечности
	UdderCount    float64 // Вымя
	SacrumCount   float64 // Крестец

	MilkTypeDescription string // Список недостатков Молочного типа. Разделитель: "/"
	BodyDescription     string // Список недостатков туловища разделитель: "/"
	LimbsDescription    string // Список недостатков конечностей разделитель: "/"
	UdderDescription    string // Список недостатков вымени разделитель: "/"
	SacrumDescription   string // Список недостатков крестца разделитель: "/"

	Summary string // Сводные данные о недостатках разделитель "/"
}
