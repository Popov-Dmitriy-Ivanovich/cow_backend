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
	if !ok { return errors.New("wrong object provided in filter filed object")}
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
		query = query.Where("birking_date BETWEEN ? AND ?", bdFrom, bdTo)
	} else if bodyData.BirkingDateFrom != nil && *bodyData.BirkingDateFrom != "" {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.BirkingDateFrom)
		if err != nil {
			return err
		}
		query = query.Where("birking_date >= ?", bdFrom)
	} else if bodyData.BirkingDateTo != nil && *bodyData.BirkingDateTo != "" {
		bdTo, err := time.Parse(time.DateOnly, *bodyData.BirkingDateTo)
		if err != nil {
			return err
		}
		query = query.Where("birking_date <= ?", bdTo)
	}
	fm.SetQuery(query)
	return nil
}