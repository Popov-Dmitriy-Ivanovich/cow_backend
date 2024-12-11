package cows_filter

import (
	"cow_backend/filters"
	"errors"
)

type ByAnyIllneses struct {
}

func (f ByAnyIllneses) Apply(fm filters.FilteredModel) error {
	query := fm.GetQuery()
	bodyData, ok := fm.GetFilterParameters()["object"].(CowsFilter)
	if !ok {
		return errors.New("wrong object provided in filter filed object")
	}
	if bodyData.HasAnyIllnes != nil && *bodyData.HasAnyIllnes {
		query = query.Where("EXISTS (SELECT 1 FROM genetics where genetics.cow_id = cows.id AND EXISTS (SELECT 1 FROM genetic_genetic_illnesses WHERE genetic_genetic_illnesses.genetic_id = genetics.id) )").Preload("Genetic").Preload("Genetic.GeneticIllnesses").Preload("Genetic.GeneticIllnesses.Illness").Preload("Genetic.GeneticIllnesses.Status")
	}
	if bodyData.HasAnyIllnes != nil && !*bodyData.HasAnyIllnes {
		query = query.Where("EXISTS (SELECT 1 FROM genetics where genetics.cow_id = cows.id AND NOT EXISTS (SELECT 1 FROM genetic_genetic_illnesses WHERE genetic_genetic_illnesses.genetic_id = genetics.id) )").Preload("Genetic").Preload("Genetic.GeneticIllnesses").Preload("Genetic.GeneticIllnesses.Illness").Preload("Genetic.GeneticIllnesses.Status")
	}
	fm.SetQuery(query)
	return nil
}
