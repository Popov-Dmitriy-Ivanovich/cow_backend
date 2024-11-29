package sexes

import (
	"cow_backend/models"
	"cow_backend/routes"

	"github.com/gin-gonic/gin"
)

// ListAccounts lists all existing accounts
//
//	@Summary      Get list of sexes
//	@Description  Возращает список полов
//	@Tags         Sexes
//	@Produce      json
//	@Success      200  {array}   models.Sex
//	@Failure      500  {object}  map[string]error
//	@Router       /sexes [get]
func (s *Sexes) Get() func(*gin.Context) {
	return routes.GenerateGetFunctionByFilters[models.Sex](true)
}
