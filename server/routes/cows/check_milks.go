package cows

import (
	"cow_backend/models"
	"cow_backend/routes"
	"errors"

	"github.com/gin-gonic/gin"
)

type ReserealizedCheckMilk struct {
	models.CheckMilk
	MilkingDays int
}

func (rcm ReserealizedCheckMilk) GetReserealizer() routes.Reserealizer {
	return &rcm
}

func (rcm *ReserealizedCheckMilk) FromBaseModel(c any) (routes.Reserealizable, error) { 
	lac := models.Lactation{}
	db := models.GetDb()

	cm, ok := c.(models.CheckMilk)
	if !ok { return	ReserealizedCheckMilk{}, errors.New("error reserealizing") }

	if err := db.First(&lac, cm.LactationId).Error; err != nil {
		return ReserealizedCheckMilk{}, err
	}

	lacDate := lac.Date
	cmDate := cm.CheckDate
	milkingDays := cmDate.Sub(lacDate.Time)
	
	
	rcm.CheckMilk = cm
	rcm.MilkingDays = int(milkingDays.Hours() / 24)
	return rcm, nil
}

// ListAccounts lists all existing accounts
//
//	@Summary      Get list of check milks
//	@Description  Возращает список всех контрольных доек для конкретной коровы.
//	@Tags         Cows
//	@Param        id   path      int  true  "ID коровы для которой ищутся контрольные дойки"
//
// @Produce      json
// @Success      200  {array}   models.CheckMilk
// @Failure      500  {object}  map[string]error
// @Router       /cows/{id}/checkMilks [get]
func (f *Cows) CheckMilks() func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		cow := models.Cow{}
		db := models.GetDb()
		if err := db.Preload("Lactation").Preload("Lactation.CheckMilks").First(&cow, id).Error; err != nil {
			c.JSON(500, err)
			return
		}
		cms := []models.CheckMilk{}
		for _, lac := range cow.Lactation {
			cms = append(cms, lac.CheckMilks...)
		}
		res := make([]routes.Reserealizable, 0, len(cms))
		for _, cm := range(cms) {
			reserealizer := &ReserealizedCheckMilk{}
			if reserealized, err := reserealizer.FromBaseModel(cm); err != nil {
				c.JSON(500, err)
				return
			} else {
				res = append(res, reserealized)
			}
		}
		c.JSON(200, res)
	}
}
