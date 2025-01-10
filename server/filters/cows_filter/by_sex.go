package cows_filter

import (
	"cow_backend/filters"
	"errors"
)

type BySex struct {
}

func (f BySex) Apply(fm filters.FilteredModel) error {
	query := fm.GetQuery()
	bodyData, ok := fm.GetFilterParameters()["object"].(CowsFilter)
	if !ok {
		return errors.New("wrong object provided in filter filed object")
	}
	if len(bodyData.Sex) != 0 {
		query = query.Where("sex_id IN ?", bodyData.Sex).Preload("Sex")
	}
	fm.SetQuery(query)
	return nil
}
