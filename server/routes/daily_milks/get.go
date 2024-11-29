package dailymilks

import (
	"cow_backend/models"
	"cow_backend/routes"

	"github.com/gin-gonic/gin"
)

// ListAccounts lists all existing accounts
//
//	@Summary      Get list of DailyMilks
//	@Description  Get list of DailyMilks.
//	@Description  DOES NOT RETURN SUBOBJECTS
//	@Tags         DailyMilks
//	@Param        id    path     int  true  "id of farm to return"
//	@Param        lactation_id    query     int  false  "id lactation to search dailimilks"
//	@Produce      json
//	@Success      200  {array}   models.DailyMilk
//	@Failure      500  {object}  map[string]error
//	@Router       /dailyMilks/{id} [get]
func (f *DailyMilk) Get() func(*gin.Context) {
	return routes.GenerateGetFunction[models.DailyMilk]("lactation_id")
}
