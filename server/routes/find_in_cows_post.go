package routes

import (
	"cow_backend/models"
	"encoding/json"
	"time"

	// "fmt"
	"io"

	"github.com/gin-gonic/gin"
)

type cowsFilter struct {
	SearchQuery            *string // used
	PageNumber             *uint   //used
	EntitiesOnPage         *uint   //used
	Sex                    []uint  //used
	FarmID                 *uint   //used
	BirthDateFrom          *string //used
	BirthDateTo            *string //used
	IsDead                 *bool   //used
	DepartDateFrom         *string //used
	DepartDateTo           *string //used
	BreedId                []uint  //used
	GenotypingDateFrom     *string
	GenotypingDateTo       *string
	ControlMilkingDateFrom *string //used
	ControlMilkingDateTo   *string //used

	Exterior             *string //????
	InseminationDateFrom *string //used
	InseminationDateTo   *string //used
	CalvingDateFrom      *string
	CalvingDateTo        *string
	IsStillBorn          *bool
	IsTwins              *bool
	IsAborted            *bool
	IsIll                *bool
	BirkingDateFrom      *string
	BirkingDateTo        *string

	InbrindingCoeffByFamilyFrom *float64
	InbrindingCoeffByFamilyTo   *float64

	InbrindingCoeffByFenotypeFrom *float64
	InbrindingCoeffByFenotypeTo   *float64

	MonogeneticIllneses []uint
}

// ListAccounts lists all existing accounts
//
//	@Summary      Get filtered list of cows
//	@Description  Get filtered list of cows
//	@Tags         accounts
//	@Param        filter    body     cowsFilter  true  "applied filters"
//	@Accept       json
//	@Produce      json
//	@Success      200  {array}   models.Cow
//	@Failure      422  {object}  map[string]error
//	@Failure      500  {object}  map[string]error
//	@Router       /CowsList [post]
func (a *Api) FindInCowsPost() func(*gin.Context) {
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
		query := db.Model(&models.Cow{})

		recordsPerPage := uint64(50)
		pageNumber := uint64(0)

		if bodyData.EntitiesOnPage != nil {
			recordsPerPage = uint64(*bodyData.EntitiesOnPage)
		}

		if bodyData.PageNumber != nil {
			pageNumber = uint64(*bodyData.PageNumber)
		}

		query = query.Limit(int(recordsPerPage)).Offset(int(recordsPerPage) * int(pageNumber))

		if searchString := bodyData.SearchQuery; searchString != nil && *searchString != "" {
			query = query.Where("name = ?", searchString).Or("rshn_number = ?", searchString).Or("inventory_number = ?", searchString)
		}

		if bodyData.Sex != nil {
			query = query.Where("sex_id IN ?", bodyData.Sex)
		}

		if bodyData.FarmID != nil {
			query = query.Where("farm_id = ?", bodyData.FarmID)
		}

		if bodyData.BirthDateFrom != nil {
			bdFrom, err := time.Parse(time.DateOnly, *bodyData.BirthDateFrom)
			if err != nil {
				c.JSON(422, err)
				return
			}
			query = query.Where("birth_date >= ?", bdFrom.UTC())
		}

		if bodyData.BirthDateTo != nil {
			bdTo, err := time.Parse(time.DateOnly, *bodyData.BirthDateTo)
			if err != nil {
				c.JSON(422, err)
				return
			}
			query = query.Where("birth_date <= ?", bdTo.UTC())
		}

		if bodyData.IsDead != nil {
			query = query.Where("is_dead = ?", bodyData.IsDead)
		}

		if bodyData.DepartDateFrom != nil {
			bdFrom, err := time.Parse(time.DateOnly, *bodyData.DepartDateFrom)
			if err != nil {
				c.JSON(422, err)
				return
			}
			query = query.Where("depart_date >= ?", bdFrom.UTC())
		}

		if bodyData.DepartDateTo != nil {
			bdTo, err := time.Parse(time.DateOnly, *bodyData.DepartDateTo)
			if err != nil {
				c.JSON(422, err)
				return
			}
			query = query.Where("depart_date <= ?", bdTo.UTC())
		}

		if bodyData.BreedId != nil {
			query = query.Where("breed_id in ?", bodyData.BreedId)
		}

		if bodyData.ControlMilkingDateFrom != nil {
			bdFrom, err := time.Parse(time.DateOnly, *bodyData.ControlMilkingDateFrom)
			if err != nil {
				c.JSON(422, err)
				return
			}
			query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND EXISTS (SELECT 1 FROM check_milks WHERE check_milks.lactation_id = lactations.id AND check_milks.check_date >= ?))", bdFrom.UTC()).Preload("Lactation").Preload("Lactation.CheckMilks")
		}

		if bodyData.ControlMilkingDateTo != nil {
			bdTo, err := time.Parse(time.DateOnly, *bodyData.ControlMilkingDateTo)
			if err != nil {
				c.JSON(422, err)
				return
			}
			query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND EXISTS (SELECT 1 FROM check_milks WHERE check_milks.lactation_id = lactations.id AND check_milks.check_date <= ?))", bdTo.UTC()).Preload("Lactation").Preload("Lactation.CheckMilks")
		}

		dbCows := []models.Cow{}
		if err := query.Debug().Find(&dbCows).Error; err != nil {
			c.JSON(500, gin.H{"error": err})
			return
		}
		// fmt.Print(query)
		c.JSON(200, gin.H{
			"recordsPerPage": recordsPerPage,
			"pageNumber":     pageNumber,
			"cows":           dbCows,
			// "query": query,
		})
	}
}
