package analitics

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AnaliticsAuthMiddleware
// used to restrict access of users to analitics
func AnaliticsAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		region := c.Param("region")
		district := c.Param("district")
		userRoleId, exists := c.Get("RoleId")
		if !exists {
			c.JSON(http.StatusInternalServerError, "RoleId не найден в контексте")
			c.Abort()
			return
		}
		// Фермер (роль 1) может видеть аналитику только по своему району
		if district != "" && (userRoleId == 1) {
			userDistrictId, exists := c.Get("DistId")
			if !exists || userDistrictId != district {
				c.JSON(421, "Нет доступа к району")
				c.Abort()
				return
			}
		}
		// Фермер (роль 1) и региональный чиновник (роль 2) могут видеть только свой регион
		if region != "" && (userRoleId == 1 || userRoleId == 2) {
			regionId, exists := c.Get("RegionId")
			if !exists || regionId != region {
				c.JSON(http.StatusInternalServerError, "нет доступа к региону")
				c.Abort()
				return
			}
		}
	}
}
