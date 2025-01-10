package cows_filter

import (
	"cow_backend/filters"
	"errors"
)

type ByInbrindingCoeffByFamily struct {
}

func (f ByInbrindingCoeffByFamily) Apply(fm filters.FilteredModel) error {
	query := fm.GetQuery()
	bodyData, ok := fm.GetFilterParameters()["object"].(CowsFilter)
	if !ok {
		return errors.New("wrong object provided in filter filed object")
	}
	if bodyData.InbrindingCoeffByFamilyFrom != nil && bodyData.InbrindingCoeffByFamilyTo != nil {
		query = query.Where("inbrinding_coeff_by_family BETWEEN ? AND ?", bodyData.InbrindingCoeffByFamilyFrom, bodyData.InbrindingCoeffByFamilyTo)
	} else if bodyData.InbrindingCoeffByFamilyFrom != nil {
		query = query.Where("inbrinding_coeff_by_family >= ?", bodyData.InbrindingCoeffByFamilyFrom)
	} else if bodyData.InbrindingCoeffByFamilyTo != nil {
		query = query.Where("inbrinding_coeff_by_family <= ?", bodyData.InbrindingCoeffByFamilyTo)
	}
	fm.SetQuery(query)
	return nil
}
