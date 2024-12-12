package analitics

import (
	"cow_backend/filters"
	"cow_backend/filters/cows_filter"
	"cow_backend/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CheckMilks struct{}

func (cm *CheckMilks) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/checkMilks")
	apiGroup.POST("/years", cm.ByYear())
	apiGroup.POST("/:year/byRegion", cm.ByRegion())
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
		*filterData.ControlMilkingDateTo = "4000-01-01"

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

type cmByRegionStatistics struct {
	Milk     float64
	Fat      float64
	Protein  float64
	CmCount  uint
	CowCount uint
	RegionId uint
}

// @Summary      Get list of years
// @Description  Возращает словарь год - количеств генотипированных коров, по ключу -1 генотипированные за все годы
// @Param        year    path     int  true  "год за который собирается статистика"
// @Param        filter    body     cows_filter.CowsFilter  true  "applied filters"
// @Tags         Analitics
// @Produce      json
// @Success      200  {array}   map[int]bool
// @Failure      500  {object}  map[string]error
// @Router       /analitics/checkMilks/{year}/byRegion [post]
func (cm CheckMilks) ByRegion() func(*gin.Context) {
	return func(c *gin.Context) {
		filterData := cows_filter.CowsFilter{}
		if err := c.ShouldBindJSON(&filterData); err != nil {
			c.JSON(422, err.Error())
		}
		filterData.ControlMilkingDateFrom = new(string)
		filterData.ControlMilkingDateTo = new(string)

		*filterData.ControlMilkingDateFrom = "0001-01-01"
		*filterData.ControlMilkingDateTo = "4000-01-01"

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
		result := map[string]cmByRegionStatistics{}
		yearStr := c.Param("year")
		yearInt, err := strconv.ParseInt(yearStr, 10, 64)
		if err != nil {
			c.JSON(422, err.Error())
			return
		}
		cowCount := uint(0)
		for _, id := range cowIds {
			dbCow := models.Cow{}
			db.Preload("Lactation").
				Preload("Lactation.CheckMilks").
				Preload("FarmGroup").
				Preload("FarmGroup.District").
				Preload("FarmGroup.District.Region").First(&dbCow, id)
			cmCount := uint(0)
			cowCount++
			for _, lac := range dbCow.Lactation {
				for _, cm := range lac.CheckMilks {
					if cm.CheckDate.Year() != int(yearInt) {
						continue
					}
					cmCount++
					val, ok := result[dbCow.FarmGroup.District.Region.Name]
					if ok {
						val.Milk += cm.Milk
						val.Fat += cm.Fat
						val.Protein += cm.Protein

					} else {
						val = cmByRegionStatistics{}
						val.Milk = cm.Milk
						val.Fat = cm.Fat
						val.Protein = cm.Protein
						val.RegionId = dbCow.FarmGroup.District.RegionId
					}
					result[dbCow.FarmGroup.District.Region.Name] = val
				}
			}
			if val, ok := result[dbCow.FarmGroup.District.Region.Name]; ok && cmCount != 0 {
				val.Milk = val.Milk / float64(cmCount)
				val.Fat = val.Fat / float64(cmCount)
				val.Protein = val.Protein / float64(cmCount)
				result[dbCow.FarmGroup.District.Region.Name] = val
			}
		}
		if cowCount != 0 {
			for key, val := range result {
				val.Milk = val.Milk / float64(cowCount)
				val.Fat = val.Fat / float64(cowCount)
				val.Protein = val.Protein / float64(cowCount)
				result[key] = val
			}
		}

		c.JSON(200, result)
	}
}
