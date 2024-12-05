package load

import (
	"cow_backend/models"
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const COW_CSV_PATH = "./csv/cows/"

var cowUniqueIndex uint64 = 0

type cowRecord struct {
	Selecs                  string
	InventoryNumber         string
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
}

func (cr *cowRecord) FromCsvRecord(rec []string) (CsvToDbLoader, error) {
	res := cowRecord{}

	res.Selecs = rec[0]

	res.InventoryNumber = rec[1]

	if rec[2] != "" {
		res.FarmNumber = &rec[2]
	}

	if rec[3] != "" {
		res.FarmName = &rec[3]
	}

	res.HozNumber = rec[4]

	res.HozName = rec[5]

	if breedId, err := strconv.ParseUint(rec[6], 10, 64); err != nil {
		return nil, err
	} else {
		res.BreedId = uint(breedId)
	}

	res.BreedName = rec[7]

	if sexId, err := strconv.ParseUint(rec[8], 10, 64); err != nil {
		return nil, err
	} else {
		res.SexId = uint(sexId)
	}

	res.SexName = rec[9]

	if rec[10] != "" {
		sel, err := strconv.ParseUint(rec[10], 10, 64)
		if err != nil {
			return nil, err
		}
		res.FatherSelecs = &sel
	}

	if rec[11] != "" {
		sel, err := strconv.ParseUint(rec[11], 10, 64)
		if err != nil {
			return nil, err
		}
		res.MotherSelecs = &sel
	}

	if rec[12] != "" {
		res.IdentificationNumber = &rec[12]
	}

	if rec[13] != "" {
		// res.InventoryNumber = &rec[13] inventory number twice
	}

	if rec[14] != "" {
		res.RSHNNumber = &rec[14]
	}

	res.Name = rec[15]

	if rec[16] != "" {
		if icbf, err := strconv.ParseFloat(rec[16], 64); err != nil {
			return nil, err
		} else {
			res.InbrindingCoeffByFamily = &icbf
		}
	}

	if birthDate, err := time.Parse("02.01.2006", rec[17]); err != nil {
		return nil, err
	} else {
		res.BirthDate = models.DateOnly{Time: birthDate}
	}

	if rec[18] != "" {
		if depDate, err := time.Parse("02.01.2006", rec[18]); err != nil {
			return nil, err
		} else {
			res.DepartDate = &models.DateOnly{Time: depDate}
		}
	}
	if rec[19] != "" {
		if deathDate, err := time.Parse("02.01.2006", rec[19]); err != nil {
			return nil, err
		} else {
			res.DepartDate = &models.DateOnly{Time: deathDate}
		}
	}

	if rec[20] != "" {
		res.OldInvNumber = &rec[20]
	}

	if rec[21] != "" {
		if birkingDate, err := time.Parse("02.01.2006", rec[21]); err != nil {
			return nil, err
		} else {
			res.DepartDate = &models.DateOnly{Time: birkingDate}
		}
	}

	if rec[22] != "" {
		res.PrevHozNumber = &rec[22]
	}

	if rec[23] != "" {
		res.PrevHozName = &rec[23]
	}

	if rec[24] != "" {
		res.BirthHozNumber = &rec[24]
	}
	if rec[25] != "" {
		res.BirthHozName = &rec[25]
	}
	if rec[26] != "" {
		res.BirthMethod = &rec[26]
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
	res.InventoryNumber = &cr.InventoryNumber
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
		csvReader := csv.NewReader(file)
		csvReader.Read()

		db := models.GetDb()
		if err := db.Transaction(func(tx *gorm.DB) error {
			// do some database operations in the transaction (use 'tx' from this point, not 'db')
			for record, err := csvReader.Read(); err != io.EOF; record, err = csvReader.Read() {
				if err != nil {
					return err
				}
				if err := LoadRecordToDb[models.Cow](&cowRecord{}, record, tx); err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
			c.JSON(500, err.Error())
			return
		}

		c.JSON(200, "OK")
	}
}
