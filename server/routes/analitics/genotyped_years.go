package analitics

import (
	"cow_backend/filters"
	"cow_backend/filters/cows_filter"
	"cow_backend/models"

	"github.com/gin-gonic/gin"
)

// YearsPost
// @Summary      Get list of years
// @Description  Возращает словарь год - количеств генотипированных коров, по ключу -1 генотипированные за все годы
// @Param        filter    body     cows_filter.CowsFilter  true  "applied filters"
// @Tags         Analytics(GeneticFilters)
// @Produce      json
// @Success      200  {array}   map[int]bool
// @Failure      500  {object}  map[string]error
// @Router       /analitics/genotyped/years [post]
func (g Genotyped) YearsPost() func(*gin.Context) {
	return func(c *gin.Context) {
		filterData := cows_filter.CowsFilter{}
		if err := c.ShouldBindJSON(&filterData); err != nil {
			c.JSON(422, err.Error())
		}
		tmp := bool(true)
		filterData.IsGenotyped = &tmp
		filterData.IsDead = new(bool)
		*filterData.IsDead = false

		db := models.GetDb()
		query := db.Model(&models.Cow{}).Where("approved <> -1")
		cfm := cows_filter.NewCowFilteredModel(filterData, query)
		if err := filters.ApplyFilters(cfm, cows_filter.ALL_FILTERS...); err != nil {
			c.JSON(422, err.Error())
			return
		}
		genotypedCowsIds := []string{}
		cfm.GetQuery().Debug().Pluck("id", &genotypedCowsIds)
		genDates := []models.DateOnly{}
		db.Model(models.Genetic{}).Debug().Where("cow_id IN ?", genotypedCowsIds).Pluck("result_date", &genDates)
		result := map[int]bool{}
		for _, date := range genDates {
			result[date.Year()] = true
		}
		c.JSON(200, result)
	}
}
