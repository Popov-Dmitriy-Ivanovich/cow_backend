package cows_filter

import (
	"cow_backend/filters"
	"errors"
	"strconv"
)

type LiveInRegion struct {
}

func (f LiveInRegion) Apply(fm filters.FilteredModel) error {
	query := fm.GetQuery()

	region, ok := fm.GetFilterParameters()["region"]
	if !ok {
		return nil
	}
	regionStr, ok := region.(string)
	if !ok {
		return errors.New("region id is not passed as string")
	}
	if RegionID, err := strconv.ParseUint(regionStr, 10, 64); err == nil {
		query = query.Where(
			"cows.farm_id IS NOT NULL AND "+
				"EXISTS(SELECT 1 FROM farms WHERE farms.id = cows.farm_id AND EXISTS "+
				"(SELECT 1 FROM districts WHERE districts.id = farms.district_id and districts.region_id = ?)) OR "+
				"cows.farm_id IS NULL AND EXISTS (SELECT 1 from farms WHERE farms.id = cows.farm_group_id AND "+
				"EXISTS(SELECT 1 FROM districts WHERE districts.id = farms.district_id AND districts.region_id = ?))",
			RegionID, RegionID)
	} else {
		return err
	}
	fm.SetQuery(query)
	return nil
}
