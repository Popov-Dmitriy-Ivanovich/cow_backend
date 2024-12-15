package admin

import (
	"cow_backend/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Admin) checkNews() func(*gin.Context) {
	return func(c *gin.Context) {
		db := models.GetDb()
		pageStr := c.Query("page") // Get the requested page number from the query string.
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			page = 1
		}

		limit := 20
		offset := (page - 1) * limit
		news := []models.News{}
		db.
			Order("Date DESC").
			Preload("Region").
			Limit(limit).
			Offset(offset).
			Find(&news)
		var count int64
		db.Model(&models.News{}).Count(&count)
		totalPages := int(math.Ceil(float64(count) / float64(limit)))
		c.HTML(http.StatusOK, "AdminNewsPage.tmpl", gin.H{
			"title":       "Таблица пользователей",
			"news":        news,
			"currentPage": page,
			"totalPages":  totalPages,
		})
	}
}

func (s *Admin) CreateNews() func(*gin.Context) {
	return func(c *gin.Context) {
		db := models.GetDb()
		regions := []models.Region{}
		db.Order("name").Find(&regions)

		c.HTML(http.StatusOK, "AdminCreateNewsPage.tmpl", gin.H{
			"title":   "Создание новости",
			"regions": regions,
		})
	}
}

func (s *Admin) NewNews() func(*gin.Context) {
	return func(c *gin.Context) {
		var request struct {
			Title    string          `json:"title"`
			RegionId *uint           `json:"region"`
			Date     models.DateOnly `json:"date"`
			Text     string          `json:"text"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		news := models.News{
			Date:     request.Date,
			RegionId: request.RegionId,
			Title:    request.Title,
			Text:     request.Text,
		}
		db := models.GetDb()

		if err := updateSequenceNews(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении последовательности: " + err.Error()})
			return
		}

		if err := db.Create(&news).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при добавлении новости: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Новость добавлена"})
	}
}

func updateSequenceNews() error {
	var maxID uint
	db := models.GetDb()
	if err := db.Model(&models.News{}).Select("max(id)").Scan(&maxID).Error; err != nil {
		return err
	}

	if err := db.Exec("SELECT setval(pg_get_serial_sequence('news', 'id'), ?)", maxID).Error; err != nil {
		return err
	}
	return nil
}

func (s *Admin) UpdateNewsPage() func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		db := models.GetDb()

		regions := []models.Region{}
		db.Find(&regions)

		news := models.News{}
		if err := db.
			Preload("Region").
			First(&news, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Новость с таким ID не найдена"})
			return
		}

		c.HTML(http.StatusOK, "AdminUpdateNewsPage.tmpl", gin.H{
			"title":   "Редактирование новости",
			"news":    news,
			"regions": regions,
		})
	}
}

func (s *Admin) UpdateNews() func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		db := models.GetDb()
		news := models.News{}
		if err := db.First(&news, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь с таким ID не найден"})
			return
		}

		var request struct {
			Title    string          `json:"title"`
			RegionId *uint           `json:"region"`
			Date     models.DateOnly `json:"date"`
			Text     string          `json:"text"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		news.Title = request.Title
		news.RegionId = request.RegionId
		news.Text = request.Text
		news.Date = request.Date

		if err := db.Save(&news).Error; err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

}
