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

// DistrictsPost
// @Summary      Get analytics by districts
// @Description  Возращает словарь район - количество живых коров, количество генотипированных
// @Tags         Analytics(GeneticFilters)
// @Param        year    path     int  true  "год за который собирается статистика"
// @Param        region    path     int  true  "регион за который собирается статистика"
// @Param        filter    body     cows_filter.CowsFilter  true  "applied filters"
// @Produce      json
// @Success      200  {array}   map[string]byDistrictStatistics
// @Failure      422  {object}  string
// @Failure      500  {object}  string
// @Failure      421  {object}  string
// @Router       /analitics/genotyped/{year}/byRegion/{region}/districts [post]
func (g Genotyped) DistrictsPost() func(*gin.Context) {
	return func(c *gin.Context) {

		region := c.Param("region")

		roleId, exists := c.Get("RoleId")
		if !exists {
			c.JSON(http.StatusInternalServerError, "RoleId не найден в контексте")
			return
		}

		if roleId != 3 && roleId != 4 {
			regionId, exists := c.Get("RegionId")
			if !exists {
				c.JSON(http.StatusInternalServerError, "RegionId не найден в контексте")
				return
			}

			log.Println(regionId, region)
			if regionId != region {
				c.JSON(421, gin.H{"error": "Нет доступа к региону"})
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

		keys := []byDistrictKeys{}
		db := models.GetDb()
		yearStr := c.Param("year")
		yearInt, err := strconv.ParseInt(yearStr, 10, 64)
		if err != nil {
			c.JSON(422, err.Error())
			return
		}
		db.Model(&models.District{}).Debug().Where(
			"region_id = ? AND EXISTS(SELECT 1 FROM farms where farms.district_id = districts.id AND "+
				" EXISTS (SELECT 1 FROM cows WHERE (cows.farm_id = farms.id OR cows.farm_group_id = farms.id) AND "+
				" (cows.death_date IS NULL OR cows.death_date < ?) AND cows.birth_date < ?  AND"+
				" EXISTS (SELECT 1 FROM genetics where genetics.cow_id = cows.id)))",
			c.Param("region"),
			time.Date(int(yearInt)+1, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(int(yearInt)+1, 1, 1, 0, 0, 0, 0, time.UTC)).Find(&keys)

		result := make(map[string]byDistrictStatistics)
		for _, key := range keys {
			aliveCowQuery := db.Model(&models.Cow{}).Where("approved <> -1")
			aliveCowFilter := cows_filter.NewCowFilteredModel(aliveFilter, aliveCowQuery)
			aliveCowFilter.Params["year"] = c.Param("year")
			aliveCowFilter.Params["district"] = strconv.FormatUint(uint64(key.ID), 10)

			genotypedCowQuery := db.Model(&models.Cow{}).Where("approved <> -1")
			genotypedCowFilter := cows_filter.NewCowFilteredModel(genotypedFilter, genotypedCowQuery)
			genotypedCowFilter.Params["year"] = c.Param("year")
			genotypedCowFilter.Params["district"] = strconv.FormatUint(uint64(key.ID), 10)

			illCowQuery := db.Model(&models.Cow{}).Where("approved <> -1")
			illCowFilter := cows_filter.NewCowFilteredModel(illFilter, illCowQuery)
			illCowFilter.Params["year"] = c.Param("year")
			illCowFilter.Params["district"] = strconv.FormatUint(uint64(key.ID), 10)

			if err := filters.ApplyFilters(aliveCowFilter, cows_filter.ALL_FILTERS...); err != nil {
				c.JSON(422, err.Error())
				return
			}
			if err := filters.ApplyFilters(genotypedCowFilter, cows_filter.ALL_FILTERS...); err != nil {
				c.JSON(422, err.Error())
				return
			}

			if err := filters.ApplyFilters(illCowFilter, cows_filter.ALL_FILTERS...); err != nil {
				c.JSON(422, err.Error())
				return
			}

			alive := int64(0)
			genotyped := int64(0)
			ill := int64(0)

			aliveCowFilter.GetQuery().Debug().Count(&alive)
			genotypedCowFilter.GetQuery().Debug().Count(&genotyped)
			illCowFilter.GetQuery().Debug().Count(&ill)

			result[key.Name] = byDistrictStatistics{
				genotypedStatistics: genotypedStatistics{
					Alive:     alive,
					Genotyped: genotyped,
					Ill:       ill,
				},
				DistrictID: key.ID,
			}
		}
		c.JSON(200, result)
	}
}
