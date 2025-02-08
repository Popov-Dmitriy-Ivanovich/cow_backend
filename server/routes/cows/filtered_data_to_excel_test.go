package cows

import (
	"cow_backend/models"
	"testing"
	"time"
)

func TestToExcelOld(t *testing.T) {
	Name := "Буренка"
	farm := "ООО Аурус"
	number := "123"
	invNumber := "321"
	breed := "Какая-нибудь порода"
	sex := "Корова"
	date := models.DateOnly{Time: time.Now()}
	coef := 3.14
	isTrue := true

	tests := []FilterSerializedCow{
		{
			ID:                        123,
			RSHNNumber:                &number,
			InventoryNumber:           &invNumber,
			Name:                      Name,
			FarmGroupName:             farm,
			BirthDate:                 date,
			Genotyped:                 true,
			Approved:                  true,
			DepartDate:                &date,
			BreedName:                 &breed,
			CheckMilkDate:             []models.DateOnly{},
			InsemenationDate:          []models.DateOnly{},
			CalvingDate:               []models.DateOnly{},
			BirkingDate:               &date,
			GenotypingDate:            &date,
			InbrindingCoeffByFamily:   &coef,
			InbrindingCoeffByGenotype: &coef,
			MonogeneticIllneses:       []models.GeneticIllnessData{},
			ExteriorRating:            &coef,
			SexName:                   &sex,
			HozName:                   &farm,

			DeathDate:   &date,
			IsDead:      &isTrue,
			IsTwins:     &isTrue,
			IsStillBorn: &isTrue,
			IsAborted:   &isTrue,
			Events:      []models.Event{},
			IsGenotyped: &isTrue,
			CreatedAt:   &date,
		},
		{
			ID:                        123,
			RSHNNumber:                nil,
			InventoryNumber:           nil,
			Name:                      Name,
			FarmGroupName:             farm,
			BirthDate:                 date,
			Genotyped:                 true,
			Approved:                  true,
			DepartDate:                &date,
			BreedName:                 &breed,
			CheckMilkDate:             []models.DateOnly{},
			InsemenationDate:          []models.DateOnly{},
			CalvingDate:               []models.DateOnly{},
			BirkingDate:               &date,
			GenotypingDate:            &date,
			InbrindingCoeffByFamily:   &coef,
			InbrindingCoeffByGenotype: &coef,
			MonogeneticIllneses:       []models.GeneticIllnessData{},
			ExteriorRating:            &coef,
			SexName:                   &sex,
			HozName:                   &farm,

			DeathDate:   &date,
			IsDead:      &isTrue,
			IsTwins:     &isTrue,
			IsStillBorn: &isTrue,
			IsAborted:   &isTrue,
			Events:      []models.Event{},
			IsGenotyped: &isTrue,
			CreatedAt:   &date,
		},
		{
			ID:                        123,
			RSHNNumber:                &number,
			InventoryNumber:           &invNumber,
			Name:                      Name,
			FarmGroupName:             farm,
			BirthDate:                 date,
			Genotyped:                 true,
			Approved:                  true,
			DepartDate:                nil,
			BreedName:                 nil,
			CheckMilkDate:             []models.DateOnly{},
			InsemenationDate:          []models.DateOnly{},
			CalvingDate:               []models.DateOnly{},
			BirkingDate:               nil,
			GenotypingDate:            nil,
			InbrindingCoeffByFamily:   nil,
			InbrindingCoeffByGenotype: nil,
			MonogeneticIllneses:       []models.GeneticIllnessData{},
			ExteriorRating:            nil,
			SexName:                   nil,
			HozName:                   nil,

			DeathDate:   nil,
			IsDead:      nil,
			IsTwins:     nil,
			IsStillBorn: nil,
			IsAborted:   nil,
			Events:      []models.Event{},
			IsGenotyped: nil,
			CreatedAt:   nil,
		},
	}

	t.Run("test", func(t *testing.T) {
		_, err := ToExcelOld(tests)
		if (err != nil) == true {
			t.Errorf("in if: %v", err)
		}
		t.Log("OK")

	})

}
