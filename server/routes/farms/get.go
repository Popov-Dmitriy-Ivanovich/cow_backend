package farms

import (
	"cow_backend/models"
	"cow_backend/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetByID
// @Summary      Get farm
// @Description  Возращает конкретную ферму (хозяйство, холдинг)
// @Tags         Farms
// @Param        id    path     int  true  "id of farm to return"
// @Produce      json
// @Success      200  {object}   models.Farm
// @Failure      422  {object}   string
// @Failure      404  {object}   string
// @Router       /farms/{id} [get]
func (f *Farms) GetByID() func(*gin.Context) {
	return routes.GenerateGetFunctionById[models.Farm]()
}

// GetByFilter
// @Summary      Get list of farms
// @Description  Возращает список ферм. Разрешает отсутсвие фильтров
// @Tags         Farms
// @Param        parrent_id    query     object  false  "ID более главной фермы, null для поиска холдингов"
// @Produce      json
// @Success      200  {array}   models.Farm
// @Failure      422  {object}   string
// @Failure      404  {object}   string
// @Router       /farms [get]
func (f *Farms) GetByFilter() func(*gin.Context) {
	return func(c *gin.Context) {
		roleId, exists := c.Get("RoleId")
		if !exists {
			c.JSON(http.StatusInternalServerError, "RoleId не найден в контексте")
			return
		}
		db := models.GetDb()
		farms := []models.Farm{}
		if roleId == 1 {
			farmId, exists := c.Get("FarmId")
			if !exists {
				c.JSON(http.StatusInternalServerError, "FarmId не найден в контексте")
				return
			}
			qres := db.Where("EXISTS (SELECT 1 FROM COWS WHERE cows.farm_group_id = farms.id) AND id = ?", farmId).Find(&farms)
			if qres.Error != nil {
				c.JSON(500, qres.Error)
			}

			c.JSON(200, farms)
		} else if roleId == 2 {
			regionId, exists := c.Get("RegionId")
			if !exists {
				c.JSON(http.StatusInternalServerError, "RegionId не найден в контексте")
				return
			}
			qres := db.
				Joins("JOIN districts AS d ON farms.district_id = d.id").
				Joins("JOIN regions AS r ON r.id = d.region_id").
				Where("EXISTS (SELECT 1 FROM COWS WHERE cows.farm_group_id = farms.id) AND r.id = ?", regionId).
				Find(&farms)
			if qres.Error != nil {
				c.JSON(500, qres.Error)
			}
			c.JSON(200, farms)
		} else {
			qres := db.Where("EXISTS (SELECT 1 FROM COWS WHERE cows.farm_group_id = farms.id)").Find(&farms)
			if qres.Error != nil {
				c.JSON(500, qres.Error)
			}
			c.JSON(200, farms)
		}
	}

	// return routes.GenerateGetFunctionByFilters[models.Farm](true, "parrent_id")
}
