package load

import (
	"cow_backend/models"
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type eventRecord struct {
	CowSelecs         uint
	GroupId           uint
	NameId            uint
	TypeId            uint
	DataResourse      *string
	DaysFromLactation uint
	Date              *models.DateOnly
	Comment1          *string
	Comment2          *string
	HeaderIndexes     map[string]int
}

const EVENT_COW_SELECS_COL = "CowSelecs"
const EVENT_GROUP_ID_COL = "GroupId"
const EVENT_NAME_ID_COL = "NameId"
const EVENT_TYPE_ID_COL = "TypeId"
const EVENT_DATA_RESOURSE_COL = "DataResourse"
const EVENT_DAYS_FROM_LAC_COL = "DaysFromLac"
const EVENT_DATE_COL = "Date"
const EVENT_COM1_COL = "Comment1"
const EVENT_COM2_COL = "Comment2"

var eventRecordParsers = map[string]func(*eventRecord, []string) error{
	EVENT_COW_SELECS_COL: func(evr *eventRecord, rec []string) error {
		selecs := rec[evr.HeaderIndexes[EVENT_COW_SELECS_COL]]
		selecsUint, err := strconv.ParseUint(selecs, 10, 64)
		if err != nil {
			return err
		}
		evr.CowSelecs = uint(selecsUint)
		return nil
	},
	EVENT_GROUP_ID_COL: func(evr *eventRecord, rec []string) error {
		groupId := rec[evr.HeaderIndexes[EVENT_GROUP_ID_COL]]
		groupIdUint, err := strconv.ParseUint(groupId, 10, 64)
		if err != nil {
			return err
		}
		evr.GroupId = uint(groupIdUint)
		return nil
	},
	EVENT_NAME_ID_COL: func(evr *eventRecord, rec []string) error {
		groupId := rec[evr.HeaderIndexes[EVENT_NAME_ID_COL]]
		eventNameId, err := strconv.ParseUint(groupId, 10, 64)
		if err != nil {
			return err
		}
		evr.NameId = uint(eventNameId)
		return nil
	},
	EVENT_TYPE_ID_COL: func(evr *eventRecord, rec []string) error {
		typeId := rec[evr.HeaderIndexes[EVENT_TYPE_ID_COL]]
		eventNameId, err := strconv.ParseUint(typeId, 10, 64)
		if err != nil {
			return err
		}
		evr.TypeId = uint(eventNameId)
		return nil
	},
	EVENT_DATA_RESOURSE_COL: func(evr *eventRecord, rec []string) error {
		dataResourse := rec[evr.HeaderIndexes[EVENT_DATA_RESOURSE_COL]]
		if dataResourse == "" {
			evr.DataResourse = nil
			return nil
		}
		evr.DataResourse = &dataResourse
		return nil
	},
	EVENT_DAYS_FROM_LAC_COL: func(evr *eventRecord, rec []string) error {
		daysStr := rec[evr.HeaderIndexes[EVENT_DAYS_FROM_LAC_COL]]
		days, err := strconv.ParseUint(daysStr, 10, 64)
		if err != nil {
			return err
		}
		evr.DaysFromLactation = uint(days)
		return nil
	},
	EVENT_DATE_COL: func(evr *eventRecord, rec []string) error {
		dateStr := rec[evr.HeaderIndexes[EVENT_DATE_COL]]
		if dateStr == "" {
			return nil
		}
		date, err := time.Parse(time.DateOnly, dateStr)
		if err != nil {
			return err
		}
		evr.Date = &models.DateOnly{Time: date}
		return nil
	},
	EVENT_COM1_COL: func(evr *eventRecord, rec []string) error {
		comment := rec[evr.HeaderIndexes[EVENT_COM1_COL]]
		if comment == "" {
			evr.Comment1 = nil
			return nil
		}
		evr.Comment1 = &comment
		return nil
	},
	EVENT_COM2_COL: func(evr *eventRecord, rec []string) error {
		comment := rec[evr.HeaderIndexes[EVENT_COM2_COL]]
		if comment == "" {
			evr.Comment2 = nil
			return nil
		}
		evr.Comment2 = &comment
		return nil
	},
}

func NewEventRecord(header []string) (*eventRecord, error) {
	evr := eventRecord{}
	columns := []string{
		EVENT_COW_SELECS_COL,
		EVENT_GROUP_ID_COL,
		EVENT_NAME_ID_COL,
		EVENT_TYPE_ID_COL,
		EVENT_DATA_RESOURSE_COL,
		EVENT_DAYS_FROM_LAC_COL,
		EVENT_DATE_COL,
		EVENT_COM1_COL,
		EVENT_COM2_COL,
	}
	evr.HeaderIndexes = make(map[string]int)
	for idx, col := range header {
		evr.HeaderIndexes[col] = idx
	}
	for _, col := range columns {
		if _, ok := evr.HeaderIndexes[col]; !ok {
			return nil, errors.New("не найдена колонка " + col)
		}
	}
	return &evr, nil
}

func (evr *eventRecord) FromCsvRecord(rec []string) (CsvToDbLoader, error) {
	for col, parser := range eventRecordParsers {
		if err := parser(evr, rec); err != nil {
			return nil, errors.New("ошибка парсинга колонки " + col + " значения " + rec[evr.HeaderIndexes[col]] + ": " + err.Error())
		}
	}
	return evr, nil
}

func (evr *eventRecord) ToDbModel(tx *gorm.DB) (any, error) {
	cmCow := models.Cow{}
	type1 := models.EventType{}
	type2 := models.EventType{}
	type3 := models.EventType{}
	db := models.GetDb()

	if err := db.First(&cmCow, map[string]any{"selecs_number": evr.CowSelecs}).Error; err != nil {
		return nil, errors.New("Не найдена корова с селексом " +
			strconv.FormatUint(uint64(evr.CowSelecs), 10))
	}
	if err := db.First(&type1, map[string]any{"type": 1, "code": evr.GroupId}).Error; err != nil {
		return nil, errors.New("Не найдена группа заболеваний с кодом " +
			strconv.FormatUint(uint64(evr.GroupId), 10))
	}
	if err := db.First(&type2, map[string]any{"type": 2, "code": evr.NameId}).Error; err != nil {
		return nil, errors.New("Не найдено заболевание с кодом " +
			strconv.FormatUint(uint64(evr.NameId), 10))
	}
	if err := db.First(&type3, map[string]any{"type": 3, "code": evr.TypeId}).Error; err != nil {
		return nil, errors.New("Не найдена разновидность заболевания с кодом " +
			strconv.FormatUint(uint64(evr.TypeId), 10))
	}
	newCm := models.Event{
		CowId:             cmCow.ID,
		EventType:         type1,
		EventType1:        type2,
		EventType2:        &type3,
		DataResourse:      evr.DataResourse,
		DaysFromLactation: evr.DaysFromLactation,
		Date:              evr.Date,
		Comment1:          evr.Comment1,
		Comment2:          evr.Comment2,
	}
	return newCm, nil
}

func (cr *eventRecord) Copy() *eventRecord {
	copy := eventRecord{}
	copy.HeaderIndexes = cr.HeaderIndexes
	return &copy
}

const EVENT_CSV_PATH = "./csv/events/"

var eventUniqueIndex uint64 = 0

func (l *Load) Event() func(*gin.Context) {
	return func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(500, err)
			return
		}
		csvField, ok := form.File["csv"]
		if !ok {
			c.JSON(500, "not found field csv")
			return
		}

		now := time.Now()
		uploadedName := EVENT_CSV_PATH + "event" + strconv.FormatInt(now.Unix(), 16) + "_" + strconv.FormatUint(eventUniqueIndex, 16) + ".csv"
		if err := c.SaveUploadedFile(csvField[0], uploadedName); err != nil {
			c.JSON(500, err)
			return
		}
		eventUniqueIndex++

		file, err := os.Open(uploadedName)
		if err != nil {
			c.JSON(500, "error opening file")
			return
		}
		defer file.Close()
		csvReader := csv.NewReader(file)
		header, err := csvReader.Read()
		if err != nil {
			c.JSON(422, err.Error())
			return
		}
		recordWithHeader, err := NewEventRecord(header)
		if err != nil {
			c.JSON(422, err.Error())
			return
		}

		errors := []string{}
		errorsMtx := sync.Mutex{}
		loaderWg := sync.WaitGroup{}
		loadChannel := make(chan loaderData)
		MakeLoadingPool(loadChannel, LoadRecordToDb[models.Event])
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		for record, err := csvReader.Read(); err != io.EOF; record, err = csvReader.Read() {
			if err != nil {
				errorsMtx.Lock()
				errors = append(errors, err.Error())
				errorsMtx.Unlock()
				continue
			}
			loaderWg.Add(1)
			loadChannel <- loaderData{
				Loader:    recordWithHeader.Copy(),
				Record:    record,
				Errors:    &errors,
				ErrorsMtx: &errorsMtx,
				WaitGroup: &loaderWg,
			}
		}
		loaderWg.Wait()
		close(loadChannel)
		c.JSON(200, errors)
	}
}
