package cows

import (
	"cow_backend/models"

	"github.com/gin-gonic/gin"
)

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
