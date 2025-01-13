package lactations

import (
	"cow_backend/models"
	"cow_backend/routes"

	"github.com/gin-gonic/gin"
)

// Get
//
//	@Summary      Get list of farms
//	@Description  Возращает конкретную лактацию
//	@Tags         Lactations
//	@Param        id    path     int  true  "id лактации"
//	@Produce      json
//	@Success      200  {object}   models.Lactation
//	@Failure      422  {object}   string
//	@Failure      404  {object}   string
//	@Router       /lactations/{id} [get]
func (f *Lactations) Get() func(*gin.Context) {
	return routes.GenerateGetFunctionById[models.Lactation]()
}
