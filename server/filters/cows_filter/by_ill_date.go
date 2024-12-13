package cows_filter

import (
	"cow_backend/filters"
	"errors"
	"time"
)

type ByIllDate struct {
}

func (f ByIllDate) Apply(fm filters.FilteredModel) error {
	query := fm.GetQuery()
	bodyData, ok := fm.GetFilterParameters()["object"].(CowsFilter)
	if !ok {
		return errors.New("wrong object provided in filter filed object")
	}

	if bodyData.IllDateFrom != nil && bodyData.IllDateTo != nil &&
		*bodyData.IllDateFrom != "" && *bodyData.IllDateTo != "" {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.IllDateFrom)
		if err != nil {
			return err
		}
		bdTo, err := time.Parse(time.DateOnly, *bodyData.IllDateTo)
		if err != nil {
			return err
		}
		query = query.Where("EXISTS( SELECT 1 FROM events where events.cow_id = cows.id AND events.event_type_id in (1, 2, 3, 4) AND events.date BETWEEN ? AND ? )",
			bdFrom.UTC(),
			bdTo.AddDate(0, 0, 1).UTC()).Preload("Events")
	} else if bodyData.IllDateFrom != nil && *bodyData.IllDateFrom != "" {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.IllDateFrom)
		if err != nil {
			return err
		}
		query = query.Where("EXISTS( SELECT 1 FROM events where events.cow_id = cows.id AND events.event_type_id in (1, 2, 3, 4) AND events.date >= ? )", bdFrom.UTC()).Preload("Events")
	} else if bodyData.IllDateTo != nil && *bodyData.IllDateTo != "" {
		bdTo, err := time.Parse(time.DateOnly, *bodyData.IllDateTo)
		if err != nil {
			return err
		}
		query = query.Where("EXISTS( SELECT 1 FROM events where events.cow_id = cows.id AND events.event_type_id in (1, 2, 3, 4) AND events.date <= ?)", bdTo.UTC()).Preload("Events")
	}
	fm.SetQuery(query)
	return nil
}
