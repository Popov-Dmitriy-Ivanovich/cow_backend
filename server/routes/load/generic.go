package load

import (
	"errors"

	"gorm.io/gorm"
)

type CsvToDbLoader interface {
	FromCsvRecord(rec []string) (CsvToDbLoader, error)
	ToDbModel(tx *gorm.DB) (any, error)
}

func LoadRecordToDb[modelType any](loader CsvToDbLoader, record []string, tx *gorm.DB) error {
	parsed, errLoad := loader.FromCsvRecord(record)
	if errLoad != nil {
		return errLoad
	}
	untypedModel, errParse := parsed.ToDbModel(tx)
	if errParse != nil {
		return errParse
	}
	if typedModel, ok := untypedModel.(modelType); !ok {
		return errors.New("wrong modelType provided to LoadRecordToDb")
	} else {
		if err := tx.Create(&typedModel).Error; err != nil {
			return err
		}
	}
	return nil
}
