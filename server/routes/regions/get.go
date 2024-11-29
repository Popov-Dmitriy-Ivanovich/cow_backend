package regions

import (
	"cow_backend/models"
	"cow_backend/routes"

	"github.com/gin-gonic/gin"
)

// @Summary      Get list of regions
// @Description   Возращает все регионы
// @Tags         Regions
// @Produce      json
// @Success      200  {array}   models.Region
// @Failure      500  {object}  map[string]error
// @Router       /regions [get]
func (r *Regions) GetByFilter() func(*gin.Context) {
	return routes.GenerateGetFunctionByFilters[models.Region](true)
}

// @Summary      Get list of regions
// @Description   Возращает регион
// @Tags         Regions
// @Param        id    path     int  true  "id региона"
// @Produce      json
// @Success      200  {object}   models.Region
// @Failure      500  {object}  map[string]error
// @Router       /regions/{id} [get]
func (r *Regions) GetByID() func(*gin.Context) {
	return routes.GenerateGetFunctionById[models.Region]()
}
