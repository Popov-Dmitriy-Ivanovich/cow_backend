package cows

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
)

// Переменные функции
// Путь к файлу
const PathToExcelFile = "./static/excel/"

var (
	cellName string
	ListName = "List1"
)

func ToExcelOld(fsc []FilterSerializedCow) (string, error) {
	f := excelize.NewFile()

	// Создаем новый лист
	index, err := f.NewSheet("List1")
	if err != nil {
		return "a", err
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
			err = f.SetCellValue(ListName, cell, "Отсутсвуют обязательные данные")
			if err != nil {
				return err
			}
			return nil
		}
		// Функция инкрементирования ячейки
		Incr := func() {
			cellName, err = excelize.CoordinatesToCellName(colNum, row+2)
			colNum++
		}
		// Проверим обязательные поля
		// if data.RSHNNumber == nil || *data.RSHNNumber == "" {
		// 	// err = writeErrorRequiredData()
		// 	if err != nil {
		// 		return "", err
		// 	}
		// 	continue
		// }
		if data.InventoryNumber == nil || *data.InventoryNumber == "" {
			err = writeErrorRequiredData()
			if err != nil {
				return "", err
			}
			continue
		}
		// if data.Name == "" {
		// 	err = writeErrorRequiredData()
		// 	if err != nil {
		// 		return "", err
		// 	}
		// 	continue
		// }
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
		Incr()
		if data.RSHNNumber != nil{ // РСХН всегда хранит номер
		    if err = f.SetCellValue(ListName, cellName, *data.RSHNNumber); err != nil {
				return "", err
			} else {
				Incr()
			}
		}else {
			if err = f.SetCellValue(ListName, cellName, ""); err != nil {
				return "", err
			} else {
				Incr()
			}
		}
		if data.InventoryNumber != nil {
		    if err = f.SetCellValue(ListName, cellName, *data.InventoryNumber); err != nil {
				return "", err
			} else {
				Incr()
			}
		}else {
			if err = f.SetCellValue(ListName, cellName, ""); err != nil {
				return "", err
			} else {
				Incr()
			}
		}
		
		if err = f.SetCellValue(ListName, cellName, data.Name); err != nil {
			return "", err
		} else {
			Incr()
		}
		if err = f.SetCellValue(ListName, cellName, data.FarmGroupName); err != nil {
			return "", err
		} else {
			Incr()
		}
		if err = f.SetCellValue(ListName, cellName, data.BirthDate.Time); err != nil {
			return "", err
		} else {
			Incr()
		}
		if err = f.SetCellValue(ListName, cellName, data.Genotyped); err != nil {
			return "", err
		} else {
			Incr()
		}
		if err = f.SetCellValue(ListName, cellName, data.Approved); err != nil {
			return "", err
		} else {
			Incr()
		}
		if data.DepartDate != nil {
			if err = f.SetCellValue(ListName, cellName, data.DepartDate.Time); err != nil {
				return "", err
			} else {
				Incr()
			}
		} else {
			if err = f.SetCellValue(ListName, cellName, ""); err != nil {
				return "", err
			} else {
				Incr()
			}
		}
		if data.BreedName != nil { // Проверка на пустой указатель
			if err = f.SetCellValue(ListName, cellName, *data.BreedName); err != nil {
				return "", err
			} else {
				Incr()
			}
		} else {
			if err = f.SetCellValue(ListName, cellName, ""); err != nil {
				return "", err
			} else {
				Incr()
			}
		}
		if data.BirkingDate != nil {
			if err = f.SetCellValue(ListName, cellName, data.BirkingDate.Time); err != nil {
				return "", err
			} else {
				Incr()
			}
		} else {
			if err = f.SetCellValue(ListName, cellName, ""); err != nil {
				return "", err
			} else {
				Incr()
			}
		}
		if data.GenotypingDate != nil {
			if err = f.SetCellValue(ListName, cellName, data.GenotypingDate.Time); err != nil {
				return "", err
			} else {
				Incr()
			}
		} else {
			if err = f.SetCellValue(ListName, cellName, ""); err != nil {
				return "", err
			} else {
				Incr()
			}
		}
		if data.InbrindingCoeffByFamily != nil {
			if err = f.SetCellValue(ListName, cellName, *data.InbrindingCoeffByFamily); err != nil {
				return "", err
			} else {
				Incr()
			}
		} else {
			if err = f.SetCellValue(ListName, cellName, ""); err != nil {
				return "", err
			} else {
				Incr()
			}
		}
		if data.InbrindingCoeffByGenotype != nil {
			if err = f.SetCellValue(ListName, cellName, *data.InbrindingCoeffByGenotype); err != nil {
				return "", err
			} else {
				Incr()
			}
		} else {
			if err = f.SetCellValue(ListName, cellName, ""); err != nil {
				return "", err
			} else {
				Incr()
			}
		}
		if data.ExteriorRating != nil {
			if err = f.SetCellValue(ListName, cellName, *data.ExteriorRating); err != nil {
				return "", err
			} else {
				Incr()
			}
		} else {
			if err = f.SetCellValue(ListName, cellName, ""); err != nil {
				return "", err
			} else {
				Incr()
			}
		}
		if data.SexName != nil {
			if err = f.SetCellValue(ListName, cellName, *data.SexName); err != nil {
				return "", err
			} else {
				Incr()
			}
		} else {
			if err = f.SetCellValue(ListName, cellName, ""); err != nil {
				return "", err
			} else {
				Incr()
			}
		}
		if data.HozName != nil {
			if err = f.SetCellValue(ListName, cellName, *data.HozName); err != nil {
				return "", err
			} else {
				Incr()
			}
		} else {
			if err = f.SetCellValue(ListName, cellName, ""); err != nil {
				return "", err
			} else {
				Incr()
			}
		}
		if data.DeathDate != nil {
			if err = f.SetCellValue(ListName, cellName, data.DeathDate.Time); err != nil {
				return "", err
			} else {
				Incr()
			}
		} else {
			if err = f.SetCellValue(ListName, cellName, ""); err != nil {
				return "", err
			} else {
				Incr()
			}
		}
		if data.IsDead != nil {
			if err = f.SetCellValue(ListName, cellName, *data.IsDead); err != nil {
				return "", err
			} else {
				Incr()
			}
		} else {
			if err = f.SetCellValue(ListName, cellName, ""); err != nil {
				return "", err
			} else {
				Incr()
			}
		}
		if data.IsTwins != nil {
			if err = f.SetCellValue(ListName, cellName, *data.IsTwins); err != nil {
				return "", err
			} else {
				Incr()
			}
		} else {
			if err = f.SetCellValue(ListName, cellName, ""); err != nil {
				return "", err
			} else {
				Incr()
			}
		}
		if data.IsStillBorn != nil {
			if err = f.SetCellValue(ListName, cellName, *data.IsStillBorn); err != nil {
				return "", err
			} else {
				Incr()
			}
		} else {
			if err = f.SetCellValue(ListName, cellName, ""); err != nil {
				return "", err
			} else {
				Incr()
			}
		}
		if data.IsAborted != nil {
			if err = f.SetCellValue(ListName, cellName, *data.IsAborted); err != nil {
				return "", err
			} else {
				Incr()
			}
		} else {
			if err = f.SetCellValue(ListName, cellName, ""); err != nil {
				return "", err
			} else {
				Incr()
			}
		}
		if data.IsGenotyped != nil {
			if err = f.SetCellValue(ListName, cellName, *data.IsGenotyped); err != nil {
				return "", err
			} else {
				Incr()
			}
		} else {
			if err = f.SetCellValue(ListName, cellName, ""); err != nil {
				return "", err
			} else {
				Incr()
			}
		}
		if data.CreatedAt != nil {
			if err = f.SetCellValue(ListName, cellName, data.CreatedAt.Time); err != nil {
				return "", err
			} else {
				Incr()
			}
		} else {
			if err = f.SetCellValue(ListName, cellName, ""); err != nil {
				return "", err
			} else {
				Incr()
			}
		}

	}

	// Сохраняем файл в cow_backend\frontend\static
	now := time.Now()
	fullPath := PathToExcelFile + "filtered_data_" + strconv.FormatInt(now.Unix(), 16) + "_" + strconv.FormatUint(uint64(len(fsc)), 16) + ".xlsx"
	if err := f.SaveAs(fullPath); err != nil {
		return "", fmt.Errorf("Ошибка создания Excel файла Error: %v", err)
	} else {
		return fullPath, nil
	}

}

func getHeaders() []string { // Получаем заголовки таблицы
	var headers []string
	t := reflect.TypeOf(FilterSerializedCow{})
	v := reflect.ValueOf(FilterSerializedCow{})

	for i := 1; i < t.NumField(); i++ { // Не берем поле id
		if v.Field(i).Kind() == reflect.Slice {
			continue
		}
		field := t.Field(i)
		headers = append(headers, field.Name)
	}

	return headers
}
