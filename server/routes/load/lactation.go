package load

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const LACTATION_CSV_PATH = "./csv/lactations/"

var lactationUniqueIndex uint64 = 0

func (l *Load) Lactation() func(*gin.Context) {
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
		uploadedName := LACTATION_CSV_PATH + "lactation_" + strconv.FormatInt(now.Unix(), 16) + "_" + strconv.FormatUint(lactationUniqueIndex, 16) + ".csv"
		if err := c.SaveUploadedFile(csv[0], uploadedName); err != nil {
			c.JSON(500, err)
			return
		}
		lactationUniqueIndex++
		c.JSON(200, "OK")
	}
}
