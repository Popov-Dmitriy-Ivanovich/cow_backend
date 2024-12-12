package cows_filter

import (
	"cow_backend/filters"
	"errors"
	"time"
)

type ByInsemenationDate struct {
}

func (f ByInsemenationDate) Apply(fm filters.FilteredModel) error {
	query := fm.GetQuery()
	bodyData, ok := fm.GetFilterParameters()["object"].(CowsFilter)
	if !ok {
		return errors.New("wrong object provided in filter filed object")
	}
	if bodyData.InseminationDateFrom != nil && bodyData.InseminationDateTo != nil {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.InseminationDateFrom)
		if err != nil {
			return err
		}
		bdTo, err := time.Parse(time.DateOnly, *bodyData.InseminationDateTo)
		if err != nil {
			return err
		}
		query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.insemenation_date BETWEEN ? and ?)", bdFrom.UTC(), bdTo.AddDate(0, 0, 1).UTC()).Preload("Lactation")
	} else if bodyData.InseminationDateFrom != nil {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.InseminationDateFrom)
		if err != nil {
			return err
		}
		query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.insemenation_date >= ?)", bdFrom.UTC()).Preload("Lactation")
	} else if bodyData.InseminationDateTo != nil {
		bdTo, err := time.Parse(time.DateOnly, *bodyData.InseminationDateTo)
		if err != nil {
			return err
		}
		query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.insemenation_date <= ?)", bdTo.AddDate(0, 0, 1).UTC()).Preload("Lactation")
	}
	fm.SetQuery(query)
	return nil
}
