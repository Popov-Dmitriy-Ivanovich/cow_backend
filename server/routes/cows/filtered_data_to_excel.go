package cows

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
)

// Путь к файлу
const PathToExcelFile = "./static/excel/filtered_data_"

var (
	cell     string
	ListName = "List1"
)

func ToExcelOld(fsc []FilterSerializedCow) (string, error) {
	f := excelize.NewFile()
	defer f.Close()
	// Создаем новый лист
	index, err := f.NewSheet("List1")
	if err != nil {
		return "", err
	}

	// Устанавливаем активный лист
	f.SetActiveSheet(index)

	// Записываем заголовки
	// headers := getHeaders()
	for i, columnName := range getHeaders() {
		if i == 0 {
			continue
		}
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		err = f.SetCellValue(ListName, cell, columnName)
		if err != nil {
			return "", err
		}
	}

	// Записываем данные
	for row, data := range fsc {
		colNum := 1
		// Объявим функция для уменьшения размера кода
		// Функция записи ошибочной строки
		writeErrorRequiredData := func() error {
			cell, err := excelize.CoordinatesToCellName(1, row+2)
			if err != nil {
				return err
			}
			err = f.SetCellValue("List1", cell, "Отсутсвуют обязательные данные")
			if err != nil {
				return err
			}
			return nil
		}
		// Функция инкрементирования ячейки
		Incr := func() {
			cell, err = excelize.CoordinatesToCellName(colNum, row+2)
			colNum++
		}
		// Проверим обязательные поля
		if *data.RSHNNumber == "" {
			err = writeErrorRequiredData()
			if err != nil {
				return "", err
			}
			continue
		}
		if *data.InventoryNumber == "" {
			err = writeErrorRequiredData()
			if err != nil {
				return "", err
			}
			continue
		}
		if data.Name == "" {
			err = writeErrorRequiredData()
			if err != nil {
				return "", err
			}
			continue
		}
		if data.FarmGroupName == "" {
			err = writeErrorRequiredData()
			if err != nil {
				return "", err
			}
			continue
		}
		if data.BirthDate.IsZero() {
			err = writeErrorRequiredData()
			if err != nil {
				return "", err
			}
			continue
		}
		// Поля Genotyped и Approved будут существовать в любом случае
		// ===== //
		// Записываем данные
		if err = f.SetCellValue(ListName, cell, *data.RSHNNumber); err != nil {
			return "", err
		} else {
			Incr()
		}
		if err = f.SetCellValue(ListName, cell, *data.InventoryNumber); err != nil {
			return "", err
		} else {
			Incr()
		}
		if err = f.SetCellValue(ListName, cell, data.Name); err != nil {
			return "", err
		} else {
			Incr()
		}
		if err = f.SetCellValue(ListName, cell, data.FarmGroupName); err != nil {
			return "", err
		} else {
			Incr()
		}
		if err = f.SetCellValue(ListName, cell, data.BirthDate.Time); err != nil {
			return "", err
		} else {
			Incr()
		}
		if err = f.SetCellValue(ListName, cell, data.Genotyped); err != nil {
			return "", err
		} else {
			Incr()
		}
		if err = f.SetCellValue(ListName, cell, data.Approved); err != nil {
			return "", err
		} else {
			Incr()
		}
		if err = f.SetCellValue(ListName, cell, data.DepartDate.Time); err != nil {
			return "", err
		} else {
			Incr()
		}
		if err = f.SetCellValue(ListName, cell, *data.BreedName); err != nil {
			return "", err
		} else {
			Incr()
		}
		if err = f.SetCellValue(ListName, cell, data.BirkingDate.Time); err != nil {
			return "", err
		} else {
			Incr()
		}
		if err = f.SetCellValue(ListName, cell, data.GenotypingDate.Time); err != nil {
			return "", err
		} else {
			Incr()
		}
		if err = f.SetCellValue(ListName, cell, *data.InbrindingCoeffByFamily); err != nil {
			return "", err
		} else {
			Incr()
		}
		if err = f.SetCellValue(ListName, cell, *data.InbrindingCoeffByGenotype); err != nil {
			return "", err
		} else {
			Incr()
		}
		if err = f.SetCellValue(ListName, cell, *data.InbrindingCoeffByFamily); err != nil {
			return "", err
		} else {
			Incr()
		}
		if err = f.SetCellValue(ListName, cell, *data.ExteriorRating); err != nil {
			return "", err
		} else {
			Incr()
		}
		if err = f.SetCellValue(ListName, cell, *data.ExteriorRating); err != nil {
			return "", err
		} else {
			Incr()
		}
		if err = f.SetCellValue(ListName, cell, *data.SexName); err != nil {
			return "", err
		} else {
			Incr()
		}
		if err = f.SetCellValue(ListName, cell, *data.HozName); err != nil {
			return "", err
		} else {
			Incr()
		}
		if err = f.SetCellValue(ListName, cell, data.DeathDate.Time); err != nil {
			return "", err
		} else {
			Incr()
		}
		if err = f.SetCellValue(ListName, cell, *data.IsDead); err != nil {
			return "", err
		} else {
			Incr()
		}
		if err = f.SetCellValue(ListName, cell, *data.IsTwins); err != nil {
			return "", err
		} else {
			Incr()
		}
		if err = f.SetCellValue(ListName, cell, *data.IsStillBorn); err != nil {
			return "", err
		} else {
			Incr()
		}
		if err = f.SetCellValue(ListName, cell, *data.IsAborted); err != nil {
			return "", err
		} else {
			Incr()
		}
		if err = f.SetCellValue(ListName, cell, *data.IsGenotyped); err != nil {
			return "", err
		} else {
			Incr()
		}
		if err = f.SetCellValue(ListName, cell, data.CreatedAt.Time); err != nil {
			return "", err
		} else {
			Incr()
		}

	}

	// Сохраняем файл в cow_backend\frontend\static
	now := time.Now()
	fullPath := PathToExcelFile + strconv.FormatInt(now.Unix(), 16) + "_" + strconv.FormatUint(uint64(len(fsc)), 16) + ".xslx"
	if err := f.SaveAs(fullPath); err != nil {
		return "", fmt.Errorf("Ошибка создания Excel файла Error: %v", err)
	} else {
		return fullPath, nil
	}

}

func isSlice(i int, data *FilterSerializedCow) bool {
	fl := reflect.ValueOf(data)
	return fl.Field(i).Kind() == reflect.Slice
}

func getHeaders() []string { // Получаем заголовки таблицы
	var headers []string
	t := reflect.TypeOf(FilterSerializedCow{})

	for i := 1; i < t.NumField(); i++ { // Не берем поле id
		if isSlice(i, &FilterSerializedCow{}) { // Пропускаем списки на стадии поиска заголовков
			continue
		}
		field := t.Field(i)
		headers = append(headers, field.Name)
	}

	return headers
}
