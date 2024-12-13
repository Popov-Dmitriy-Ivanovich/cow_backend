package load

import (
	"cow_backend/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const DOCUMENT_PATH = "./static/documents/"

var documentUniqueIndex uint64 = 0

func (l *Load) Document() func(*gin.Context) {
	return func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(500, err)
			return
		}
		document, ok := form.File["Document"]
		if !ok || len(document) == 0 {
			c.JSON(500, "not found field csv")
			return
		}

		now := time.Now()
		fileName := "doc" + strconv.FormatInt(now.Unix(), 16) + "_" + strconv.FormatUint(documentUniqueIndex, 16) + document[0].Filename
		uploadedName := DOCUMENT_PATH + fileName

		if err := c.SaveUploadedFile(document[0], uploadedName); err != nil {
			c.JSON(500, err)
			return
		}
		documentUniqueIndex++

		cowId, err := strconv.ParseUint(form.Value["CowID"][0], 10, 64)
		if err != nil {
			c.JSON(422, err.Error())
			return
		}
		dbCow := models.Cow{}
		db := models.GetDb()
		if err := db.Preload("Documents").First(&dbCow, cowId).Error; err != nil {
			c.JSON(422, err.Error())
			return
		}
		if err := db.Model(&dbCow).Association("Documents").Append(&models.Document{Path: fileName}); err != nil {
			c.JSON(500, err.Error())
			return
		}
		c.JSON(200, "ok")
	}
}
