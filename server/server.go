package main

import (
	"cow_backend/models"
	"cow_backend/routes"
	"fmt"
	// "net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	models.GetDb()
	
	r := gin.Default()
	apiHandlers := &routes.Api{}
	apiHandlers.WriteRoutes(&r.RouterGroup)
	r.Run()

	
	fmt.Println("Hello world")
}
