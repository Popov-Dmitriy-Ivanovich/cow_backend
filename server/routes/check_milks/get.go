package checkmilks

import (
	"cow_backend/models"
	"cow_backend/routes"

	"github.com/gin-gonic/gin"
)

// ListAccounts lists all existing accounts
//
//	@Summary      Get list of checkMilks
//	@Description  Get list of CheckMilks.
//	@Description  DOES NOT RETURN SUBOBJECTS
//	@Tags         CheckMilks
//	@Param        id    query     int  false  "id of farm to return"
//	@Produce      json
//	@Success      200  {array}   models.CheckMilk
//	@Failure      500  {object}  map[string]error
//	@Router       /checkMilks/get [get]
func (f *CheckMilks) Get() func(*gin.Context) {
	return routes.GenerateGetFunction[models.CheckMilk]()
}
