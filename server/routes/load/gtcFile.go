package load

import (
	"cow_backend/models"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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

		db := models.GetDb()
		errors := []string{}
		for fileName, filePath := range filesNaming {
			cow := models.Cow{}
			selecs := strings.Split(fileName, ".")[0]
			if err := db.First(&cow, map[string]any{"selecs_number": selecs}).Error; err != nil {
				errors = append(errors, err.Error())
				continue
			}
			genetic := models.Genetic{}
			if err := db.FirstOrCreate(&genetic, map[string]any{"cow_id": cow.ID}).Error; err != nil {
				errors = append(errors, err.Error())
				continue
			}
			genetic.GtcFilePath = &filePath

			if err := db.Save(&genetic).Error; err != nil {
				errors = append(errors, err.Error())
				continue
			}
		}
		c.JSON(200, errors)
	}
}
