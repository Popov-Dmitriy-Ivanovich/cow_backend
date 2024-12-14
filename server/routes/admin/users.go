package admin

import (
	"cow_backend/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Admin) CheckUsersTable() func(*gin.Context) {
	return func(c *gin.Context) {
		db := models.GetDb()
		pageStr := c.Query("page") // Get the requested page number from the query string.
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			page = 1
		}

		limit := 20
		offset := (page - 1) * limit
		users := []models.User{}
		db.
			Order("id").
			Where("role_id != 4").
			Preload("Farm").
			Preload("Region").
			Preload("Role").
			Limit(limit).
			Offset(offset).
			Find(&users)
		var count int64
		db.Model(&models.User{}).Count(&count)
		totalPages := int(math.Ceil(float64(count) / float64(limit)))
		c.HTML(http.StatusOK, "AdminUsersPage.tmpl", gin.H{
			"title":       "Таблица пользователей",
			"users":       users,
			"currentPage": page,
			"totalPages":  totalPages,
		})
	}
}
