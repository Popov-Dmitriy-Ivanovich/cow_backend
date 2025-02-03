package cows_filter

import (
	"cow_backend/filters"
)

type ByIncludeOnlyFilter struct{}

func (f ByIncludeOnlyFilter) Apply(fm filters.FilteredModel) error {
	query := fm.GetQuery()

	filterData, ok := fm.GetFilterParameters()["object"].(CowsFilter)
	if !ok {
		return nil
	}

	if len(filterData.IncludeOnly) == 0 {
		return nil
	}

	query = query.Where(map[string]any{"cows.id": filterData.IncludeOnly})

	fm.SetQuery(query)
	return nil
}
