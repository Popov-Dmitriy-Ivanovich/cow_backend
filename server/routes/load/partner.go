package load

import (
	"cow_backend/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const PARTNER_LOGO_PATH = "./static/partners/"

var partnerLogoUniqueIndex uint64 = 0

func (l *Load) Partner() func(*gin.Context) {
	return func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(500, err)
			return
		}
		logo, ok := form.File["Logo"]
		if !ok {
			c.JSON(500, "not found field logo")
			return
		}

		now := time.Now()
		filname := "partnerLogo" + strconv.FormatInt(now.Unix(), 16) + "_" + strconv.FormatUint(partnerLogoUniqueIndex, 16) + logo[0].Filename
		uploadedName := PARTNER_LOGO_PATH + filname
		if err := c.SaveUploadedFile(logo[0], uploadedName); err != nil {
			c.JSON(500, err)
			return
		}
		db := models.GetDb()
		partner := models.Partner{
			Name:        form.Value["Name"][0],
			Address:     &form.Value["Addres"][0],
			Phone:       &form.Value["Phone"][0],
			Email:       &form.Value["Email"][0],
			Description: form.Value["Description"][0],
			LogoPath:    &filname,
		}
		if err := db.Create(&partner).Error; err != nil {
			c.JSON(500, err.Error())
		}
		partnerLogoUniqueIndex++
		c.JSON(200, "OK")
	}
}
