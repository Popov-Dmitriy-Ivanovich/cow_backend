package analitics

import (
	"cow_backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Total struct{}

type RegionalResponse struct {
	models.GaussianStatistics
}

type GaussianResponse struct {
	models.GaussianStatistics
}

// RegionalStatistics
// @Summary      Get statistics for region
// @Description  Еще не придумал что возвращает
// @Param        region_id    path     int  true  "регион по которому собиается статистика"
// @Tags         NEW_ANALITICS
// @Produce      json
// @Success      200  {object}   map[string]RegionalResponse
// @Failure      422  {object}   string
// @Router       /analitics/total/{region_id}/regionalStatistics [get]
func (t Total) RegionalStatistics() gin.HandlerFunc {
	return func(c *gin.Context) {
		farmIds := []uint{}
		db := models.GetDb()
		farmsQuery := `EXISTS (SELECT 1 FROM districts WHERE farms.district_id = districts.id AND districts.region_id = ?)`

		if err := db.Model(&models.Farm{}).Where(farmsQuery, c.Param("region_id")).Pluck("id", &farmIds).Error; err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}

		statistics := []models.GaussianStatistics{}
		regIdUint, err := strconv.ParseUint(c.Param("region_id"), 10, 64)
		if err != nil {
			c.JSON(422, err.Error())
		}
		if err := db.Where("farm_id in ? or region_id = ?", farmIds, regIdUint).Preload("Farm").Preload("Region").Find(&statistics).Error; err != nil {
			c.JSON(500, err.Error())
		}

		c.JSON(200, statistics)
	}
}
