package load

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const CM_CSV_PATH = "./csv/check_milks/"

var cmUniqueIndex uint64 = 0

func (l *Load) CheckMilk() func(*gin.Context) {
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
		uploadedName := CM_CSV_PATH + "check_milk_" + now.Format(time.Stamp) + "_" + strconv.FormatUint(cmUniqueIndex, 10) + ".csv"
		if err := c.SaveUploadedFile(csv[0], uploadedName); err != nil {
			c.JSON(500, err)
			return
		}
		cmUniqueIndex++
		c.JSON(200, "OK")
	}
}
