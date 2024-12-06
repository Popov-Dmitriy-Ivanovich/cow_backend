package admin

import (
	"cow_backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Admin) CheckCowTable() func(*gin.Context) {
	return func(c *gin.Context) {
		db := models.GetDb()
		cows := []models.Cow{}
		db.Where("approved = 0").Find(&cows)
		c.HTML(http.StatusOK, "AdminCowTablePage.tmpl", gin.H{"title": "Таблица коров", "cows": cows})
	}
}
