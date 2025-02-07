package load

import (
	"cow_backend/models"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const DOCUMENT_PATH = "./static/documents/"

var documentUniqueIndex uint64 = 0

func (l *Load) Document() func(*gin.Context) {
	return func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(500, err.Error())
			return
		}
		doc, ok := form.File["Document"]
		if !ok {
			c.JSON(500, "not found field Document")
			return
		}

		now := time.Now()
		filename := "doc_" + strconv.FormatInt(now.Unix(), 16) + "_" + strconv.FormatUint(documentUniqueIndex, 16)
		uploadFolder := DOCUMENT_PATH + filename

		filesNaming := map[string]string{}

		for _, file := range doc {
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
			dbDoc := models.Document{
				CowID: cow.ID,
				Path:  filePath,
			}
			if err := db.Create(&dbDoc).Error; err != nil {
				errors = append(errors, err.Error())
			}
		}
		c.JSON(200, errors)
	}
}
