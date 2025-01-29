package cows

import (
	"cow_backend/models"

	"github.com/gin-gonic/gin"
)

// Grades
// @Summary      Get grades
// @Description  Возращает словарь с ключами:
// @Description  1. ByRegion - Значения оценок EBV по региону
// @Description  2. ByHoz - Значения оценок EBV по хозяйству
// @Description  3. ByCountry - Значения оценок EBV по стране
// @Description  4. Average - Средние значения оценок EBV
// @Description  5. PercentsRegion - Отклонение оценок от среднего значения для региона
// @Tags         Cows
// @Param        id   path      int  true  "ID коровы для которой ищутся оценки"
// @Produce      json
// @Success      200  {object}   map[string]models.Grade
// @Failure      500  {object}   string
// @Router       /cows/{id}/grades [get]
func (c *Cows) Grades() func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		db := models.GetDb()
		cow := models.Cow{}
		percentageRegion := models.Grade{}
		if err := db.
			Preload("GradeRegion").
			Preload("GradeHoz").
			Preload("GradeCountry").
			First(&cow, id).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}

		blupStat := models.BlupStatistics{}
		if res := db.Limit(1).Order("created_at desc").Find(&blupStat); res.Error != nil {
			c.JSON(500, res.Error.Error())
			return
		} else if res.RowsAffected == 0 {
			c.JSON(500, "не найдена статистика blup")
		}
		if cow.GradeRegion != nil {
			percentageRegion = cow.GradeRegion.Grade
			if cow.GradeRegion.EbvFat != nil {
				*percentageRegion.EbvFat = *cow.GradeRegion.EbvFat/blupStat.AverageEbvFatRegion*100 - 100
			}
			if cow.GradeRegion.EbvMilk != nil {
				*percentageRegion.EbvMilk = *cow.GradeRegion.EbvMilk/blupStat.AverageEbvMilkRegion*100 - 100
			}
			if cow.GradeRegion.EbvProtein != nil {
				*percentageRegion.EbvProtein = *cow.GradeRegion.EbvProtein/blupStat.AverageEbvProteinRegion*100 - 100
			}
			if cow.GradeRegion.EbvInsemenation != nil {
				*percentageRegion.EbvInsemenation = *cow.GradeRegion.EbvInsemenation/blupStat.AverageEbvInsemenationRegion*100 - 100
			}
			if cow.GradeRegion.EbvService != nil {
				*percentageRegion.EbvService = *cow.GradeRegion.EbvService/blupStat.AverageEbvServiceRegion*100 - 100
			}
			if cow.GradeRegion.GeneralValue != nil {
				*percentageRegion.GeneralValue = *cow.GradeRegion.GeneralValue/blupStat.AverageEbvGeneralValueRegion*100 - 100
			}
		}

		c.JSON(200, gin.H{
			"ByRegion":       cow.GradeRegion,
			"ByHoz":          cow.GradeHoz,
			"ByCountry":      cow.GradeCountry,
			"Average":        blupStat,
			"PercentsRegion": percentageRegion,
		})
	}
}
