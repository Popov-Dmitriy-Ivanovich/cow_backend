package updates

import (
	"cow_backend/models"

	"github.com/gin-gonic/gin"
)

type Update struct {
}

//	@Summary      Get update date and ID
//	@Description  Возращает дату апдейта БД
//	@Tags         Sexes
//	@Produce      json
//	@Success      200  {array}   models.Update
//	@Failure      500  {object}  map[string]error
//	@Router       /updates [get]
func (u *Update) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/updates")
	apiGroup.GET("/", func(c *gin.Context) {
		db := models.GetDb()
		dbUpdate := models.Update{}
		db.Order("date desc").First(&dbUpdate)
		c.JSON(200, dbUpdate)
	})
}
