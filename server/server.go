package main

import (
	"cow_backend/models"
	"cow_backend/routes"
	"cow_backend/routes/analitics"
	"cow_backend/routes/breeds"
	checkmilks "cow_backend/routes/check_milks"
	"cow_backend/routes/cows"
	dailymilks "cow_backend/routes/daily_milks"
	"cow_backend/routes/districts"
	"cow_backend/routes/farms"
	"cow_backend/routes/gui"
	"cow_backend/routes/lactations"
	"cow_backend/routes/load"
	"cow_backend/routes/monogenetic_illnesses"
	"cow_backend/routes/regions"
	"cow_backend/routes/sexes"
	"fmt"

	// "net/http"
	_ "cow_backend/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// "github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title           GenMilk API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      genmilk.ru
// @BasePath  /api

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	models.GetDb()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	apiGroup := r.Group("/api")
	routes.WriteRoutes(apiGroup, &routes.Api{}, &regions.Regions{}, &farms.Farms{}, &breeds.Breeds{}, &checkmilks.CheckMilks{},
		&cows.Cows{}, &dailymilks.DailyMilk{}, &districts.Districts{}, &lactations.Lactations{}, &sexes.Sexes{}, &analitics.Analitics{},
		&monogenetic_illnesses.MonogeneticIllneses{}, &gui.Gui{}, &load.Load{})

	apiGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()

	fmt.Println("Hello world")
}
