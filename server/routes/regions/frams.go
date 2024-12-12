package regions

import (
	"cow_backend/models"

	"github.com/gin-gonic/gin"
)

// ListAccounts lists all existing accounts
//
//	@Summary      Get farm by region id
//	@Description  Возращает все фермы в регионе
//
// @Tags         Regions
// @Param        id    path     int  true  "id of region"
// @Produce      json
// @Success      200  {object}   models.Farm
// @Failure      500  {object}  map[string]error
// @Router       /regions/{id}/getFarms [get]
func (f *Regions) GetFarms() func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		db := models.GetDb()
		farms := []models.Farm{}
		dist := models.District{}
		if err := db.Where("region_id = ?", id).First(&dist).Error; err != nil {
			c.JSON(500, err.Error)
			return
		}
		if err := db.Where("district_id = ?", dist.ID).Find(&farms).Error; err != nil {
			c.JSON(500, err.Error)
			return
		}
		c.JSON(200, gin.H{"farms": farms})
	}
}
