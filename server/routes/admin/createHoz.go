package admin

import (
	"cow_backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Admin) CreateHoz() func(*gin.Context) {
	return func(c *gin.Context) {

		db := models.GetDb()
		regions := []models.Region{}
		districts := []models.District{}
		holds := []models.Farm{}
		db.Find(&regions)
		db.Find(&districts)
		db.Where("type = 1").Find(&holds)
		c.HTML(http.StatusOK, "AdminCreateHozPage.tmpl", gin.H{
			"title":     "Создание хозяйства",
			"holds":     holds,
			"regions":   regions,
			"districts": districts})
	}
}
