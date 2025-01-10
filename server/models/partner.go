package models

type Partner struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  // Название партнера
	Address     *string // Адрес
	Phone       *string // телефон
	Email       *string // эл. почта
	Description string  // описание партнера
	LogoPath    *string // путь к логотипу партнера относительно genmlik.ru/api/static/partners
}
