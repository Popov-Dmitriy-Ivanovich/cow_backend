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
// @Success      200  {array}   RegionalResponse
// @Failure      422  {object}   string
// @Router       /analitics/total/{region_id}/regionalStatistics [get]
func (t Total) RegionalStatistics() gin.HandlerFunc {
	return func(c *gin.Context) {
		farmIds := []uint{}
		db := models.GetDb()
		farmsQuery := `EXISTS (SELECT 1 FROM districts WHERE farms.district_id = districts.id AND districts.region_id = ?)`

		if err := db.Model(&models.Farm{}).Where(farmsQuery, c.Param("region_id")).Pluck("id", &farmIds).Error; err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		statistics := []models.GaussianStatistics{}
		regIdUint, err := strconv.ParseUint(c.Param("region_id"), 10, 64)
		if err != nil {
			c.JSON(422, err.Error())
			return
		}
		if err := db.Where("farm_id in ? or region_id = ?", farmIds, regIdUint).Preload("Farm").Preload("Region").Find(&statistics).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}

		c.JSON(200, statistics)
	}
}

// RegionStatistics
// @Summary      Get statistics for region
// @Description  Еще не придумал что возвращает
// @Param        region_id    path     int  true  "регион по которому собиается статистика"
// @Tags         NEW_ANALITICS
// @Produce      json
// @Success      200  {object}   RegionalResponse
// @Failure      422  {object}   string
// @Router       /analitics/total/region/{region_id} [get]
func (t Total) RegionStatistics() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := models.GetDb()
		statistics := models.GaussianStatistics{}

		if err := db.First(&statistics, map[string]any{"region_id": c.Param("region_id")}).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}
		c.JSON(200, statistics)
	}
}

// FarmStatistics
// @Summary      Get statistics for region
// @Description  Еще не придумал что возвращает
// @Param        farm_id    path     int  true  "холдинг/хозяйство по которому собиается статистика"
// @Tags         NEW_ANALITICS
// @Produce      json
// @Success      200  {object}   RegionalResponse
// @Failure      422  {object}   string
// @Router       /analitics/total/farm/{farm_id} [get]
func (t Total) FarmStatistics() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := models.GetDb()
		statistics := models.GaussianStatistics{}

		if err := db.First(&statistics, map[string]any{"farm_id": c.Param("farm_id")}).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}
		c.JSON(200, statistics)
	}
}
