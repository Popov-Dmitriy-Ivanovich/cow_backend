package cows_filter

import (
	"cow_backend/filters"
	"errors"
)

type ByHoz struct {
}

func (f ByHoz) Apply(fm filters.FilteredModel) error {
	query := fm.GetQuery()
	bodyData, ok := fm.GetFilterParameters()["object"].(CowsFilter)
	if !ok {
		return errors.New("wrong object provided in filter filed object")
	}
	if bodyData.HozId != nil {
		query = query.Where("farm_group_id = ?", bodyData.HozId).Preload("Farm")
	}
	fm.SetQuery(query)
	return nil
}
