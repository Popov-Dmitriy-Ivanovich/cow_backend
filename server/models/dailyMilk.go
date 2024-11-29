package models

type DailyMilk struct {
	ID             uint     `gorm:"primaryKey"`
	LactationId    *uint    // ID лактации во время котороый была дойка `example:"1"`
	Date           DateOnly // Дата дойки
	Milk           int      // Параметр дойки `example:"12"`
	MilkMorning    int      // Параметр дойки `example:"12"`
	MilkNoon       int      // Параметр дойки `example:"12"`
	MilkEvening    int      // Параметр дойки `example:"12"`
	Fat            int      // Параметр дойки `example:"12"`
	FatMorning     int      // Параметр дойки `example:"12"`
	FatNoon        int      // Параметр дойки `example:"12"`
	FatEvening     int      // Параметр дойки `example:"12"`
	Protein        int      // Параметр дойки `example:"12"`
	ProteinMorning int      // Параметр дойки `example:"12"`
	ProteinNoon    int      // Параметр дойки `example:"12"`
	ProteinEvening int      // Параметр дойки `example:"12"`
}
