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
	
	holding, ok := fm.GetFilterParameters()["holding"]
	if !ok {
		return nil
	}
	holdingStr, ok := holding.(string); 
	if !ok {
		return errors.New("region id is not passed as string")
	}
	if holdingID, err := strconv.ParseUint(holdingStr,10,64); err == nil {
		query = query.Where("holding_id = ?", holdingID)
	} else {
		return err
	}
	fm.SetQuery(query)
	return nil
}