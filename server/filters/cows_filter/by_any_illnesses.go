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
		query = query.Where("EXISTS (SELECT 1 FROM genetics where genetics.cow_id = cows.id AND  " +
			"EXISTS (SELECT 1 FROM genetic_illness_data WHERE genetic_illness_data.genetic_id = genetics.id AND (genetic_illness_data.status_id is NULL OR " +
			"EXISTS (SELECT 1 FROM genetic_illness_statuses WHERE genetic_illness_statuses.id = genetic_illness_data.status_id AND genetic_illness_statuses.status <> 'FREE'))))").
			Preload("Genetic").Preload("Genetic.GeneticIllnessesData").Preload("Genetic.GeneticIllnessesData.Illness").
			Preload("Genetic.GeneticIllnessesData.Status")
	}
	if bodyData.HasAnyIllnes != nil && !*bodyData.HasAnyIllnes {
		query = query.Where("EXISTS (SELECT 1 FROM genetics where genetics.cow_id = cows.id AND  " +
			"NOT EXISTS (SELECT 1 FROM genetic_illness_data WHERE genetic_illness_data.genetic_id = genetics.id AND (genetic_illness_data.status_id is NULL OR " +
			"EXISTS (SELECT 1 FROM genetic_illness_statuses WHERE genetic_illness_statuses.id = genetic_illness_data.status_id AND genetic_illness_statuses.status <> 'FREE'))))").
			Preload("Genetic").Preload("Genetic.GeneticIllnessesData").Preload("Genetic.GeneticIllnessesData.Illness").
			Preload("Genetic.GeneticIllnessesData.Status")
	}
	fm.SetQuery(query)
	return nil
}
