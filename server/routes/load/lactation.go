package load

import (
	"cow_backend/models"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const LACTATION_CSV_PATH = "./csv/lactations/"

var lactationUniqueIndex uint64 = 0

type lactationRecord struct {
	CowSelecs        uint
	Number           uint
	InsemenationNum  int
	InsemenationDate models.DateOnly
	CalvingCount     int
	CalvingDate      models.DateOnly
	Abort            bool
	MilkAll          *float64
	FatAll           *float64
	ProteinAll       *float64
	Milk305          *float64
	Fat305           *float64
	Protein305       *float64
	Days             *int
	ServicePeriod    *uint
	HeaderIndexes    map[string]int
}

const COW_SELECS_COL = "CowSelecs"
const LAC_NUMBER_COL = "Number"
const INSEMENATION_NUMBER_COL = "InsemenationNum"
const INSEMENATION_DATE_COL = "InsemenationDate"
const CALVING_COUNT_COL = "CalvingCount"
const CALVING_DATE_COL = "CalvingDate"
const ABORT_COL = "Abort"
const MILK_ALL_COL = "MilkAll"
const MILK_305_COL = "Milk305"
const FAT_ALL_COL = "FatAll"
const FAT_305_COL = "Fat305"
const PROTEIN_ALL_COL = "ProteinAll"
const PROTEIN_305_COL = "Protein305"
const DAYS_COL = "Days"
const SERVICE_PERIOD_COL = "ServicePeriod"

var lactationRecordParsers = map[string]func(*lactationRecord, []string) error{
	COW_SELECS_COL: func(lr *lactationRecord, record []string) error {
		selecs := record[lr.HeaderIndexes[COW_SELECS_COL]]
		selecsUint, err := strconv.ParseUint(selecs, 10, 64)
		if err != nil {
			return err
		}
		lr.CowSelecs = uint(selecsUint)
		return nil
	},

	LAC_NUMBER_COL: func(lr *lactationRecord, record []string) error {
		number := record[lr.HeaderIndexes[LAC_NUMBER_COL]]
		numberUint, err := strconv.ParseUint(number, 10, 64)
		if err != nil {
			return err
		}
		lr.Number = uint(numberUint)
		return nil
	},

	INSEMENATION_NUMBER_COL: func(lr *lactationRecord, record []string) error {
		number := record[lr.HeaderIndexes[INSEMENATION_NUMBER_COL]]
		numberInt, err := strconv.ParseInt(number, 10, 64)
		if err != nil {
			return err
		}
		lr.InsemenationNum = int(numberInt)
		return nil
	},

	INSEMENATION_DATE_COL: func(lr *lactationRecord, record []string) error {
		dateStr := record[lr.HeaderIndexes[INSEMENATION_DATE_COL]]
		date, err := ParseTime(dateStr)
		if err != nil {
			return err
		}
		lr.InsemenationDate = models.DateOnly{Time: date}
		return nil
	},

	CALVING_COUNT_COL: func(lr *lactationRecord, record []string) error {
		countStr := record[lr.HeaderIndexes[CALVING_COUNT_COL]]
		if countStr == "Мертвый теленок" {
			lr.CalvingCount = 0
		} else if countStr == "бычок и телочка" {
			lr.CalvingCount = 2
		} else {
			lr.CalvingCount = 1
		}
		return nil
	},

	CALVING_DATE_COL: func(lr *lactationRecord, record []string) error {
		dateStr := record[lr.HeaderIndexes[CALVING_DATE_COL]]
		date, err := ParseTime(dateStr)
		if err != nil {
			return err
		}
		lr.CalvingDate = models.DateOnly{Time: date}
		return nil
	},

	ABORT_COL: func(lr *lactationRecord, record []string) error {
		abortStr := record[lr.HeaderIndexes[ABORT_COL]]
		abortInt, err := strconv.ParseInt(abortStr, 10, 64)
		if err != nil {
			return err
		}
		lr.Abort = abortInt == 1
		return nil
	},

	MILK_ALL_COL: func(lr *lactationRecord, record []string) error {
		milkStr := record[lr.HeaderIndexes[MILK_ALL_COL]]
		if milkStr == "" {
			lr.MilkAll = nil
			return nil
		}
		milk, err := strconv.ParseFloat(milkStr, 64)
		if err != nil {
			return err
		}
		lr.MilkAll = &milk
		return nil
	},

	MILK_305_COL: func(lr *lactationRecord, record []string) error {
		milkStr := record[lr.HeaderIndexes[MILK_305_COL]]
		if milkStr == "" {
			lr.Milk305 = nil
			return nil
		}
		milk, err := strconv.ParseFloat(milkStr, 64)
		if err != nil {
			return err
		}
		lr.Milk305 = &milk
		return nil
	},

	FAT_ALL_COL: func(lr *lactationRecord, record []string) error {
		fatStr := record[lr.HeaderIndexes[FAT_ALL_COL]]
		if fatStr == "" {
			lr.FatAll = nil
			return nil
		}
		fat, err := strconv.ParseFloat(fatStr, 64)
		if err != nil {
			return err
		}
		lr.FatAll = &fat
		return nil
	},

	FAT_305_COL: func(lr *lactationRecord, record []string) error {
		fatStr := record[lr.HeaderIndexes[FAT_305_COL]]
		if fatStr == "" {
			lr.Fat305 = nil
			return nil
		}
		fat, err := strconv.ParseFloat(fatStr, 64)
		if err != nil {
			return err
		}
		lr.Fat305 = &fat
		return nil
	},

	PROTEIN_ALL_COL: func(lr *lactationRecord, record []string) error {
		proteinStr := record[lr.HeaderIndexes[PROTEIN_ALL_COL]]
		if proteinStr == "" {
			lr.ProteinAll = nil
			return nil
		}
		protein, err := strconv.ParseFloat(proteinStr, 64)
		if err != nil {
			return err
		}
		lr.ProteinAll = &protein
		return nil
	},

	PROTEIN_305_COL: func(lr *lactationRecord, record []string) error {
		proteinStr := record[lr.HeaderIndexes[PROTEIN_ALL_COL]]
		if proteinStr == "" {
			lr.Protein305 = nil
			return nil
		}
		protein, err := strconv.ParseFloat(proteinStr, 64)
		if err != nil {
			return err
		}
		lr.Protein305 = &protein
		return nil
	},

	DAYS_COL: func(lr *lactationRecord, record []string) error {
		daysStr := record[lr.HeaderIndexes[DAYS_COL]]
		if daysStr == "" {
			lr.Days = nil
			return nil
		}
		days, err := strconv.ParseInt(daysStr, 10, 64)
		if err != nil {
			return err
		}
		daysInt := int(days)
		lr.Days = &daysInt
		return nil
	},

	SERVICE_PERIOD_COL: func(lr *lactationRecord, record []string) error {
		servStr := record[lr.HeaderIndexes[SERVICE_PERIOD_COL]]
		if servStr == "" {
			lr.ServicePeriod = nil
			return nil
		}
		serv, err := strconv.ParseUint(servStr, 10, 64)
		if err != nil {
			return err
		}
		suint := uint(serv)
		lr.ServicePeriod = &suint
		return nil
	},
}

func NewLactationRecord(header []string) (*lactationRecord, error) {
	lr := lactationRecord{}
	columns := []string{
		COW_SELECS_COL,
		LAC_NUMBER_COL,
		INSEMENATION_NUMBER_COL,
		INSEMENATION_DATE_COL,
		CALVING_COUNT_COL,
		CALVING_DATE_COL,
		ABORT_COL,
		MILK_ALL_COL,
		MILK_305_COL,
		FAT_ALL_COL,
		FAT_305_COL,
		PROTEIN_ALL_COL,
		PROTEIN_305_COL,
		DAYS_COL,
		SERVICE_PERIOD_COL,
	}
	lr.HeaderIndexes = make(map[string]int)
	for idx, col := range header {
		lr.HeaderIndexes[col] = idx
	}
	for _, col := range columns {
		if _, ok := lr.HeaderIndexes[col]; !ok {
			return nil, errors.New("не найдена колонка " + col)
		}
	}
	return &lr, nil
}

func (lr *lactationRecord) FromCsvRecord(rec []string) (CsvToDbLoader, error) {
	for col, parser := range lactationRecordParsers {
		if err := parser(lr, rec); err != nil {
			return nil, errors.New("ошибка парсинга колонки " + col + " значения " + rec[lr.HeaderIndexes[col]] + ": " + err.Error())
		}
	}
	return lr, nil
}

func (lr *lactationRecord) ToDbModel(tx *gorm.DB) (any, error) {
	//return nil, errors.New("преобразование в модель БД отключено")
	cow := models.Cow{}
	db := models.GetDb()
	if err := db.First(&cow, map[string]any{"selecs_number": lr.CowSelecs}).Error; err != nil {
		log.Printf("ошибка поиска коровы: %q", err.Error())
		return nil, errors.New("Не найдена корова с селексом " + strconv.FormatUint(uint64(lr.CowSelecs), 10))
	}
	lactationCount := int64(0)
	if err := db.Model(&models.Lactation{}).Where(map[string]any{"cow_id": cow.ID, "number": lr.Number}).Count(&lactationCount).Error; err != nil {
		log.Printf("Произошла ошибка при поиске лактаций: %q", err.Error())
		return nil, err
	}
	if lactationCount != 0 {
		return nil, errors.New("Лактация с номером " + strconv.FormatUint(uint64(lr.Number), 10) + " коровы с селексом " + strconv.FormatUint(uint64(lr.CowSelecs), 10) + " уже существует")
	}

	lac := models.Lactation{
		CowId:            cow.ID,
		Number:           lr.Number,
		InsemenationNum:  lr.InsemenationNum,
		InsemenationDate: lr.InsemenationDate,
		CalvingCount:     lr.CalvingCount,
		CalvingDate:      lr.CalvingDate,
		ServicePeriod:    lr.ServicePeriod,
		Abort:            lr.Abort,
		MilkAll:          lr.MilkAll,
		Milk305:          lr.Milk305,
		FatAll:           lr.FatAll,
		Fat305:           lr.Fat305,
		ProteinAll:       lr.ProteinAll,
		Protein305:       lr.Protein305,
		Days:             lr.Days, // количество дней, когда корова дает молоко
	}

	return lac, nil
}

func (cr *lactationRecord) Copy() *lactationRecord {
	copy := lactationRecord{}
	copy.HeaderIndexes = cr.HeaderIndexes
	return &copy
}

func (l *Load) Lactation() func(*gin.Context) {
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
		uploadedName := LACTATION_CSV_PATH + "lactation_" + strconv.FormatInt(now.Unix(), 16) + "_" + strconv.FormatUint(lactationUniqueIndex, 16) + ".csv"
		if err := c.SaveUploadedFile(csvField[0], uploadedName); err != nil {
			c.JSON(500, err)
			return
		}
		lactationUniqueIndex++
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
		recordWithHeader, err := NewLactationRecord(header)
		if err != nil {
			c.JSON(422, err.Error())
			return
		}
		errors := []string{}
		errorsMtx := sync.Mutex{}
		loaderWg := sync.WaitGroup{}
		loadChannel := make(chan loaderData)
		MakeLoadingPool(loadChannel, LoadRecordToDb[models.Lactation])
		log.Printf("[INFO] START PARSING CSV FILE")
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		for record, err := csvReader.Read(); err != io.EOF; record, err = csvReader.Read() {
			if err != nil {
				errorsMtx.Lock()
				errors = append(errors, err.Error())
				log.Printf("[ERROR] PARSING FILE %q", err.Error())
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
		log.Printf("[INFO] LOADED ALL DATA FROM CSV TO PROCESSING CHANNEL")
		loaderWg.Wait()
		close(loadChannel)
		c.JSON(200, errors)
	}
}
