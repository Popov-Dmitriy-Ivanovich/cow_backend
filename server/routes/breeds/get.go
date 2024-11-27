package breeds

import (
	"cow_backend/models"
	"cow_backend/routes"

	"github.com/gin-gonic/gin"
)

// ListAccounts lists all existing accounts
//
//	@Summary      Get list of breeds
//	@Description  Get list of breeds.
//	@Description  DOES NOT RETURN SUBOBJECTS
//	@Tags         Breeds
//	@Param        id    query     int  false  "id of farm to return"
//	@Produce      json
//	@Success      200  {array}   models.Breed
//	@Failure      500  {object}  map[string]error
//	@Router       /breeds/get [get]
func (f *Breeds) Get() func(*gin.Context) {
	return routes.GenerateGetFunction[models.Breed]()
}
