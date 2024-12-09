package admin

import (
	"cow_backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Admin) CreateHolding() func(*gin.Context) {
	return func(c *gin.Context) {

		db := models.GetDb()
		regions := []models.Region{}
		districts := []models.District{}
		db.Order("name").Find(&regions)
		db.Order("name").Find(&districts)

		c.HTML(http.StatusOK, "AdminCreateHoldingPage.tmpl", gin.H{
			"title":     "Создание холдинга",
			"regions":   regions,
			"districts": districts})
	}
}
