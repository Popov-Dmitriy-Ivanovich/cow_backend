package cows

import (
	"cow_backend/models"
	"cow_backend/routes"
	"errors"
	"sort"

	"github.com/gin-gonic/gin"
)

type ReserealizedCheckMilk struct {
	models.CheckMilk
	MilkingDays     int  // День доения относительно начала лактации
	LactationNumber uint // Номер лактации
}

func (rcm ReserealizedCheckMilk) GetReserealizer() routes.Reserealizer {
	return &rcm
}

func (rcm *ReserealizedCheckMilk) FromBaseModel(c any) (routes.Reserealizable, error) {
	lac := models.Lactation{}
	db := models.GetDb()

	cm, ok := c.(models.CheckMilk)
	if !ok {
		return ReserealizedCheckMilk{}, errors.New("error reserealizing")
	}

	if err := db.First(&lac, cm.LactationId).Error; err != nil {
		return ReserealizedCheckMilk{}, err
	}

	lacDate := lac.CalvingDate
	cmDate := cm.CheckDate
	milkingDays := cmDate.Sub(lacDate.Time)

	rcm.LactationNumber = lac.Number
	rcm.CheckMilk = cm
	rcm.MilkingDays = int(milkingDays.Hours() / 24)
	return *rcm, nil
}

// CheckMilks
//
//	@Summary      Get list of check milks
//	@Description  Возращает список всех контрольных доек для конкретной коровы.
//	@Tags         Cows
//	@Param        id   path      int  true  "ID коровы для которой ищутся контрольные дойки"
//
// @Produce      json
// @Success      200  {array}   ReserealizedCheckMilk
// @Failure      500  {object}  string
// @Router       /cows/{id}/checkMilks [get]
func (f *Cows) CheckMilks() func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		cow := models.Cow{}
		db := models.GetDb()
		if err := db.Preload("Lactation").Preload("Lactation.CheckMilks").First(&cow, id).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}
		cms := []models.CheckMilk{}
		for _, lac := range cow.Lactation {
			cms = append(cms, lac.CheckMilks...)
		}
		res := make([]ReserealizedCheckMilk, 0, len(cms))
		for _, cm := range cms {
			reserealizer := &ReserealizedCheckMilk{}
			if reserealized, err := reserealizer.FromBaseModel(cm); err != nil {
				c.JSON(500, err.Error())
				return
			} else {
				appended, ok := reserealized.(ReserealizedCheckMilk)
				if !ok {
					c.JSON(500, "could not resrerealize")
					return
				}
				res = append(res, appended)
			}
		}
		sort.Slice(res, func(i, j int) bool {
			if res[i].LactationNumber == res[j].LactationNumber {
				return res[i].CheckDate.Before(res[j].CheckDate.Time)
			}
			return res[i].LactationNumber < res[j].LactationNumber
		})
		c.JSON(200, res)
	}
}
