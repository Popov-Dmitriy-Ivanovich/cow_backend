package cows_filter

import (
	"cow_backend/filters"
	"errors"
	"strings"
)

type BySearchString struct {
}

func (f BySearchString) Apply(fm filters.FilteredModel) error {
	query := fm.GetQuery()
	bodyData, ok := fm.GetFilterParameters()["object"].(CowsFilter)
	if !ok {
		return errors.New("wrong object provided in filter filed object")
	}
	if searchString := bodyData.SearchQuery; searchString != nil && *searchString != "" {
		*searchString = "%" + *searchString + "%"
		*searchString = strings.ToLower(*searchString)
		query = query.Where("LOWER(name) LIKE ? or LOWER(rshn_number) LIKE ? or LOWER(inventory_number) LIKE ? or LOWER(CAST(selecs_number AS TEXT)) like ?", searchString, searchString, searchString, searchString)
	}
	fm.SetQuery(query)
	return nil
}
