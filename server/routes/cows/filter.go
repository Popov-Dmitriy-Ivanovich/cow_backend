package cows

import (
	"cow_backend/filters/cows_filter"
	"cow_backend/models"
	"encoding/json"
	"net/http"
	"strconv"

	// "sort"
	"time"

	// "fmt"
	"io"

	"github.com/gin-gonic/gin"
)

type FilterSerializedCow struct {
	ID                        uint                        `validate:"required" example:"123"`                                   // ID коровы
	RSHNNumber                *string                     `validate:"required" example:"123"`                                   // РСХН номер коровы
	InventoryNumber           *string                     `validate:"required" example:"321"`                                   // Инвентарный номер коровы
	Name                      string                      `validate:"required" example:"Буренка"`                               // Кличка коровы
	FarmGroupName             string                      `validate:"required" example:"ООО Аурус"`                             // Название хозяйства, в котором корова
	BirthDate                 models.DateOnly             `validate:"required"`                                                 // Дата рождения коровы
	Genotyped                 bool                        `validate:"required" example:"true"`                                  // Факт генотипирования коровы
	Approved                  bool                        `validate:"required" example:"true"`                                  // Подтверждена ли админом
	DepartDate                *models.DateOnly            `json:",omitempty" validate:"optional"`                               // Дата выбытия
	BreedName                 *string                     `json:",omitempty" validate:"optional" example:"Какая-нибудь порода"` // Название породы коровы
	CheckMilkDate             []models.DateOnly           `json:",omitempty" validate:"optional"`                               // Дата контрольного доения
	InsemenationDate          []models.DateOnly           `json:",omitempty" validate:"optional"`                               // Дата осеменения
	CalvingDate               []models.DateOnly           `json:",omitempty" validate:"optional"`                               // Дата отела
	BirkingDate               *models.DateOnly            `json:",omitempty" validate:"optional"`                               // Дата перебирковки
	GenotypingDate            *models.DateOnly            `json:",omitempty" validate:"optional"`                               // Дата генотипирования
	InbrindingCoeffByFamily   *float64                    `json:",omitempty" validate:"optional" example:"3.14"`                // Коэф. инбриндинга по родословной
	InbrindingCoeffByGenotype *float64                    `json:",omitempty" validate:"optional" example:"3.14"`                // Коэф. инбриндинга по генотипу
	MonogeneticIllneses       []models.GeneticIllnessData `json:",omitempty" validate:"optional"`                               // моногенные заболевания
	ExteriorRating            *float64                    `json:",omitempty" validate:"optional"`                               // Оценка экстерьера
	SexName                   *string                     `json:",omitempty" validate:"optional"`                               // Название породы
	HozName                   *string                     `json:",omitempty" validate:"optional"`                               // Название хозяйства

	DeathDate   *models.DateOnly `json:",omitempty" validate:"optional"` // Дата смерти
	IsDead      *bool            `json:",omitempty" validate:"optional"` // Факт смерти
	IsTwins     *bool            `json:",omitempty" validate:"optional"` // Факт рождения близнецов
	IsStillBorn *bool            `json:",omitempty" validate:"optional"` // Факт мертворождения
	IsAborted   *bool            `json:",omitempty" validate:"optional"` // Факт аборта
	Events      []models.Event   `json:",omitempty" validate:"optional"` // Вет события
	IsGenotyped *bool            `json:",omitempty" validate:"optional"` // Факт генотипирования
	CreatedAt   *models.DateOnly `json:",omitempty" validate:"optional"` // Дата внесения информации о корове в БД
}

func serializeByFilter(c *models.Cow, filter *cows_filter.CowsFilter) FilterSerializedCow {
	res := FilterSerializedCow{
		ID:              c.ID,
		RSHNNumber:      c.RSHNNumber,
		InventoryNumber: c.InventoryNumber,
		Name:            c.Name,
		FarmGroupName:   c.FarmGroup.Name,
		BirthDate:       c.BirthDate,
		Genotyped:       c.Genetic != nil,
		Approved:        c.Approved != 0,
	}
	if filter.DepartDateTo != nil && *filter.DepartDateTo != "" ||
		filter.DepartDateFrom != nil && *filter.DepartDateFrom != "" {
		res.DepartDate = c.DepartDate
	}

	if len(filter.BreedId) != 0 {
		res.BreedName = &c.Breed.Name
	}

	if filter.InbrindingCoeffByFamilyFrom != nil || filter.InbrindingCoeffByFamilyTo != nil {
		res.InbrindingCoeffByFamily = c.InbrindingCoeffByFamily
	}
	if filter.InbrindingCoeffByGenotypeFrom != nil || filter.InbrindingCoeffByGenotypeTo != nil {
		res.InbrindingCoeffByGenotype = c.Genetic.InbrindingCoeffByGenotype
	}
	if filter.GenotypingDateFrom != nil && *filter.GenotypingDateFrom != "" ||
		filter.GenotypingDateTo != nil && *filter.GenotypingDateTo != "" {
		res.GenotypingDate = c.Genetic.ResultDate
	}
	if len(filter.MonogeneticIllneses) != 0 || filter.HasAnyIllnes != nil {
		res.MonogeneticIllneses = c.Genetic.GeneticIllnessesData
	}
	if filter.ControlMilkingDateFrom != nil && *filter.ControlMilkingDateFrom != "" ||
		filter.ControlMilkingDateTo != nil && *filter.ControlMilkingDateTo != "" {
		for _, lactation := range c.Lactation {
			for _, cm := range lactation.CheckMilks {
				if filter.ControlMilkingDateFrom != nil {
					date, err := time.Parse(time.DateOnly, *filter.ControlMilkingDateFrom)
					if err != nil {
						continue
					}
					if !date.Equal(cm.CheckDate.Time) && date.After(cm.CheckDate.Time) {
						continue
					}
				}
				if filter.ControlMilkingDateTo != nil {
					date, err := time.Parse(time.DateOnly, *filter.ControlMilkingDateTo)
					if err != nil {
						continue
					}
					if !date.Equal(cm.CheckDate.Time) && date.Before(cm.CheckDate.Time) {
						continue
					}
				}
				res.CheckMilkDate = append(res.CheckMilkDate, cm.CheckDate)
			}
		}
	}

	if filter.InseminationDateFrom != nil && *filter.InseminationDateFrom != "" ||
		filter.InseminationDateTo != nil && *filter.InseminationDateTo != "" {
		for _, lac := range c.Lactation {
			if filter.InseminationDateFrom != nil {
				date, err := time.Parse(time.DateOnly, *filter.InseminationDateFrom)
				if err != nil {
					continue
				}
				if !date.Equal(lac.InsemenationDate.Time) && date.After(lac.InsemenationDate.Time) {
					continue
				}
			}
			if filter.InseminationDateTo != nil {
				date, err := time.Parse(time.DateOnly, *filter.InseminationDateTo)
				if err != nil {
					continue
				}
				if !date.Equal(lac.InsemenationDate.Time) && date.Before(lac.InsemenationDate.Time) {
					continue
				}
			}
			res.InsemenationDate = append(res.InsemenationDate, lac.InsemenationDate)

		}
	}

	if filter.CalvingDateFrom != nil && *filter.CalvingDateFrom != "" ||
		filter.CalvingDateTo != nil && *filter.CalvingDateTo != "" {
		for _, lac := range c.Lactation {
			if filter.CalvingDateFrom != nil && *filter.CalvingDateFrom != "" {
				date, err := time.Parse(time.DateOnly, *filter.CalvingDateFrom)
				if err != nil {
					continue
				}
				if !date.Equal(lac.CalvingDate.Time) && date.After(lac.CalvingDate.Time) {
					continue
				}
			}
			if filter.CalvingDateTo != nil && *filter.CalvingDateTo != "" {
				date, err := time.Parse(time.DateOnly, *filter.CalvingDateTo)
				if err != nil {
					continue
				}
				if !date.Equal(lac.CalvingDate.Time) && date.Before(lac.CalvingDate.Time) {
					continue
				}
			}
			res.CalvingDate = append(res.CalvingDate, lac.CalvingDate)
		}
	}

	if filter.BirkingDateFrom != nil && *filter.BirkingDateFrom != "" ||
		filter.BirkingDateTo != nil && *filter.BirkingDateTo != "" {
		res.BirkingDate = c.BirkingDate
	}
	if filter.IllDateFrom != nil && *filter.IllDateFrom != "" ||
		filter.IllDateTo != nil && *filter.IllDateTo != "" {
		for _, event := range c.Events {
			if event.EventTypeId > 4 {
				continue
			}
			eventDate := event.Date.Time
			if filter.IllDateFrom != nil && *filter.IllDateFrom != "" {
				dateFrom, err := time.Parse(time.DateOnly, *filter.IllDateFrom)
				if err != nil {
					continue
				}
				if dateFrom.After(eventDate) && !dateFrom.Equal(eventDate) {
					continue
				}
			}
			if filter.IllDateTo != nil && *filter.IllDateTo != "" {
				dateTo, err := time.Parse(time.DateOnly, *filter.IllDateTo)
				if err != nil {
					continue
				}
				if dateTo.Before(eventDate) && !dateTo.Equal(eventDate) {
					continue
				}
			}
			res.Events = append(res.Events, event)
		}
	}
	if len(filter.Sex) != 0 {
		res.SexName = &c.Sex.Name
	}
	if filter.HozId != nil {
		res.HozName = &c.FarmGroup.Name
	}
	if filter.IsDead != nil {
		res.DeathDate = c.DeathDate
		res.IsDead = filter.IsDead
	}
	if filter.IsStillBorn != nil {
		res.IsStillBorn = filter.IsStillBorn
	}
	if filter.IsAborted != nil {
		res.IsAborted = filter.IsAborted
	}
	if filter.IsTwins != nil {
		res.IsTwins = filter.IsTwins
	}
	if filter.ExteriorFrom != nil || filter.ExteriorTo != nil {
		res.ExteriorRating = &c.Exterior.Rating
	}
	if filter.IsGenotyped != nil {
		res.IsGenotyped = filter.IsGenotyped
	}
	if filter.CreatedAtFrom != nil && *filter.CreatedAtFrom != "" ||
		filter.CreatedAtTo != nil && *filter.CreatedAtTo != "" {
		res.CreatedAt = &models.DateOnly{Time: c.CreatedAt}
	}
	return res
}

// Filter
// @Summary      Get filtered list of cows
// @Description  Возращает словарь с двумя ключами "N", "LST". По ключу "N" - общее кол-во найденных коров,
// @Description  по ключу "LST" массив объектов filterSerealizedCow (см. Models)
// @Tags         Cows
// @Param        filter    body     cows_filter.CowsFilter  true  "applied filters"
// @Accept       json
// @Produce      json
// @Success      200  {array}   map[string]FilterSerializedCow
// @Failure      422  {object}  string
// @Failure      500  {object}  string
// @Router       /cows/filter [post]
func (c *Cows) Filter() func(*gin.Context) {
	return func(c *gin.Context) {

		roleId, exists := c.Get("RoleId")
		if !exists {
			c.JSON(http.StatusInternalServerError, "RoleId не найден в контексте")
			return
		}

		jsonData, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(500, err.Error())
			return
		}

		bodyData := cows_filter.CowsFilter{}
		if len(jsonData) != 0 {
			err = json.Unmarshal(jsonData, &bodyData)
			if err != nil {
				c.JSON(422, err.Error())
				return
			}
		}

		if roleId == 1 {
			farmIdStr, exists := c.Get("FarmId")
			if !exists {
				c.JSON(http.StatusInternalServerError, "FarmId не найден в контексте")
				return
			}

			farmIdUint64, err := strconv.ParseUint(farmIdStr.(string), 10, 0)
			if err != nil {
				c.JSON(http.StatusInternalServerError, "Ошибка преобразования FarmId: "+err.Error())
				return
			}

			farmId := uint(farmIdUint64)
			bodyData.HozId = &farmId
		}

		if roleId == 2 {
			regionIdStr, exists := c.Get("RegionId")
			if !exists {
				c.JSON(http.StatusInternalServerError, "RegionId не найден в контексте")
				return
			}

			regionIdUint64, err := strconv.ParseUint(regionIdStr.(string), 10, 0)
			if err != nil {
				c.JSON(http.StatusInternalServerError, "Ошибка преобразования RegionId: "+err.Error())
				return
			}

			regionId := uint(regionIdUint64)
			bodyData.RegionId = &regionId
		}

		db := models.GetDb()
		query := db.Model(&models.Cow{}).Preload("FarmGroup").Preload("Genetic").Where("approved <> -1")
		if nQuery, err := AddFiltersToQuery(bodyData, query); err != nil {
			c.JSON(422, err.Error())
			return
		} else {
			query = nQuery
		}
		// ====================================================================================================
		// ======================================= PAGINATION =================================================
		// ====================================================================================================
		recordsPerPage := uint64(50)
		pageNumber := uint64(1)
		if bodyData.EntitiesOnPage != nil {
			recordsPerPage = uint64(*bodyData.EntitiesOnPage)
		}

		if bodyData.PageNumber != nil {
			pageNumber = uint64(*bodyData.PageNumber)
		}
		// ====================================================================================================
		// ==================================Get final query result ===========================================
		// ====================================================================================================
		resCount := int64(0)
		if err := query.Count(&resCount).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}

		query = query.Limit(int(recordsPerPage)).Offset(int(recordsPerPage) * int(pageNumber-1)).Order("inventory_number")
		dbCows := []models.Cow{}
		if err := query.Debug().Find(&dbCows).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}

		res := make([]FilterSerializedCow, 0, len(dbCows))
		for _, c := range dbCows {
			res = append(res, serializeByFilter(&c, &bodyData))
		}

		// fmt.Print(query)
		c.JSON(200, gin.H{
			"N":   resCount,
			"LST": res,
			// "query": query,
		})
	}
}
