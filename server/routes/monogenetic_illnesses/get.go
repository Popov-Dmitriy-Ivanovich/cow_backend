package monogenetic_illnesses

import (
	"cow_backend/models"
	"cow_backend/routes"

	"github.com/gin-gonic/gin"
)

// ListAccounts lists all existing accounts
//
//	@Summary      Get list of monogenetic illnesses
//	@Description  Возращает список генетических заболеваний
//	@Tags         MonogeneticIllnessses
//	@Produce      json
//	@Success      200  {array}   models.GeneticIllness
//	@Failure      500  {object}  map[string]error
//	@Router       /monogenetic_illnesses [get]
func (s *MonogeneticIllneses) Get() func(*gin.Context) {
	return routes.GenerateGetFunctionByFilters[models.GeneticIllness](true)
}
