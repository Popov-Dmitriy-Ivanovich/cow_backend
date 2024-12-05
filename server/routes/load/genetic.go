package load

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const GENETIC_CSV_PATH = "./csv/genetics/"

var geneticUniqueIndex uint64 = 0

func (l *Load) Genetic() func(*gin.Context) {
	return func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(500, err)
			return
		}
		csv, ok := form.File["csv"]
		if !ok {
			c.JSON(500, "not found field csv")
			return
		}

		now := time.Now()
		uploadedName := GENETIC_CSV_PATH + "genetic_" + strconv.FormatInt(now.Unix(), 16) + "_" + strconv.FormatUint(geneticUniqueIndex, 16) + ".csv"
		if err := c.SaveUploadedFile(csv[0], uploadedName); err != nil {
			c.JSON(500, err)
			return
		}
		geneticUniqueIndex++
		c.JSON(200, "OK")
	}
}
