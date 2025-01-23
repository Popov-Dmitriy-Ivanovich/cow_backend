package models

type District struct {
	ID       uint   `gorm:"primaryKey"` // ID района
	Region   Region `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	RegionId uint   `gorm:"index"` // ID региона. в котором район
	Name     string // Название района
}
