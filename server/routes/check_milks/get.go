package checkmilks

import (
	"cow_backend/models"
	"cow_backend/routes"

	"github.com/gin-gonic/gin"
)

// GetByID
// @Summary      Get checkMilk
// @Description  Возращает контрольную дойку
// @Tags         CheckMilks
// @Param        id    path     int  false  "id контрольной дойки"
// @Produce      json
// @Success      200  {object}   models.CheckMilk
// @Failure      422  {object}   string
// @Failure      404  {object}   string
// @Router       /checkMilks/{id} [get]
func (f *CheckMilks) GetByID() func(*gin.Context) {
	return routes.GenerateGetFunctionById[models.CheckMilk]()
}

// GetByFilter
// @Summary      Get list of checkMilks
// @Description  Возращает список контрольных доек. Без фильтра нельзя, т.к. слишком много контрольных доений
// @Tags         CheckMilks
// @Param 		 lactation_id query int false "id лактации, для корой ищутся котнольные дойки"
// @Produce      json
// @Success      200  {array}   models.CheckMilk
// @Failure      422  {object}   string
// @Failure      404  {object}   string
// @Router       /checkMilks [get]
func (f *CheckMilks) GetByFilter() func(*gin.Context) {
	return routes.GenerateGetFunctionByFilters[models.CheckMilk](false, "lactation_id")
}
