package admin

import (
	"cow_backend/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Admin) CheckHozTable(typeHoz int) func(*gin.Context) {
	return func(c *gin.Context) {
		var count int64
		db := models.GetDb()
		pageStr := c.Query("page") // Get the requested page number from the query string.
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			page = 1
		}
		limit := 20
		offset := (page - 1) * limit

		hoz := []models.Farm{}
		db.
			Order("id").
			Where("type= ?", typeHoz).
			Preload("Parrent").
			Preload("District.Region").
			Limit(limit).
			Offset(offset).
			Find(&hoz)
		db.Model(&models.Farm{}).Where("type = ?", typeHoz).Count(&count)

		totalPages := int(math.Ceil(float64(count) / float64(limit)))

		AdminPages := map[int]string{
			1: "AdminHoldingsPage.tmpl",
			2: "AdminHozPage.tmpl",
			3: "AdminFarmsPage.tmpl",
		}

		c.HTML(http.StatusOK, AdminPages[typeHoz], gin.H{
			"title":       "Таблица холдингов",
			"hoz":         hoz,
			"currentPage": page,
			"totalPages":  totalPages})
	}
}
