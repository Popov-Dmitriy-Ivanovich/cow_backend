package main

import (
	"cow_backend/models"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	models.GetDb()
	fmt.Println("Hello world")
}
