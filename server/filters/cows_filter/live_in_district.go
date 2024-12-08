package cows_filter

import (
	"cow_backend/filters"
	"errors"
	"strconv"
)

type LiveInDistrict struct {
}

func (f LiveInDistrict) Apply(fm filters.FilteredModel) error {
	query := fm.GetQuery()

	district, ok := fm.GetFilterParameters()["district"]
	if !ok {
		return nil
	}
	districtStr, ok := district.(string)
	if !ok {
		return errors.New("district id is not passed as string")
	}
	if districtID, err := strconv.ParseUint(districtStr, 10, 64); err == nil {
		query = query.Where("farm_id IS NOT NULL AND EXISTS(SELECT 1 FROM farms WHERE farms.id = cows.farm_id AND farms.district_id = ?) OR "+
			"farm_id is NULL and EXISTS (SELECT 1 FROM farms where farms.id = cows.farm_group_id AND farms.district_id = ?)", districtID, districtID)
	} else {
		return err
	}
	fm.SetQuery(query)
	return nil
}
