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
// @Deprecated   Рут является deprecated, никакие изменения в его работу вноситься не будут.
// @Deprecated   Если появится необходимость как либо модифицировать его работу - лучше написать новый
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
			//qres := db.Where("EXISTS (SELECT 1 FROM COWS WHERE cows.farm_group_id = farms.id) AND id = ?", farmId).Find(&farms)
			qres := db.Where(map[string]any{"parrent_id": nil, "id": farmId, "type": []uint{1, 2}}).Find(&farms)
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
				Where("(parrent_id is NULL) AND r.id = ? AND type in (1,2)", regionId).
				Find(&farms)
			if qres.Error != nil {
				c.JSON(500, qres.Error)
			}
			c.JSON(200, farms)
		} else {
			qres := db.Where("parrent_id is NULL AND type in (1,2)").Find(&farms)
			if qres.Error != nil {
				c.JSON(500, qres.Error)
			}
			c.JSON(200, farms)
		}
	}

	// return routes.GenerateGetFunctionByFilters[models.Farm](true, "parrent_id")
}

// GetFarms
// @Summary Get list of Farms
// @Description Возвращает список всех ферм
// @Tags Farms
// @Produce json
// @Success 200 {array} models.Farm
// @Failure      500  {object}   string
// @Router       /farms/farm [get]
func (f *Farms) GetFarms() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := models.GetDb()
		farms := []models.Farm{}
		if err := db.Find(&farms, map[string]any{"type": 3}).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}
		c.JSON(200, farms)
	}
}

// GetHoz
// @Summary Get list of Hoz
// @Description Возвращает список всех хозяйств
// @Tags Farms
// @Produce json
// @Success 200 {array} models.Farm
// @Failure      500  {object}   string
// @Router       /farms/hoz [get]
func (f *Farms) GetHoz() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := models.GetDb()
		farms := []models.Farm{}
		if err := db.Find(&farms, map[string]any{"type": 2}).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}
		c.JSON(200, farms)
	}
}

// GetHoldings
// @Summary Get list of holdings
// @Description Возвращает список всех холдингов
// @Tags Farms
// @Produce json
// @Success 200 {array} models.Farm
// @Failure      500  {object}   string
// @Router       /farms/hold [get]
func (f *Farms) GetHoldings() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := models.GetDb()
		farms := []models.Farm{}
		if err := db.Find(&farms, map[string]any{"type": 1}).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}
		c.JSON(200, farms)
	}
}
