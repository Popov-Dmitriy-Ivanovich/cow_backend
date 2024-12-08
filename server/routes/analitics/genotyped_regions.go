package analitics

import (
	"cow_backend/filters"
	"cow_backend/filters/cows_filter"
	"cow_backend/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// type byRegionStatistics struct {
// 	genotypedStatistics
// 	RegionID uint
// }

// type byRegionKeys struct {
// 	Name string
// 	ID   uint
// }

// @Summary      Get list of years
// @Description  Возращает словарь регион - количество живых коров, количество генотипированных
// @Tags         Analitics
// @Param        year    path     int  true  "год за который собирается статистика"
// @Param        filter    body     cows_filter.CowsFilter  true  "applied filters"
// @Produce      json
// @Success      200  {array}   map[string]byRegionStatistics
// @Failure      500  {object}  map[string]error
// @Router       /analitics/genotyped/{year}/regions [post]
func (g Genotyped) RegionsPost() func(*gin.Context) {
	return func(c *gin.Context) {
		filterData := cows_filter.CowsFilter{}
		if err := c.ShouldBindJSON(&filterData); err != nil {
			c.JSON(422, err.Error())
		}
		aliveFilter := filterData
		genotypedFilter := filterData

		genotypedFilter.IsGenotyped = new(bool)
		*genotypedFilter.IsGenotyped = true
		keys := []byRegionKeys{}
		db := models.GetDb()
		yearStr := c.Param("year")
		yearInt, err := strconv.ParseInt(yearStr,10,64)
		if err != nil {
			c.JSON(422, err.Error())
			return
		}
		db.Model(&models.Region{}).Debug().Where(
			"EXISTS(SELECT 1 FROM districts where districts.region_id = regions.id AND " +
				"EXISTS(SELECT 1 FROM farms where farms.district_id = districts.id AND " +
				" EXISTS (SELECT 1 FROM cows WHERE (cows.farm_id = farms.id OR cows.farm_group_id = farms.id) AND " +
				" (cows.death_date IS NULL OR cows.death_date < ?) AND cows.birth_date < ? AND",
				" EXISTS (SELECT 1 FROM genetics where genetics.cow_id = cows.id))))",
				time.Date(int(yearInt)+1,1,1,0,0,0,0,time.UTC),
				time.Date(int(yearInt)+1,1,1,0,0,0,0,time.UTC)).Find(&keys)

		result := make(map[string]byRegionStatistics)
		for _, key := range keys {
			aliveCowQuery := db.Model(&models.Cow{})
			aliveCowFilter := cows_filter.NewCowFilteredModel(aliveFilter, aliveCowQuery)
			aliveCowFilter.Params["year"] = c.Param("year")
			aliveCowFilter.Params["region"] = strconv.FormatUint(uint64(key.ID), 10)

			genotypedCowQuery := db.Model(&models.Cow{})
			genotypedCowFilter := cows_filter.NewCowFilteredModel(genotypedFilter, genotypedCowQuery)
			genotypedCowFilter.Params["year"] = c.Param("year")
			genotypedCowFilter.Params["region"] = strconv.FormatUint(uint64(key.ID), 10)
			if err := filters.ApplyFilters(aliveCowFilter,
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
				cows_filter.AliveInYear{},
				cows_filter.LiveInRegion{}); err != nil {
				c.JSON(422, err.Error())
				return
			}
			if err := filters.ApplyFilters(genotypedCowFilter,
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
				cows_filter.AliveInYear{},
				cows_filter.LiveInRegion{}); err != nil {
				c.JSON(422, err.Error())
				return
			}

			alive := int64(0)
			genotyped := int64(0)

			aliveCowFilter.GetQuery().Debug().Count(&alive)
			genotypedCowFilter.GetQuery().Debug().Count(&genotyped)
			result[key.Name] = byRegionStatistics{
				genotypedStatistics: genotypedStatistics{
					Alive:     alive,
					Genotyped: genotyped,
				},
				RegionID: key.ID,
			}
		}
		c.JSON(200, result)
	}
}