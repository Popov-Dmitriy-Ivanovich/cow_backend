package admin

import (
	"cow_backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Admin) UpdateFarmPage() func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		db := models.GetDb()
		regions := []models.Region{}
		districts := []models.District{}
		db.Find(&regions)
		db.Find(&districts)

		farm := models.Farm{}
		if err := db.
			Preload("Parrent").
			Preload("District.Region").
			First(&farm, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Холдинг с таким ID не найден"})
			return
		}

		c.HTML(http.StatusOK, "AdminUpdateHoldingPage.tmpl", gin.H{
			"title":     "Редактирование холдинга",
			"farm":      farm,
			"regions":   regions,
			"districts": districts})
	}
}

func (s *Admin) UpdateFarm() func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		db := models.GetDb()
		farm := models.Farm{}
		if err := db.First(&farm, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Ферма с таким ID не найдена"})
			return
		}

		var request struct {
			HozNumber   string `json:"holding_number"`
			DistrictId  string `json:"district"`
			Fullname    string `json:"fullname"`
			Name        string `json:"name"`
			Inn         string `json:"inn"`
			Address     string `json:"address"`
			Phone       string `json:"phone"`
			Email       string `json:"email"`
			Description string `json:"description"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if request.HozNumber != "" {
			farm.HozNumber = &request.HozNumber
		}
		if request.DistrictId != "" {
			distr, err := strconv.ParseUint(request.DistrictId, 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID фермы"})
			}
			farm.DistrictId = uint(distr)
		}
		if request.Fullname != "" {
			farm.Name = request.Fullname
		}
		if request.Name != "" {
			farm.NameShort = request.Name
		}
		if request.Inn != "" {
			farm.Inn = &request.Inn
		}

		if request.Address != "" {
			farm.Address = request.Address
		}
		if request.Phone != "" {
			farm.Phone = &request.Phone
		}
		if request.Email != "" {
			farm.Email = &request.Email
		}
		if request.Description != "" {
			farm.Description = &request.Description
		}

		if err := db.Save(&farm).Error; err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

}
