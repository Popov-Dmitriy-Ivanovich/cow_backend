package cows_filter

import (
	"cow_backend/filters"
	"errors"
)

type ByMonogeneticIllnesses struct {

}

func (f ByMonogeneticIllnesses) Apply(fm filters.FilteredModel) error {
	query := fm.GetQuery()
	bodyData, ok := fm.GetFilterParameters()["object"].(CowsFilter)
	if !ok { return errors.New("wrong object provided in filter filed object")}
	if bodyData.IsIll != nil && *bodyData.IsIll {
		if len(bodyData.MonogeneticIllneses) != 0 {
			query = query.Where("EXISTS (SELECT 1 FROM genetics where genetics.cow_id = cows.id AND EXISTS (SELECT 1 FROM genetic_genetic_illnesses WHERE genetic_genetic_illnesses.genetic_id = genetics.id AND genetic_illness_id IN ?) )", bodyData.MonogeneticIllneses).Preload("Genetic").Preload("Genetic.GeneticIllnesses")
		}
	}
	if bodyData.IsIll != nil && !*bodyData.IsIll {
		if len(bodyData.MonogeneticIllneses) != 0 {
			query = query.Where("EXISTS (SELECT 1 FROM genetics where genetics.cow_id = cows.id AND NOT EXISTS (SELECT 1 FROM genetic_genetic_illnesses WHERE genetic_genetic_illnesses.genetic_id = genetics.id AND genetic_illness_id IN ?) )", bodyData.MonogeneticIllneses).Preload("Genetic").Preload("Genetic.GeneticIllnesses")
		}
	}
	fm.SetQuery(query)
	return nil
}