package analitics

import (
	"cow_backend/models"
	"cow_backend/routes/auth"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Genotyped struct {
}

type genotypedStatistics struct {
	Alive     int64
	Genotyped int64
	Ill       int64
}

func (g Genotyped) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/genotyped")
	apiGroup.Use(auth.AuthMiddleware(auth.Farmer, auth.RegionalOff, auth.FederalOff))
	apiGroup.GET("/years", g.Years())
	apiGroup.GET("/:year/regions", g.Regions())
	apiGroup.GET("/:year/byRegion/:region/districts", g.Districts())
	apiGroup.GET("/:year/byDistrict/:district/hold", g.Hold())
	apiGroup.GET("/:year/byHold/:hold/hoz", g.Hoz())
	apiGroup.POST("/years", g.YearsPost())
	apiGroup.POST("/:year/regions", g.RegionsPost())
	apiGroup.POST("/:year/byRegion/:region/districts", g.DistrictsPost())
	apiGroup.POST("/:year/byDistrict/:district/hoz", g.HozPost())
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
	genotypedStatistics
	RegionID uint
}

type byRegionKeys struct {
	Name string
	ID   uint
}

// @Summary      Get list of years
// @Description  Возращает словарь регион - количество живых коров, количество генотипированных
// @Tags         Analitics
// @Param        year    path     int  true  "год за который собирается статистика"
// @Produce      json
// @Success      200  {array}   map[string]byRegionStatistics
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

		regionKeys := []byRegionKeys{} // regions where alive cows are registered
		db.Debug().Model(&models.Region{}).Where("EXISTS(SELECT 1 FROM districts WHERE districts.region_id = regions.id AND "+
			" EXISTS (SELECT 1 FROM farms WHERE farms.district_id = districts.id AND "+
			" EXISTS (SELECT 1 FROM cows WHERE (cows.farm_id = farms.id OR cows.farm_id is NULL AND cows.farm_group_id = farms.id) AND ((cows.death_date IS NULL OR cows.death_date < ?) AND cows.birth_date < ?))))",
			time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC), time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC)).Find(&regionKeys)

		res := make(map[string]byRegionStatistics)
		for _, key := range regionKeys {
			aliveWithoutFarm := int64(0)
			aliveWithFarm := int64(0)
			genotypedWithFarm := int64(0)
			genotypedWithoutFarm := int64(0)

			db.Debug().Model(&models.Cow{}).Where("farm_id IS NOT NULL AND ((death_date IS NULL OR death_date < ?) AND cows.birth_date < ?) AND EXISTS "+
				" (SELECT 1 FROM farms WHERE farms.id = cows.farm_id AND "+" EXISTS(SELECT 1 FROM districts WHERE districts.id = farms.district_id AND "+
				" EXISTS(SELECT 1 FROM regions WHERE regions.id = districts.region_id and regions.id = ?)))",
				time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC), time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC),
				key.ID).Count(&aliveWithFarm)
			db.Debug().Model(&models.Cow{}).Where("farm_id IS NULL AND ((death_date IS NULL OR death_date < ?) AND cows.birth_date < ?) AND EXISTS "+
				" (SELECT 1 FROM farms WHERE farms.id = cows.farm_group_id AND "+" EXISTS(SELECT 1 FROM districts WHERE districts.id = farms.district_id AND "+
				" EXISTS(SELECT 1 FROM regions WHERE regions.id = districts.region_id and regions.id = ?)))",
				time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC), time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC),
				key.ID).Count(&aliveWithoutFarm)

			db.Debug().Model(&models.Genetic{}).Where("EXISTS(SELECT 1 FROM cows WHERE cows.id = genetics.cow_id AND ((cows.death_date IS NULL OR cows.death_date < ?) AND cows.birth_date < ?) AND cows.farm_id IS NOT NULL AND "+
				" EXISTS(SELECT 1 FROM farms WHERE farms.id = cows.farm_id AND EXISTS (SELECT 1 FROM districts WHERE districts.id = farms.district_id AND EXISTS(SELECT 1 FROM regions WHERE regions.id = districts.region_id and regions.id = ? ))))",
				time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC), time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC), key.ID).Count(&genotypedWithFarm)
			db.Debug().Model(&models.Genetic{}).Where("EXISTS(SELECT 1 FROM cows WHERE cows.id = genetics.cow_id AND ((cows.death_date IS NULL OR cows.death_date < ?) AND cows.birth_date < ?) AND cows.farm_id IS NULL AND "+
				" EXISTS(SELECT 1 FROM farms WHERE farms.id = cows.farm_group_id AND EXISTS (SELECT 1 FROM districts WHERE districts.id = farms.district_id AND EXISTS(SELECT 1 FROM regions WHERE regions.id = districts.region_id and regions.id = ? ))))",
				time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC), time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC), key.ID).Count(&genotypedWithoutFarm)
			res[key.Name] = byRegionStatistics{
				genotypedStatistics: genotypedStatistics{
					Alive:     aliveWithFarm + aliveWithoutFarm,
					Genotyped: genotypedWithFarm + genotypedWithoutFarm},
				RegionID: key.ID,
			}
		}

		c.JSON(200, res)
	}
}

type byDistrictStatistics struct {
	genotypedStatistics
	DistrictID uint
}

type byDistrictKeys struct {
	Name string
	ID   uint
}

// @Summary      Get list of years
// @Description  Возращает словарь район - количество живых коров, количество генотипированных
// @Tags         Analitics
// @Param        year    path     int  true  "год за который собирается статистика"
// @Param        region    path     int  true  "регион за который собирается статистика"
// @Produce      json
// @Success      200  {array}   map[string]byDistrictStatistics
// @Failure      500  {object}  map[string]error
// @Router       /analitics/genotyped/{year}/byRegion/{region}/districts [get]
func (g Genotyped) Districts() func(*gin.Context) {
	return func(c *gin.Context) {

		region := c.Param("region")
		year := c.Param("year")

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

		yearInt, err := strconv.ParseInt(year, 10, 64)
		if err != nil {
			c.JSON(422, err.Error())
			return
		}
		districtKeys := []byDistrictKeys{}
		db := models.GetDb()

		db.Model(&models.District{}).Where(" region_id = ? AND EXISTS (SELECT 1 FROM farms WHERE farms.district_id = districts.id AND "+
			" EXISTS (SELECT 1 FROM cows WHERE (cows.farm_id = farms.id OR cows.farm_id is NULL AND cows.farm_group_id = farms.id) AND ((cows.death_date IS NULL OR cows.death_date < ?) AND cows.birth_date < ?)))",
			region,
			time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC), time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC)).Find(&districtKeys)

		res := make(map[string]byDistrictStatistics)
		for _, key := range districtKeys {
			genotypedWithFarm := int64(0)
			genotypedWithoutFarm := int64(0)
			aliveWithFarm := int64(0)
			aliveWithoutFarm := int64(0)

			db.Model(&models.Cow{}).Where("farm_id IS NOT NULL AND ((death_date IS NULL OR death_date < ?) AND cows.birth_date < ?) AND EXISTS (SELECT 1 FROM farms WHERE farms.id = cows.farm_id AND "+
				" EXISTS (SELECT 1 FROM districts WHERE districts.id = farms.district_id AND districts.id = ?))", time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC), time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC),
				key.ID,
			).Count(&aliveWithFarm)
			db.Model(&models.Cow{}).Where("farm_id IS NULL AND ((death_date IS NULL OR death_date < ?) AND cows.birth_date < ?) AND EXISTS (SELECT 1 FROM farms WHERE farms.id = cows.farm_group_id AND "+
				" EXISTS (SELECT 1 FROM districts WHERE districts.id = farms.district_id AND districts.id = ?))", time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC), time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC),
				key.ID,
			).Count(&aliveWithoutFarm)

			db.Model(&models.Genetic{}).Where("EXISTS(SELECT 1 FROM cows WHERE cows.id = genetics.cow_id AND ((cows.death_date IS NULL OR cows.death_date < ?) AND cows.birth_date < ?) AND cows.farm_id IS NOT NULL AND "+
				"EXISTS (SELECT 1 FROM farms WHERE farms.id = cows.farm_id AND EXISTS(SELECT 1 FROM districts WHERE districts.id = farms.district_id AND districts.id = ?)))",
				time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC), time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC),
				key.ID).Count(&genotypedWithFarm)
			db.Model(&models.Genetic{}).Where("EXISTS(SELECT 1 FROM cows WHERE cows.id = genetics.cow_id AND ((cows.death_date IS NULL OR cows.death_date < ?) AND cows.birth_date < ?) AND cows.farm_id IS NULL AND "+
				"EXISTS (SELECT 1 FROM farms WHERE farms.id = cows.farm_group_id AND EXISTS(SELECT 1 FROM districts WHERE districts.id = farms.district_id AND districts.id = ?)))",
				time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC), time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC),
				key.ID).Count(&genotypedWithoutFarm)
			res[key.Name] = byDistrictStatistics{
				genotypedStatistics: genotypedStatistics{
					Alive:     aliveWithFarm + aliveWithoutFarm,
					Genotyped: genotypedWithFarm + genotypedWithoutFarm,
				},
				DistrictID: key.ID,
			}
		}
		c.JSON(200, res)
	}
}

type byHoldStatistics struct {
	genotypedStatistics
	HoldID *uint
}

// @Summary      Get list of years
// @Description  Возращает словарь хозяйство - количество живых коров, количество генотипированных
// @Tags         Analitics
// @Param        year    path     int  true  "год за который собирается статистика"
// @Param        district    path     int  true  "район за который собирается статистика"
// @Produce      json
// @Success      200  {array}   map[string]byHoldStatistics
// @Failure      500  {object}  map[string]error
// @Router       /analitics/genotyped/{year}/byDistrict/{district}/hold [get]
func (g Genotyped) Hold() func(*gin.Context) {
	return func(c *gin.Context) {

		year := c.Param("year")
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

		db := models.GetDb()
		yearInt, err := strconv.ParseUint(year, 10, 64)
		if err != nil {
			c.JSON(422, err.Error())
			return
		}

		cows := []models.Cow{}

		db.Preload("FarmGroup").Preload("FarmGroup.Parrent").Preload("Genetic").Where("(cows.death_date IS NULL OR cows.death_date < ?) AND cows.birth_date < ? AND EXISTS (SELECT 1 FROM farms WHERE cows.farm_group_id = farms.id and farms.district_id = ?)",
			time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC), time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC), district).Find(&cows)
		res := make(map[string]byHoldStatistics)
		for _, cow := range cows {

			val := byHoldStatistics{HoldID: &cow.FarmGroup.ID}
			if _, ok := res[cow.FarmGroup.Name]; ok {
				val = res[cow.FarmGroup.Name]
			}
			val.Alive += 1
			if cow.Genetic != nil {
				val.Genotyped += 1
			}
			res[cow.FarmGroup.Name] = val
		}

		c.JSON(200, res)
	}
}

type byHozStatistics struct {
	genotypedStatistics
}

type byHozKeys struct {
	Name string
	ID   uint
}

// @Summary      Get list of years
// @Description  Возращает словарь хозяйство - количество живых коров, количество генотипированных
// @Tags         Analitics
// @Param        year    path     int  true  "год за который собирается статистика"
// @Param        hold    path     int  true  "холдинг за который собирается статистика"
// @Produce      json
// @Success      200  {array}   map[string]byHoldStatistics
// @Failure      500  {object}  map[string]error
// @Router       /analitics/genotyped/{year}/byHold/{hold}/hoz [get]
func (g Genotyped) Hoz() func(*gin.Context) {
	return func(c *gin.Context) {
		db := models.GetDb()

		year := c.Param("year")
		hold := c.Param("hold")
		yearInt, err := strconv.ParseInt(year, 10, 64)
		if err != nil {
			c.JSON(422, err.Error())
			return
		}

		keys := []byHozKeys{}
		db.Model(&models.Farm{}).Where("parrent_id = ? AND EXISTS (SELECT 1 FROM cows WHERE cows.farm_group_id = farms.id AND (cows.death_date IS NULL OR cows.death_date < ?) AND cows.birth_date < ?)",
			hold, time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC), time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC)).Find(&keys)
		res := make(map[string]byHozStatistics)
		for _, key := range keys {
			alive := int64(0)
			genotyped := int64(0)

			db.Debug().Model(&models.Cow{}).Where("EXISTS (SELECT 1 FROM farms WHERE farms.id = cows.farm_group_id AND farms.parrent_id = ? AND farms.id = ?) AND (cows.death_date IS NULL OR cows.death_date < ?) AND cows.birth_date < ? ", hold, key.ID, time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC), time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC)).Count(&alive)
			db.Debug().Model(&models.Cow{}).Where("EXISTS (SELECT 1 FROM farms WHERE farms.id = cows.farm_group_id AND farms.parrent_id = ? AND farms.id = ?) AND (cows.death_date IS NULL OR cows.death_date < ?) AND cows.birth_date < ? AND EXISTS (SELECT 1 FROM genetics where genetics.cow_id = cows.id)", hold, key.ID, time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC), time.Date(int(yearInt+1), 1, 1, 0, 0, 0, 0, time.UTC)).Count(&genotyped)
			res[key.Name] = byHozStatistics{
				genotypedStatistics: genotypedStatistics{
					Alive:     alive,
					Genotyped: genotyped,
				},
			}
		}
		c.JSON(200, res)
	}
}
