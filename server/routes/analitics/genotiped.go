package analitics

import (
	"cow_backend/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Genotyped struct {
}

func (g Genotyped) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/genotyped")
	apiGroup.GET("/years", g.Years())
	apiGroup.GET("/:year/regions", g.Regions())
}

// @Summary      Get list of years
// @Description  Возращает словарь год - количеств генотипированных коров, по ключу -1 генотипированные за все годы
// @Tags         Analitics
// @Produce      json
// @Success      200  {array}   map[int]uint
// @Failure      500  {object}  map[string]error
// @Router       /analitics/genotyped/years [get]
func (g Genotyped) Years() func(*gin.Context) {
	return func(c *gin.Context) {
		db := models.GetDb()
		dates := []models.DateOnly{}
		db.Model(&models.Genetic{}).Pluck("result_date", &dates)
		result := make(map[int]uint)
		allYears := 0
		for _, date := range dates {
			if _, ok := result[date.Year()]; ok {
				result[date.Year()] += 1
			} else {
				result[date.Year()] = 1
			}
			allYears++
		}
		result[-1] = uint(allYears)
		c.JSON(200, result)
	}
}

type byRegionStatistics struct {
	Alive     int64
	Genotyped int64
}

// @Summary      Get list of years
// @Description  Возращает словарь регион - количество живых коров, количество генотипированных
// @Tags         Analitics
// @Param        year    path     int  true  "год за который собирается статистика"
// @Produce      json
// @Success      200  {array}   map[int]uint
// @Failure      500  {object}  map[string]error
// @Router       /analitics/genotyped/{year}/regions [get]
func (g Genotyped) Regions() func(*gin.Context) {
	return func(c *gin.Context) {
		year := c.Param("year")
		db := models.GetDb()
		yearInt, err := strconv.ParseUint(year, 10, 64)
		if err != nil {
			c.JSON(422, "wrong year")
			return
		}
		// farmIds := []*uint{}
		// db.Model(&models.Cow{}).Where("EXISTS (SELECT 1 FROM genetics where genetics.cow_id = cows.id and genetics.result_date BETWEEN ? AND ?)",
		// 	time.Date(int(yearInt), 1, 1, 0, 0, 0, 0, time.UTC),
		// 	time.Date(int(yearInt), 12, 31, 0, 0, 0, 0, time.UTC)).Pluck("farm_id", farmIds)

		regionNames := []string{} // regions where alive cows are registered
		db.Model(&models.Region{}).Where("EXISTS(SELECT 1 FROM districts WHERE districts.region_id = regions.id AND "+
			" EXISTS (SELECT 1 FROM farms WHERE farms.district_id = districts.id AND "+
			" EXISTS (SELECT 1 FROM cows WHERE (cows.farm_id = farms.id OR cows.farm_group_id = farms.id) AND (cows.death_date IS NULL OR cows.death_date < ? AND cows.birth_date < ?))))",
			time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC), time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC)).Pluck("name", &regionNames)

		res := make(map[string]byRegionStatistics)
		for _, rName := range regionNames {
			aliveWithoutFarm := int64(0)
			aliveWithFarm := int64(0)
			genotypedWithFarm := int64(0)
			genotypedWithoutFarm := int64(0)

			db.Debug().Model(&models.Cow{}).Where("farm_id IS NOT NULL AND (death_date IS NULL OR death_date < ? AND cows.birth_date < ?) AND EXISTS "+
				" (SELECT 1 FROM farms WHERE farms.id = cows.farm_id AND "+" EXISTS(SELECT 1 FROM districts WHERE districts.id = farms.district_id AND "+
				" EXISTS(SELECT 1 FROM regions WHERE regions.id = districts.region_id and regions.NAME = ?)))",
				time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC), time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC),
				rName).Count(&aliveWithFarm)
			db.Debug().Model(&models.Cow{}).Where("farm_id IS NULL AND (death_date IS NULL OR death_date < ? AND cows.birth_date < ?) AND EXISTS "+
				" (SELECT 1 FROM farms WHERE farms.id = cows.farm_group_id AND "+" EXISTS(SELECT 1 FROM districts WHERE districts.id = farms.district_id AND "+
				" EXISTS(SELECT 1 FROM regions WHERE regions.id = districts.region_id and regions.NAME = ?)))",
				time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC), time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC),
				rName).Count(&aliveWithoutFarm)

			db.Debug().Model(&models.Genetic{}).Where("EXISTS(SELECT 1 FROM cows WHERE cows.id = genetics.cow_id AND (death_date IS NULL OR death_date < ? AND cows.birth_date < ?) AND cows.farm_id IS NOT NULL AND "+
				" EXISTS(SELECT 1 FROM farms WHERE farms.id = cows.farm_id AND EXISTS (SELECT 1 FROM districts WHERE districts.id = farms.district_id AND EXISTS(SELECT 1 FROM regions WHERE regions.id = districts.region_id and regions.NAME = ? ))))",
				time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC), time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC), rName).Count(&genotypedWithFarm)
			db.Debug().Model(&models.Genetic{}).Where("EXISTS(SELECT 1 FROM cows WHERE cows.id = genetics.cow_id AND (death_date IS NULL OR death_date < ? AND cows.birth_date < ?) AND cows.farm_id IS NULL AND "+
				" EXISTS(SELECT 1 FROM farms WHERE farms.id = cows.farm_group_id AND EXISTS (SELECT 1 FROM districts WHERE districts.id = farms.district_id AND EXISTS(SELECT 1 FROM regions WHERE regions.id = districts.region_id and regions.NAME = ? ))))",
				time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC), time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC), rName).Count(&genotypedWithoutFarm)
			res[rName] = byRegionStatistics{Alive: aliveWithFarm + aliveWithoutFarm, Genotyped: genotypedWithFarm + genotypedWithoutFarm}
		}

		c.JSON(200, res)
	}
}
