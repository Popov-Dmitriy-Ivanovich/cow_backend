package breeds

import (
	"cow_backend/models"
	"cow_backend/routes"

	"github.com/gin-gonic/gin"
)

// ListAccounts lists all existing accounts
//
//	@Summary      Get list of breeds
//	@Description  Возращает список пород
//	@Tags         Breeds
//	@Param        id    query     int  false  "ID конкретной породы, если нужно вернуть одну"
//	@Produce      json
//	@Success      200  {array}   models.Breed
//	@Failure      500  {object}  map[string]error
//	@Router       /breeds/get [get]
func (f *Breeds) Get() func(*gin.Context) {
	return routes.GenerateGetFunction[models.Breed]()
}
