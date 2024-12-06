package analitics

import "github.com/gin-gonic/gin"

type Analitics struct {
}

func (b *Analitics) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/analitics")
	genotypedWriter := Genotyped{}
	checkMilksWriter := CheckMilks{}
	checkMilksWriter.WriteRoutes(apiGroup)
	genotypedWriter.WriteRoutes(apiGroup)
}
