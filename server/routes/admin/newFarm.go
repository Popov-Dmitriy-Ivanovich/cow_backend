package admin

import (
	"cow_backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Admin) NewFarm() func(*gin.Context) {
	return func(c *gin.Context) {
		var request struct {
			HozNuber   string `json:"hoznumber"`
			DistrictID string `json:"district"`
			ParrentId  string `json:"parentid"`
			Fullname   string `json:"fullname"`
			Address    string `json:"address"`
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

		var parrId *uint
		if request.ParrentId != "" {
			parr, err := strconv.ParseUint(request.ParrentId, 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID хозяйства"})
				return
			}
			value := uint(parr)
			parrId = &value
		} else {
			parrId = nil
		}

		farms := models.Farm{
			HozNumber:  &request.HozNuber,
			DistrictId: distId,
			ParrentId:  parrId,
			Type:       3,
			Name:       request.Fullname,
			NameShort:  request.Fullname,
			Address:    request.Address,
		}
		db := models.GetDb()

		// обновление последовательности
		if err := updateSequenceFarms(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении последовательности: " + err.Error()})
			return
		}

		if err := db.Create(&farms).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при добавлении фермы: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Новая ферма создана"})
	}
}
