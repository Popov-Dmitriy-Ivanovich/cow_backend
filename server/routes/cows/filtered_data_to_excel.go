package cows

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/xuri/excelize/v2"
)



func ToExcel(fsc []FilterSerializedCow) (error) {
	f := excelize.NewFile()
	defer f.Close()
	// Создаем новый лист
	index, err := f.NewSheet("List1")
	if err != nil {
		return err
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
		f.SetCellValue("List1", cell, columnName)
	}

	// Записываем данные
	for row, data := range fsc {
		if err := checkFieldIsReuireNoEmpty(&data); err != nil{
			return err
		}
		v := reflect.ValueOf(data)
		for i := 0; i < v.NumField(); i++ {

			// Пропускаем запись срезов
			if isSlice(i, &data) {
				continue
			}
			
		    cell, err := excelize.CoordinatesToCellName(i+1, row+2)
			if err != nil {
				return err
			}
			cellValue := v.Field(i).Interface
			field := reflect.ValueOf(cellValue)
			switch field.Kind() {
			case reflect.Ptr:
				if field.IsNil() {
					f.SetCellValue("List1", cell, nil)
				} else {
					f.SetCellValue("List1", cell, field.Elem()) // Разыменование указателя
				}
			default:
				f.SetCellValue("List1", cell, field.Interface())	
			}
			
		}
	}

	// Сохраняем файл в cow_backend\frontend\static
	fullPath := "frontend/static/" + "filtered_data.xlsx"
	if err := f.SaveAs(fullPath); err != nil {
		return err
	}else {
		return nil
	}

}

func isSlice(i int, data *FilterSerializedCow) bool {
	fl := reflect.ValueOf(data)
	return fl.Field(i).Kind() == reflect.Slice
}

func getHeaders() []string { // Получаем заголовки таблицы
    var headers []string
	t := reflect.TypeOf(FilterSerializedCow{})

	for i:=1; i<t.NumField(); i++ { // Не берем поле id
		if isSlice(i, &FilterSerializedCow{}) { // Пропускаем списки на стадии поиска заголовков
			continue
		}
	    field := t.Field(i)
		headers = append(headers, field.Name)
	}

	return headers
}

func checkFieldIsReuireNoEmpty(data *FilterSerializedCow) error {

	validate := validator.New()
	err := validate.Struct(data) // Если обязательные поля не заполнены, то валидация не пройдет
	if err != nil {
		return err
	}

	return nil
}

