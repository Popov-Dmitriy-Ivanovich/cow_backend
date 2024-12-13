package cows_filter

import (
	"cow_backend/filters"
	"errors"
	"time"
)

type ByControlMilkingDate struct {
}

func (f ByControlMilkingDate) Apply(fm filters.FilteredModel) error {
	query := fm.GetQuery()
	bodyData, ok := fm.GetFilterParameters()["object"].(CowsFilter)
	if !ok {
		return errors.New("wrong object provided in filter filed object")
	}
	if bodyData.ControlMilkingDateFrom != nil && bodyData.ControlMilkingDateTo != nil &&
		*bodyData.ControlMilkingDateFrom != "" && *bodyData.ControlMilkingDateTo != "" {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.ControlMilkingDateFrom)
		if err != nil {
			return err
		}
		bdTo, err := time.Parse(time.DateOnly, *bodyData.ControlMilkingDateTo)
		if err != nil {
			return err
		}
		query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND EXISTS (SELECT 1 FROM check_milks WHERE check_milks.lactation_id = lactations.id AND check_milks.check_date BETWEEN ? AND ?))", bdFrom.UTC(), bdTo.AddDate(0, 0, 1).UTC()).Preload("Lactation").Preload("Lactation.CheckMilks")
	} else if bodyData.ControlMilkingDateFrom != nil && *bodyData.ControlMilkingDateFrom != "" {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.ControlMilkingDateFrom)
		if err != nil {
			return err
		}
		query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND EXISTS (SELECT 1 FROM check_milks WHERE check_milks.lactation_id = lactations.id AND check_milks.check_date >= ?))", bdFrom.UTC()).Preload("Lactation").Preload("Lactation.CheckMilks")
	} else if bodyData.ControlMilkingDateTo != nil && *bodyData.ControlMilkingDateTo != "" {
		bdTo, err := time.Parse(time.DateOnly, *bodyData.ControlMilkingDateTo)
		if err != nil {
			return err
		}
		query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND EXISTS (SELECT 1 FROM check_milks WHERE check_milks.lactation_id = lactations.id AND check_milks.check_date <= ?))", bdTo.AddDate(0, 0, 1).UTC()).Preload("Lactation").Preload("Lactation.CheckMilks")
	}
	fm.SetQuery(query)
	return nil
}
