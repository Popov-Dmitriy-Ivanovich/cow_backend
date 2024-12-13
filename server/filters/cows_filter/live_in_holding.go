package cows_filter

import (
	"cow_backend/filters"
	"errors"
	"strconv"
)

type LiveInHolding struct {
}

func (f LiveInHolding) Apply(fm filters.FilteredModel) error {
	query := fm.GetQuery()

	hoz, ok := fm.GetFilterParameters()["hoz"]
	if !ok {
		return nil
	}
	hozStr, ok := hoz.(string)
	if !ok {
		return errors.New("region id is not passed as string")
	}
	if hozId, err := strconv.ParseUint(hozStr, 10, 64); err == nil {
		query = query.Where("farm_group_id = ?", hozId)
	} else {
		return err
	}
	fm.SetQuery(query)
	return nil
}
