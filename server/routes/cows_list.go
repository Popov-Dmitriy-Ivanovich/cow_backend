package routes

// import (
// 	"cow_backend/models"
// 	"encoding/json"
// 	"time"

// 	// "fmt"
// 	"io"

// 	"github.com/gin-gonic/gin"
// )

// type cowsFilter struct {
// 	SearchQuery            *string `example:"Буренка"` // used
// 	PageNumber             *uint   `default:"1"`       //used
// 	EntitiesOnPage         *uint   `default:"50"`      //used
// 	Sex                    []uint  //used
// 	FarmID                 *uint   `default:"1"`          //used
// 	BirthDateFrom          *string `example:"1800-01-21"` //used
// 	BirthDateTo            *string `example:"2800-01-21"` //used
// 	IsDead                 *bool   `default:"false"`      //used
// 	DepartDateFrom         *string `example:"1800-01-21"` //used
// 	DepartDateTo           *string `example:"2800-01-21"` //used
// 	BreedId                []uint  //used
// 	GenotypingDateFrom     *string `example:"1800-01-21"` //??? Genotiping needed
// 	GenotypingDateTo       *string `example:"2800-01-21"` //??? Genotiping needed
// 	ControlMilkingDateFrom *string `example:"1800-01-21"` //used
// 	ControlMilkingDateTo   *string `example:"2800-01-21"` //used

// 	Exterior             *float64 `default:"3.14"` //used
// 	InseminationDateFrom *string  `example:"1800-01-21"`
// 	InseminationDateTo   *string  `example:"2800-01-21"`
// 	CalvingDateFrom      *string  `example:"1800-01-21"` //used
// 	CalvingDateTo        *string  `example:"2800-01-21"` //used
// 	IsStillBorn          *bool    `default:"false"`      //used
// 	IsTwins              *bool    `default:"false"`      //used
// 	IsAborted            *bool    `default:"false"`      //used
// 	IsIll                *bool    `default:"false"`      //??? Genotiping needed
// 	BirkingDateFrom      *string  `example:"1800-01-21"` // date field
// 	BirkingDateTo        *string  `example:"2800-01-21"` // date field

// 	InbrindingCoeffByFamilyFrom *float64 `default:"3.14"` // used
// 	InbrindingCoeffByFamilyTo   *float64 `default:"3.14"` // used

// 	InbrindingCoeffByFenotypeFrom *float64 `default:"3.14"` //??? Genotiping needed
// 	InbrindingCoeffByFenotypeTo   *float64 `default:"3.14"` //??? Genotiping needed

// 	MonogeneticIllneses []uint //??? Genotiping needed
// }

// // ListAccounts lists all existing accounts
// //
// //	@Summary      Get filtered list of cows
// //	@Description  Get filtered list of cows.
// //	@Description  SearchQuery - имя, номер РСХН или инвентарный номер
// //	@Description  PageNumber - номер страницы для отображения
// //	@Description  EntitiesOnPage - количество коров на каждой странице
// //	@Description  Sex - массив полов для поиска (можно выбрать несколько)
// //	@Description  FarmID - ID фермы на которой живет корова
// //	@Description  BirthDateFrom - Отображает коров, родившихся после этой даты
// //	@Description  BirthDateTo - Отображает коров, родившихся до этой даты
// //	@Description  IsDead - Если флаг истина - ищет мертвых коров, иначе живых
// //	@Description  DepartDateFrom - Ищет коров отбывших из коровника после данной даты
// //	@Description  DepartDateTo - Ищет коров отбывших из коровника до данной даты
// //	@Description  BreedId - ищет коров имеющих одну из пород по BreedId
// //	@Description	GenotypingDateFrom - НЕ ИСПОЛЬЗУЕТСЯ
// //	@Description	GenotypingDateTo - НЕ ИСПОЛЬЗУЕТСЯ
// //	@Description	ControlMilkingDateFrom - ищет коров у которых была хотябы одна контрольная дойка после этой даты
// //	@Description	ControlMilkingDateTo - ищет коров у которых была хотябы одна контрольная дойка до этой даты
// //	@Description
// //	@Description	Exterior - Ищет коров с оценкой экстерьера равной этому значению
// //	@Description	InseminationDateFrom - Ищет коров которые были хотябы раз осеменены после данной даты
// //	@Description	InseminationDateTo - Ищет коров которые были хотябы раз осеменены до данной даты
// //	@Description	CalvingDateFrom  - Ищет коров у которых был отел хотябы раз после данной даты
// //	@Description	CalvingDateTo - Ищет коров у которых был отел хотябы раз до данной даты
// //	@Description	IsStillBorn  - Ищет коров у которых хотябы раз было мертворождение
// //	@Description	IsTwins - Ищет коров у которых хотябы раз родились близнецы/двойняшки
// //	@Description	IsAborted - Ищет коров, которым хотябы раз сделали аборт
// //	@Description	IsIll - НЕ ИСОПЛЬЗУЕТСЯ
// //	@Description	BirkingDateFrom - Ищет коров у которых дата перебирковки больше
// //	@Description	BirkingDateTo - Ищет коров у которых дата перебирковки меньше
// //	@Description
// //	@Description	InbrindingCoeffByFamilyFrom Ищет коров, у которых коэф. инбриндинга по роду больше
// //	@Description	InbrindingCoeffByFamilyTo   - Ищет коров у которых дата перебирковки меньше
// //	@Description
// //	@Description	InbrindingCoeffByFenotypeFrom Genotiping needed
// //	@Description	InbrindingCoeffByFenotypeTo    Genotiping needed
// //	@Description
// //	@Description	MonogeneticIllneses []uint Genotiping needed
// //	@Tags         CowList
// //	@Param        filter    body     cowsFilter  true  "applied filters"
// //	@Accept       json
// //	@Produce      json
// //	@Success      200  {array}   models.Cow
// //	@Failure      422  {object}  map[string]error
// //	@Failure      500  {object}  map[string]error
// //	@Router       /CowsList [post]
// func (a *Api) CowsList() func(*gin.Context) {
// 	return func(c *gin.Context) {
// 		jsonData, err := io.ReadAll(c.Request.Body)
// 		if err != nil {
// 			c.JSON(500, gin.H{"error": err})
// 			return
// 		}

// 		bodyData := cowsFilter{}
// 		if len(jsonData) != 0 {
// 			err = json.Unmarshal(jsonData, &bodyData)
// 			if err != nil {
// 				c.JSON(422, gin.H{"error": err})
// 				return
// 			}
// 		}

// 		db := models.GetDb()
// 		query := db.Model(&models.Cow{})

// 		recordsPerPage := uint64(50)
// 		pageNumber := uint64(1)

// 		if bodyData.EntitiesOnPage != nil {
// 			recordsPerPage = uint64(*bodyData.EntitiesOnPage)
// 		}

// 		if bodyData.PageNumber != nil {
// 			pageNumber = uint64(*bodyData.PageNumber)
// 		}

// 		query = query.Limit(int(recordsPerPage)).Offset(int(recordsPerPage) * int(pageNumber-1))

// 		if searchString := bodyData.SearchQuery; searchString != nil && *searchString != "" {
// 			query = query.Where("name = ?", searchString).Or("rshn_number = ?", searchString).Or("inventory_number = ?", searchString)
// 		}

// 		if bodyData.Sex != nil {
// 			query = query.Where("sex_id IN ?", bodyData.Sex)
// 		}

// 		if bodyData.FarmID != nil {
// 			query = query.Where("farm_id = ?", bodyData.FarmID).Preload("Farm")
// 		}

// 		if bodyData.BirthDateFrom != nil {
// 			bdFrom, err := time.Parse(time.DateOnly, *bodyData.BirthDateFrom)
// 			if err != nil {
// 				c.JSON(422, err)
// 				return
// 			}
// 			query = query.Where("birth_date >= ?", bdFrom.UTC())
// 		}

// 		if bodyData.BirthDateTo != nil {
// 			bdTo, err := time.Parse(time.DateOnly, *bodyData.BirthDateTo)
// 			if err != nil {
// 				c.JSON(422, err)
// 				return
// 			}
// 			query = query.Where("birth_date <= ?", bdTo.UTC())
// 		}

// 		if bodyData.IsDead != nil {
// 			query = query.Where("is_dead = ?", bodyData.IsDead)
// 		}

// 		if bodyData.DepartDateFrom != nil {
// 			bdFrom, err := time.Parse(time.DateOnly, *bodyData.DepartDateFrom)
// 			if err != nil {
// 				c.JSON(422, err)
// 				return
// 			}
// 			query = query.Where("depart_date >= ?", bdFrom.UTC())
// 		}

// 		if bodyData.DepartDateTo != nil {
// 			bdTo, err := time.Parse(time.DateOnly, *bodyData.DepartDateTo)
// 			if err != nil {
// 				c.JSON(422, err)
// 				return
// 			}
// 			query = query.Where("depart_date <= ?", bdTo.UTC())
// 		}

// 		if bodyData.BreedId != nil {
// 			query = query.Where("breed_id in ?", bodyData.BreedId)
// 		}

// 		if bodyData.ControlMilkingDateFrom != nil {
// 			bdFrom, err := time.Parse(time.DateOnly, *bodyData.ControlMilkingDateFrom)
// 			if err != nil {
// 				c.JSON(422, err)
// 				return
// 			}
// 			query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND EXISTS (SELECT 1 FROM check_milks WHERE check_milks.lactation_id = lactations.id AND check_milks.check_date >= ?))", bdFrom.UTC()).Preload("Lactation").Preload("Lactation.CheckMilks")
// 		}

// 		if bodyData.ControlMilkingDateTo != nil {
// 			bdTo, err := time.Parse(time.DateOnly, *bodyData.ControlMilkingDateTo)
// 			if err != nil {
// 				c.JSON(422, err)
// 				return
// 			}
// 			query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND EXISTS (SELECT 1 FROM check_milks WHERE check_milks.lactation_id = lactations.id AND check_milks.check_date <= ?))", bdTo.UTC()).Preload("Lactation").Preload("Lactation.CheckMilks")
// 		}

// 		if bodyData.CalvingDateFrom != nil {
// 			bdFrom, err := time.Parse(time.DateOnly, *bodyData.CalvingDateFrom)
// 			if err != nil {
// 				c.JSON(422, err)
// 				return
// 			}
// 			query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.calving_date >= ?)", bdFrom.UTC()).Preload("Lactation")
// 		}

// 		if bodyData.CalvingDateTo != nil {
// 			bdTo, err := time.Parse(time.DateOnly, *bodyData.CalvingDateTo)
// 			if err != nil {
// 				c.JSON(422, err)
// 				return
// 			}
// 			query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.calving_date <= ?)", bdTo.UTC()).Preload("Lactation")
// 		}

// 		if bodyData.IsStillBorn != nil && *bodyData.IsStillBorn { // stillborn means, that 0 cows are born
// 			query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.calving_count = ?)", 0).Preload("Lactation")
// 		}

// 		if bodyData.IsTwins != nil && *bodyData.IsTwins { // twins means, that 2 cows are born
// 			query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.calving_count = ?)", 2).Preload("Lactation")
// 		}

// 		if bodyData.IsAborted != nil && *bodyData.IsAborted { // abort is marked by flag for some reason
// 			query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.abort = ?)", true).Preload("Lactation")
// 		}

// 		if bodyData.Exterior != nil {
// 			query = query.Where("exterior = ?", bodyData.Exterior)
// 		}

// 		if bodyData.BirkingDateFrom != nil {
// 			bdFrom, err := time.Parse(time.DateOnly, *bodyData.BirkingDateFrom)
// 			if err != nil {
// 				c.JSON(422, err)
// 				return
// 			}
// 			query = query.Where("birking_date >= ?", bdFrom)
// 		}

// 		if bodyData.BirkingDateTo != nil {
// 			bdTo, err := time.Parse(time.DateOnly, *bodyData.BirkingDateTo)
// 			if err != nil {
// 				c.JSON(422, err)
// 				return
// 			}
// 			query = query.Where("birking_date <= ?", bdTo)
// 		}

// 		if bodyData.InbrindingCoeffByFamilyFrom != nil {
// 			query = query.Where("inbrinding_coeff_by_family >= ?", bodyData.InbrindingCoeffByFamilyFrom)
// 		}

// 		if bodyData.InbrindingCoeffByFamilyTo != nil {
// 			query = query.Where("inbrinding_coeff_by_family <= ?", bodyData.InbrindingCoeffByFamilyTo)
// 		}

// 		if bodyData.InseminationDateFrom != nil {
// 			bdFrom, err := time.Parse(time.DateOnly, *bodyData.BirkingDateFrom)
// 			if err != nil {
// 				c.JSON(422, err)
// 				return
// 			}
// 			query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.insemenation_date >= ?)", bdFrom)
// 		}

// 		if bodyData.InseminationDateTo != nil {
// 			bdTo, err := time.Parse(time.DateOnly, *bodyData.InseminationDateTo)
// 			if err != nil {
// 				c.JSON(422, err)
// 				return
// 			}
// 			query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.insemenation_date <= ?)", bdTo)
// 		}
// 		dbCows := []models.Cow{}
// 		if err := query.Debug().Find(&dbCows).Error; err != nil {
// 			c.JSON(500, gin.H{"error": err})
// 			return
// 		}
// 		// fmt.Print(query)
// 		c.JSON(200, gin.H{
// 			"N":   len(dbCows),
// 			"LST": dbCows,
// 			// "query": query,
// 		})
// 	}
// }
