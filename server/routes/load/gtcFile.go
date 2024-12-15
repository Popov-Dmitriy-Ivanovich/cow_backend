package load

import (
	"cow_backend/models"
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const GTC_FILE_PATH = "./static/gtc/"

var gtcUniqueIndex uint64 = 0

func (l *Load) GtcFile() func(*gin.Context) {
	return func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(500, err.Error())
			return
		}
		gtc, ok := form.File["gtc"]
		if !ok {
			c.JSON(500, "not found field gtc")
			return
		}

		now := time.Now()
		filename := "gtc_" + strconv.FormatInt(now.Unix(), 16) + "_" + strconv.FormatUint(gtcUniqueIndex, 16)
		uploadFolder := GTC_FILE_PATH + filename

		filesNaming := map[string]string{}

		for _, file := range gtc {
			uploadPath := uploadFolder + "/" + file.Filename
			filesNaming[file.Filename] = filename + "/" + file.Filename
			if err := c.SaveUploadedFile(file, uploadPath); err != nil {
				c.JSON(500, err.Error())
				return
			}
		}

		csvFile := form.File["csv"]
		csvFilePath := "./csv/gtcCsv" + strconv.FormatInt(now.Unix(), 16) + "_" + strconv.FormatUint(gtcUniqueIndex, 16) + ".csv"
		if err := c.SaveUploadedFile(csvFile[0], csvFilePath); err != nil {
			c.JSON(500, err.Error())
		}

		file, err := os.Open(csvFilePath)
		if err != nil {
			c.JSON(500, "error opening file")
			return
		}
		defer file.Close()
		csvReader := csv.NewReader(file)
		errors := []string{}
		for record, err := csvReader.Read(); err != io.EOF; record, err = csvReader.Read() {
			selecsStr := record[0]
			fileName := record[1]
			filePath, ok := filesNaming[fileName]
			if !ok {
				errors = append(errors, "не загружен файл "+fileName)
				continue
			}
			dbCow := models.Cow{}
			db := models.GetDb()
			if err := db.Preload("Genetic").First(&dbCow, map[string]any{"selecs_number": selecsStr}).Error; err != nil {
				errors = append(errors, "не удалось найти корову с селексом "+selecsStr)
				continue
			}
			if dbCow.Genetic == nil {
				dbCow.Genetic = new(models.Genetic)
				dbCow.Genetic.ResultDate = &models.DateOnly{Time: time.Now().UTC()}
				dbCow.Genetic.BloodDate = &models.DateOnly{Time: time.Now().UTC()}
			}
			dbCow.Genetic.GtcFilePath = &filePath

			if err := db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&dbCow).Error; err != nil {
				errors = append(errors, err.Error())
				continue
			}
		}
		c.JSON(200, errors)
	}
}
