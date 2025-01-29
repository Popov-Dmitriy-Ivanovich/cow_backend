package load

import (
	"cow_backend/models"
	"encoding/csv"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io"
	"os"
	"strconv"
	"sync"
	"time"
)

const EXT_ASSESSMENT_DATE_COL = "assesment_date"
const EXT_BODY_STRUCTURE_COL = "body_structure"
const EXT_MILK_STRENGTH_COL = "milk_strength"
const EXT_LIMBS_COL = "limbs"
const EXT_UDDER_COL = "udder"
const EXT_SUCRUM_COL = "sacrum"
const EXT_SELECS_COL = "selecs"
const EXT_RATING = "rating"

type exteriorRecord struct {
	AssessmentDate *models.DateOnly
	BodyStructure  *float64
	MilkStrength   *float64
	Limbs          *float64
	Udder          *float64
	Sacrum         *float64
	Rating         float64
	Selecs         uint64
	HeaderIndexes  map[string]int
}

var exteriorDataParsers = map[string]func(*exteriorRecord, []string) error{
	EXT_ASSESSMENT_DATE_COL: func(er *exteriorRecord, rec []string) error {
		dateStr := rec[er.HeaderIndexes[EXT_ASSESSMENT_DATE_COL]]
		if dateStr == "" {
			return nil
		}
		date, err := time.Parse(time.DateOnly, dateStr)
		if err != nil {
			return errors.New("ошибка прасинга даты оценки: " + err.Error())
		}
		er.AssessmentDate = &models.DateOnly{Time: date}
		return nil
	},
	EXT_BODY_STRUCTURE_COL: func(er *exteriorRecord, rec []string) error {
		valStr := rec[er.HeaderIndexes[EXT_BODY_STRUCTURE_COL]]
		if valStr == "" {
			return nil
		}
		val, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			return errors.New("ошибка парсинга body_structure: " + err.Error())
		}
		er.BodyStructure = &val
		return nil
	},
	EXT_MILK_STRENGTH_COL: func(er *exteriorRecord, rec []string) error {
		valStr := rec[er.HeaderIndexes[EXT_MILK_STRENGTH_COL]]
		if valStr == "" {
			return nil
		}
		val, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			return errors.New("ошибка парсинга milk_strength: " + err.Error())
		}
		er.MilkStrength = &val
		return nil
	},
	EXT_LIMBS_COL: func(er *exteriorRecord, rec []string) error {
		valStr := rec[er.HeaderIndexes[EXT_LIMBS_COL]]
		if valStr == "" {
			return nil
		}
		val, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			return errors.New("ошибка парсинга limbs: " + err.Error())
		}
		er.Limbs = &val
		return nil
	},
	EXT_UDDER_COL: func(er *exteriorRecord, rec []string) error {
		valStr := rec[er.HeaderIndexes[EXT_UDDER_COL]]
		if valStr == "" {
			return nil
		}
		val, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			return errors.New("ошибка парсинга udder: " + err.Error())
		}
		er.Udder = &val
		return nil
	},
	EXT_SUCRUM_COL: func(er *exteriorRecord, rec []string) error {
		valStr := rec[er.HeaderIndexes[EXT_SUCRUM_COL]]
		if valStr == "" {
			return nil
		}
		val, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			return errors.New("ошибка парсинга sacrum: " + err.Error())
		}
		er.Sacrum = &val
		return nil
	},
	EXT_SELECS_COL: func(er *exteriorRecord, rec []string) error {
		valStr := rec[er.HeaderIndexes[EXT_SELECS_COL]]
		if valStr == "" {
			return errors.New("нельзя загрузить экстерьер без селекса коровы")
		}
		val, err := strconv.ParseUint(valStr, 10, 64)
		if err != nil {
			return errors.New("ошибка парсинга селекса " + err.Error())
		}
		er.Selecs = val
		return nil
	},
	EXT_RATING: func(er *exteriorRecord, rec []string) error {
		valStr := rec[er.HeaderIndexes[EXT_RATING]]
		if valStr == "" {
			return errors.New("rating - обязательная колонка")
		}
		val, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			return errors.New("ошибка парсинга рейтинга " + err.Error())
		}
		er.Rating = val
		return nil
	},
}

func newExteriorRecord(header []string) (*exteriorRecord, error) {
	exteriorRecord := &exteriorRecord{}
	columns := []string{
		EXT_ASSESSMENT_DATE_COL,
		EXT_BODY_STRUCTURE_COL,
		EXT_MILK_STRENGTH_COL,
		EXT_LIMBS_COL,
		EXT_UDDER_COL,
		EXT_SUCRUM_COL,
		EXT_SELECS_COL,
		EXT_RATING,
	}
	exteriorRecord.HeaderIndexes = make(map[string]int)
	for i, col := range header {
		exteriorRecord.HeaderIndexes[col] = i
	}
	for _, col := range columns {
		if _, ok := exteriorRecord.HeaderIndexes[col]; !ok {
			return nil, errors.New("не найдена колонка " + col)
		}
	}
	return exteriorRecord, nil
}

func (extr *exteriorRecord) FromCsvRecord(rec []string) (CsvToDbLoader, error) {
	for col, parser := range exteriorDataParsers {
		if err := parser(extr, rec); err != nil {
			return nil, errors.New("ошибка парсинга колонки " + col + " значения " + rec[extr.HeaderIndexes[col]] + ": " + err.Error())
		}
	}
	return extr, nil
}

func (extr *exteriorRecord) ToDbModel(db *gorm.DB) (any, error) {

	cow := models.Cow{}
	if err := db.First(&cow, map[string]any{"selecs_number": extr.Selecs}).Error; err != nil {
		return nil, errors.New("Не найдена корова с селексом " + strconv.FormatUint(uint64(extr.Selecs), 10))
	}

	exterior := models.Exterior{}
	if err := db.FirstOrCreate(&exterior, map[string]any{"cow_id": cow.ID}).Error; err != nil {
		return nil, errors.New("ошибка получения экстерьера " + err.Error())
	}

	exterior.MilkStrength = extr.MilkStrength
	exterior.Limbs = extr.Limbs
	exterior.Udder = extr.Udder
	exterior.Sacrum = extr.Sacrum
	exterior.BodyStructure = extr.BodyStructure
	exterior.AssessmentDate = extr.AssessmentDate
	exterior.Rating = extr.Rating
	return exterior, nil
}

func (extr *exteriorRecord) Copy() *exteriorRecord {
	copy := exteriorRecord{}
	copy.HeaderIndexes = extr.HeaderIndexes
	return &copy
}

const EXT_CSV_PATH = "./csv/exterior_data/"

var exterior_unique_index = uint64(0)

func (l *Load) ExteriorData() func(*gin.Context) {
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
		uploadedName := EXT_CSV_PATH + "exterior_data" + strconv.FormatInt(now.Unix(), 16) + "_" + strconv.FormatUint(exterior_unique_index, 16) + ".csv"
		if err := c.SaveUploadedFile(csvField[0], uploadedName); err != nil {
			c.JSON(500, err)
			return
		}
		exterior_unique_index++

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
		recordWithHeader, err := newExteriorRecord(header)
		if err != nil {
			c.JSON(422, err.Error())
			return
		}

		errors := []string{}
		errorsMtx := sync.Mutex{}
		loaderWg := sync.WaitGroup{}
		loadChannel := make(chan loaderData)
		MakeLoadingPool(loadChannel, SaveRecordToDb[models.Exterior])
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
