package cows

import (
	"cow_backend/models"
	"cow_backend/routes"
	"errors"

	"github.com/gin-gonic/gin"
)

type ReserealizedCow struct {
	models.Cow
	BreedName *string // порода, null если нет
	SexName   *string // пол, null если нет
	FarmName  *string // ферма на которой живет, null если нет
	HozHame   *string // хозяйство на котором живет, null, если нет

	Father *models.Cow // Отец
	Mother *models.Cow // Мать
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
	father := &models.Cow{}
	mother := &models.Cow{}
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

	if cow.FatherSelecs != nil {
		if err := db.Limit(1).Order("depart_date desc").Find(father,
			map[string]any{"selecs_number": cow.FatherSelecs}).Error; err != nil {
			return ReserealizedCow{}, err
		}
		rc.Father = father
	}

	if cow.MotherSelecs != nil {
		if err := db.Limit(1).Order("depart_date desc").Find(mother,
			map[string]any{"selecs_number": cow.MotherSelecs}).Error; err != nil {
			return ReserealizedCow{}, err
		}
		rc.Mother = mother
	}

	rc.Cow = cow
	rc.BreedName = &breed.Name
	rc.SexName = &sex.Name
	rc.FarmName = &farm.Name
	rc.HozHame = &hoz.Name
	return *rc, nil
}

// GetByID
// @Summary      Get concrete cow
// @Description  Возращает конкретную корову. Поля Father и Mother, имеют FatherId и MotherID null всегда, это неправильно, но так надо
// @Tags         Cows
// @Param        id   path      int  true  "ID конкретной коровы, чтобы ее вернуть"
// @Produce      json
// @Success      200  {object}   ReserealizedCow
// @Failure      500  {object}   string
// @Router       /cows/{id} [get]
func (f *Cows) GetByID() func(*gin.Context) {
	return routes.GenerateReserealizingGetFunctionByID[models.Cow, ReserealizedCow](ReserealizedCow{})
}

// GetByFilter
// @Summary      Get list of cows
// @Description  Возращает коров удовлетворяющих условиям фильтрации.
// @Tags         Cows
// @Param        farm_id    query     int  false  "ID фермы (НЕ хозяйства), к которой принадлежит корова"
// @Param 		 farm_group_id query int false "ID хозяйства (НЕ фермы), к которому принадлежит корова"
// @Success      200  {array}   ReserealizedCow
// @Failure      500  {object}  string
// @Router       /cows [get]
func (f *Cows) GetByFilter() func(*gin.Context) {
	return routes.GenerateReserealizingGetFunctionByFilters[models.Cow, ReserealizedCow](ReserealizedCow{}, "farm_id", "farm_group_id")
}
