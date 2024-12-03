package load

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const EVENT_CSV_PATH = "./csv/events/"

var eventUniqueIndex uint64 = 0

func (l *Load) Event() func(*gin.Context) {
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
		uploadedName := EVENT_CSV_PATH + "event" + now.Format(time.Stamp) + "_" + strconv.FormatUint(eventUniqueIndex, 10) + ".csv"
		if err := c.SaveUploadedFile(csv[0], uploadedName); err != nil {
			c.JSON(500, err)
			return
		}
		eventUniqueIndex++
		c.JSON(200, "OK")
	}
}
