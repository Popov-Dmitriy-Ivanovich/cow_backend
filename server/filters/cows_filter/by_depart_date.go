package cows_filter

import (
	"cow_backend/filters"
	"errors"
	"time"
)

type ByDepartDate struct {
}

func (f ByDepartDate) Apply(fm filters.FilteredModel) error {
	query := fm.GetQuery()
	bodyData, ok := fm.GetFilterParameters()["object"].(CowsFilter)
	if !ok {
		return errors.New("wrong object provided in filter filed object")
	}
	if bodyData.DepartDateFrom != nil && bodyData.DepartDateTo != nil &&
		*bodyData.DepartDateFrom != "" && *bodyData.DepartDateTo != "" {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.DepartDateFrom)
		if err != nil {
			return err
		}
		bdTo, err := time.Parse(time.DateOnly, *bodyData.DepartDateTo)
		if err != nil {
			return err
		}
		query = query.Where("depart_date BETWEEN ? AND ?", bdFrom.UTC(), bdTo.AddDate(0, 0, 1).UTC())
	} else if bodyData.DepartDateFrom != nil && *bodyData.DepartDateFrom != "" {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.DepartDateFrom)
		if err != nil {
			return err
		}
		query = query.Where("depart_date >= ?", bdFrom.UTC())
	} else if bodyData.DepartDateTo != nil && *bodyData.DepartDateTo != "" {
		bdTo, err := time.Parse(time.DateOnly, *bodyData.DepartDateTo)
		if err != nil {
			return err
		}
		query = query.Where("depart_date <= ?", bdTo.AddDate(0, 0, 1).UTC())
	}
	fm.SetQuery(query)
	return nil
}
