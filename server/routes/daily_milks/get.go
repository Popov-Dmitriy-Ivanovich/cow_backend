package dailymilks

import (
	"cow_backend/models"
	"cow_backend/routes"

	"github.com/gin-gonic/gin"
)

// ListAccounts lists all existing accounts
//
//	@Summary      Get list of DailyMilks
//	@Description  Возращает конкретную дойку коровы.
//	@Tags         DailyMilks
//	@Param        id    path     int  true  "id of farm to return"
//	@Produce      json
//	@Success      200  {object}   models.DailyMilk
//	@Failure      500  {object}  map[string]error
//	@Router       /dailyMilks/{id} [get]
func (f *DailyMilk) GetByID() func(*gin.Context) {
	return routes.GenerateGetFunctionById[models.DailyMilk]()
}

// @Summary      Get list of DailyMilks
// @Description  Возвращает дойки удовлетворяющие фильтрам
// @Tags         DailyMilks
// @Param        lactation_id    query     int  false  "id lactation to search dailimilks"
// @Produce      json
// @Success      200  {array}   models.DailyMilk
// @Failure      500  {object}  map[string]error
// @Router       /dailyMilks [get]
func (f *DailyMilk) GetByFilter() func(*gin.Context) {
	return routes.GenerateGetFunctionByFilters[models.DailyMilk](false, "lactation_id")
}
