package gui

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Gui struct {
}

func (s *Gui) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/gui")
	apiGroup.GET("/cowLoad", s.CowLoad())
	apiGroup.GET("/checkMilkLoad", s.CheckMilkLoad())
	apiGroup.GET("/eventLoad", s.EventLoad())
	apiGroup.GET("/geneticLoad", s.GeneticLoad())
	apiGroup.GET("/gradeLoad", s.GradeLoad())
	apiGroup.GET("/lactationLoad", s.LactationLoad())
	apiGroup.GET("/exteriorLoad", func(c *gin.Context) { c.HTML(200, "ExteriorLoadPage.tmpl", gin.H{"title": "экстерьер"}) })
	apiGroup.GET("/gtcLoad", func(c *gin.Context) { c.HTML(200, "GtcLoadPage.tmpl", gin.H{"title": "gtc"}) })
	apiGroup.GET("/partnerLoad", func(c *gin.Context) { c.HTML(200, "PartnerLoadPage.tmpl", gin.H{"title": "gtc"}) })
	apiGroup.GET("/documentLoad", func(c *gin.Context) { c.HTML(200, "DocumentLoadPage.tmpl", gin.H{"title": "document"}) })
	apiGroup.GET("/exteriorDataLoad", func(c *gin.Context) { c.HTML(200, "ExteriorDataLoadPage.tmpl", gin.H{"title": "document"}) })
	apiGroup.GET("", s.Index())
}

func (s *Gui) CheckMilkLoad() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "CheckMilkLoadPage.tmpl", gin.H{"title": "контрольные доения"})
	}
}

func (s *Gui) EventLoad() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "EventLoadPage.tmpl", gin.H{"title": "вет. события"})
	}
}

func (s *Gui) GeneticLoad() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "GeneticLoadPage.tmpl", gin.H{"title": "генотипирование"})
	}
}

func (s *Gui) GradeLoad() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "GradeLoadPage.tmpl", gin.H{"title": "оценки"})
	}
}

func (s *Gui) LactationLoad() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "LactationLoadPage.tmpl", gin.H{"title": "лактации"})
	}
}
