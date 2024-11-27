package regions

import (
	"cow_backend/models"
	"cow_backend/routes"

	"github.com/gin-gonic/gin"
)

// ListAccounts lists all existing accounts
//
//	@Summary      Get list of regions
//	@Description  Get list of regions.
//	@Tags         Regions
//	@Param        id    query     int  false  "id of region to return"
//	@Produce      json
//	@Success      200  {array}   models.Region
//	@Failure      500  {object}  map[string]error
//	@Router       /regions/get [get]
func (r *Regions) Get() func(*gin.Context) {
	return routes.GenerateGetFunction[models.Region]()
}
