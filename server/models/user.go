package models

type User struct {
	ID                    uint `gorm:"primaryKey"`
	NameSurnamePatronimic string
	Role                  Role
	RoleId                int
	Email                 string
	Phone                 string
	Password              []byte `json:"-"`
	Farm                  *Farm  `json:"-"`
	FarmId                *uint
	Region                Region `json:"-"`
	RegionId              uint   `example:"1"`
}
