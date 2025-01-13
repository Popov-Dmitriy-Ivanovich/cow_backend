package cows

import (
	"cow_backend/models"

	"github.com/gin-gonic/gin"
)

// Exterior
// @Summary      Get exterior
// @Description  Возращает информацию об экстерьере, null, если нет
// @Tags         Cows
// @Param        id   path      int  true  "ID коровы для которой ищется экстерьер"
// @Produce      json
// @Success      200  {object}   models.Exterior
// @Failure      500  {object}   string
// @Router       /cows/{id}/exterior [get]
func (f *Cows) Exterior() func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		cow := models.Cow{}
		db := models.GetDb()
		if err := db.Preload("Exterior").First(&cow, id).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}
		c.JSON(200, cow.Exterior)
	}
}
