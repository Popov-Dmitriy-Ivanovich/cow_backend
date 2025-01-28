package cows_filter

import (
	"cow_backend/filters"
	"errors"
)

type ByEbvRegionGeneralValue struct {
}

func (f ByEbvRegionGeneralValue) Apply(fm filters.FilteredModel) error {
	query := fm.GetQuery()
	bodyData, ok := fm.GetFilterParameters()["object"].(CowsFilter)
	if !ok {
		return errors.New("wrong object provided in filter filed object")
	}
	if bodyData.EbvGeneralValueRegionTo != nil && bodyData.EbvGeneralValueRegionFrom != nil {
		query = query.Where("EXISTS (SELECT 1 from grade_regions where cow_id = cows.id AND general_value BETWEEN ? AND ?)",
			*bodyData.EbvGeneralValueRegionFrom,
			*bodyData.EbvGeneralValueRegionTo).Preload("GradeRegion")
	} else if bodyData.EbvGeneralValueRegionTo != nil {
		query = query.Where("EXISTS (SELECT 1 from grade_regions where cow_id = cows.id AND general_value <= ?)", *bodyData.EbvGeneralValueRegionTo).Preload("GradeRegion")
	} else if bodyData.EbvGeneralValueRegionFrom != nil {
		query = query.Where("EXISTS (SELECT 1 from grade_regions where cow_id = cows.id AND general_value >= ?)", *bodyData.EbvGeneralValueRegionFrom).Preload("GradeRegion")
	}
	fm.SetQuery(query)
	return nil
}
