package routes

import (
	"cow_backend/models"
	"regexp"

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

func GenerateGetFunctionById[T any]() func(c *gin.Context) {
	return func(c *gin.Context) {
		db := models.GetDb()
		id := c.Param("id")
		objs := new(T)

		if m, _ := regexp.MatchString("^[0-9]+$", id); id == "" || !m {
			c.JSON(422, "wrong ID provided")
			return
		}

		if err := db.First(&objs, id).Error; err != nil {
			c.JSON(404, "record not found")
			return
		}
		c.JSON(200, objs)
	}

}

func GenerateGetFunctionByFilters[T any](allowEmpty bool, filters ...string) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := models.GetDb()

		objs := []T{}

		query := db.Where("true")
		emptyConds := true
		for _, filter := range filters {
			value, _ := c.GetQuery(filter)
			if value == "" {
				continue
			}
			emptyConds = false
			if value == "null" {
				query = query.Where(map[string]any{filter: nil})
			} else {
				query = query.Where(map[string]string{filter: value})
			}

		}
		if emptyConds && !allowEmpty {
			c.JSON(422, "all filters empty")
			return
		}
		if err := query.Find(&objs).Error; err != nil {
			c.JSON(404, "record not found")
			return
		}
		c.JSON(200, objs)
	}

}

type Reserealizer interface {
	FromBaseModel(value any) (Reserealizable, error)
}

type Reserealizable interface {
	GetReserealizer() Reserealizer
}

func GenerateReserealizingGetFunctionByID[DbModel any, R Reserealizable](value R) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := models.GetDb()
		id := c.Param("id")
		obj := new(DbModel)

		if m, _ := regexp.MatchString("^[0-9]+$", id); id == "" || !m {
			c.JSON(422, "id empty")
		}

		if err := db.First(&obj, id).Error; err != nil {
			c.JSON(404, "record not found")
			return
		}
		if obj == nil {
			c.JSON(404, "Found object is nil")
			return
		}
		reserealizer := value.GetReserealizer()
		if res, err := reserealizer.FromBaseModel(*obj); err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, res)
		}

	}

}

func GenerateReserealizingGetFunctionByFilters[DbModel any, R Reserealizable](value R, filters ...string) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := models.GetDb()
		objs := []DbModel{}
		query := db.Where("true")
		emptyConds := true

		for _, filter := range filters {
			value, _ := c.GetQuery(filter)
			if value == "" {
				continue
			}
			emptyConds = false
			query = query.Where(map[string]string{filter: value})
		}

		if emptyConds {
			c.JSON(422, "all filters empty")
			return
		}

		if err := query.Find(&objs).Error; err != nil {
			c.JSON(404, "record not found")
			return
		}

		res := []R{}
		reserealizer := value.GetReserealizer()

		for _, obj := range objs {
			if reserealized, err := reserealizer.FromBaseModel(obj); err != nil {
				c.JSON(500, err.Error())
				return
			} else if typed, ok := reserealized.(R); !ok {
				c.JSON(500, "wrong type in reserealizer")
				return
			} else {
				res = append(res, typed)
			}
		}
		c.JSON(200, res)
	}

}
