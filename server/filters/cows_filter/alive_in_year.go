package cows_filter

import (
	"cow_backend/filters"
	"errors"
	"strconv"
	"time"
)

type AliveInYear struct {
}

func (f AliveInYear) Apply(fm filters.FilteredModel) error {
	query := fm.GetQuery()

	year, ok := fm.GetFilterParameters()["year"]
	if !ok {
		return nil
	}
	yearStr, ok := year.(string)
	if !ok {
		return errors.New("year is not passed as string")
	}
	if yearInt, err := strconv.ParseInt(yearStr, 10, 64); err == nil {
		query = query.Where("birth_date <= ? AND (death_date is NULL or death_date <= ?)",
			time.Date(int(yearInt)+1, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(int(yearInt)+1, 1, 1, 0, 0, 0, 0, time.UTC))
	} else {
		return err
	}
	fm.SetQuery(query)
	return nil
}
