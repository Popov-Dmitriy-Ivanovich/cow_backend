package cows_filter

import (
	"cow_backend/filters"
	"errors"
	"gorm.io/gorm"
)

type OrderBy struct {
}

var orderingsDesc = map[string]string{
	"RSHN":             "rshn_number desc",
	"InventoryNumber":  "inventory_number desc",
	"Name":             "name desc",
	"BirthDate":        "birth_date desc",
	"GeneralEbvRegion": "grade_regions.general_value desc",
}
var orderingsAsc = map[string]string{
	"RSHN":             "rshn_number asc",
	"InventoryNumber":  "inventory_number asc",
	"Name":             "name asc",
	"BirthDate":        "birth_date asc",
	"GeneralEbvRegion": "grade_regions.general_value asc",
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
		query = query.Preload("GradeRegion", func(db *gorm.DB) *gorm.DB {
			return db.Order(orderStr)
		})
	}

	fm.SetQuery(query)
	return nil
}
