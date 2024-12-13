package cows_filter

import (
	"cow_backend/filters"
	"errors"
	"time"
)

type ByCreatedAt struct {
}

func (f ByCreatedAt) Apply(fm filters.FilteredModel) error {
	query := fm.GetQuery()
	bodyData, ok := fm.GetFilterParameters()["object"].(CowsFilter)
	if !ok {
		return errors.New("wrong object provided in filter filed object")
	}
	if bodyData.CreatedAtFrom != nil && bodyData.CreatedAtTo != nil &&
		*bodyData.CreatedAtFrom != "" && *bodyData.CreatedAtTo != "" {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.CreatedAtFrom)
		if err != nil {
			return err
		}
		bdTo, err := time.Parse(time.DateOnly, *bodyData.CreatedAtTo)
		if err != nil {
			return err
		}
		query = query.Where("created_at BETWEEN ? AND ?", bdFrom.UTC(), bdTo.AddDate(0, 0, 1).UTC())
	} else if bodyData.CreatedAtFrom != nil && *bodyData.CreatedAtFrom != "" {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.CreatedAtFrom)
		if err != nil {
			return err
		}
		query = query.Where("created_at >= ?", bdFrom.UTC())
	} else if bodyData.CreatedAtTo != nil && *bodyData.CreatedAtTo != "" {
		bdTo, err := time.Parse(time.DateOnly, *bodyData.CreatedAtTo)
		if err != nil {
			return err
		}
		query = query.Where("created_at <= ?", bdTo.AddDate(0, 0, 1).UTC())
	}
	fm.SetQuery(query)
	return nil
}
