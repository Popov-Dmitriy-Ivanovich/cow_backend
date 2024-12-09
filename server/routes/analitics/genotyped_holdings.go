package analitics

import (
	"cow_backend/filters"
	"cow_backend/filters/cows_filter"
	"cow_backend/models"
	"log"
	// "strconv"

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
// @Param        district    path     int  true  "район за который собирается статистика"
// @Param        filter    body     cows_filter.CowsFilter  true  "applied filters"
// @Produce      json
// @Success      200  {array}   map[string]byHoldStatistics
// @Failure      500  {object}  map[string]error
// @Router       /analitics/genotyped/{year}/byDistrict/{district}/hold [post]
func (g Genotyped) HoldingsPost() func(*gin.Context) {
	return func(c *gin.Context) {
		filterData := cows_filter.CowsFilter{}
		if err := c.ShouldBindJSON(&filterData); err != nil {
			c.JSON(422, err.Error())
		}
		aliveFilter := filterData
		genotypedFilter := filterData

		genotypedFilter.IsGenotyped = new(bool)
		*genotypedFilter.IsGenotyped = true

		db := models.GetDb()

		aliveCowQuery := db.Model(&models.Cow{}).Preload("Holding").Preload("FarmGroup")
		aliveCowFilter := cows_filter.NewCowFilteredModel(aliveFilter, aliveCowQuery)
		aliveCowFilter.Params["year"] = c.Param("year")
		aliveCowFilter.Params["district"] = c.Param("district")

		genotypedCowQuery := db.Model(&models.Cow{}).Preload("Holding").Preload("FarmGroup")
		genotypedCowFilter := cows_filter.NewCowFilteredModel(genotypedFilter, genotypedCowQuery)
		genotypedCowFilter.Params["year"] = c.Param("year")
		genotypedCowFilter.Params["district"] = c.Param("district")

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
			cows_filter.LiveInDistrict{}); err != nil {
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
			cows_filter.LiveInDistrict{}); err != nil {
			c.JSON(422, err.Error())
			return
		}

		aliveCows := []models.Cow{}
		genotypedCows := []models.Cow{}

		aliveCowFilter.GetQuery().Find(&aliveCows)
		genotypedCowFilter.GetQuery().Find(&genotypedCows)

		result := make(map[string]byHoldStatistics)
		for _, cow := range aliveCows {
			key := ""
			holdId := new(uint)

			if cow.Holding == nil {
				holdId = nil
				key = cow.FarmGroup.Name
			} else {
				holdId = &cow.Holding.ID
				key = cow.Holding.Name
			}

			if stat, ok := result[key]; ok {
				stat.Alive+=1
				result[key] = stat
			} else {
				result[key] = byHoldStatistics{
					genotypedStatistics: genotypedStatistics{
						Alive: 1,
						Genotyped: 0,
					},
					HoldID: holdId,
				}
			}
		}

		for _, cow := range genotypedCows {
			key := ""
			if cow.Holding == nil {
				key = cow.FarmGroup.Name
			} else {
				key = cow.Holding.Name
			}

			if stat, ok := result[key]; ok {
				stat.Genotyped+=1
				result[key] = stat
			} else {
				log.Println("Found genotyped cow which is not included into alive cows")
			}
		}

		c.JSON(200, result)
	}
}