package routes

import (
	"cow_backend/models"

	"github.com/gin-gonic/gin"
)

func WriteRoutes(rg *gin.RouterGroup, rw ...RouteWriter) {
	for _, rw := range rw {
		rw.WriteRoutes(rg)
	}
}

type RouteWriter interface {
	WriteRoutes(*gin.RouterGroup)
}

func GenerateGetFunction[T any](filters ...string) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := models.GetDb()
		id, _ := c.GetQuery("id")
		objs := []T{}
		query := db.Where("true")
		if id != "" {
			query = db.Where("id = ?", id)
		}
		for _, filter := range filters {
			value, _ := c.GetQuery(filter)
			if value == "" {
				continue
			}
			query = query.Where(map[string]string{filter: value})
		}
		if err := query.Find(&objs, id).Error; err != nil {
			c.JSON(500, err)
			return
		}
		c.JSON(200, objs)
	}

}
