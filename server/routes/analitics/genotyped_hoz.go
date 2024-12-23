package analitics

import (
	"cow_backend/filters"
	"cow_backend/filters/cows_filter"
	"cow_backend/models"
	"log"
	"net/http"
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
// @Description  Возращает словарь хозяйство - количество живых коров, количество генотипированных
// @Tags         Analitics
// @Param        year    path     int  true  "год за который собирается статистика"
// @Param        district    path     int  true  "район за который собирается статистика"
// @Param        filter    body     cows_filter.CowsFilter  true  "applied filters"
// @Produce      json
// @Success      200  {array}   map[string]byHozStatistics
// @Failure      500  {object}  map[string]error
// @Router       /analitics/genotyped/{year}/byDistrict/{district}/hoz [post]
func (g Genotyped) HozPost() func(*gin.Context) {
	return func(c *gin.Context) {
		

		district := c.Param("district")
		roleId, exists := c.Get("RoleId")
		if !exists {
			c.JSON(http.StatusInternalServerError, "RoleId не найден в контексте")
			return
		}

		if roleId == 1 {
			distId, exists := c.Get("DistId")
			if !exists {
				c.JSON(http.StatusInternalServerError, "DistId не найден в контексте")
				return
			}

			log.Println(distId, district)
			if distId != district {
				c.JSON(421, gin.H{"error": "Нет доступа к округу"})
				c.Abort()
				return
			}
		}

		filterData := cows_filter.CowsFilter{}
		if err := c.ShouldBindJSON(&filterData); err != nil {
			c.JSON(422, err.Error())
		}
		aliveFilter := filterData
		genotypedFilter := filterData
		illFilter := filterData

		aliveFilter.IsDead = new(bool)
		*aliveFilter.IsDead = false

		genotypedFilter.IsGenotyped = new(bool)
		*genotypedFilter.IsGenotyped = true
		genotypedFilter.IsDead = new(bool)
		*genotypedFilter.IsDead = false

		illFilter.HasAnyIllnes = new(bool)
		*illFilter.HasAnyIllnes = true
		illFilter.IsDead = new(bool)
		*illFilter.IsDead = false

		keys := []byHozKeys{}
		db := models.GetDb()
		yearStr := c.Param("year")
		yearInt, err := strconv.ParseInt(yearStr, 10, 64)
		if err != nil {
			c.JSON(422, err.Error())
			return
		}
		db.Model(&models.Farm{}).Debug().Where(
			"district_id = ? AND EXISTS (SELECT 1 FROM cows where (cows.farm_group_id = farms.id) AND "+
				"(cows.death_date IS NULL OR cows.death_date < ?) AND cows.birth_date < ?)",
			c.Param("district"),
			time.Date(int(yearInt)+1, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(int(yearInt)+1, 1, 1, 0, 0, 0, 0, time.UTC)).Find(&keys)

		result := make(map[string]byHozStatistics)
		for _, key := range keys {
			aliveCowQuery := db.Model(&models.Cow{}).Where("approved <> -1")
			aliveCowFilter := cows_filter.NewCowFilteredModel(aliveFilter, aliveCowQuery)
			aliveCowFilter.Params["year"] = c.Param("year")
			aliveCowFilter.Params["district"] = c.Param("district")
			aliveCowFilter.Params["hoz"] = strconv.FormatUint(uint64(key.ID), 10)

			genotypedCowQuery := db.Model(&models.Cow{}).Where("approved <> -1")
			genotypedCowFilter := cows_filter.NewCowFilteredModel(genotypedFilter, genotypedCowQuery)
			genotypedCowFilter.Params["year"] = c.Param("year")
			genotypedCowFilter.Params["district"] = c.Param("district")
			genotypedCowFilter.Params["hoz"] = strconv.FormatUint(uint64(key.ID), 10)

			illCowQuery := db.Model(&models.Cow{}).Where("approved <> -1")
			illCowFilter := cows_filter.NewCowFilteredModel(illFilter, illCowQuery)
			illCowFilter.Params["year"] = c.Param("year")
			illCowFilter.Params["district"] = c.Param("district")
			illCowFilter.Params["hoz"] = strconv.FormatUint(uint64(key.ID), 10)

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
				cows_filter.LiveInDistrict{},
				cows_filter.LiveInHolding{}); err != nil {
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
				cows_filter.LiveInDistrict{},
				cows_filter.LiveInHolding{}); err != nil {
				c.JSON(422, err.Error())
				return
			}
			if err := filters.ApplyFilters(illCowFilter,
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
				cows_filter.LiveInDistrict{},
				cows_filter.LiveInHolding{}); err != nil {
				c.JSON(422, err.Error())
				return
			}
			alive := int64(0)
			genotyped := int64(0)
			ill := int64(0)

			aliveCowFilter.GetQuery().Debug().Count(&alive)
			genotypedCowFilter.GetQuery().Debug().Count(&genotyped)
			illCowFilter.GetQuery().Debug().Count(&ill)
			result[key.Name] = byHozStatistics{
				genotypedStatistics: genotypedStatistics{
					Alive:     alive,
					Genotyped: genotyped,
					Ill:       ill,
				},
			}
		}
		c.JSON(200, result)
	}
}
