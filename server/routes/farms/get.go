package farms

import (
	"cow_backend/models"
	"cow_backend/routes"

	"github.com/gin-gonic/gin"
)

// ListAccounts lists all existing accounts
//
//	@Summary      Get list of farms
//	@Description  Get list of farms.
//	@Description  DOES NOT RETURN SUBOBJECTS
//	@Tags         Farms
//	@Param        id    path     int  true  "id of farm to return"
//	@Produce      json
//	@Success      200  {array}   models.Farm
//	@Failure      500  {object}  map[string]error
//	@Router       /farms/{id} [get]
func (f *Farms) Get() func(*gin.Context) {
	return routes.GenerateGetFunction[models.Farm]()
}
