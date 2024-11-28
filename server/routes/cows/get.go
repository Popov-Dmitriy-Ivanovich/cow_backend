package cows

import (
	"cow_backend/models"
	"cow_backend/routes"

	"github.com/gin-gonic/gin"
)

// ListAccounts lists all existing accounts
//
//	@Summary      Get list of cows
//	@Description  Возращает список коров.
//	@Tags         Cows
//	@Param        id    query     int  false  "ID конкретной коровы, чтобы ее вернуть"
//	@Param        farm_id    query     int  false  "ID фермы (НЕ хозяйства), к которой принадлежит корова"
//	@Param 		  farm_group_id query int false "ID хозяйства (НЕ фермы), к которому принадлежит корова"
//	@Produce      json
//	@Success      200  {array}   models.Cow
//	@Failure      500  {object}  map[string]error
//	@Router       /cows/get [get]
func (f *Cows) Get() func(*gin.Context) {
	return routes.GenerateGetFunction[models.Cow]("farm_id", "farm_group_id")
}
