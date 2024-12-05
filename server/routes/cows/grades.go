package cows

import (
	"cow_backend/models"

	"github.com/gin-gonic/gin"
)

//	@Summary      Get list of children
//	@Description  Возращает список всех оценок конкретной коровы.
//	@Tags         Cows
//	@Param        id   path      int  true  "ID коровы для которой ищутся оценки"
//
// @Produce      json
// @Success      200  {object}   map[string]models.Grade
// @Failure      500  {object}  map[string]error
// @Router       /cows/{id}/grades [get]
func (c *Cows) Grades() func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		db := models.GetDb()
		cow := models.Cow{}
		if err := db.Preload("GradeRegion").Preload("GradeHoz").First(&cow, id).Error; err != nil {
			c.JSON(500, err)
			return
		}
		c.JSON(200, gin.H{
			"ByRegion": cow.GradeRegion,
			"ByHoz":    cow.GradeHoz,
		})
	}
}
