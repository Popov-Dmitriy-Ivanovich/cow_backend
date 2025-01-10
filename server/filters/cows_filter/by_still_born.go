package cows_filter

import (
	"cow_backend/filters"
	"errors"
)

type ByStillBorn struct {
}

func (f ByStillBorn) Apply(fm filters.FilteredModel) error {
	query := fm.GetQuery()
	bodyData, ok := fm.GetFilterParameters()["object"].(CowsFilter)
	if !ok {
		return errors.New("wrong object provided in filter filed object")
	}
	if bodyData.IsStillBorn != nil && *bodyData.IsStillBorn { // stillborn means, that 0 cows are born
		query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.calving_count = ?)", 0).Preload("Lactation")
	}
	if bodyData.IsStillBorn != nil && !*bodyData.IsStillBorn { // stillborn means, that 0 cows are born
		query = query.Where("NOT EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.calving_count = ?)", 0).Preload("Lactation")
	}
	fm.SetQuery(query)
	return nil
}
