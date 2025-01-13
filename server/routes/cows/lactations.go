package cows

import (
	"cow_backend/models"
	"sort"

	"github.com/gin-gonic/gin"
)

// Lactations
// @Summary      Get list of lactations
// @Description  Возращает список всех лактаций для конкретной коровы.
// @Tags         Cows
// @Param        id   path      int  true  "ID коровы для которой ищутся лактации"
// @Produce      json
// @Success      200  {array}   models.Lactation
// @Failure      500  {object}  string
// @Router       /cows/{id}/lactations [get]
func (f *Cows) Lactations() func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		cow := models.Cow{}
		db := models.GetDb()
		if err := db.Preload("Lactation").First(&cow, id).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}
		sort.Slice(cow.Lactation, func(i, j int) bool {
			return cow.Lactation[i].Number < cow.Lactation[j].Number
		})
		c.JSON(200, cow.Lactation)
	}
}
