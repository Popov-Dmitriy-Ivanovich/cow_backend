package analitics

import (
	"cow_backend/filters"
	"cow_backend/filters/cows_filter"
	"cow_backend/models"

	"github.com/gin-gonic/gin"
)

type CheckMilks struct{}

func (cm *CheckMilks) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/checkMilks")
	apiGroup.POST("/years", cm.ByYear())
}

// @Summary      Get list of years
// @Description  Возращает словарь год - количеств генотипированных коров, по ключу -1 генотипированные за все годы
// @Param        filter    body     cows_filter.CowsFilter  true  "applied filters"
// @Tags         Analitics
// @Produce      json
// @Success      200  {array}   map[int]bool
// @Failure      500  {object}  map[string]error
// @Router       /analitics/checkMilks/years [post]
func (cm CheckMilks) ByYear() func(*gin.Context) {
	return func(c *gin.Context) {
		filterData := cows_filter.CowsFilter{}
		if err := c.ShouldBindJSON(&filterData); err != nil {
			c.JSON(422, err.Error())
		}
		filterData.ControlMilkingDateFrom = new(string)
		filterData.ControlMilkingDateTo = new(string)

		*filterData.ControlMilkingDateFrom = "0001-01-01"
		*filterData.ControlMilkingDateFrom = "40000-01-01"

		db := models.GetDb()
		cmCowQuery := db.Model(models.Cow{})
		cmCowFilter := cows_filter.NewCowFilteredModel(filterData, cmCowQuery)
		if err := filters.ApplyFilters(cmCowFilter,
			cows_filter.ByAbort{},
			cows_filter.ByAnyIllneses{},
			cows_filter.ByBirkingDate{},
			cows_filter.ByBreed{},
			cows_filter.ByBrithDate{},
			cows_filter.ByCalvingDate{},
			cows_filter.ByControlMilkingDate{},
			cows_filter.ByCreatedAt{},
			cows_filter.ByDeath{},
			cows_filter.ByDepartDate{},
			cows_filter.ByExterior{},
			cows_filter.ByHoz{},
			cows_filter.ByIllDate{},
			cows_filter.ByInbrindingCoeffByFamily{},
			cows_filter.ByInbrindingCoeffByGenotype{},
			cows_filter.ByIsGenotyped{},
			cows_filter.ByInsemenationDate{},
			cows_filter.BySearchString{},
			cows_filter.BySex{},
			cows_filter.ByStillBorn{},
			cows_filter.ByTwins{},
			cows_filter.ByMonogeneticIllnesses{},
		); err != nil {
			c.JSON(422, err.Error())
			return
		}
		cowIds := []uint{}
		cmCowFilter.GetQuery().Pluck("id", &cowIds)
		result := map[int]bool{}

		for _, id := range cowIds {
			dbCow := models.Cow{}
			db.Preload("Lactation").Preload("Lactation.CheckMilks").First(&dbCow, id)
			for _, lac := range dbCow.Lactation {
				for _, cm := range lac.CheckMilks {
					result[cm.CheckDate.Year()] = true
				}
			}
		}
		c.JSON(200, result)
	}
}
