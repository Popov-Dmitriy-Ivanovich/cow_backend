package cows

import (
	"cow_backend/models"

	"github.com/gin-gonic/gin"
)

// Genetic
// @Summary      Get list of check milks
// @Description  Возращает генетическую информацию для коровы, null, если нет
// @Tags         Cows
// @Param        id   path      int  true  "ID коровы для которой ищется генетическая информация"
// @Produce      json
// @Success      200  {object}   models.Genetic
// @Failure      500  {object}   string
// @Router       /cows/{id}/genetic [get]
func (f *Cows) Genetic() func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		cow := models.Cow{}
		db := models.GetDb()
		if err := db.
			Preload("Genetic").
			Preload("Genetic.GeneticIllnessesData").
			Preload("Genetic.GeneticIllnessesData.Status").
			Preload("Genetic.GeneticIllnessesData.Illness").
			First(&cow, id).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}
		c.JSON(200, cow.Genetic)
	}
}
