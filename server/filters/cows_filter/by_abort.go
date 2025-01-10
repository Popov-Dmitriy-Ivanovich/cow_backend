package cows_filter

import (
	"cow_backend/filters"
	"errors"
)

type ByAbort struct {
}

func (f ByAbort) Apply(fm filters.FilteredModel) error {
	query := fm.GetQuery()
	bodyData, ok := fm.GetFilterParameters()["object"].(CowsFilter)
	if !ok {
		return errors.New("wrong object provided in filter filed object")
	}
	if bodyData.IsAborted != nil && *bodyData.IsAborted { // abort is marked by flag for some reason
		query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.abort = ?)", true).Preload("Lactation")
	}
	if bodyData.IsAborted != nil && !*bodyData.IsAborted { // abort is marked by flag for some reason
		query = query.Where("NOT EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.abort = ?)", true).Preload("Lactation")
	}
	fm.SetQuery(query)
	return nil
}
