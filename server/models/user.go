package models

type User struct {
	ID                    uint   `gorm:"primaryKey"`
	NameSurnamePatronimic string // ФИО
	Role                  Role
	RoleId                int    // ID роли
	Email                 string // Почта
	Phone                 string // телефон
	Password              []byte `json:"-"`
	Farm                  *Farm  `json:"-"`
	FarmId                *uint  // ID хозяйства
	Region                Region `json:"-"`
	RegionId              uint   `example:"1"` // ID региона
}
