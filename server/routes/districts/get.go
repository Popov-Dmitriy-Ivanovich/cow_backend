package districts

import (
	"cow_backend/models"
	"cow_backend/routes"

	"github.com/gin-gonic/gin"
)

// ListAccounts lists all existing accounts
//
//	@Summary      Get list of Districts
//	@Description  Возращает все районы. Разрешает отсутсвие фильтров
//	@Tags         Districtts
//	@Produce      json
//	@Success      200  {array}   models.DailyMilk
//	@Failure      500  {object}  map[string]error
//	@Router       /districts [get]
func (f *Districts) Get() func(*gin.Context) {
	return routes.GenerateGetFunctionByFilters[models.District](true)
}
