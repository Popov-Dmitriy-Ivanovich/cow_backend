package cows

import (
	"cow_backend/models"
	"cow_backend/routes"

	"github.com/gin-gonic/gin"
)

// ListAccounts lists all existing accounts
//
//	@Summary      Get list of farms
//	@Description  Get list of farms.
//	@Description  DOES NOT RETURN SUBOBJECTS
//	@Tags         Cows
//	@Param        id    query     int  false  "id of farm to return"
//	@Produce      json
//	@Success      200  {array}   models.Cow
//	@Failure      500  {object}  map[string]error
//	@Router       /cows/get [get]
func (f *Cows) Get() func(*gin.Context) {
	return routes.GenerateGetFunction[models.Cow]()
}
