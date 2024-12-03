package load

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const COW_CSV_PATH = "./csv/cow/"

var uniqueIndex uint64 = 0

func (l *Load) Cow() func(*gin.Context) {
	return func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(500, err)
			return
		}
		csv, ok := form.File["csvCow"]
		if !ok {
			c.JSON(500, "not found field csvCow")
			return
		}

		now := time.Now()
		uploadedName := COW_CSV_PATH + "cow_" + now.Format(time.Stamp) + "_" + strconv.FormatUint(uniqueIndex, 10) + ".csv"
		if err := c.SaveUploadedFile(csv[0], uploadedName); err != nil {
			c.JSON(500, err)
			return
		}
		uniqueIndex++
		c.JSON(200, "OK")
	}
}
