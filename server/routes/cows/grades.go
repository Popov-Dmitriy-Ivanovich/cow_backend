package cows

import (
	"cow_backend/models"

	"github.com/gin-gonic/gin"
)

// Grades
// @Summary      Get grades
// @Description  Возращает словарь с двумя ключам "ByRegion" - оценки по региону и "ByHoz" - оценки по хозяйству
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
		if err := db.Preload("GradeRegion").Preload("GradeHoz").First(&cow, id).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}
		c.JSON(200, gin.H{
			"ByRegion": cow.GradeRegion,
			"ByHoz":    cow.GradeHoz,
		})
	}
}
