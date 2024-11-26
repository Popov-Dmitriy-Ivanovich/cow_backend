package routes

import (
	"cow_backend/models"
	"encoding/json"
	// "fmt"
	"io"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (a *Api) FindInCowsPost() func (*gin.Context) {
	return func (c *gin.Context) {
		jsonData, err := io.ReadAll(c.Request.Body)
		if err != nil {
		    c.JSON(500, gin.H{"error": err})
			return
		}
		

		bodyData := map[string]string{}
		if len(jsonData) != 0 {
			err = json.Unmarshal(jsonData,&bodyData)
			if err != nil {
				c.JSON(500, gin.H{"error": err})
				return
			}
		}
		

		
		db := models.GetDb()
		query := db.Model(&models.Cow{})


		recordsPerPage := uint64(50)
		pageNumber := uint64(0)
		searchString := ""

		if recPerPage, ok := bodyData["n_on_page"]; ok && recPerPage != "" {
			if uintRecPerPage, err := strconv.ParseUint(recPerPage, 10, 64); err != nil {
				c.JSON(500, gin.H{"error":err})
				return
			} else {
				recordsPerPage = uintRecPerPage
			}			
		}
		

		if strPageNumber, ok := bodyData["npage"]; ok && strPageNumber != "" {
			if uintPageNumber, err := strconv.ParseUint(strPageNumber, 10, 64); err != nil {
				c.JSON(500, gin.H{"error": err})
				return
			} else {
				pageNumber = uintPageNumber - 1
			}
		}


		query = query.Limit(int(recordsPerPage)).Offset(int(recordsPerPage)*int(pageNumber))

		if strSearchString, ok := bodyData["fnd_str"]; ok && strSearchString != "" {
			searchString = strSearchString
			query = query.Where("name = ?", searchString).Or("rshn_number = ?", searchString).Or("inventory_number = ?", searchString)
		}

		

		dbCows := []models.Cow{}
		if err := query.Debug().Find(&dbCows).Error; err != nil {
			c.JSON(500, gin.H{"error": err})
			return
		}
		// fmt.Print(query)
		c.JSON(200, gin.H{
			"recordsPerPage":recordsPerPage,
			"pageNumber": pageNumber,
			"cows": dbCows,
			// "query": query,
		})
	}
}