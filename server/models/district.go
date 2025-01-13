package models

type District struct {
	ID       uint   `gorm:"primaryKey"` // ID района
	Region   Region `json:"-"`
	RegionId uint   `gorm:"index"` // ID региона. в котором район
	Name     string // Название района
}
