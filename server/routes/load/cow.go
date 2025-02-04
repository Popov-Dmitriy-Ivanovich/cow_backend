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

const COW_CSV_PATH = "./csv/cows/"

var cowUniqueIndex uint64 = 0

type cowRecord struct {
	Selecs                  string
	InventoryNumber         *string
	FarmNumber              *string
	FarmName                *string
	HozNumber               string
	HozName                 string
	BreedId                 uint
	BreedName               string
	SexId                   uint
	SexName                 string
	FatherSelecs            *uint64
	MotherSelecs            *uint64
	IdentificationNumber    *string
	RSHNNumber              *string
	Name                    string
	InbrindingCoeffByFamily *float64

	BirthDate   models.DateOnly
	DepartDate  *models.DateOnly
	DeathDate   *models.DateOnly
	BirkingDate *models.DateOnly

	OldInvNumber *string

	PrevHozNumber  *string
	PrevHozName    *string
	BirthHozNumber *string
	BirthHozName   *string
	BirthMethod    *string
	HeaderIndexes  map[string]int
}

func NewCowRecord(header []string) (*cowRecord, error) {
	cr := cowRecord{}
	headerIndexes := make(map[string]int)
	for idx, colName := range header {
		headerIndexes[colName] = idx
	}
	cr.HeaderIndexes = headerIndexes
	requiredColumns := []string{
		"CowSelecs",
		"CowInvNumber",
		"FarmID",
		"FarmName",
		"HozID",
		"HozName",
		"BreedID",
		"BreedName",
		"SexID",
		"SexName",
		"FatherSelecs",
		"MotherSelecs",
		"IdentificationNumber",
		"InventoryNumber",
		"RSHNNumber",
		"Name",
		"InbrindingCoeffByFamily",
		"BirthDate",
		"DepartDate",
		"DeathDate",
		"OldInvNumber",
		"BirkingDate",
		"PrevHozId",
		"PrevHozName",
		"BirthHozID",
		"BirthHozName",
		"BirthWay",
	}
	for _, col := range requiredColumns {
		if _, ok := cr.HeaderIndexes[col]; !ok {
			return nil, errors.New("Column " + col + " not found in CSV")
		}
	}
	return &cr, nil
}

func (cr *cowRecord) FromCsvRecord(rec []string) (CsvToDbLoader, error) {
	res := cowRecord{}

	res.Selecs = rec[cr.HeaderIndexes["CowSelecs"]]

	if rec[cr.HeaderIndexes["CowInvNumber"]] != "" {
		res.InventoryNumber = &rec[cr.HeaderIndexes["CowInvNumber"]]
	}

	if rec[cr.HeaderIndexes["FarmID"]] != "" {
		res.FarmNumber = &rec[cr.HeaderIndexes["FarmID"]]
	}

	if rec[cr.HeaderIndexes["FarmID"]] != "" {
		res.FarmName = &rec[cr.HeaderIndexes["FarmName"]]
	}

	res.HozNumber = rec[cr.HeaderIndexes["HozID"]]

	res.HozName = rec[cr.HeaderIndexes["HozName"]]

	if breedId, err := strconv.ParseUint(rec[cr.HeaderIndexes["BreedID"]], 10, 64); err != nil {
		return nil, errors.New("Не удалось распарсить ID породы " + rec[cr.HeaderIndexes["BreedID"]] + err.Error())
	} else {
		res.BreedId = uint(breedId)
	}

	res.BreedName = rec[cr.HeaderIndexes["BreedName"]]

	if sexId, err := strconv.ParseUint(rec[cr.HeaderIndexes["SexID"]], 10, 64); err != nil {
		return nil, errors.New("Не удалось распарсить ID пола " + rec[cr.HeaderIndexes["SexID"]] + err.Error())
	} else {
		res.SexId = uint(sexId)
	}

	res.SexName = rec[cr.HeaderIndexes["SexName"]]

	if rec[cr.HeaderIndexes["FatherSelecs"]] != "" {
		sel, err := strconv.ParseUint(rec[cr.HeaderIndexes["FatherSelecs"]], 10, 64)
		if err != nil {
			return nil, err
		}
		res.FatherSelecs = &sel
	}

	if rec[cr.HeaderIndexes["MotherSelecs"]] != "" {
		sel, err := strconv.ParseUint(rec[cr.HeaderIndexes["MotherSelecs"]], 10, 64)
		if err != nil {
			return nil, err
		}
		res.MotherSelecs = &sel
	}

	if rec[cr.HeaderIndexes["IdentificationNumber"]] != "" {
		res.IdentificationNumber = &rec[cr.HeaderIndexes["IdentificationNumber"]]
	}

	if rec[cr.HeaderIndexes["RSHNNumber"]] != "" {
		res.RSHNNumber = &rec[cr.HeaderIndexes["RSHNNumber"]]
	}

	res.Name = rec[cr.HeaderIndexes["Name"]]

	if rec[cr.HeaderIndexes["InbrindingCoeffByFamily"]] != "" {
		if icbf, err := strconv.ParseFloat(rec[cr.HeaderIndexes["InbrindingCoeffByFamily"]], 64); err != nil {
			return nil, err
		} else {
			res.InbrindingCoeffByFamily = &icbf
		}
	}

	if birthDate, err := ParseTime(rec[cr.HeaderIndexes["BirthDate"]]); err != nil {
		return nil, err
	} else {
		res.BirthDate = models.DateOnly{Time: birthDate}
	}

	if rec[cr.HeaderIndexes["DepartDate"]] != "" {
		if depDate, err := ParseTime(rec[cr.HeaderIndexes["DepartDate"]]); err != nil {
			return nil, err
		} else {
			res.DepartDate = &models.DateOnly{Time: depDate}
		}
	}
	if rec[cr.HeaderIndexes["DeathDate"]] != "" {
		if deathDate, err := ParseTime(rec[cr.HeaderIndexes["DeathDate"]]); err != nil {
			return nil, err
		} else {
			res.DepartDate = &models.DateOnly{Time: deathDate}
		}
	}

	if rec[cr.HeaderIndexes["OldInvNumber"]] != "" {
		res.OldInvNumber = &rec[cr.HeaderIndexes["OldInvNumber"]]
	}

	if rec[cr.HeaderIndexes["BirkingDate"]] != "" {
		if birkingDate, err := ParseTime(rec[cr.HeaderIndexes["BirkingDate"]]); err != nil {
			return nil, err
		} else {
			res.BirkingDate = &models.DateOnly{Time: birkingDate}
		}
	}

	if rec[cr.HeaderIndexes["PrevHozId"]] != "" {
		res.PrevHozNumber = &rec[cr.HeaderIndexes["PrevHozId"]]
	}

	if rec[cr.HeaderIndexes["PrevHozName"]] != "" {
		res.PrevHozName = &rec[cr.HeaderIndexes["PrevHozName"]]
	}

	if rec[cr.HeaderIndexes["BirthHozID"]] != "" {
		res.BirthHozNumber = &rec[cr.HeaderIndexes["BirthHozID"]]
	}
	if rec[cr.HeaderIndexes["BirthHozName"]] != "" {
		res.BirthHozName = &rec[cr.HeaderIndexes["BirthHozName"]]
	}
	if rec[cr.HeaderIndexes["BirthWay"]] != "" {
		res.BirthMethod = &rec[cr.HeaderIndexes["BirthWay"]]
	}

	return &res, nil
}

func (cr *cowRecord) ToDbModel(tx *gorm.DB) (any, error) {
	res := models.Cow{}

	sameCowCount := int64(0)
	if tx.Model(&models.Cow{}).Where("inventory_number = ? AND selecs_number = ?", cr.InventoryNumber, cr.Selecs).Count(&sameCowCount); sameCowCount != 0 {
		return nil, errors.New("that cow already exists")
	}

	if cr.FarmNumber != nil {
		if err := tx.First(&res.Farm, map[string]any{"hoz_number": cr.FarmNumber}).Error; err != nil {
			return nil, errors.New("не удалось найти ферму с hoz_number = " + *cr.FarmNumber)
		}
	}
	if err := tx.First(&res.FarmGroup, map[string]any{"hoz_number": cr.HozNumber}).Error; err != nil {
		return nil, errors.New("не удалось найти хозяйство с hoz_number = " + cr.HozNumber)
	}
	if err := tx.First(&res.Breed, map[string]any{"name": cr.BreedName}).Error; err != nil {
		return nil, errors.New("не удалось найти породу с BreedName = " + cr.BreedName)
	}
	if err := tx.First(&res.Sex, map[string]any{"id": cr.SexId}).Error; err != nil {
		return nil, errors.New("не удалось найти пол с ID = " + strconv.FormatUint(uint64(cr.SexId), 10))
	}
	if res.Sex.Name != cr.SexName {
		return nil, errors.New("Wrong sex name, sex with id has name " + res.Sex.Name + " but " + cr.SexName + " provided")
	}

	res.FatherSelecs = cr.FatherSelecs
	res.MotherSelecs = cr.MotherSelecs
	res.IdentificationNumber = cr.IdentificationNumber

	if sel, err := strconv.ParseUint(cr.Selecs, 10, 64); err != nil {
		return nil, err
	} else {
		res.SelecsNumber = &sel
	}
	res.InventoryNumber = cr.InventoryNumber
	// res.SelecsNumber = &cr.Selecs
	res.RSHNNumber = cr.RSHNNumber
	res.Name = cr.Name
	res.InbrindingCoeffByFamily = cr.InbrindingCoeffByFamily
	res.Approved = 0
	res.BirthDate = cr.BirthDate
	res.DepartDate = cr.DepartDate
	res.DeathDate = cr.DeathDate
	res.BirkingDate = cr.BirkingDate
	res.BirthMethod = cr.BirthMethod
	if cr.PrevHozNumber != nil {
		if err := tx.First(&res.PreviousHoz, map[string]any{"hoz_number": cr.PrevHozNumber}).Error; err != nil {
			return nil, errors.New("не удалось найти хозяйство с номером " + *cr.PrevHozNumber)
		}
	}
	if cr.BirthHozNumber != nil {
		if err := tx.First(&res.BirthHoz, map[string]any{"hoz_number": cr.BirthHozNumber}).Error; err != nil {
			return nil, errors.New("не удалось найти хозяйство с номером " + *cr.BirthHozNumber)
		}
	}
	if cr.OldInvNumber != nil {
		res.PreviousInventoryNumber = cr.OldInvNumber
	}
	return res, nil
}

func (cr *cowRecord) Copy() *cowRecord {
	copy := cowRecord{}
	copy.HeaderIndexes = cr.HeaderIndexes
	return &copy
}
func (l *Load) Cow() func(*gin.Context) {
	return func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(500, "ошибка чтения формы")
			return
		}
		csvField, ok := form.File["csv"]
		if !ok {
			c.JSON(500, "not found field csv")
			return
		}

		now := time.Now()
		uploadedName := COW_CSV_PATH + "cow_" + strconv.FormatInt(now.Unix(), 16) + "_" + strconv.FormatUint(cowUniqueIndex, 16) + ".csv"
		if err := c.SaveUploadedFile(csvField[0], uploadedName); err != nil {
			c.JSON(500, "ошибка сохранения загруженного файла")
			return
		}
		cowUniqueIndex++

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
		recordWithHeader, err := NewCowRecord(header)
		if err != nil {
			c.JSON(422, err.Error())
			return
		}

		errors := []string{}
		errorsMtx := sync.Mutex{}
		loaderWg := sync.WaitGroup{}
		loadChannel := make(chan loaderData)
		MakeLoadingPool(loadChannel, LoadRecordToDb[models.Cow])

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
