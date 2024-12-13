package admin

import (
	"cow_backend/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Admin) CheckCowTable() func(*gin.Context) {
	return func(c *gin.Context) {
		db := models.GetDb()
		pageStr := c.Query("page")
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			page = 1
		}

		limit := 20
		offset := (page - 1) * limit

		cows := []models.Cow{}
		db.Where("approved = 0").
			Order("id").
			Preload("Farm").
			Preload("FarmGroup").
			Preload("PreviousHoz").
			Preload("BirthHoz").
			Preload("Breed").
			Preload("Sex").
			Preload("BirthHoz").
			Limit(limit).
			Offset(offset).
			Find(&cows)

		var count int64
		db.Model(&models.Cow{}).Where("approved = 0").Count(&count)

		totalPages := int(math.Ceil(float64(count) / float64(limit)))

		c.HTML(http.StatusOK, "AdminCowTablePage.tmpl", gin.H{
			"title":       "Таблица коров",
			"cows":        cows,
			"currentPage": page,
			"totalPages":  totalPages,
		})
	}
}
