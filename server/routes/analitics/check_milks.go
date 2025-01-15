package analitics

import (
	"cow_backend/filters"
	"cow_backend/filters/cows_filter"
	"cow_backend/models"
	"cow_backend/routes/auth"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CheckMilks struct{}

func (cm *CheckMilks) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/checkMilks")
	apiGroup.Use(auth.AuthMiddleware(auth.Farmer, auth.RegionalOff, auth.FederalOff))
	apiGroup.POST("/years", cm.ByYear())
	apiGroup.POST("/:year/byRegion", cm.ByRegion())
	apiGroup.POST("/:year/byRegion/:region/byDistrict", cm.ByDistrict())
	apiGroup.POST("/:year/byDistrict/:district/byHoz", cm.ByHoz())
}

func iterateCows(cowIds []uint, year int, regionId *uint, districtId *uint,
	resultsWriter func(cow *models.Cow, cmCount uint, milk float64, fat float64, protein float64)) {
	db := models.GetDb()
	for _, id := range cowIds {
		dbCow := models.Cow{}
		db.Preload("Lactation").
			Preload("Lactation.CheckMilks").
			Preload("FarmGroup").
			Preload("FarmGroup.District").
			Preload("FarmGroup.District.Region").First(&dbCow, id)
		cmCount := uint(0)
		milk := float64(0)
		fat := float64(0)
		protein := float64(0)
		if regionId != nil {
			if dbCow.FarmGroup.District.RegionId != *regionId {
				continue
			}
		}
		if districtId != nil {
			if dbCow.FarmGroup.DistrictId != *districtId {
				continue
			}
		}
		for _, lac := range dbCow.Lactation {
			for _, cm := range lac.CheckMilks {
				if cm.CheckDate.Year() != year {
					continue
				}
				cmCount++
				milk += cm.Milk
				fat += cm.Fat
				protein += cm.Protein
			}
		}
		resultsWriter(&dbCow, cmCount, milk, fat, protein)
	}
}

func createDefaultCmCowsFilter(filterData cows_filter.CowsFilter) *cows_filter.CowFilteredModel {
	filterData.ControlMilkingDateFrom = new(string)
	filterData.ControlMilkingDateTo = new(string)

	*filterData.ControlMilkingDateFrom = "0001-01-01"
	*filterData.ControlMilkingDateTo = "4000-01-01"

	filterData.IsDead = new(bool)
	*filterData.IsDead = false

	db := models.GetDb()
	cmCowQuery := db.Model(models.Cow{}).Where("approved <> -1")
	return cows_filter.NewCowFilteredModel(filterData, cmCowQuery)
}

// ByYear
// @Summary      Get list of years
// @Description  Возращает словарь год - истина. Ключи словаря - это годы по которым есть аналитика
// @Param        filter    body     cows_filter.CowsFilter  true  "applied filters"
// @Tags         Analytics (Milk)
// @Produce      json
// @Success      200  {object}   map[int]bool
// @Failure      422  {object}   string
// @Router       /analitics/checkMilks/years [post]
func (cm CheckMilks) ByYear() func(*gin.Context) {
	return func(c *gin.Context) {
		filterData := cows_filter.CowsFilter{}
		if err := c.ShouldBindJSON(&filterData); err != nil {
			c.JSON(422, err.Error())
		}
		cmCowFilter := createDefaultCmCowsFilter(filterData)
		if err := filters.ApplyFilters(cmCowFilter, cows_filter.ALL_FILTERS...); err != nil {
			c.JSON(422, err.Error())
			return
		}
		cowIds := []uint{}
		cmCowFilter.GetQuery().Pluck("id", &cowIds)
		result := map[int]bool{}
		db := models.GetDb()
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
	Milk     float64 // надой на корову
	Fat      float64 // жир на корову
	Protein  float64 // белок на корову
	CmCount  uint    // служебная штука (кол-во контрольных доений)
	CowCount uint    // служебная штука (кол-во коров)
	RegionId uint    // ID региона по которому собрана аналитика
}

// ByRegion
// @Summary      Get by region analytics
// @Description  Возращает словарь название региона - статистика по региону
// @Param        year    path     int  true  "год за который собирается статистика"
// @Param        filter    body     cows_filter.CowsFilter  true  "applied filters"
// @Tags         Analytics (Milk)
// @Produce      json
// @Success      200  {array}   map[string]cmByRegionStatistics
// @Failure      422  {object}  string
// @Router       /analitics/checkMilks/{year}/byRegion [post]
func (cm CheckMilks) ByRegion() func(*gin.Context) {
	return func(c *gin.Context) {
		filterData := cows_filter.CowsFilter{}
		if err := c.ShouldBindJSON(&filterData); err != nil {
			c.JSON(422, err.Error())
		}
		cmCowFilter := createDefaultCmCowsFilter(filterData)
		if err := filters.ApplyFilters(cmCowFilter, cows_filter.ALL_FILTERS...); err != nil {
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
		iterateCows(cowIds, int(yearInt), nil, nil,
			func(dbCow *models.Cow, cmCount uint, milk float64, fat float64, protein float64) {
				if val, ok := result[dbCow.FarmGroup.District.Region.Name]; ok && cmCount != 0 {
					val.Milk += milk / float64(cmCount)
					val.Fat += fat / float64(cmCount)
					val.Protein += protein / float64(cmCount)
					val.CowCount += 1
					result[dbCow.FarmGroup.District.Region.Name] = val
				} else if !ok && cmCount != 0 {
					val := cmByRegionStatistics{}
					val.Milk = milk / float64(cmCount)
					val.Fat = fat / float64(cmCount)
					val.Protein = protein / float64(cmCount)
					val.CowCount = 1
					val.RegionId = dbCow.FarmGroup.District.RegionId
					result[dbCow.FarmGroup.District.Region.Name] = val
				}
			})

		for key, val := range result {
			val.Milk = val.Milk / float64(val.CowCount)
			val.Fat = val.Fat / float64(val.CowCount)
			val.Protein = val.Protein / float64(val.CowCount)
			result[key] = val
		}

		c.JSON(200, result)
	}
}

type cmByDistrictStatistics struct {
	Milk       float64 // удой на корову
	Fat        float64 // жир на корову
	Protein    float64 // белок на корову
	CmCount    uint    // количество обработанных контрольных доений (служебная)
	CowCount   uint    // количество обработанных коров (служебная)
	DistrictId uint    // ID района
}

// ByDistrict
// @Summary      Get by district analytics
// @Description  Возращает словарь название района - аналитика по району
// @Param        year    path     int  true  "год за который собирается статистика"
// @Param        region    path     int  true  "регион за который собирается статистика"
// @Param        filter    body     cows_filter.CowsFilter  true  "applied filters"
// @Tags         Analytics (Milk)
// @Produce      json
// @Success      200  {array}   map[string]cmByDistrictStatistics
// @Failure      422  {object}  string
// @Router       /analitics/checkMilks/{year}/byRegion/{region}/byDistrict [post]
func (cm CheckMilks) ByDistrict() func(*gin.Context) {
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
		cmCowFilter := createDefaultCmCowsFilter(filterData)
		if err := filters.ApplyFilters(cmCowFilter, cows_filter.ALL_FILTERS...); err != nil {
			c.JSON(422, err.Error())
			return
		}
		regionIdStr := c.Param("region")
		regionIdInt, err := strconv.ParseUint(regionIdStr, 10, 64)
		if err != nil {
			c.JSON(422, err.Error())
			return
		}
		cowIds := []uint{}
		cmCowFilter.GetQuery().Pluck("id", &cowIds)
		result := map[string]cmByDistrictStatistics{}
		yearStr := c.Param("year")
		yearInt, err := strconv.ParseInt(yearStr, 10, 64)
		if err != nil {
			c.JSON(422, err.Error())
			return
		}
		regIdUint := uint(regionIdInt)
		iterateCows(cowIds, int(yearInt), &regIdUint, nil,
			func(dbCow *models.Cow, cmCount uint, milk float64, fat float64, protein float64) {
				if val, ok := result[dbCow.FarmGroup.District.Name]; ok && cmCount != 0 {
					val.Milk += milk / float64(cmCount)
					val.Fat += fat / float64(cmCount)
					val.Protein += protein / float64(cmCount)
					val.CowCount += 1
					result[dbCow.FarmGroup.District.Name] = val
				} else if !ok && cmCount != 0 {
					val := cmByDistrictStatistics{}
					val.Milk = milk / float64(cmCount)
					val.Fat = fat / float64(cmCount)
					val.Protein = protein / float64(cmCount)
					val.CowCount = 1
					val.DistrictId = dbCow.FarmGroup.District.ID
					result[dbCow.FarmGroup.District.Name] = val
				}
			})

		for key, val := range result {
			val.Milk = val.Milk / float64(val.CowCount)
			val.Fat = val.Fat / float64(val.CowCount)
			val.Protein = val.Protein / float64(val.CowCount)
			result[key] = val
		}

		c.JSON(200, result)
	}
}

type cmByHozStatistics struct {
	Milk     float64 // удой на корову
	Fat      float64 // жир на корову
	Protein  float64 // белок на корову
	CmCount  uint    // служебная штука
	CowCount uint    // служебная штука
}

// ByHoz
// @Summary      Получить аналитику по хозяйству
// @Description  Возращает словарь название хозяйства - аналитика по хозяйству
// @Param        year    path     int  true  "год за который собирается статистика"
// @Param        district    path     int  true  "район за который собирается статистика"
// @Param        filter    body     cows_filter.CowsFilter  true  "applied filters"
// @Tags         Analytics (Milk)
// @Produce      json
// @Success      200  {array}   map[string]cmByHozStatistics
// @Failure      422  {object}  string
// @Router       /analitics/checkMilks/{year}/byDistrict/{district}/byHoz [post]
func (cm CheckMilks) ByHoz() func(*gin.Context) {
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

		cmCowFilter := createDefaultCmCowsFilter(filterData)
		if err := filters.ApplyFilters(cmCowFilter, cows_filter.ALL_FILTERS...); err != nil {
			c.JSON(422, err.Error())
			return
		}
		districtIdStr := c.Param("district")
		districtIdInt, err := strconv.ParseUint(districtIdStr, 10, 64)
		if err != nil {
			c.JSON(422, err.Error())
			return
		}
		cowIds := []uint{}
		cmCowFilter.GetQuery().Pluck("id", &cowIds)
		result := map[string]cmByHozStatistics{}
		yearStr := c.Param("year")
		yearInt, err := strconv.ParseInt(yearStr, 10, 64)
		if err != nil {
			c.JSON(422, err.Error())
			return
		}
		districtIdUint := uint(districtIdInt)
		iterateCows(cowIds, int(yearInt), nil, &districtIdUint,
			func(dbCow *models.Cow, cmCount uint, milk float64, fat float64, protein float64) {
				if val, ok := result[dbCow.FarmGroup.Name]; ok && cmCount != 0 {
					val.Milk += milk / float64(cmCount)
					val.Fat += fat / float64(cmCount)
					val.Protein += protein / float64(cmCount)
					val.CowCount += 1
					result[dbCow.FarmGroup.Name] = val
				} else if !ok && cmCount != 0 {
					val := cmByHozStatistics{}
					val.Milk = milk / float64(cmCount)
					val.Fat = fat / float64(cmCount)
					val.Protein = protein / float64(cmCount)
					val.CowCount = 1
					result[dbCow.FarmGroup.Name] = val
				}
			})

		for key, val := range result {
			val.Milk = val.Milk / float64(val.CowCount)
			val.Fat = val.Fat / float64(val.CowCount)
			val.Protein = val.Protein / float64(val.CowCount)
			result[key] = val
		}

		c.JSON(200, result)
	}
}
