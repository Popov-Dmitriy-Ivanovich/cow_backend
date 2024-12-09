package admin

import (
	"cow_backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Admin) NewHolding() func(*gin.Context) {
	return func(c *gin.Context) {
		var request struct {
			HozNuber    string `json:"hoznumber"`
			DistrictID  string `json:"district"`
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
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID фермы"})
			return
		}
		distId := uint(dist)

		hold := models.Farm{
			HozNumber:   &request.HozNuber,
			DistrictId:  distId,
			Type:        1,
			Name:        request.Fullname,
			NameShort:   request.Name,
			Inn:         &request.INN,
			Address:     request.Address,
			Phone:       &request.Phone,
			Email:       &request.Email,
			Description: &request.Description,
		}
		db := models.GetDb()

		if err := db.Create(&hold).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании холдинга: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Новый холдинг создан"})
	}
}
