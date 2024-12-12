package cows_filter

import (
	"cow_backend/filters"
	"errors"
	"time"
)

type ByCalvingDate struct {
}

func (f ByCalvingDate) Apply(fm filters.FilteredModel) error {
	query := fm.GetQuery()
	bodyData, ok := fm.GetFilterParameters()["object"].(CowsFilter)
	if !ok {
		return errors.New("wrong object provided in filter filed object")
	}
	if bodyData.CalvingDateFrom != nil && bodyData.CalvingDateTo != nil &&
		*bodyData.CalvingDateFrom != "" && *bodyData.CalvingDateTo != "" {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.CalvingDateFrom)
		if err != nil {
			return err
		}
		bdTo, err := time.Parse(time.DateOnly, *bodyData.CalvingDateTo)
		if err != nil {
			return err
		}
		query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.calving_date BETWEEN ? AND ?)", bdFrom.UTC(), bdTo.AddDate(0, 0, 1).UTC()).Preload("Lactation")
	} else if bodyData.CalvingDateFrom != nil && *bodyData.CalvingDateFrom != "" {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.CalvingDateFrom)
		if err != nil {
			return err
		}
		query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.calving_date >= ?)", bdFrom.UTC()).Preload("Lactation")
	} else if bodyData.CalvingDateTo != nil && *bodyData.CalvingDateTo != "" {
		bdTo, err := time.Parse(time.DateOnly, *bodyData.CalvingDateTo)
		if err != nil {
			return err
		}
		query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.calving_date <= ?)", bdTo.AddDate(0, 0, 1).UTC()).Preload("Lactation")
	}
	fm.SetQuery(query)
	return nil
}
