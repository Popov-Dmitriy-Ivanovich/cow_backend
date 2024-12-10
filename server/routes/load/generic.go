package load

import (
	"cow_backend/models"

	"gorm.io/gorm"
)

type CsvToDbLoader interface {
	FromCsvRecord(rec []string) (CsvToDbLoader, error)
	ToDbModel(tx *gorm.DB) (any, error)
}

func LoadRecordToDb[modelType any](loader CsvToDbLoader, record []string) error {
	parsed, errLoad := loader.FromCsvRecord(record)
	if errLoad != nil {
		return errLoad
	}
	db := models.GetDb()
	untypedModel, errParse := parsed.ToDbModel(db)
	if errParse != nil {
		return errParse
	}
	if err := db.Create(&untypedModel).Error; err != nil {
		return err
	}

	return nil
}
