package cows_filter

import (
	"cow_backend/filters"
	"errors"
)

type ByExterior struct {
}

func (f ByExterior) Apply(fm filters.FilteredModel) error {
	query := fm.GetQuery()
	bodyData, ok := fm.GetFilterParameters()["object"].(CowsFilter)
	if !ok {
		return errors.New("wrong object provided in filter filed object")
	}
	if bodyData.ExteriorFrom != nil && bodyData.ExteriorTo != nil {
		query = query.Where("EXISTS(SELECT 1 FROM exteriors WHERE exteriors.cow_id = cows.id AND exteriors.rating BETWEEN ? AND ?)", bodyData.ExteriorFrom, bodyData.ExteriorTo).Preload("Exterior")
	} else if bodyData.ExteriorFrom != nil {
		query = query.Where("EXISTS(SELECT 1 FROM exteriors WHERE exteriors.cow_id = cows.id AND exteriors.rating >= ?)", bodyData.ExteriorFrom).Preload("Exterior")
	} else if bodyData.ExteriorTo != nil {
		query = query.Where("EXISTS(SELECT 1 FROM exteriors WHERE exteriors.cow_id = cows.id AND exteriors.rating <= ?)", bodyData.ExteriorTo).Preload("Exterior")
	}
	fm.SetQuery(query)
	return nil
}
