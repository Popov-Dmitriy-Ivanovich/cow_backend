package load

import (
	"cow_backend/models"
	"errors"
	"io"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const CM_COW_SELECS_COL = "CowSelecs"
const CM_LACTATION_DATE_COL = "LactationDate"
const CM_CHECK_DATE_COL = "CheckDate"
const MILK_COL = "Milk"
const FAT_COL = "Fat"
const PROTEIN_COL = "Protein"
const PROBE_NUMBER_COL = "ProbeNumber"
const DRY_MATTER_COL = "DryMatter"
const SOMATIC_NUC_COUNT_COL = "SomaticNucCount"

type cmRecord struct {
	DryMatter       *float64
	SomaticNucCount *float64
	ProbeNumber     *uint
	Protein         float64
	Fat             float64
	Milk            float64
	CowSelecs       uint
	LactationDate   models.DateOnly
	CheckDate       models.DateOnly
	HeaderIndexes   map[string]int
}

var cmRecordParsers = map[string]func(*cmRecord, []string) error{
	CM_COW_SELECS_COL: func(cmr *cmRecord, record []string) error {
		selecs := record[cmr.HeaderIndexes[CM_COW_SELECS_COL]]
		selecsUint, err := strconv.ParseUint(selecs, 10, 64)
		if err != nil {
			return err
		}
		cmr.CowSelecs = uint(selecsUint)
		return nil
	},

	CM_LACTATION_DATE_COL: func(cmr *cmRecord, record []string) error {
		dateStr := record[cmr.HeaderIndexes[CM_LACTATION_DATE_COL]]
		date, err := ParseTime(dateStr)
		if err != nil {
			return err
		}
		cmr.LactationDate = models.DateOnly{Time: date}
		return nil
	},

	MILK_COL: func(cmr *cmRecord, record []string) error {
		milkStr := record[cmr.HeaderIndexes[MILK_COL]]
		if milkStr == "" {
			return errors.New("значение milk не может отсутсвовать")
		}
		milk, err := strconv.ParseFloat(milkStr, 64)
		if err != nil {
			return err
		}
		cmr.Milk = milk
		return nil
	},

	FAT_COL: func(cmr *cmRecord, record []string) error {
		fatStr := record[cmr.HeaderIndexes[FAT_COL]]
		if fatStr == "" {
			return errors.New("значение Fat не может отсутсвовать")
		}
		fat, err := strconv.ParseFloat(fatStr, 64)
		if err != nil {
			return err
		}
		cmr.Fat = fat
		return nil
	},

	PROTEIN_COL: func(cmr *cmRecord, record []string) error {
		proteinStr := record[cmr.HeaderIndexes[PROTEIN_COL]]
		if proteinStr == "" {
			return errors.New("значение Protein не может отсутсвовать")
		}
		protein, err := strconv.ParseFloat(proteinStr, 64)
		if err != nil {
			return err
		}
		cmr.Protein = protein
		return nil
	},

	PROBE_NUMBER_COL: func(cmr *cmRecord, record []string) error {
		probeNStr := record[cmr.HeaderIndexes[PROBE_NUMBER_COL]]
		if probeNStr == "" {
			cmr.ProbeNumber = nil
			return nil
		}
		probeN, err := strconv.ParseUint(probeNStr, 10, 64)
		if err != nil {
			return err
		}
		probeNUint := uint(probeN)
		cmr.ProbeNumber = &probeNUint
		return nil
	},

	SOMATIC_NUC_COUNT_COL: func(cmr *cmRecord, record []string) error {
		sncStr := record[cmr.HeaderIndexes[SOMATIC_NUC_COUNT_COL]]
		if sncStr == "" {
			cmr.SomaticNucCount = nil
			return nil
		}
		snc, err := strconv.ParseFloat(sncStr, 64)
		if err != nil {
			return err
		}

		cmr.SomaticNucCount = &snc
		return nil
	},

	DRY_MATTER_COL: func(cmr *cmRecord, record []string) error {
		dmStr := record[cmr.HeaderIndexes[DRY_MATTER_COL]]
		if dmStr == "" {
			cmr.DryMatter = nil
			return nil
		}
		dm, err := strconv.ParseFloat(dmStr, 64)
		if err != nil {
			return err
		}

		cmr.SomaticNucCount = &dm
		return nil
	},
	CM_CHECK_DATE_COL: func(cmr *cmRecord, record []string) error {
		dateStr := record[cmr.HeaderIndexes[CM_CHECK_DATE_COL]]
		date, err := ParseTime(dateStr)
		if err != nil {
			return err
		}
		cmr.CheckDate = models.DateOnly{Time: date}
		return nil
	},
}

func NewCmRecord(header []string) (*cmRecord, error) {
	cmr := cmRecord{}
	columns := []string{
		CM_COW_SELECS_COL,
		CM_LACTATION_DATE_COL,
		MILK_COL,
		FAT_COL,
		PROTEIN_COL,
		PROBE_NUMBER_COL,
		DRY_MATTER_COL,
		SOMATIC_NUC_COUNT_COL,
	}
	cmr.HeaderIndexes = make(map[string]int)
	for idx, col := range header {
		cmr.HeaderIndexes[col] = idx
	}
	for _, col := range columns {
		if _, ok := cmr.HeaderIndexes[col]; !ok {
			return nil, errors.New("не найдена колонка " + col)
		}
	}
	return &cmr, nil
}

func (cmr *cmRecord) FromCsvRecord(rec []string) (CsvToDbLoader, error) {
	for col, parser := range cmRecordParsers {
		if err := parser(cmr, rec); err != nil {
			return nil, errors.New("ошибка парсинга колонки " + col + " значения " + rec[cmr.HeaderIndexes[col]] + ": " + err.Error())
		}
	}
	return cmr, nil
}

func (cmr *cmRecord) ToDbModel(tx *gorm.DB) (any, error) {
	cmCow := models.Cow{}
	cmLac := models.Lactation{}
	db := models.GetDb()

	if err := db.First(&cmCow, map[string]any{"selecs_number": cmr.CowSelecs}).Error; err != nil {
		return nil, errors.New("Не найдена корова с селексом " + strconv.FormatUint(uint64(cmr.CowSelecs), 10))
	}
	if err := db.First(&cmLac, map[string]any{"cow_id": cmCow.ID, "calving_date": cmr.LactationDate}).Error; err != nil {
		return nil, errors.New("Не найдена лактация для коровы с селексом " + strconv.FormatUint(uint64(cmr.CowSelecs), 10) + " и датой отела " + cmr.LactationDate.Format(time.DateOnly))
	}

	newCm := models.CheckMilk{
		LactationId:     cmLac.ID,
		CheckDate:       cmr.CheckDate,
		Milk:            cmr.Milk,
		Fat:             cmr.Fat,
		Protein:         cmr.Protein,
		SomaticNucCount: cmr.SomaticNucCount,
		ProbeNumber:     cmr.ProbeNumber,
		DryMatter:       cmr.DryMatter,
	}
	return newCm, nil
}

func (cr *cmRecord) Copy() *cmRecord {
	copy := cmRecord{}
	copy.HeaderIndexes = cr.HeaderIndexes
	return &copy
}

const CM_CSV_PATH = "./csv/check_milks/"

var cmUniqueIndex uint64 = 0

func (l *Load) CheckMilk() func(*gin.Context) {
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
		uploadedName := CM_CSV_PATH + "check_milk_" + strconv.FormatInt(now.Unix(), 16) + "_" + strconv.FormatUint(cmUniqueIndex, 16) + ".csv"
		if err := c.SaveUploadedFile(csvField[0], uploadedName); err != nil {
			c.JSON(500, err)
			return
		}
		cmUniqueIndex++

		file, err := os.Open(uploadedName)
		if err != nil {
			c.JSON(500, "error opening file")
			return
		}
		defer file.Close()
		csvReader, header, err := GetCsvReader(file)
		if err != nil {
			c.JSON(422, err.Error())
			return
		}
		recordWithHeader, err := NewCmRecord(header)
		if err != nil {
			c.JSON(422, err.Error())
			return
		}

		errors := []string{}
		errorsMtx := sync.Mutex{}
		loaderWg := sync.WaitGroup{}
		loadChannel := make(chan loaderData)
		MakeLoadingPool(loadChannel, LoadRecordToDb[models.CheckMilk])
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
