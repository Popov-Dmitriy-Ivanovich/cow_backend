package models

type User struct {
	ID                    uint `gorm:"primaryKey"`
	NameSurnamePatronimic *string
	Role                  int
	Email                 string
	Phone                 *string
	Password              []byte
	Farm                  *Farm
	FarmId                *uint
}
