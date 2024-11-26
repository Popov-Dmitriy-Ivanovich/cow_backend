package routes

import "github.com/gin-gonic/gin"

type Api struct {

}

func (a *Api) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/api")
	apiGroup.POST("/find_in_cows_post", a.FindInCowsPost())
}