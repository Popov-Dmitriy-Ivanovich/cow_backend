package breeds

import (
	"cow_backend/models"
	"cow_backend/routes"

	"github.com/gin-gonic/gin"
)

// GetByID
// @Summary      Get breed
// @Description  Возращает породу.
// @Tags         Breeds
// @Param        id    path     int  true  "ID конкретной породы, если нужно вернуть одну."
// @Produce      json
// @Success      200  {object}   models.Breed
// @Failure      422  {object}   string
// @Failure      404  {object}   string
// @Router       /breeds/{id} [get]
func (f *Breeds) GetByID() func(*gin.Context) {
	return routes.GenerateGetFunctionById[models.Breed]()
}

// GetByFilter
// @Summary      Get list of breeds
// @Description  Возращает список всех пород
// @Tags         Breeds
// @Produce      json
// @Success      200  {array}   models.Breed
// @Failure      422  {object}   string
// @Failure      404  {object}   string
// @Router       /breeds [get]
func (f *Breeds) GetByFilter() func(*gin.Context) {
	return routes.GenerateGetFunctionByFilters[models.Breed](true)
}
