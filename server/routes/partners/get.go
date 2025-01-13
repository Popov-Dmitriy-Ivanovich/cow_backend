package partners

import (
	"cow_backend/models"
	"cow_backend/routes"

	"github.com/gin-gonic/gin"
)

// Get
// @Summary      Get list of partners
// @Description  Возращает список партнеров
// @Tags         Partners
// @Produce      json
// @Success      200  {array}   models.Partner
// @Failure      422  {object}   string
// @Failure      404  {object}   string
// @Router       /partners [get]
func (s *Partners) Get() func(*gin.Context) {
	return routes.GenerateGetFunctionByFilters[models.Partner](true)
}
