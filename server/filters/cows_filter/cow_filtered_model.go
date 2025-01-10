package cows_filter

import (
	"cow_backend/filters"

	"gorm.io/gorm"
)

type CowFilteredModel struct {
	filters.BaseFilteredModel
}

func (cfm *CowFilteredModel) GetQuery() *gorm.DB {
	return cfm.BaseFilteredModel.Query
}
func (cfm *CowFilteredModel) GetFilterParameters() map[string]any {
	return cfm.BaseFilteredModel.Params
}
func (cfm *CowFilteredModel) SetQuery(q *gorm.DB) {
	cfm.BaseFilteredModel.Query = q
}

func NewCowFilteredModel(object CowsFilter, q *gorm.DB) *CowFilteredModel {
	cfm := CowFilteredModel{
		BaseFilteredModel: filters.BaseFilteredModel{
			Params: map[string]any{"object": object},
			Query:  q,
		},
	}
	return &cfm
}
