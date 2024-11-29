package lactations

import (
	"cow_backend/models"
	"cow_backend/routes"

	"github.com/gin-gonic/gin"
)

// ListAccounts lists all existing accounts
//
//	@Summary      Get list of farms
//	@Description  Возращает конкретную лактацию
//	@Description  DOES NOT RETURN SUBOBJECTS
//	@Tags         Lactations
//	@Param        id    path     int  true  "id лактации"
//	@Produce      json
//	@Success      200  {object}   models.Lactation
//	@Failure      500  {object}  map[string]error
//	@Router       /lactations/{id} [get]
func (f *Lactations) Get() func(*gin.Context) {
	return routes.GenerateGetFunctionById[models.Lactation]()
}
