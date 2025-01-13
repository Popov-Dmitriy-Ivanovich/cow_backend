package auth

import (
	"cow_backend/models"

	"github.com/gin-gonic/gin"
)

// Roles
// @Summary      LOGIN
// @Description  Возращает список ролей
// @Tags         LOGIN
// @Produce      json
// @Success      200  {array}   models.Role
// @Failure      400  {object}  string
// @Failure      401  {object}  string
// @Failure      500  {object}  string
// @Router       /auth/roles [get]
func (s *Auth) Roles() func(*gin.Context) {
	return func(c *gin.Context) {
		roles := []models.Role{}
		db := models.GetDb()
		if err := db.Find(&roles).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}
		c.JSON(200, roles)
	}
}
