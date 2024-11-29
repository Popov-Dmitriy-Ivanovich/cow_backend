package districts

import (
	"cow_backend/models"
	"cow_backend/routes"

	"github.com/gin-gonic/gin"
)

// ListAccounts lists all existing accounts
//
//	@Summary      Get list of Districts
//	@Description  Get list of Districts.
//	@Description  DOES NOT RETURN SUBOBJECTS
//	@Tags         Districtts
//	@Param        id    path     int  true  "id of farm to return"
//	@Produce      json
//	@Success      200  {array}   models.DailyMilk
//	@Failure      500  {object}  map[string]error
//	@Router       /districts/{id} [get]
func (f *Districts) Get() func(*gin.Context) {
	return routes.GenerateGetFunction[models.District]()
}
