package regions

import (
	"cow_backend/models"
	"cow_backend/routes"
	"sort"

	"github.com/gin-gonic/gin"
)

// GetByFilter
// @Summary      Get list of regions
// @Description   Возращает все регионы
// @Tags         Regions
// @Produce      json
// @Success      200  {array}   models.Region
// @Failure      500  {object}  string
// @Router       /regions [get]
func (r *Regions) GetByFilter() func(*gin.Context) {
	return routes.GenerateGetFunctionByFilters[models.Region](true)
}

// GetByID
// @Summary      Get concrete region
// @Description   Возращает регион
// @Tags         Regions
// @Param        id    path     int  true  "id региона"
// @Produce      json
// @Success      200  {object}   models.Region
// @Failure      500  {object}   string
// @Router       /regions/{id} [get]
func (r *Regions) GetByID() func(*gin.Context) {
	return routes.GenerateGetFunctionById[models.Region]()
}

// News
// @Summary      Get region's news
// @Description  Возращает новости региона
// @Tags         Regions
// @Param        id    path     int  true  "id региона"
// @Produce      json
// @Success      200  {array}   models.News
// @Failure      500  {object}  string
// @Router       /regions/{id}/news [get]
func (r *Regions) News() func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		db := models.GetDb()
		region := models.Region{}
		if err := db.Preload("News").First(&region, id).Error; err != nil {
			c.JSON(404, err.Error())
		}
		sort.Slice(region.News, func(i, j int) bool {
			return region.News[i].Date.After(region.News[j].Date.Time)
		})
		c.JSON(200, region.News[0:5])
	}
}
