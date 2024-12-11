package admin

import (
	"cow_backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Admin) NewHoz() func(*gin.Context) {
	return func(c *gin.Context) {
		var request struct {
			HozNuber    string `json:"hoznumber"`
			DistrictID  string `json:"district"`
			ParrentId   string `json:"parentid"`
			Fullname    string `json:"fullname"`
			Name        string `json:"name"`
			INN         string `json:"inn"`
			Address     string `json:"address"`
			Phone       string `json:"phone"`
			Email       string `json:"email"`
			Description string `json:"description"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		dist, err := strconv.ParseUint(request.DistrictID, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID района"})
		}
		distId := uint(dist)

		parr, err := strconv.ParseUint(request.ParrentId, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID холдинга"})
		}
		parrId := uint(parr)

		hozs := models.Farm{
			HozNumber:   &request.HozNuber,
			DistrictId:  distId,
			ParrentId:   &parrId,
			Type:        2,
			Name:        request.Fullname,
			NameShort:   request.Name,
			Inn:         &request.INN,
			Address:     request.Address,
			Phone:       &request.Phone,
			Email:       &request.Email,
			Description: &request.Description,
		}
		db := models.GetDb()

		// обновление последовательности
		if err := updateSequenceFarms(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении последовательности: " + err.Error()})
			return
		}

		if err := db.Create(&hozs).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при добавлении хозяйства: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Новое хозяйство создано"})
	}
}
