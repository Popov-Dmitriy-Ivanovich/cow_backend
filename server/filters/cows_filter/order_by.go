package cows_filter

import (
	"cow_backend/filters"
	"errors"
)

type OrderBy struct {
}

var orderingsDesc = map[string]string{
	"RSHN":             "rshn_number desc NULLS LAST",
	"InventoryNumber":  "inventory_number desc NULLS LAST",
	"Name":             "name desc NULLS LAST",
	"BirthDate":        "birth_date desc NULLS LAST",
	"GeneralEbvRegion": "\"GradeRegion\".general_value desc NULLS LAST",
}
var orderingsAsc = map[string]string{
	"RSHN":             "rshn_number asc NULLS LAST",
	"InventoryNumber":  "inventory_number asc NULLS LAST",
	"Name":             "name asc NULLS LAST",
	"BirthDate":        "birth_date asc NULLS LAST",
	"GeneralEbvRegion": "\"GradeRegion\".general_value asc NULLS LAST",
}

func (f OrderBy) Apply(fm filters.FilteredModel) error {
	query := fm.GetQuery()

	bodyData, ok := fm.GetFilterParameters()["object"].(CowsFilter)
	if !ok {
		return errors.New("wrong object provided in filter filed object")
	}

	if bodyData.OrderBy != nil && bodyData.OrderByDesc != nil {
		orderStr := ""
		if *bodyData.OrderByDesc {
			orderStr, ok = orderingsDesc[*bodyData.OrderBy]
			if !ok {
				return nil
			}
		} else {
			orderStr, ok = orderingsAsc[*bodyData.OrderBy]
			if !ok {
				return nil
			}
		}
		query = query.Joins("GradeRegion").Order(orderStr)
	}

	fm.SetQuery(query)
	return nil
}
