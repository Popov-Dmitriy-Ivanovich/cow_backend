package cows_filter

import (
	"cow_backend/filters"
	"cow_backend/models"
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
		hoz := []uint{}
		db := models.GetDb()
		if err := db.Model(&models.Farm{}).Where(map[string]any{
			"parrent_id": bodyData.HozId,
			"type":       []uint{1, 2},
		}).Pluck("id", &hoz).Error; err != nil {
			return err
		}
		hoz = append(hoz, *bodyData.HozId)
		query = query.Where(map[string]any{"farm_group_id": hoz}).Preload("Farm")
	}
	fm.SetQuery(query)
	return nil
}
