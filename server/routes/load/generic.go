package load

import (
	"cow_backend/models"
	"errors"
	"sync"

	"gorm.io/gorm"
)

const MAX_CONCURENT_LOADERS = 64

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
	typedModel, ok := untypedModel.(modelType)
	if !ok {
		return errors.New("wrong type provided to load record to db")
	}
	if err := db.Debug().Create(&typedModel).Error; err != nil {
		return err
	}

	return nil
}

func SaveRecordToDb[modelType any](loader CsvToDbLoader, record []string) error {
	parsed, errLoad := loader.FromCsvRecord(record)
	if errLoad != nil {
		return errLoad
	}
	db := models.GetDb()
	untypedModel, errParse := parsed.ToDbModel(db)
	if errParse != nil {
		return errParse
	}
	typedModel, ok := untypedModel.(modelType)
	if !ok {
		return errors.New("wrong type provided to load record to db")
	}
	if err := db.Save(&typedModel).Error; err != nil {
		return err
	}

	return nil
}

type loaderData struct {
	Loader    CsvToDbLoader
	Record    []string
	Errors    *[]string
	ErrorsMtx *sync.Mutex
	WaitGroup *sync.WaitGroup
}

func MakeLoadingPool(ch chan loaderData, loaderFunc func(CsvToDbLoader, []string) error) {
	for i := 0; i < MAX_CONCURENT_LOADERS; i++ {
		go func() {
			for lr := range ch {
				if err := loaderFunc(lr.Loader, lr.Record); err != nil {
					lr.ErrorsMtx.Lock()
					*lr.Errors = append(*lr.Errors, err.Error())
					lr.ErrorsMtx.Unlock()
				}
				lr.WaitGroup.Done()
			}
		}()
	}
}
