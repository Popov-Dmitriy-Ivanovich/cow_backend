package lactations

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
//	@Tags         Lactations
//	@Param        id    query     int  false  "id of farm to return"
//	@Produce      json
//	@Success      200  {array}   models.Lactation
//	@Failure      500  {object}  map[string]error
//	@Router       /lactations/get [get]
func (f *Lactations) Get() func(*gin.Context) {
	return routes.GenerateGetFunction[models.Lactation]()
}
