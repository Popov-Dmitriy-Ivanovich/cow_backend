package admin

import (
	"cow_backend/models"
	"fmt"
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
			Where("type= ?", typeHoz).
			Limit(limit).
			Offset(offset).
			Find(&hoz)
		db.Model(&models.Farm{}).Where("type = ?", typeHoz).Count(&count)
		AdminPages := map[int]string{
			1: "AdminHoldingsPage.tmpl",
			2: "AdminHozPage.tmpl",
			3: "AdminFarmsPage.tmpl",
		}
		totalPages := int(math.Ceil(float64(count) / float64(limit)))
		fmt.Println(AdminPages[typeHoz])
		c.HTML(http.StatusOK, AdminPages[typeHoz], gin.H{
			"title":       "Таблица холдингов",
			"hoz":         hoz,
			"currentPage": page,
			"totalPages":  totalPages})
	}
}
