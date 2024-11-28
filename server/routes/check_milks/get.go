package checkmilks

import (
	"cow_backend/models"
	"cow_backend/routes"

	"github.com/gin-gonic/gin"
)

// ListAccounts lists all existing accounts
//
//	@Summary      Get list of checkMilks
//	@Description  Возращает список контрольный доек
//	@Tags         CheckMilks
//	@Param        id    query     int  false  "id контрольной дойки"
//	@Param 		  lactation_id query int false "id лактации, для корой ищутся котнольные дойки"
//	@Produce      json
//	@Success      200  {array}   models.CheckMilk
//	@Failure      500  {object}  map[string]error
//	@Router       /checkMilks/get [get]
func (f *CheckMilks) Get() func(*gin.Context) {
	return routes.GenerateGetFunction[models.CheckMilk]("lactation_id")
}
