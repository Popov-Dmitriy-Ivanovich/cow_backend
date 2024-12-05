package cows

import (
	"cow_backend/models"

	"github.com/gin-gonic/gin"
)

// ListAccounts lists all existing accounts
//
//	@Summary      Get list of children
//	@Description  Возращает список всех детей для конкретной коровы.
//	@Tags         Cows
//	@Param        id   path      int  true  "ID коровы для которой ищутся лактации"
//
// @Produce      json
// @Success      200  {array}   models.Lactation
// @Failure      500  {object}  map[string]error
// @Router       /cows/{id}/children [get]
func (f *Cows) Children() func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		cow := models.Cow{}
		db := models.GetDb()
		if err := db.First(&cow, id).Error; err != nil {
			c.JSON(500, err)
			return
		}

		selecs := cow.SelecsNumber
		cld := []models.Cow{}
		db.Where("mother_selecs = ? OR father_selecs = ?", selecs, selecs).Find(&cld)

		c.JSON(200, cld)
	}
}
