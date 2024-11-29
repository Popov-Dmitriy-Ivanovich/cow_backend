package cows

import (
	"cow_backend/models"
	"cow_backend/routes"
	"errors"

	"github.com/gin-gonic/gin"
)

type ReserealizedCow struct {
	models.Cow
	BreedName *string `json:",omitempty"`
	SexName   *string `json:",omitempty"`
}

func (rc ReserealizedCow) GetReserealizer() routes.Reserealizer {
	return &rc
}
func (rc *ReserealizedCow) FromBaseModel(c any) (routes.Reserealizable, error) {
	cow, ok := c.(models.Cow)
	if !ok {
		return ReserealizedCow{}, errors.New("wrong type provided to get new cow from db cow")
	}
	db := models.GetDb()
	breed := models.Breed{}
	sex := models.Sex{}
	if err := db.First(&breed, cow.BreedId).Error; err != nil {
		return ReserealizedCow{}, err
	}
	if err := db.First(&sex, cow.SexId).Error; err != nil {
		return ReserealizedCow{}, err
	}
	rc.Cow = cow
	rc.BreedName = &breed.Name
	rc.SexName = &sex.Name
	return *rc, nil
}

// ListAccounts lists all existing accounts
//
//	@Summary      Get list of cows
//	@Description  Возращает конкретную корову.
//	@Tags         Cows
//	@Param        id   path      int  true  "ID конкретной коровы, чтобы ее вернуть"
//	@Produce      json
//	@Success      200  {object}   models.Cow
//	@Failure      500  {object}  map[string]error
//	@Router       /cows/{id} [get]
func (f *Cows) GetByID() func(*gin.Context) {
	return routes.GenerateReserealizingGetFunctionByID[models.Cow, ReserealizedCow](ReserealizedCow{})
	// return routes.GenerateGetFunction[models.Cow]("farm_id", "farm_group_id")
}

// @Summary      Get list of cows
// @Description  Возращает коров удовлетворяющих условиям фильтрации.
// @Tags         Cows
// @Param        farm_id    query     int  false  "ID фермы (НЕ хозяйства), к которой принадлежит корова"
// @Param 		 farm_group_id query int false "ID хозяйства (НЕ фермы), к которому принадлежит корова"
// @Success      200  {array}   models.Cow
// @Failure      500  {object}  map[string]error
// @Router       /cows [get]
func (f *Cows) GetByFilter() func(*gin.Context) {
	return routes.GenerateReserealizingGetFunctionByFilters[models.Cow, ReserealizedCow](ReserealizedCow{}, "farm_id", "farm_group_id")
}
