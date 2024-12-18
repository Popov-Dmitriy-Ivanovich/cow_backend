package admin

import (
	"cow_backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Admin) UpdateFarmPage(typeHoz int) func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		db := models.GetDb()
		regions := []models.Region{}
		districts := []models.District{}
		db.Order("name").Find(&regions)
		db.Order("name").Find(&districts)

		farm := models.Farm{}
		if err := db.
			Order("name").
			Preload("Parrent").
			Preload("District.Region").
			First(&farm, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Холдинг с таким ID не найден"})
			return
		}

		holds := []models.Farm{}
		hoz := []models.Farm{}
		db.Order("name").Where("type = 1").Find(&holds)
		db.Order("name").Where("type = 2").Find(&hoz)

		AdminPages := map[int]string{
			1: "AdminUpdateHoldingPage.tmpl",
			2: "AdminUpdateHozPage.tmpl",
			3: "AdminUpdateFarmPage.tmpl",
		}

		c.HTML(http.StatusOK, AdminPages[typeHoz], gin.H{
			"title":     "Редактирование",
			"farm":      farm,
			"holds":     holds,
			"hoz":       hoz,
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
			ParrentId   string `json:"parrent"`
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

		if request.ParrentId != "" {
			parrent, err := strconv.ParseUint(request.ParrentId, 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID холдинга"})
			}
			parrId := uint(parrent)
			farm.ParrentId = &parrId
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
