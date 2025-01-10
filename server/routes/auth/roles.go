package auth

import (
	"cow_backend/models"

	"github.com/gin-gonic/gin"
)

// @Summary      LOGIN
// @Tags         LOGIN
// @Produce      json
// @Success      200  {array}   models.Role
// @Failure      400  {object}  map[string]error
// @Failure      401  {object}  map[string]error
// @Failure      500  {object}  map[string]error
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
