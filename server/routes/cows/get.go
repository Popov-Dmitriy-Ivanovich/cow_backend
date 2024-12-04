package cows

import (
	"cow_backend/models"
	"cow_backend/routes"
	"errors"

	"github.com/gin-gonic/gin"
)

type ReserealizedCow struct {
	models.Cow
	BreedName *string          // порода, null если нет
	SexName   *string          // пол, null если нет
	FarmName  *string          // ферма на которой живет, null если нет
	HozHame   *string          // хозяйство на котором живет, null, если нет
	Father    *ReserealizedCow // Отец, null если нет
	Mother    *ReserealizedCow // Мать, null, если нет
	Genetic   *models.Genetic  // Информация о генотипировании, null если нет
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
	farm := models.Farm{}
	hoz := models.Farm{}

	if err := db.First(&breed, cow.BreedId).Error; err != nil {
		return ReserealizedCow{}, err
	}
	if err := db.First(&sex, cow.SexId).Error; err != nil {
		return ReserealizedCow{}, err
	}
	if cow.FarmID != nil {
		if err := db.First(&farm, cow.FarmID).Error; err != nil {
			return ReserealizedCow{}, err
		}
	}
	if err := db.First(&hoz, cow.FarmGroupId).Error; err != nil {
		return ReserealizedCow{}, err
	}

	
	rc.Cow = cow
	genetic := &models.Genetic{}
	if qRes := db.Preload("GeneticIllnesses").Limit(1).Find(genetic, map[string]any{"cow_id": cow.ID}); qRes.Error != nil {
		return ReserealizedCow{}, qRes.Error
	} else if qRes.RowsAffected != 0 {
		rc.Genetic = genetic
	}
	exterior := &models.Exterior{}
	if qRes := db.Limit(1).Find(exterior, map[string]any{"cow_id": cow.ID}); qRes.Error != nil {
		return ReserealizedCow{}, qRes.Error
	} else if qRes.RowsAffected != 0 {
		rc.Exterior = exterior
	}
	
	rc.BreedName = &breed.Name
	rc.SexName = &sex.Name
	rc.FarmName = &farm.Name
	rc.HozHame = &hoz.Name
	return *rc, nil
}

// ListAccounts lists all existing accounts
//
//	@Summary      Get list of cows
//	@Description  Возращает конкретную корову. Поля Father и Mother, имеют FatherId и MotherID null всегда, это неправильно, но так надо
//	@Tags         Cows
//	@Param        id   path      int  true  "ID конкретной коровы, чтобы ее вернуть"
//	@Produce      json
//	@Success      200  {object}   ReserealizedCow
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
// @Success      200  {object}   ReserealizedCow
// @Failure      500  {object}  map[string]error
// @Router       /cows [get]
func (f *Cows) GetByFilter() func(*gin.Context) {
	return routes.GenerateReserealizingGetFunctionByFilters[models.Cow, ReserealizedCow](ReserealizedCow{}, "farm_id", "farm_group_id")
}
