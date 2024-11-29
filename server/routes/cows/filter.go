package cows

import (
	"cow_backend/models"
	"encoding/json"
	"time"

	// "fmt"
	"io"

	"github.com/gin-gonic/gin"
)

type cowsFilter struct { // Фильтр коров
	SearchQuery            *string `example:"Буренка" validate:"optional"` // Имя, номер РСХН или инвентарный номер, по которым ищется корова
	PageNumber             *uint   `default:"1" validate:"optional"`       // Номер страницы для отображения
	EntitiesOnPage         *uint   `default:"50" validate:"optional"`      // Количество сущностей на странице
	Sex                    []uint  //ID пола коровы (если нужно несколько разных полов - несколько ID)
	HozId                  *uint   `example:"1" validate:"optional"`          //ID фермы, для которой ищутся коровы
	BirthDateFrom          *string `example:"1800-01-21" validate:"optional"` //Фильтр по дню рождения коровы ОТ (возращает всех кто родился в эту дату или позднее)
	BirthDateTo            *string `example:"2800-01-21" validate:"optional"` //Фильтр по дню рождения коровы ОТ (возращает всех кто родился в эту дату или раньше)
	IsDead                 *bool   `default:"false" validate:"optional"`      //Фильтр живых/мертвых коров (true - ищет мертвых, false - живых)
	DepartDateFrom         *string `example:"1800-01-21" validate:"optional"` //Фильтр по дате открепления коровы ищет всех коров открепленных от коровника в эту дату или позднее
	DepartDateTo           *string `example:"2800-01-21" validate:"optional"` //Фильтр по дате открепления коровы ищет всех коров открепленных от коровника в эту дату или раньше
	BreedId                []uint  //Фильтр по ID пород несколько ID пород - возращает всех коров, ID пород которых в списке
	GenotypingDateFrom     *string `example:"1800-01-21" validate:"optional"` //??? Не реализован
	GenotypingDateTo       *string `example:"2800-01-21" validate:"optional"` //??? Не реализован
	ControlMilkingDateFrom *string `example:"1800-01-21" validate:"optional"` // Фильтр по дате контрольной дойки, ищет коров у которых была контрольная дойка в эту дату или позднее
	ControlMilkingDateTo   *string `example:"2800-01-21" validate:"optional"` // Фильтр по дате контрольной дойки, ищет коров у которых была контрольная дойка в эту дату или ранее

	Exterior             *float64 `default:"3.14" validate:"optional"`       // Фильтр по оценке экстерьера коровы, будет переработан
	InseminationDateFrom *string  `example:"1800-01-21" validate:"optional"` // Фильтр по дате осеменения коровы, ищет коров у которых было осеменение в эту дату или позднее
	InseminationDateTo   *string  `example:"2800-01-21" validate:"optional"` // Фильтр по дате осеменения коровы, ищет коров у которых было осеменение в эту дату или ранее
	CalvingDateFrom      *string  `example:"1800-01-21" validate:"optional"` // Фильтр по дате отела коровы, ищет коров у которых был отел в эту дату или позднее
	CalvingDateTo        *string  `example:"2800-01-21" validate:"optional"` // Фильтр по дате осеменения коровы, ищет коров у которых был отел в эту дату или позднее
	IsStillBorn          *bool    `default:"false" validate:"optional"`      // Фильтр по мертворождению Было/не было
	IsTwins              *bool    `default:"false" validate:"optional"`      // Фильтр по родам двойняшек Было / не было
	IsAborted            *bool    `default:"false" validate:"optional"`      // Фильтр по абортам Был/ не был
	IsIll                *bool    `default:"false" validate:"optional"`      //??? Не реализован
	BirkingDateFrom      *string  `example:"1800-01-21" validate:"optional"` // Фильтр по дате перебирковки коровы, ищет коров у которых быа перебирковка в эту дату или позднее
	BirkingDateTo        *string  `example:"2800-01-21" validate:"optional"` // Фильтр по дате осеменения коровы, ищет коров у которых была перебирковка в эту дату или позднее

	InbrindingCoeffByFamilyFrom *float64 `default:"3.14" validate:"optional"` // Не реализован
	InbrindingCoeffByFamilyTo   *float64 `default:"3.14" validate:"optional"` // Не реализован

	InbrindingCoeffByFenotypeFrom *float64 `default:"3.14" validate:"optional"` //?Не реализован
	InbrindingCoeffByFenotypeTo   *float64 `default:"3.14" validate:"optional"` //?Не реализован

	MonogeneticIllneses []uint //?Не реализован
}

type FilterSerializedCow struct {
	RSHNNumber              string            `validate:"required" example:"123"`
	InventoryNumber         string            `validate:"required" example:"321"`
	Name                    string            `validate:"required" example:"Буренка"`
	FarmGroupName           string            `validate:"required" example:"ООО Аурус"`
	BirthDate               models.DateOnly   `validate:"required"`
	Genotyped               bool              `validate:"required" example:"true"`
	DepartDate              *models.DateOnly  `json:",omitempty" validate:"optional"`
	BreedName               *string           `json:",omitempty" validate:"optional" example:"Какая-нибудь порода"`
	CheckMilkDate           []models.DateOnly `json:",omitempty" validate:"optional"`
	InsemenationDate        []models.DateOnly `json:",omitempty" validate:"optional"`
	CalvingDate             []models.DateOnly `json:",omitempty" validate:"optional"`
	BirkingDate             *models.DateOnly  `json:",omitempty" validate:"optional"`
	InbrindingCoeffByFamily *float64          `json:",omitempty" validate:"optional" example:"3.14"`
}

func serializeByFilter(c *models.Cow, filter *cowsFilter) FilterSerializedCow {
	res := FilterSerializedCow{
		RSHNNumber:      c.RSHNNumber,
		InventoryNumber: c.InventoryNumber,
		Name:            c.Name,
		FarmGroupName:   c.FarmGroup.Name,
		BirthDate:       c.BirthDate,
		Genotyped:       true,
	}
	if filter.DepartDateTo != nil || filter.DepartDateFrom != nil {
		res.DepartDate = c.DepartDate
	}
	if len(filter.BreedId) != 0 {
		res.BreedName = &c.Breed.Name
	}
	if filter.InbrindingCoeffByFamilyFrom != nil || filter.InbrindingCoeffByFamilyTo != nil {
		res.InbrindingCoeffByFamily = &c.InbrindingCoeffByFamily
	}
	if filter.ControlMilkingDateFrom != nil || filter.ControlMilkingDateTo != nil {
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

	if filter.InseminationDateFrom != nil || filter.InseminationDateTo != nil {
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

	if filter.CalvingDateFrom != nil || filter.CalvingDateTo != nil {
		for _, lac := range c.Lactation {
			if filter.CalvingDateFrom != nil {
				date, err := time.Parse(time.DateOnly, *filter.CalvingDateFrom)
				if err != nil {
					continue
				}
				if !date.Equal(lac.CalvingDate.Time) && date.After(lac.CalvingDate.Time) {
					continue
				}
			}
			if filter.CalvingDateTo != nil {
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

	if filter.BirkingDateFrom != nil || filter.BirkingDateTo != nil {
		res.BirkingDate = &c.BirkingDate
	}
	return res
}

// ListAccounts lists all existing accounts
//
//	@Summary      Get filtered list of cows
//	@Tags         Cows
//	@Param        filter    body     cowsFilter  true  "applied filters"
//	@Accept       json
//	@Produce      json
//	@Success      200  {array}   FilterSerializedCow
//	@Failure      422  {object}  map[string]error
//	@Failure      500  {object}  map[string]error
//	@Router       /cows/filter [post]
func (c *Cows) Filter() func(*gin.Context) {
	return func(c *gin.Context) {
		jsonData, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(500, gin.H{"error": err})
			return
		}

		bodyData := cowsFilter{}
		if len(jsonData) != 0 {
			err = json.Unmarshal(jsonData, &bodyData)
			if err != nil {
				c.JSON(422, gin.H{"error": err})
				return
			}
		}

		db := models.GetDb()
		query := db.Model(&models.Cow{}).Preload("FarmGroup")

		recordsPerPage := uint64(50)
		pageNumber := uint64(1)

		if bodyData.EntitiesOnPage != nil {
			recordsPerPage = uint64(*bodyData.EntitiesOnPage)
		}

		if bodyData.PageNumber != nil {
			pageNumber = uint64(*bodyData.PageNumber)
		}

		if searchString := bodyData.SearchQuery; searchString != nil && *searchString != "" {
			*searchString = "%" + *searchString + "%"
			query = query.Where("name LIKE ?", searchString).Or("rshn_number LIKE ?", searchString).Or("inventory_number LIKE ?", searchString)
		}

		if len(bodyData.Sex) != 0 {
			query = query.Where("sex_id IN ?", bodyData.Sex).Preload("Sex")
		}

		if bodyData.HozId != nil {
			query = query.Where("farm_group_id = ?", bodyData.HozId).Preload("Farm")
		}

		if len(bodyData.BreedId) != 0 {
			query = query.Where("breed_id in ?", bodyData.BreedId).Preload("Breed")
		}

		if bodyData.IsDead != nil && *bodyData.IsDead {
			query = query.Where("death_date IS NOT NULL")
		}

		// ====================================================================================================
		// =================================== Filter by brithday date ========================================
		// ====================================================================================================
		if bodyData.BirthDateFrom != nil && bodyData.BirthDateTo != nil {
			bdFrom, err := time.Parse(time.DateOnly, *bodyData.BirthDateFrom)
			if err != nil {
				c.JSON(422, err)
				return
			}
			bdTo, err := time.Parse(time.DateOnly, *bodyData.BirthDateTo)
			if err != nil {
				c.JSON(422, err)
				return
			}
			query = query.Where("birth_date BETWEEN ? AND ?", bdFrom.UTC(), bdTo.UTC())
		} else if bodyData.BirthDateFrom != nil {
			bdFrom, err := time.Parse(time.DateOnly, *bodyData.BirthDateFrom)
			if err != nil {
				c.JSON(422, err)
				return
			}
			query = query.Where("birth_date >= ?", bdFrom.UTC())
		} else if bodyData.BirthDateTo != nil {
			bdTo, err := time.Parse(time.DateOnly, *bodyData.BirthDateTo)
			if err != nil {
				c.JSON(422, err)
				return
			}
			query = query.Where("birth_date <= ?", bdTo.UTC())
		}

		// ====================================================================================================
		// ================================ Filter by departation date ========================================
		// ====================================================================================================
		if bodyData.DepartDateFrom != nil && bodyData.DepartDateTo != nil {
			bdFrom, err := time.Parse(time.DateOnly, *bodyData.DepartDateFrom)
			if err != nil {
				c.JSON(422, err)
				return
			}
			bdTo, err := time.Parse(time.DateOnly, *bodyData.DepartDateTo)
			if err != nil {
				c.JSON(422, err)
				return
			}
			query = query.Where("depart_date BETWEEN ? AND ?", bdFrom.UTC(), bdTo.UTC())
		} else if bodyData.DepartDateFrom != nil {
			bdFrom, err := time.Parse(time.DateOnly, *bodyData.DepartDateFrom)
			if err != nil {
				c.JSON(422, err)
				return
			}
			query = query.Where("depart_date >= ?", bdFrom.UTC())
		} else if bodyData.DepartDateTo != nil {
			bdTo, err := time.Parse(time.DateOnly, *bodyData.DepartDateTo)
			if err != nil {
				c.JSON(422, err)
				return
			}
			query = query.Where("depart_date <= ?", bdTo.UTC())
		}

		// ====================================================================================================
		// ============================ Filter by control milking date ========================================
		// ====================================================================================================
		if bodyData.ControlMilkingDateFrom != nil && bodyData.ControlMilkingDateTo != nil {
			bdFrom, err := time.Parse(time.DateOnly, *bodyData.ControlMilkingDateFrom)
			if err != nil {
				c.JSON(422, err)
				return
			}
			bdTo, err := time.Parse(time.DateOnly, *bodyData.ControlMilkingDateTo)
			if err != nil {
				c.JSON(422, err)
				return
			}
			query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND EXISTS (SELECT 1 FROM check_milks WHERE check_milks.lactation_id = lactations.id AND check_milks.check_date BETWEEN ? AND ?))", bdFrom.UTC(), bdTo.UTC()).Preload("Lactation").Preload("Lactation.CheckMilks")
		} else if bodyData.ControlMilkingDateFrom != nil {
			bdFrom, err := time.Parse(time.DateOnly, *bodyData.ControlMilkingDateFrom)
			if err != nil {
				c.JSON(422, err)
				return
			}
			query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND EXISTS (SELECT 1 FROM check_milks WHERE check_milks.lactation_id = lactations.id AND check_milks.check_date >= ?))", bdFrom.UTC()).Preload("Lactation").Preload("Lactation.CheckMilks")
		} else if bodyData.ControlMilkingDateTo != nil {
			bdTo, err := time.Parse(time.DateOnly, *bodyData.ControlMilkingDateTo)
			if err != nil {
				c.JSON(422, err)
				return
			}
			query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND EXISTS (SELECT 1 FROM check_milks WHERE check_milks.lactation_id = lactations.id AND check_milks.check_date <= ?))", bdTo.UTC()).Preload("Lactation").Preload("Lactation.CheckMilks")
		}

		// ====================================================================================================
		// ==================================== Filter by calving date ========================================
		// ====================================================================================================
		if bodyData.CalvingDateFrom != nil && bodyData.CalvingDateTo != nil {
			bdFrom, err := time.Parse(time.DateOnly, *bodyData.CalvingDateFrom)
			if err != nil {
				c.JSON(422, err)
				return
			}
			bdTo, err := time.Parse(time.DateOnly, *bodyData.CalvingDateTo)
			if err != nil {
				c.JSON(422, err)
				return
			}
			query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.calving_date BETWEEN ? AND ?)", bdFrom.UTC(), bdTo.UTC()).Preload("Lactation")
		} else if bodyData.CalvingDateFrom != nil {
			bdFrom, err := time.Parse(time.DateOnly, *bodyData.CalvingDateFrom)
			if err != nil {
				c.JSON(422, err)
				return
			}
			query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.calving_date >= ?)", bdFrom.UTC()).Preload("Lactation")
		} else if bodyData.CalvingDateTo != nil {
			bdTo, err := time.Parse(time.DateOnly, *bodyData.CalvingDateTo)
			if err != nil {
				c.JSON(422, err)
				return
			}
			query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.calving_date <= ?)", bdTo.UTC()).Preload("Lactation")
		}

		// ====================================================================================================
		// ======================= Filter by birth parameters (e/g stillborn, twins ...)=======================
		// ====================================================================================================
		if bodyData.IsStillBorn != nil && *bodyData.IsStillBorn { // stillborn means, that 0 cows are born
			query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.calving_count = ?)", 0).Preload("Lactation")
		}

		if bodyData.IsTwins != nil && *bodyData.IsTwins { // twins means, that 2 cows are born
			query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.calving_count = ?)", 2).Preload("Lactation")
		}

		if bodyData.IsAborted != nil && *bodyData.IsAborted { // abort is marked by flag for some reason
			query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.abort = ?)", true).Preload("Lactation")
		}

		if bodyData.Exterior != nil {
			query = query.Where("exterior = ?", bodyData.Exterior)
		}

		// ====================================================================================================
		// ===================================  Filter by birinkg date ========================================
		// ====================================================================================================
		if bodyData.BirkingDateFrom != nil && bodyData.BirkingDateTo != nil {
			bdFrom, err := time.Parse(time.DateOnly, *bodyData.BirkingDateFrom)
			if err != nil {
				c.JSON(422, err)
				return
			}
			bdTo, err := time.Parse(time.DateOnly, *bodyData.BirkingDateTo)
			if err != nil {
				c.JSON(422, err)
				return
			}
			query = query.Where("birking_date BETWEEN ? AND ?", bdFrom, bdTo)
		} else if bodyData.BirkingDateFrom != nil {
			bdFrom, err := time.Parse(time.DateOnly, *bodyData.BirkingDateFrom)
			if err != nil {
				c.JSON(422, err)
				return
			}
			query = query.Where("birking_date >= ?", bdFrom)
		} else if bodyData.BirkingDateTo != nil {
			bdTo, err := time.Parse(time.DateOnly, *bodyData.BirkingDateTo)
			if err != nil {
				c.JSON(422, err)
				return
			}
			query = query.Where("birking_date <= ?", bdTo)
		}

		// ====================================================================================================
		// ============= ============= FILTER BY INBRINDING COEF ============= ============= ============= ===
		// ====================================================================================================
		if bodyData.InbrindingCoeffByFamilyFrom != nil && bodyData.InbrindingCoeffByFamilyTo != nil {
			query = query.Where("inbrinding_coeff_by_family BETWEEN ? AND ?", bodyData.InbrindingCoeffByFamilyFrom, bodyData.InbrindingCoeffByFamilyTo)
		} else if bodyData.InbrindingCoeffByFamilyFrom != nil {
			query = query.Where("inbrinding_coeff_by_family >= ?", bodyData.InbrindingCoeffByFamilyFrom)
		} else if bodyData.InbrindingCoeffByFamilyTo != nil {
			query = query.Where("inbrinding_coeff_by_family <= ?", bodyData.InbrindingCoeffByFamilyTo)
		}

		// ====================================================================================================
		// ================================== FITLER BY INSEMENTAION DATE =====================================
		// ====================================================================================================
		if bodyData.InseminationDateFrom != nil && bodyData.InseminationDateTo != nil {
			bdFrom, err := time.Parse(time.DateOnly, *bodyData.InseminationDateFrom)
			if err != nil {
				c.JSON(422, err)
				return
			}
			bdTo, err := time.Parse(time.DateOnly, *bodyData.InseminationDateTo)
			if err != nil {
				c.JSON(422, err)
				return
			}
			query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.insemenation_date BETWEEN ? and ?)", bdFrom.UTC(), bdTo.UTC()).Preload("Lactation")
		} else if bodyData.InseminationDateFrom != nil {
			bdFrom, err := time.Parse(time.DateOnly, *bodyData.InseminationDateFrom)
			if err != nil {
				c.JSON(422, err)
				return
			}
			query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.insemenation_date >= ?)", bdFrom.UTC()).Preload("Lactation")
		} else if bodyData.InseminationDateTo != nil {
			bdTo, err := time.Parse(time.DateOnly, *bodyData.InseminationDateTo)
			if err != nil {
				c.JSON(422, err)
				return
			}
			query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.insemenation_date <= ?)", bdTo.UTC()).Preload("Lactation")
		}

		// ====================================================================================================
		// ==================================Get final query result ===========================================
		// ====================================================================================================
		resCount := int64(0)
		if err := query.Count(&resCount).Error; err != nil {
			c.JSON(500, err)
			return
		}

		query = query.Limit(int(recordsPerPage)).Offset(int(recordsPerPage) * int(pageNumber-1))
		dbCows := []models.Cow{}
		if err := query.Debug().Find(&dbCows).Error; err != nil {
			c.JSON(500, gin.H{"error": err})
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
