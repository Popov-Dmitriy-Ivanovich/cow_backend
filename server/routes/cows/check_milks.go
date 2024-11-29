package cows

import (
	"cow_backend/models"

	"github.com/gin-gonic/gin"
)

// ListAccounts lists all existing accounts
//
//	@Summary      Get list of check milks
//	@Description  Возращает список всех контрольных доек для конкретной коровы.
//	@Tags         Cows
//	@Param        id   path      int  true  "ID конкретной коровы, чтобы ее вернуть"
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
		c.JSON(200, cms)
	}
}
