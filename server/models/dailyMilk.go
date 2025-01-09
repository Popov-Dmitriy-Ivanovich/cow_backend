package models

type DailyMilk struct {
	ID             uint     `gorm:"primaryKey"`
	LactationId    uint    // ID лактации во время котороый была дойка `example:"1"`
	Date           DateOnly // Дата дойки
	Milk           int      // Суммарный надой `example:"12"`
	MilkMorning    *int      // Надой утром `example:"12"`
	MilkNoon       *int      // Надой днем `example:"12"`
	MilkEvening    *int      // Надой вечером `example:"12"`
	Fat            int      // Суммарный жир `example:"12"`
	FatMorning     *int      // Жир утром`example:"12"`
	FatNoon        *int      // Жир днем `example:"12"`
	FatEvening     *int      // Жир вечером `example:"12"`
	Protein        int      // Суммарный белок `example:"12"`
	ProteinMorning *int      // Белок утром `example:"12"`
	ProteinNoon    *int      // Белок днем `example:"12"`
	ProteinEvening *int      // Белок вечером `example:"12"`
}
