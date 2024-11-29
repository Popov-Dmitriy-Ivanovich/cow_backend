package sexes

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
//	@Tags         Sexes
//	@Param        id    path     int  true  "id of farm to return"
//	@Produce      json
//	@Success      200  {array}   models.Sex
//	@Failure      500  {object}  map[string]error
//	@Router       /sexes/{id} [get]
func (s *Sexes) Get() func(*gin.Context) {
	return routes.GenerateGetFunction[models.Sex]()
}
