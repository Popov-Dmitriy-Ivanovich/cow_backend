package cows

import (
	"cow_backend/models"
	"os"
	"testing"
	"time"
)

func TestToExcelOld(t *testing.T) {
	// Тестовый путь
	testPathToExcelFile := "./static/excel/"
	forDeleteTestPath := "./static"

	Name := "Буренка"
	farm := "ООО Аурус"
	number := "123"
	invNumber := "321"
	breed := "Какая-нибудь порода"
	sex := "Корова"
	now := time.Now()
	date := models.DateOnly{Time: time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)}
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
	err := os.MkdirAll(testPathToExcelFile, 0777)
	if err != nil {
		t.Fatalf("Ошибка создания директории: %v", err)
	}
	defer os.RemoveAll(forDeleteTestPath)

	// Формируем путь к тестовому файлу

	t.Run("test", func(t *testing.T) {
		filepath, err := ToExcelOld(tests)
		if (err != nil) == true {
			t.Errorf("ошибка: %v", err)
		}
		t.Log("OK, функция выполнена")

		// Окончательная проверка что файл существует
		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			t.Errorf("Файл не был создан: %v", err)
		} else {
			t.Logf("Файл создан: %v", filepath)
		}

	})

}
