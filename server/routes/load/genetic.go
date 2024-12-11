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

type geneticRecord struct {
	CowSelecs                 uint
	ProbeNumber               string
	BloodDate                 *models.DateOnly
	ResultDate                *models.DateOnly
	InbrindingCoeffByGenotype *float64
	GeneticIllnesses          []models.GeneticIllnessData
	HeaderIndexes             map[string]int
}

const COW_SELECS_COLUMN = "CowSelecs"
const PROBE_NUMBER_COLUMN = "ProbeNumber"
const BLOOD_DATE_COLUMN = "BloodDate"
const RESULT_DATE_COLUMN = "ResultDate"
const INBRINDING_COEFF_BY_GENOTYPE_COLUMN = "InbrindingCoeffByGenotype"
const HSD_COLUMN = "HCD"
const HH1_COLUMN = "HH1"
const HH3_COLUMN = "HH3"
const HH4_COLUMN = "HH4"
const HH5_COLUMN = "HH5"
const HH6_COLUMN = "HH6"
const BLAD_COLUMN = "BLAD"
const CVM_COLUMN = "CVM"
const DUMPS_COLUMN = "DUMPS"
const BC_COLUMN = "BC"
const FXID_COLUMN = "FXID"
const MF_COLUMN = "MF"
const FGFR2_COLUMN = "FGFR2"
const IH_COLUMN = "IH"

const HSD_DESC = "HCD DESCRIPTION"
const HH1_DESC = "HH1 DESCRIPTION"
const HH3_DESC = "HH3 DESCRIPTION"
const HH4_DESC = "HH4 DESCRIPTION"
const HH5_DESC = "HH5 DESCRIPTION"
const HH6_DESC = "HH6 DESCRIPTION"
const BLAD_DESC = "BLAD DESCRIPTION"
const CVM_DESC = "CVM DESCRIPTION"
const DUMPS_DESC = "DUMPS DESCRIPTION"
const BC_DESC = "BC DESCRIPTION"
const FXID_DESC = "FXID DESCRIPTION"
const MF_DESC = "MF DESCRIPTION"
const FGFR2_DESC = "FGFR2 DESCRIPTION"
const IH_DESC = "IH DESCRIPTION"

var HSD_OMIA = "HCD OMIA"
var HH1_OMIA = "HH1 OMIA"
var HH3_OMIA = "HH3 OMIA"
var HH4_OMIA = "HH4 OMIA"
var HH5_OMIA = "HH5 OMIA"
var HH6_OMIA = "HH6 OMIA"
var BLAD_OMIA = "BLAD OMIA"
var CVM_OMIA = "CVM OMIA"
var DUMPS_OMIA = "DUMPS OMIA"
var BC_OMIA = "BC OMIA"
var FXID_OMIA = "FXID OMIA"
var MF_OMIA = "MF OMIA"
var FGFR2_OMIA = "FGFR2 OMIA"
var IH_OMIA = "IH OMIA"

var MONOGENETIC_ILLNESSES = map[string]models.GeneticIllness{
	HSD_COLUMN: {
		Name:        "HCD",
		Description: HSD_DESC,
		OMIA:        &HSD_OMIA,
	},
	HH1_COLUMN: {
		Name:        "HH1",
		Description: HH1_DESC,
		OMIA:        &HH1_OMIA,
	},
	HH3_COLUMN: {
		Name:        "HH3",
		Description: HH3_DESC,
		OMIA:        &HH3_OMIA,
	},
	HH4_COLUMN: {
		Name:        "HH4",
		Description: HH4_DESC,
		OMIA:        &HH4_OMIA,
	},
	HH5_COLUMN: {
		Name:        "HH5",
		Description: HH5_DESC,
		OMIA:        &HH5_OMIA,
	},
	HH6_COLUMN: {
		Name:        "HH6",
		Description: HH6_DESC,
		OMIA:        &HH6_OMIA,
	},
	BLAD_COLUMN: {
		Name:        "BLAD",
		Description: BLAD_DESC,
		OMIA:        &BLAD_OMIA,
	},
	CVM_COLUMN: {
		Name:        "CVM",
		Description: CVM_DESC,
		OMIA:        &CVM_OMIA,
	},
	DUMPS_COLUMN: {
		Name:        "DUMPS",
		Description: DUMPS_DESC,
		OMIA:        &DUMPS_OMIA,
	},
	BC_COLUMN: {
		Name:        "BC",
		Description: BC_DESC,
		OMIA:        &BC_OMIA,
	},
	FXID_COLUMN: {
		Name:        "FXID",
		Description: FXID_DESC,
		OMIA:        &FXID_OMIA,
	},
	MF_COLUMN: {
		Name:        "MF",
		Description: MF_DESC,
		OMIA:        &MF_OMIA,
	},
	FGFR2_COLUMN: {
		Name:        "FGFR2",
		Description: FGFR2_DESC,
		OMIA:        &FGFR2_OMIA,
	},
	IH_COLUMN: {
		Name:        "IH",
		Description: IH_DESC,
		OMIA:        &IH_OMIA,
	},
}

func NewGeneticRecord(header []string) (*geneticRecord, error) {
	columns := []string{
		COW_SELECS_COLUMN,
		PROBE_NUMBER_COLUMN,
		BLOOD_DATE_COLUMN,
		RESULT_DATE_COLUMN,
		INBRINDING_COEFF_BY_GENOTYPE_COLUMN,
		HSD_COLUMN,
		HH1_COLUMN,
		HH3_COLUMN,
		HH4_COLUMN,
		HH5_COLUMN,
		HH6_COLUMN,
		BLAD_COLUMN,
		CVM_COLUMN,
		DUMPS_COLUMN,
		BC_COLUMN,
		FXID_COLUMN,
		MF_COLUMN,
		FGFR2_COLUMN,
		IH_COLUMN,
	}
	gr := geneticRecord{}
	hi := make(map[string]int)
	for idx, col := range header {
		hi[col] = idx
	}
	for _, col := range columns {
		if _, ok := hi[col]; !ok {
			return nil, errors.New("Колонки " + col + " не хватает")
		}
	}
	gr.HeaderIndexes = hi
	return &gr, nil
}

func (gr *geneticRecord) FromCsvRecord(rec []string) (CsvToDbLoader, error) {

	if sel, err := strconv.ParseUint(rec[gr.HeaderIndexes[COW_SELECS_COLUMN]],
		10, 64); err == nil {
		gr.CowSelecs = uint(sel)
	} else {
		return nil, errors.New("Не удалось распарсить селекс " + rec[gr.HeaderIndexes[COW_SELECS_COLUMN]])
	}

	gr.ProbeNumber = rec[gr.HeaderIndexes[PROBE_NUMBER_COLUMN]]

	if dateStr := rec[gr.HeaderIndexes[BLOOD_DATE_COLUMN]]; dateStr != "" {
		if date, err := time.Parse(time.DateOnly, dateStr); err == nil {
			gr.BloodDate = &models.DateOnly{Time: date}
		} else {
			return nil, errors.New("Не удалось распарсить дату " + dateStr)
		}
	}

	if dateStr := rec[gr.HeaderIndexes[RESULT_DATE_COLUMN]]; dateStr != "" {
		if date, err := time.Parse(time.DateOnly, dateStr); err == nil {
			gr.ResultDate = &models.DateOnly{Time: date}
		} else {
			return nil, errors.New("Не удалось распарсить дату " + dateStr)
		}
	}

	if inbrStr := rec[gr.HeaderIndexes[INBRINDING_COEFF_BY_GENOTYPE_COLUMN]]; inbrStr != "" {
		if inbr, err := strconv.ParseFloat(inbrStr, 64); err == nil {
			gr.InbrindingCoeffByGenotype = &inbr
		} else {
			return nil, errors.New("Не удалось распарсить коеф. инбриндинга " + inbrStr)
		}
	}
	db := models.GetDb()
	geneticIllnesses := make([]models.GeneticIllnessData, 0, 15)
	for col, val := range MONOGENETIC_ILLNESSES {
		status := rec[gr.HeaderIndexes[col]]
		data := models.GeneticIllnessData{}
		dbIllness := models.GeneticIllness{}

		if err := db.First(&dbIllness, map[string]any{"name": val.Name}).Error; err != nil {
			return nil, errors.New("Не удалось найти заболевание " + val.Name)
		}

		data.Illness = dbIllness

		if status == "" {
			data.Status = nil
		} else {
			dbStatus := models.GeneticIllnessStatus{}

			if err := db.First(&dbStatus, map[string]any{"status": status}).Error; err != nil {
				return nil, errors.New("Не удалось найти статус заболевания " + status)
			}
			data.Status = &dbStatus
		}
		geneticIllnesses = append(geneticIllnesses, data)
	}
	gr.GeneticIllnesses = geneticIllnesses
	return gr, nil
}

func (cr *geneticRecord) ToDbModel(tx *gorm.DB) (any, error) {
	cow := models.Cow{}
	db := models.GetDb()
	if err := db.Preload("Genetic").First(&cow, map[string]any{"selecs_number": cr.CowSelecs}).Error; err != nil {
		return nil, errors.New("Не найдена корова с селексом " + strconv.FormatUint(uint64(cr.CowSelecs), 10))
	}
	if cow.Genetic != nil {
		return nil, errors.New("Уже загружена генетическая информация для коровы с селексом " + strconv.FormatUint(uint64(cr.CowSelecs), 10))
	}
	cow.Genetic = new(models.Genetic)
	cow.Genetic.BloodDate = cr.BloodDate
	cow.Genetic.ResultDate = cr.ResultDate
	cow.Genetic.InbrindingCoeffByGenotype = cr.InbrindingCoeffByGenotype
	cow.Genetic.ProbeNumber = cr.ProbeNumber
	cow.Genetic.GeneticIllnessesData = cr.GeneticIllnesses
	cow.Genetic.CowID = cow.ID
	return *cow.Genetic, nil
}

func (cr *geneticRecord) Copy() *geneticRecord {
	copy := geneticRecord{}
	copy.HeaderIndexes = cr.HeaderIndexes
	return &copy
}

const GENETIC_CSV_PATH = "./csv/genetics/"

var geneticUniqueIndex uint64 = 0

func (l *Load) Genetic() func(*gin.Context) {
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
		uploadedName := GENETIC_CSV_PATH + "genetic_" + strconv.FormatInt(now.Unix(), 16) + "_" + strconv.FormatUint(geneticUniqueIndex, 16) + ".csv"
		if err := c.SaveUploadedFile(csvField[0], uploadedName); err != nil {
			c.JSON(500, err)
			return
		}
		geneticUniqueIndex++
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
		recordWithHeader, err := NewGeneticRecord(header)
		if err != nil {
			c.JSON(422, err.Error())
			return
		}

		errors := []string{}
		errorsMtx := sync.Mutex{}
		loaderWg := sync.WaitGroup{}
		loadChannel := make(chan loaderData)
		MakeLoadingPool(loadChannel, LoadRecordToDb[models.Genetic])

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
