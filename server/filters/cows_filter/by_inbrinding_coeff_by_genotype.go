package cows_filter

import (
	"cow_backend/filters"
	"errors"
)

type ByInbrindingCoeffByGenotype struct {
}

func (f ByInbrindingCoeffByGenotype) Apply(fm filters.FilteredModel) error {
	query := fm.GetQuery()
	bodyData, ok := fm.GetFilterParameters()["object"].(CowsFilter)
	if !ok {
		return errors.New("wrong object provided in filter filed object")
	}
	if bodyData.InbrindingCoeffByGenotypeFrom != nil && bodyData.InbrindingCoeffByGenotypeTo != nil {
		query = query.Where("EXISTS( SELECT 1 FROM genetics where genetics.cow_id = cows.id AND genetics.inbrinding_coeff_by_genotype BETWEEN ? AND ?)",
			bodyData.InbrindingCoeffByGenotypeFrom,
			bodyData.InbrindingCoeffByGenotypeTo).Preload("Genetic")
	} else if bodyData.InbrindingCoeffByGenotypeFrom != nil {
		query = query.Where("EXISTS( SELECT 1 FROM genetics where genetics.cow_id = cows.id AND genetics.inbrinding_coeff_by_genotype >= ?)",
			bodyData.InbrindingCoeffByGenotypeFrom).Preload("Genetic")
	} else if bodyData.InbrindingCoeffByGenotypeTo != nil {
		query = query.Where("EXISTS( SELECT 1 FROM genetics where genetics.cow_id = cows.id AND genetics.inbrinding_coeff_by_genotype <= ?)",
			bodyData.InbrindingCoeffByGenotypeTo).Preload("Genetic")
	}
	fm.SetQuery(query)
	return nil
}
