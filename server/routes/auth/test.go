package auth

import (
	"github.com/gin-gonic/gin"
)

func (a *Auth) Test() func(*gin.Context) {
	return func(c *gin.Context) {

		roleId, _ := c.Get("RoleId")
		regionId, _ := c.Get("RegionId")
		farmId, _ := c.Get("FarmId")
		res := gin.H{"roleId": roleId, "regionId": regionId, "farmId": farmId}
		c.JSON(200, res)
	}
}
