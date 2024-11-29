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
//	@Param        id    path     int  true  "id of farm to return"
//	@Param        cow_id    query     int  false  "id of cow for wich lactations should be find"
//	@Produce      json
//	@Success      200  {array}   models.Lactation
//	@Failure      500  {object}  map[string]error
//	@Router       /lactations/{id} [get]
func (f *Lactations) Get() func(*gin.Context) {
	return routes.GenerateGetFunction[models.Lactation]("cow_id")
}
