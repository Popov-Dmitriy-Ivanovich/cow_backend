package cows_filter

import (
	"cow_backend/filters"
	"errors"
)

type ByTwins struct {
}

func (f ByTwins) Apply(fm filters.FilteredModel) error {
	query := fm.GetQuery()
	bodyData, ok := fm.GetFilterParameters()["object"].(CowsFilter)
	if !ok {
		return errors.New("wrong object provided in filter filed object")
	}
	if bodyData.IsTwins != nil && *bodyData.IsTwins { // twins means, that 2 cows are born
		query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.calving_count = ?)", 2).Preload("Lactation")
	}
	if bodyData.IsTwins != nil && !*bodyData.IsTwins { // twins means, that 2 cows are born
		query = query.Where("NOT EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.calving_count = ?)", 2).Preload("Lactation")
	}
	fm.SetQuery(query)
	return nil
}
