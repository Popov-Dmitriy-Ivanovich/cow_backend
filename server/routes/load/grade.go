package load

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const GRADE_CSV_PATH = "./csv/grades/"

var gradeUniqueIndex uint64 = 0

func (l *Load) Grade() func(*gin.Context) {
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
		uploadedName := GRADE_CSV_PATH + "grade_" + now.Format(time.Stamp) + "_" + strconv.FormatUint(gradeUniqueIndex, 10) + ".csv"
		if err := c.SaveUploadedFile(csv[0], uploadedName); err != nil {
			c.JSON(500, err)
			return
		}
		gradeUniqueIndex++
		c.JSON(200, "OK")
	}
}
