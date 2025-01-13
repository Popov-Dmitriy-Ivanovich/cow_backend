package dailymilks

import (
	"cow_backend/models"
	"cow_backend/routes"

	"github.com/gin-gonic/gin"
)

// GetByID
// @Summary      Get dailyMilk
// @Description  Возращает конкретную дойку коровы.
// @Tags         DailyMilks
// @Param        id    path     int  true  "id of dailymilk"
// @Produce      json
// @Success      200  {object}   models.DailyMilk
// @Failure      422  {object}   string
// @Failure      404  {object}   string
// @Router       /dailyMilks/{id} [get]
func (f *DailyMilk) GetByID() func(*gin.Context) {
	return routes.GenerateGetFunctionById[models.DailyMilk]()
}

// GetByFilter
// @Summary      Get list of DailyMilks
// @Description  Возвращает дойки удовлетворяющие фильтрам. Без фильтра не работает, т.к. слишком много доений
// @Tags         DailyMilks
// @Param        lactation_id    query     int  false  "id lactation to search dailimilks"
// @Produce      json
// @Success      200  {array}   models.DailyMilk
// @Failure      422  {object}   string
// @Failure      404  {object}   string
// @Router       /dailyMilks [get]
func (f *DailyMilk) GetByFilter() func(*gin.Context) {
	return routes.GenerateGetFunctionByFilters[models.DailyMilk](false, "lactation_id")
}
