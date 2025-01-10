package filters

import "gorm.io/gorm"

type Filter interface {
	Apply(FilteredModel) error
}

type FilteredModel interface {
	GetQuery() *gorm.DB
	GetFilterParameters() map[string]any
	SetQuery(*gorm.DB)
}

type BaseFilteredModel struct {
	Params map[string]any
	Query  *gorm.DB
}

func ApplyFilters(model FilteredModel, filters ...Filter) error {
	for _, filter := range filters {
		if err := filter.Apply(model); err != nil {
			return err
		}
	}
	return nil
}
