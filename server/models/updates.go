package models

import "time"

type Update struct {
	ID   uint      `gorm:"primaryKey"`
	Date time.Time // Дата обновления базы данных
}
