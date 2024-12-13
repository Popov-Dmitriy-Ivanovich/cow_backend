package cows_filter

import (
	"cow_backend/filters"
	"errors"
	"time"
)

type ByBirkingDate struct {
}

func (f ByBirkingDate) Apply(fm filters.FilteredModel) error {
	query := fm.GetQuery()
	bodyData, ok := fm.GetFilterParameters()["object"].(CowsFilter)
	if !ok {
		return errors.New("wrong object provided in filter filed object")
	}
	if bodyData.BirkingDateFrom != nil && bodyData.BirkingDateTo != nil &&
		*bodyData.BirkingDateFrom != "" && *bodyData.BirkingDateTo != "" {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.BirkingDateFrom)
		if err != nil {
			return err
		}
		bdTo, err := time.Parse(time.DateOnly, *bodyData.BirkingDateTo)
		if err != nil {
			return err
		}
		query = query.Where("birking_date BETWEEN ? AND ?", bdFrom.UTC(), bdTo.AddDate(0, 0, 1).UTC())
	} else if bodyData.BirkingDateFrom != nil && *bodyData.BirkingDateFrom != "" {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.BirkingDateFrom)
		if err != nil {
			return err
		}
		query = query.Where("birking_date >= ?", bdFrom.UTC())
	} else if bodyData.BirkingDateTo != nil && *bodyData.BirkingDateTo != "" {
		bdTo, err := time.Parse(time.DateOnly, *bodyData.BirkingDateTo)
		if err != nil {
			return err
		}
		query = query.Where("birking_date <= ?", bdTo.AddDate(0, 0, 1).UTC())
	}
	fm.SetQuery(query)
	return nil
}
