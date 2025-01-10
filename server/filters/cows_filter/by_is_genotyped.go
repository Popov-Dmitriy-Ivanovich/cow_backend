package cows_filter

import (
	"cow_backend/filters"
	"errors"
)

type ByIsGenotyped struct {
}

func (f ByIsGenotyped) Apply(fm filters.FilteredModel) error {
	query := fm.GetQuery()
	bodyData, ok := fm.GetFilterParameters()["object"].(CowsFilter)
	if !ok {
		return errors.New("wrong object provided in filter filed object")
	}
	if bodyData.IsGenotyped != nil && *bodyData.IsGenotyped {
		query = query.Where("EXISTS (SELECT 1 FROM genetics where genetics.cow_id = cows.id)")
	}
	if bodyData.IsGenotyped != nil && !*bodyData.IsGenotyped {
		query = query.Where("NOT EXISTS (SELECT 1 FROM genetics where genetics.cow_id = cows.id)")
	}
	fm.SetQuery(query)
	return nil
}
