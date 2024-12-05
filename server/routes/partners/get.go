package partners

import (
	"cow_backend/models"
	"cow_backend/routes"

	"github.com/gin-gonic/gin"
)

// ListAccounts lists all existing accounts
//
//	@Summary      Get list of partners
//	@Description  Возращает список партнеров
//	@Tags         Partners
//	@Produce      json
//	@Success      200  {array}   models.Partner
//	@Failure      500  {object}  map[string]error
//	@Router       /partners [get]
func (s *Partners) Get() func(*gin.Context) {
	return routes.GenerateGetFunctionByFilters[models.Partner](true)
}
