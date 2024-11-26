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
//	@Param        id    query     int  false  "id of farm to return"
//	@Produce      json
//	@Success      200  {array}   models.DailyMilk
//	@Failure      500  {object}  map[string]error
//	@Router       /dailyMilks/get [get]
func (f *DailyMilk) Get() func(*gin.Context) {
	return routes.GenerateGetFunction[models.DailyMilk]()
}
