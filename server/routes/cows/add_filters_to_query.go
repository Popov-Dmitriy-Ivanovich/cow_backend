package cows

import (
	"cow_backend/filters"
	"cow_backend/filters/cows_filter"

	"gorm.io/gorm"
)

type CowsFilter struct { // Фильтр коров
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
	GenotypingDateFrom     *string `example:"1800-01-21" validate:"optional"` //фильтр по дате генотипирования ОТ
	GenotypingDateTo       *string `example:"2800-01-21" validate:"optional"` //фильтр по дате генотипирования ДО
	ControlMilkingDateFrom *string `example:"1800-01-21" validate:"optional"` // Фильтр по дате контрольной дойки, ищет коров у которых была контрольная дойка в эту дату или позднее
	ControlMilkingDateTo   *string `example:"2800-01-21" validate:"optional"` // Фильтр по дате контрольной дойки, ищет коров у которых была контрольная дойка в эту дату или ранее

	ExteriorFrom *float64 `default:"3.14" validate:"optional"` // Фильтр по оценке экстерьера коровы ОТ
	ExteriorTo   *float64 `default:"3.14" validate:"optional"` // Фильтр по оценке экстерьера коровы ДО
	// Exterior             *float64 `default:"3.14" validate:"optional"`       // Фильтр по оценке экстерьера коровы, будет переработан
	InseminationDateFrom *string `example:"1800-01-21" validate:"optional"` // Фильтр по дате осеменения коровы, ищет коров у которых было осеменение в эту дату или позднее
	InseminationDateTo   *string `example:"2800-01-21" validate:"optional"` // Фильтр по дате осеменения коровы, ищет коров у которых было осеменение в эту дату или ранее
	CalvingDateFrom      *string `example:"1800-01-21" validate:"optional"` // Фильтр по дате отела коровы, ищет коров у которых был отел в эту дату или позднее
	CalvingDateTo        *string `example:"2800-01-21" validate:"optional"` // Фильтр по дате осеменения коровы, ищет коров у которых был отел в эту дату или позднее
	IsStillBorn          *bool   `default:"false" validate:"optional"`      // Фильтр по мертворождению Было/не было
	IsTwins              *bool   `default:"false" validate:"optional"`      // Фильтр по родам двойняшек Было / не было
	IsAborted            *bool   `default:"false" validate:"optional"`      // Фильтр по абортам Был/ не был

	BirkingDateFrom *string `example:"1800-01-21" validate:"optional"` // Фильтр по дате перебирковки коровы, ищет коров у которых быа перебирковка в эту дату или позднее
	BirkingDateTo   *string `example:"2800-01-21" validate:"optional"` // Фильтр по дате осеменения коровы, ищет коров у которых была перебирковка в эту дату или позднее

	InbrindingCoeffByFamilyFrom *float64 `default:"3.14" validate:"optional"` // фильтр по коэф. инбриндинга по роду ОТ
	InbrindingCoeffByFamilyTo   *float64 `default:"3.14" validate:"optional"` // фильтр по коэф. инбриндинга по роду ДО

	InbrindingCoeffByGenotypeFrom *float64 `default:"3.14" validate:"optional"` //фильтр по коэф. инбриндинга по генотипу ОТ
	InbrindingCoeffByGenotypeTo   *float64 `default:"3.14" validate:"optional"` //фильтр по коэф. инбриндинга по генотипу ДО

	HasAnyIllnes        *bool  `default:"false" validate:"optional"` //Флаг true - возращает коров у которых есть хотябы одно заболевение, false - возращает коров, у которых нет ни одного
	IsIll               *bool  `default:"false" validate:"optional"` //??? Не реализован
	MonogeneticIllneses []uint // ID ген. заболеваний их /api/mongenetic_illnesses

	IllDateFrom *string `example:"1800-01-21" validate:"optional"` // Фильтр по дате начала болезни ОТ
	IllDateTo   *string `example:"1800-01-21" validate:"optional"` // Фильтр по дате начала болезни ДО

	IsGenotyped   *bool   `validate:"optional"`
	CreatedAtFrom *string `validate:"optional"`
	CreatedAtTo   *string `validate:"optional"`
}

func AddFiltersToQuery(bodyData cows_filter.CowsFilter, query *gorm.DB) (*gorm.DB, error) {
	cfm := cows_filter.NewCowFilteredModel(bodyData, query)
	if err := filters.ApplyFilters(cfm,
		cows_filter.ALL_FILTERS...); err != nil {
		return nil, err
	}
	return cfm.GetQuery(), nil
}
