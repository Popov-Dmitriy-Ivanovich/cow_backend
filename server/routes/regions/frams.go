package regions

import (
	"cow_backend/models"

	"github.com/gin-gonic/gin"
)

// GetFarms
//
//	@Summary      Get farm by region id
//	@Description  Возращает все фермы в регионе (Данные представлены как словарь с единственным ключем "farms")
//
// @Tags         Regions
// @Param        id    path     int  true  "id of region"
// @Produce      json
// @Success      200  {array}   models.Farm
// @Failure      500  {object}  string
// @Router       /regions/{id}/getFarms [get]
func (f *Regions) GetFarms() func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		db := models.GetDb()
		farms := []models.Farm{}
		dists := []models.District{}
		if err := db.Where("region_id = ?", id).Find(&dists).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}

		if len(dists) == 0 {
			c.JSON(200, gin.H{"farms": farms})
			return
		}

		var districtIDs []uint

		for _, dist := range dists {
			districtIDs = append(districtIDs, dist.ID)
		}

		if err := db.Where("district_id IN ?", districtIDs).Find(&farms).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}
		c.JSON(200, gin.H{"farms": farms})
	}
}
