package models

type UserRegisterRequest struct {
	ID        uint
	HozNumber string // номер хоз-ва к которому привязвыается пользователь: либо существует, либо newHoz

	NameSurnamePatronimic string
	RoleId                uint
	Email                 string
	Phone                 string
	Password              string

	RegionId uint
}
type HozRegisterRequest struct {
	ID uint

	HoldNumber string

	HozNumber  string
	DistrictId uint

	Name        string
	ShortName   string
	Inn         *string
	Address     *string
	Phone       *string
	Email       *string
	Description *string
	CowsCount   *uint
}
type HoldRegisterRequest struct {
	ID          uint
	HozNumber   string
	DistrictId  string
	Name        string
	ShortName   string
	Inn         *string
	Address     *string
	Phone       *string
	Email       *string
	Description *string
	CowsCount   *uint
}
