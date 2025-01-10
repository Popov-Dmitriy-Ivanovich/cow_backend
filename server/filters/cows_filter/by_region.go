package cows_filter

import (
	"cow_backend/filters"
	"errors"
)

type ByRegion struct {
}

func (f ByRegion) Apply(fm filters.FilteredModel) error {
	query := fm.GetQuery()
	bodyData, ok := fm.GetFilterParameters()["object"].(CowsFilter)
	if !ok {
		return errors.New("wrong object provided in filter filed object")
	}

	if bodyData.RegionId == nil {
		return nil
	}

	query = query.Where("EXIST (SELECT 1 FROM farms WHERE farms.id = cows.farm_id OR farms.id = cows.farm_group_id AND "+
		"EXISTS(SELECT 1 FROM districts where districts.id = farms.district_id AND districts.region_id = ?))", bodyData.RegionId)
	fm.SetQuery(query)
	return nil
}
