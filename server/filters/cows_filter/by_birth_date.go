package cows_filter

import (
	"cow_backend/filters"
	"errors"
	"time"
)

type ByBrithDate struct {
}

func (f ByBrithDate) Apply(fm filters.FilteredModel) error {
	query := fm.GetQuery()
	bodyData, ok := fm.GetFilterParameters()["object"].(CowsFilter)
	if !ok {
		return errors.New("wrong object provided in filter filed object")
	}
	if bodyData.BirthDateFrom != nil && bodyData.BirthDateTo != nil &&
		*bodyData.BirthDateFrom != "" && *bodyData.BirthDateTo != "" {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.BirthDateFrom)
		if err != nil {
			return err
		}
		bdTo, err := time.Parse(time.DateOnly, *bodyData.BirthDateTo)
		if err != nil {
			return err
		}
		query = query.Where("birth_date BETWEEN ? AND ?", bdFrom.UTC(), bdTo.AddDate(0, 0, 1).UTC())
	} else if bodyData.BirthDateFrom != nil && *bodyData.BirthDateFrom != "" {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.BirthDateFrom)
		if err != nil {
			return err
		}
		query = query.Where("birth_date >= ?", bdFrom.UTC())
	} else if bodyData.BirthDateTo != nil && *bodyData.BirthDateTo != "" {
		bdTo, err := time.Parse(time.DateOnly, *bodyData.BirthDateTo)
		if err != nil {
			return err
		}
		query = query.Where("birth_date <= ?", bdTo.AddDate(0, 0, 1).UTC())
	}
	fm.SetQuery(query)
	return nil
}
