package load

import (
	"cow_backend/models"
	"strconv"
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
		filename := "gtc_" + strconv.FormatInt(now.Unix(), 16) + "_" + strconv.FormatUint(gtcUniqueIndex, 16) + ".gtc"
		uploadedName := GTC_FILE_PATH + filename
		if err := c.SaveUploadedFile(gtc[0], uploadedName); err != nil {
			c.JSON(500, err.Error())
			return
		}

		db := models.GetDb()
		genetic := models.Genetic{}
		if err := db.First(&genetic, form.Value["GeneticID"][0]).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}
		genetic.GtcFilePath = &filename
		if err := db.Save(&genetic).Error; err != nil {
			c.JSON(500, err.Error())
		}
		gtcUniqueIndex++
		c.JSON(200, "OK")
	}
}
