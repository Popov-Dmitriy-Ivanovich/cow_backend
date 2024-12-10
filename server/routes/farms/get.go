package farms

import (
	"cow_backend/models"
	"cow_backend/routes"

	"github.com/gin-gonic/gin"
)

// ListAccounts lists all existing accounts
//
//	@Summary      Get farm
//	@Description  GВозращает конкретную ферму
//
// @Tags         Farms
// @Param        id    path     int  true  "id of farm to return"
// @Produce      json
// @Success      200  {object}   models.Farm
// @Failure      500  {object}  map[string]error
// @Router       /farms/{id} [get]
func (f *Farms) GetByID() func(*gin.Context) {
	return routes.GenerateGetFunctionById[models.Farm]()
}

//	@Summary      Get list of farms
//	@Description  Возращает список ферм. Разрешает отсутсвие фильтров
//
// @Tags         Farms
// @Param        parrent_id    query     object  false  "ID более главной фермы, null для поиска хозяйств"
// @Produce      json
// @Success      200  {array}   models.Farm
// @Failure      500  {object}  map[string]error
// @Router       /farms [get]
func (f *Farms) GetByFilter() func(*gin.Context) {
	return func(c *gin.Context) {
		db := models.GetDb()
		farms := []models.Farm{}
		qres := db.Where("EXISTS (SELECT 1 FROM COWS WHERE cows.farm_group_id = farms.id)").Find(&farms)
		if qres.Error != nil {
			c.JSON(500, qres.Error)
		}
		c.JSON(200, farms)
	}

	// return routes.GenerateGetFunctionByFilters[models.Farm](true, "parrent_id")
}
