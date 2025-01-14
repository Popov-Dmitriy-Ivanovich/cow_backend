package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// DateOnly
// Структура, сохраняющаяся в БД и представляющая дату без времени
type DateOnly struct {
	time.Time
}

func (do DateOnly) Value() (driver.Value, error) {
	return do.Add(0), nil
}

func (do *DateOnly) Scan(src any) error {
	time, ok := src.(time.Time)
	if !ok {
		return errors.New("not time provided to DateOnly")
	}
	*do = DateOnly{time}
	return nil
}

func (do DateOnly) MarshalJSON() ([]byte, error) {
	timeStr := do.Format(time.DateOnly)
	return json.Marshal(timeStr)
}

func (do *DateOnly) ToTime() time.Time {
	return do.Add(0)
}
