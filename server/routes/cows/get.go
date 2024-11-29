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
	return routes.GenerateReserealizingGetFunction[models.Cow, ReserealizedCow]("farm_id", "farm_group_id")
	// return routes.GenerateGetFunction[models.Cow]("farm_id", "farm_group_id")
}
