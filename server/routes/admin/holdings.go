package admin

import (
	"cow_backend/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Admin) CheckHoldingsTable() func(*gin.Context) {
	return func(c *gin.Context) {
		db := models.GetDb()
		pageStr := c.Query("page") // Get the requested page number from the query string.
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			page = 1
		}
		limit := 20
		offset := (page - 1) * limit

		holds := []models.Farm{}
		db.
			Where("parrent_id is Null").
			Limit(limit).
			Offset(offset).
			Find(&holds)

		var count int64
		db.Model(&models.Farm{}).Where("parrent_id is Null").Count(&count)
		totalPages := int(math.Ceil(float64(count) / float64(limit)))
		c.HTML(http.StatusOK, "AdminHoldingsPage.tmpl", gin.H{
			"title":       "Таблица холдингов",
			"holds":       holds,
			"currentPage": page,
			"totalPages":  totalPages})
	}
}
