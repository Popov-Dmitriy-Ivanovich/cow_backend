package cows_filter

import (
	"cow_backend/filters"
	"errors"
)

type ByBreed struct {
}

func (f ByBreed) Apply(fm filters.FilteredModel) error {
	query := fm.GetQuery()
	bodyData, ok := fm.GetFilterParameters()["object"].(CowsFilter)
	if !ok {
		return errors.New("wrong object provided in filter filed object")
	}
	if len(bodyData.BreedId) != 0 {
		query = query.Where("breed_id in ?", bodyData.BreedId).Preload("Breed")
	}
	fm.SetQuery(query)
	return nil
}
