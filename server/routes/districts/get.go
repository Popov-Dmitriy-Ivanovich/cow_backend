package districts

import (
	"cow_backend/models"
	"cow_backend/routes"

	"github.com/gin-gonic/gin"
)

// Get
// @Summary      Get list of Districts
// @Description  Возращает все районы. Разрешает отсутсвие фильтров
// @Tags         Districts
// @Produce      json
// @Success      200  {array}   models.DailyMilk
// @Failure      422  {object}   string
// @Failure      404  {object}   string
// @Router       /districts [get]
func (f *Districts) Get() func(*gin.Context) {
	return routes.GenerateGetFunctionByFilters[models.District](true)
}
