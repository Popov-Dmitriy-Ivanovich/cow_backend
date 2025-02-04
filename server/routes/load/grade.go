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

const GR_COW_SELECS_COL = "CowSelecs"

const GR_GENERAL_VALUE_HOZ_COL = "GeneralValueHoz"
const GR_EBV_MILK_HOZ_COL = "ebv_milkHoz"
const GR_EBV_FAT_HOZ_COL = "evb_fatHoz"
const GR_EBV_PROTEIN_HOZ_COL = "ebv_proteinHoz"
const GR_EBV_INSEMENATION_HOZ_COL = "ebv_insemenationHoz"
const GR_EBV_SERVICE_HOZ_COL = "ebv_serviceHoz"

const GR_GENERAL_VALUE_REG_COL = "GeneralValueReg"
const GR_EBV_MILK_REG_COL = "ebv_milkReg"
const GR_EBV_FAT_REG_COL = "evb_fatReg"
const GR_EBV_PROTEIN_REG_COL = "ebv_proteinReg"
const GR_EBV_INSEMENATION_REG_COL = "ebv_insemenationReg"
const GR_EBV_SERVICE_REG_COL = "ebv_serviceReg"

type gradeRecord struct {
	CowSelecs uint

	EbvMilkHoz *float64
	EbvMilkReg *float64

	EbvFatHoz *float64
	EbvFatReg *float64

	EbvProteinHoz *float64
	EbvProteinReg *float64

	EbvInsemenationHoz *float64
	EbvInsemenationReg *float64

	EbvServiceHoz *float64
	EbvServiceReg *float64

	GeneralValueHoz *float64
	GeneralValueReg *float64
	HeaderIndexes   map[string]int
}

var gradeRecordParsers = map[string]func(*gradeRecord, []string) error{
	GR_COW_SELECS_COL: func(gr *gradeRecord, rec []string) error {
		if rec[gr.HeaderIndexes[GR_COW_SELECS_COL]] == "" {
			return errors.New("колонка CowSelecs не может быть пустой")
		}
		selecs, err := strconv.ParseUint(rec[gr.HeaderIndexes[GR_COW_SELECS_COL]], 10, 64)
		if err != nil {
			return err
		}
		gr.CowSelecs = uint(selecs)
		return nil
	},
	GR_GENERAL_VALUE_HOZ_COL: func(gr *gradeRecord, rec []string) error {
		generalValue := rec[gr.HeaderIndexes[GR_GENERAL_VALUE_HOZ_COL]]
		if generalValue == "" {
			gr.GeneralValueHoz = nil
			return nil
		}
		generalValueFloat, err := strconv.ParseFloat(generalValue, 64)
		if err != nil {
			return errors.New("не удалось распарсить general value hoz = " + generalValue + " " + err.Error())
		}
		gr.GeneralValueHoz = &generalValueFloat
		return nil
	},
	GR_GENERAL_VALUE_REG_COL: func(gr *gradeRecord, rec []string) error {
		generalValue := rec[gr.HeaderIndexes[GR_GENERAL_VALUE_REG_COL]]
		if generalValue == "" {
			gr.GeneralValueReg = nil
			return nil
		}
		generalValueFloat, err := strconv.ParseFloat(generalValue, 64)
		if err != nil {
			return errors.New("не удалось распарсить general value hoz = " + generalValue + " " + err.Error())
		}
		gr.GeneralValueReg = &generalValueFloat
		return nil
	},

	GR_EBV_FAT_HOZ_COL: func(gr *gradeRecord, rec []string) error {
		ebvValue := rec[gr.HeaderIndexes[GR_EBV_FAT_HOZ_COL]]
		if ebvValue == "" {
			gr.EbvFatHoz = nil
			return nil
		}
		ebvFloat, err := strconv.ParseFloat(ebvValue, 64)
		if err != nil {
			return errors.New("не удалось распарсить general value hoz = " + ebvValue + " " + err.Error())
		}
		gr.EbvFatHoz = &ebvFloat
		return nil
	},
	GR_EBV_FAT_REG_COL: func(gr *gradeRecord, rec []string) error {
		ebvValue := rec[gr.HeaderIndexes[GR_EBV_FAT_REG_COL]]
		if ebvValue == "" {
			gr.EbvFatReg = nil
			return nil
		}
		ebvFloat, err := strconv.ParseFloat(ebvValue, 64)
		if err != nil {
			return errors.New("не удалось распарсить general value hoz = " + ebvValue + " " + err.Error())
		}
		gr.EbvFatReg = &ebvFloat
		return nil
	},

	GR_EBV_MILK_HOZ_COL: func(gr *gradeRecord, rec []string) error {
		ebvValue := rec[gr.HeaderIndexes[GR_EBV_MILK_HOZ_COL]]
		if ebvValue == "" {
			gr.EbvMilkHoz = nil
			return nil
		}
		ebvFloat, err := strconv.ParseFloat(ebvValue, 64)
		if err != nil {
			return errors.New("не удалось распарсить general value hoz = " + ebvValue + " " + err.Error())
		}
		gr.EbvMilkHoz = &ebvFloat
		return nil
	},
	GR_EBV_MILK_REG_COL: func(gr *gradeRecord, rec []string) error {
		ebvValue := rec[gr.HeaderIndexes[GR_EBV_MILK_REG_COL]]
		if ebvValue == "" {
			gr.EbvMilkReg = nil
			return nil
		}
		ebvFloat, err := strconv.ParseFloat(ebvValue, 64)
		if err != nil {
			return errors.New("не удалось распарсить general value hoz = " + ebvValue + " " + err.Error())
		}
		gr.EbvMilkReg = &ebvFloat
		return nil
	},

	GR_EBV_PROTEIN_HOZ_COL: func(gr *gradeRecord, rec []string) error {
		ebvValue := rec[gr.HeaderIndexes[GR_EBV_PROTEIN_HOZ_COL]]
		if ebvValue == "" {
			gr.EbvProteinHoz = nil
			return nil
		}
		ebvFloat, err := strconv.ParseFloat(ebvValue, 64)
		if err != nil {
			return errors.New("не удалось распарсить general value hoz = " + ebvValue + " " + err.Error())
		}
		gr.EbvProteinHoz = &ebvFloat
		return nil
	},
	GR_EBV_PROTEIN_REG_COL: func(gr *gradeRecord, rec []string) error {
		ebvValue := rec[gr.HeaderIndexes[GR_EBV_PROTEIN_REG_COL]]
		if ebvValue == "" {
			gr.EbvProteinReg = nil
			return nil
		}
		ebvFloat, err := strconv.ParseFloat(ebvValue, 64)
		if err != nil {
			return errors.New("не удалось распарсить general value hoz = " + ebvValue + " " + err.Error())
		}
		gr.EbvProteinReg = &ebvFloat
		return nil
	},

	GR_EBV_INSEMENATION_HOZ_COL: func(gr *gradeRecord, rec []string) error {
		ebvValue := rec[gr.HeaderIndexes[GR_EBV_PROTEIN_HOZ_COL]]
		if ebvValue == "" {
			gr.EbvInsemenationHoz = nil
			return nil
		}
		ebvFloat, err := strconv.ParseFloat(ebvValue, 64)
		if err != nil {
			return errors.New("не удалось распарсить general value hoz = " + ebvValue + " " + err.Error())
		}
		gr.EbvInsemenationHoz = &ebvFloat
		return nil
	},
	GR_EBV_INSEMENATION_REG_COL: func(gr *gradeRecord, rec []string) error {
		ebvValue := rec[gr.HeaderIndexes[GR_EBV_PROTEIN_REG_COL]]
		if ebvValue == "" {
			gr.EbvInsemenationReg = nil
			return nil
		}
		ebvFloat, err := strconv.ParseFloat(ebvValue, 64)
		if err != nil {
			return errors.New("не удалось распарсить general value hoz = " + ebvValue + " " + err.Error())
		}
		gr.EbvInsemenationReg = &ebvFloat
		return nil
	},

	GR_EBV_SERVICE_HOZ_COL: func(gr *gradeRecord, rec []string) error {
		ebvValue := rec[gr.HeaderIndexes[GR_EBV_SERVICE_HOZ_COL]]
		if ebvValue == "" {
			gr.EbvServiceHoz = nil
			return nil
		}
		ebvFloat, err := strconv.ParseFloat(ebvValue, 64)
		if err != nil {
			return errors.New("не удалось распарсить general value hoz = " + ebvValue + " " + err.Error())
		}
		gr.EbvServiceHoz = &ebvFloat
		return nil
	},
	GR_EBV_SERVICE_REG_COL: func(gr *gradeRecord, rec []string) error {
		ebvValue := rec[gr.HeaderIndexes[GR_EBV_SERVICE_REG_COL]]
		if ebvValue == "" {
			gr.EbvServiceReg = nil
			return nil
		}
		ebvFloat, err := strconv.ParseFloat(ebvValue, 64)
		if err != nil {
			return errors.New("не удалось распарсить general value hoz = " + ebvValue + " " + err.Error())
		}
		gr.EbvServiceReg = &ebvFloat
		return nil
	},
}

func NewGradeRecord(header []string) (*gradeRecord, error) {
	gr := gradeRecord{}
	columns := []string{
		GR_COW_SELECS_COL,

		GR_GENERAL_VALUE_HOZ_COL,
		GR_EBV_MILK_HOZ_COL,
		GR_EBV_FAT_HOZ_COL,
		GR_EBV_PROTEIN_HOZ_COL,
		GR_EBV_INSEMENATION_HOZ_COL,
		GR_EBV_SERVICE_HOZ_COL,

		GR_GENERAL_VALUE_REG_COL,
		GR_EBV_MILK_REG_COL,
		GR_EBV_FAT_REG_COL,
		GR_EBV_PROTEIN_REG_COL,
		GR_EBV_INSEMENATION_REG_COL,
		GR_EBV_SERVICE_REG_COL,
	}
	gr.HeaderIndexes = make(map[string]int)
	for idx, col := range header {
		gr.HeaderIndexes[col] = idx
	}
	for _, col := range columns {
		if _, ok := gr.HeaderIndexes[col]; !ok {
			return nil, errors.New("не найдена колонка " + col)
		}
	}
	return &gr, nil
}

func (gr *gradeRecord) FromCsvRecord(rec []string) (CsvToDbLoader, error) {
	for col, parser := range gradeRecordParsers {
		if err := parser(gr, rec); err != nil {
			return nil, errors.New("ошибка парсинга колонки " + col + " значения " + rec[gr.HeaderIndexes[col]] + ": " + err.Error())
		}
	}
	return gr, nil
}

func (gr *gradeRecord) ToDbModel(tx *gorm.DB) (any, error) {
	grCow := models.Cow{}

	db := models.GetDb()
	if err := db.Preload("GradeRegion").Preload("GradeHoz").First(&grCow, map[string]any{"selecs_number": gr.CowSelecs}).Error; err != nil {
		return nil, errors.New("Не найдена корова с селексом " + strconv.FormatUint(uint64(gr.CowSelecs), 10))
	}

	if grCow.GradeHoz != nil {
		return nil, errors.New("Для коровы с селексом " + strconv.FormatUint(uint64(gr.CowSelecs), 10) + " Уже существует оценка по хозяйству")
	}
	if grCow.GradeRegion != nil {
		return nil, errors.New("Для коровы с селексом " + strconv.FormatUint(uint64(gr.CowSelecs), 10) + " Уже существует оценка по региону")
	}

	grCow.GradeHoz = new(models.Grade)
	grCow.GradeRegion = new(models.Grade)

	grCow.GradeHoz.EbvFat = gr.EbvFatHoz
	grCow.GradeRegion.EbvFat = gr.EbvFatReg

	grCow.GradeHoz.EbvInsemenation = gr.EbvInsemenationHoz
	grCow.GradeRegion.EbvInsemenation = gr.EbvInsemenationReg

	grCow.GradeHoz.EbvMilk = gr.EbvMilkHoz
	grCow.GradeRegion.EbvMilk = gr.EbvMilkReg

	grCow.GradeHoz.EbvProtein = gr.EbvProteinHoz
	grCow.GradeRegion.EbvProtein = gr.EbvProteinReg

	grCow.GradeHoz.EvbService = gr.EbvServiceHoz
	grCow.GradeRegion.EvbService = gr.EbvServiceReg

	grCow.GradeHoz.GeneralValue = gr.GeneralValueHoz
	grCow.GradeRegion.GeneralValue = gr.GeneralValueReg
	return grCow, nil
}

func (gr *gradeRecord) Copy() *gradeRecord {
	copy := gradeRecord{}
	copy.HeaderIndexes = gr.HeaderIndexes
	return &copy
}

const GRADE_CSV_PATH = "./csv/grades/"

var gradeUniqueIndex uint64 = 0

func (l *Load) Grade() func(*gin.Context) {
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
		uploadedName := GRADE_CSV_PATH + "grade_" + strconv.FormatInt(now.Unix(), 16) + "_" + strconv.FormatUint(gradeUniqueIndex, 16) + ".csv"
		if err := c.SaveUploadedFile(csvField[0], uploadedName); err != nil {
			c.JSON(500, err)
			return
		}
		gradeUniqueIndex++

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
		recordWithHeader, err := NewGradeRecord(header)
		if err != nil {
			c.JSON(422, err.Error())
			return
		}

		errors := []string{}
		errorsMtx := sync.Mutex{}
		loaderWg := sync.WaitGroup{}
		loadChannel := make(chan loaderData)
		MakeLoadingPool(loadChannel, SaveRecordToDb[models.Cow])
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
