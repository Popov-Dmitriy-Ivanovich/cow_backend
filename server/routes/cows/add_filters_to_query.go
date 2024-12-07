package cows

import (
	"time"

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

func AddFiltersToQuery(bodyData CowsFilter, query *gorm.DB) (*gorm.DB, error) {

	if bodyData.IsGenotyped != nil && *bodyData.IsGenotyped {
		query = query.Where("EXISTS (SELECT 1 FROM genetics where genetics.cow_id = cows.id)")
	}
	if bodyData.IsGenotyped != nil && !*bodyData.IsGenotyped {
		query = query.Where("NOT EXISTS (SELECT 1 FROM genetics where genetics.cow_id = cows.id)")
	}

	if searchString := bodyData.SearchQuery; searchString != nil && *searchString != "" {
		*searchString = "%" + *searchString + "%"
		query = query.Where("name LIKE ? or rshn_number LIKE ? or inventory_number LIKE ? or selecs_number like ?", searchString, searchString, searchString, searchString)
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
	if bodyData.IsIll != nil && *bodyData.IsIll {
		if len(bodyData.MonogeneticIllneses) != 0 {
			query = query.Where("EXISTS (SELECT 1 FROM genetics where genetics.cow_id = cows.id AND EXISTS (SELECT 1 FROM genetic_genetic_illnesses WHERE genetic_genetic_illnesses.genetic_id = genetics.id AND genetic_illness_id IN ?) )", bodyData.MonogeneticIllneses).Preload("Genetic").Preload("Genetic.GeneticIllnesses")
		}
	}
	if bodyData.IsIll != nil && !*bodyData.IsIll {
		if len(bodyData.MonogeneticIllneses) != 0 {
			query = query.Where("EXISTS (SELECT 1 FROM genetics where genetics.cow_id = cows.id AND NOT EXISTS (SELECT 1 FROM genetic_genetic_illnesses WHERE genetic_genetic_illnesses.genetic_id = genetics.id AND genetic_illness_id IN ?) )", bodyData.MonogeneticIllneses).Preload("Genetic").Preload("Genetic.GeneticIllnesses")
		}
	}
	if bodyData.HasAnyIllnes != nil && *bodyData.HasAnyIllnes {
		query = query.Where("EXISTS (SELECT 1 FROM genetics where genetics.cow_id = cows.id AND EXISTS (SELECT 1 FROM genetic_genetic_illnesses WHERE genetic_genetic_illnesses.genetic_id = genetics.id) )").Preload("Genetic").Preload("Genetic.GeneticIllnesses")
	}
	if bodyData.HasAnyIllnes != nil && !*bodyData.HasAnyIllnes {
		query = query.Where("EXISTS (SELECT 1 FROM genetics where genetics.cow_id = cows.id AND NOT EXISTS (SELECT 1 FROM genetic_genetic_illnesses WHERE genetic_genetic_illnesses.genetic_id = genetics.id) )").Preload("Genetic").Preload("Genetic.GeneticIllnesses")
	}
	if bodyData.IsDead != nil && *bodyData.IsDead {
		query = query.Where("death_date IS NOT NULL")
	}
	if bodyData.IsDead != nil && !*bodyData.IsDead {
		query = query.Where("death_date IS NULL")
	}

	// ====================================================================================================
	// =============================== Filter by ill date from ============================================
	// ====================================================================================================

	if bodyData.CreatedAtFrom != nil && bodyData.CreatedAtTo != nil &&
		*bodyData.CreatedAtFrom != "" && *bodyData.CreatedAtTo != "" {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.CreatedAtFrom)
		if err != nil {
			return nil, err
		}
		bdTo, err := time.Parse(time.DateOnly, *bodyData.CreatedAtTo)
		if err != nil {
			return nil, err
		}
		query = query.Where("created_at BETWEEN ? AND ?", bdFrom, bdTo)
	} else if bodyData.CreatedAtFrom != nil && *bodyData.CreatedAtFrom != "" {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.CreatedAtFrom)
		if err != nil {
			return nil, err
		}
		query = query.Where("created_at >= ?", bdFrom.UTC())
	} else if bodyData.CreatedAtTo != nil && *bodyData.CreatedAtTo != "" {
		bdTo, err := time.Parse(time.DateOnly, *bodyData.CreatedAtTo)
		if err != nil {
			return nil, err
		}
		query = query.Where("created_at <= ?)", bdTo.UTC())
	}

	// ====================================================================================================
	// ============================== Filter by Ill date from events  =====================================
	// ====================================================================================================

	if bodyData.IllDateFrom != nil && bodyData.IllDateTo != nil &&
		*bodyData.IllDateFrom != "" && *bodyData.IllDateTo != "" {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.IllDateFrom)
		if err != nil {
			return nil, err
		}
		bdTo, err := time.Parse(time.DateOnly, *bodyData.IllDateTo)
		if err != nil {
			return nil, err
		}
		query = query.Where("EXISTS( SELECT 1 FROM events where events.cow_id = cows.id AND events.event_type_id in (1, 2, 3, 4) AND events.date BETWEEN ? AND ? )",
			bdFrom,
			bdTo).Preload("Events")
	} else if bodyData.IllDateFrom != nil && *bodyData.IllDateFrom != "" {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.IllDateFrom)
		if err != nil {
			return nil, err
		}
		query = query.Where("EXISTS( SELECT 1 FROM events where events.cow_id = cows.id AND events.event_type_id in (1, 2, 3, 4) AND events.date >= ? )", bdFrom.UTC()).Preload("Events")
	} else if bodyData.IllDateTo != nil && *bodyData.IllDateTo != "" {
		bdTo, err := time.Parse(time.DateOnly, *bodyData.IllDateTo)
		if err != nil {
			return nil, err
		}
		query = query.Where("EXISTS( SELECT 1 FROM events where events.cow_id = cows.id AND events.event_type_id in (1, 2, 3, 4) AND events.date <= ?)", bdTo.UTC()).Preload("Events")
	}

	// ====================================================================================================
	// ========================= Filter by inbrinding coeff by date of genotyping =========================
	// ====================================================================================================
	if bodyData.GenotypingDateFrom != nil && bodyData.GenotypingDateTo != nil &&
		*bodyData.GenotypingDateFrom != "" && *bodyData.GenotypingDateTo != "" {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.GenotypingDateFrom)
		if err != nil {
			return nil, err
		}
		bdTo, err := time.Parse(time.DateOnly, *bodyData.GenotypingDateTo)
		if err != nil {
			return nil, err
		}
		query = query.Where("EXISTS( SELECT 1 FROM genetics where genetics.cow_id = cows.id AND genetics.result_date BETWEEN ? AND ?)",
			bdFrom,
			bdTo).Preload("Genetic")
	} else if bodyData.GenotypingDateFrom != nil && *bodyData.GenotypingDateFrom != "" {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.GenotypingDateFrom)
		if err != nil {
			return nil, err
		}
		query = query.Where("EXISTS( SELECT 1 FROM genetics where genetics.cow_id = cows.id AND genetics.result_date >= ?)", bdFrom.UTC()).Preload("Genetic")
	} else if bodyData.GenotypingDateTo != nil && *bodyData.GenotypingDateTo != "" {
		bdTo, err := time.Parse(time.DateOnly, *bodyData.GenotypingDateTo)
		if err != nil {
			return nil, err
		}
		query = query.Where("EXISTS( SELECT 1 FROM genetics where genetics.cow_id = cows.id AND genetics.result_date <= ?)", bdTo.UTC()).Preload("Genetic")
	}
	// ====================================================================================================
	// =================================== Filter by inbrinding coeff by genotype =========================
	// ====================================================================================================
	if bodyData.InbrindingCoeffByGenotypeFrom != nil && bodyData.InbrindingCoeffByGenotypeTo != nil {
		query = query.Where("EXISTS( SELECT 1 FROM genetics where genetics.cow_id = cows.id AND genetics.inbrinding_coeff_by_genotype BETWEEN ? AND ?)",
			bodyData.InbrindingCoeffByGenotypeFrom,
			bodyData.InbrindingCoeffByGenotypeTo).Preload("Genetic")
	} else if bodyData.InbrindingCoeffByGenotypeFrom != nil {
		query = query.Where("EXISTS( SELECT 1 FROM genetics where genetics.cow_id = cows.id AND genetics.inbrinding_coeff_by_genotype >= ?)",
			bodyData.InbrindingCoeffByGenotypeFrom).Preload("Genetic")
	} else if bodyData.InbrindingCoeffByGenotypeTo != nil {
		query = query.Where("EXISTS( SELECT 1 FROM genetics where genetics.cow_id = cows.id AND genetics.inbrinding_coeff_by_genotype >= ?)",
			bodyData.InbrindingCoeffByGenotypeTo).Preload("Genetic")
	}

	// ====================================================================================================
	// =================================== Filter by brithday date ========================================
	// ====================================================================================================
	if bodyData.BirthDateFrom != nil && bodyData.BirthDateTo != nil &&
		*bodyData.BirthDateFrom != "" && *bodyData.BirthDateTo != "" {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.BirthDateFrom)
		if err != nil {
			return nil, err
		}
		bdTo, err := time.Parse(time.DateOnly, *bodyData.BirthDateTo)
		if err != nil {
			return nil, err
		}
		query = query.Where("birth_date BETWEEN ? AND ?", bdFrom.UTC(), bdTo.UTC())
	} else if bodyData.BirthDateFrom != nil && *bodyData.BirthDateFrom != "" {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.BirthDateFrom)
		if err != nil {
			return nil, err
		}
		query = query.Where("birth_date >= ?", bdFrom.UTC())
	} else if bodyData.BirthDateTo != nil && *bodyData.BirthDateTo != "" {
		bdTo, err := time.Parse(time.DateOnly, *bodyData.BirthDateTo)
		if err != nil {
			return nil, err
		}
		query = query.Where("birth_date <= ?", bdTo.UTC())
	}

	// ====================================================================================================
	// ================================ Filter by departation date ========================================
	// ====================================================================================================
	if bodyData.DepartDateFrom != nil && bodyData.DepartDateTo != nil &&
		*bodyData.DepartDateFrom != "" && *bodyData.DepartDateTo != "" {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.DepartDateFrom)
		if err != nil {
			return nil, err
		}
		bdTo, err := time.Parse(time.DateOnly, *bodyData.DepartDateTo)
		if err != nil {
			return nil, err
		}
		query = query.Where("depart_date BETWEEN ? AND ?", bdFrom.UTC(), bdTo.UTC())
	} else if bodyData.DepartDateFrom != nil && *bodyData.DepartDateFrom != "" {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.DepartDateFrom)
		if err != nil {
			return nil, err
		}
		query = query.Where("depart_date >= ?", bdFrom.UTC())
	} else if bodyData.DepartDateTo != nil && *bodyData.DepartDateTo != "" {
		bdTo, err := time.Parse(time.DateOnly, *bodyData.DepartDateTo)
		if err != nil {
			return nil, err
		}
		query = query.Where("depart_date <= ?", bdTo.UTC())
	}

	// ====================================================================================================
	// ============================ Filter by control milking date ========================================
	// ====================================================================================================
	if bodyData.ControlMilkingDateFrom != nil && bodyData.ControlMilkingDateTo != nil &&
		*bodyData.ControlMilkingDateFrom != "" && *bodyData.ControlMilkingDateTo != "" {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.ControlMilkingDateFrom)
		if err != nil {
			return nil, err
		}
		bdTo, err := time.Parse(time.DateOnly, *bodyData.ControlMilkingDateTo)
		if err != nil {
			return nil, err
		}
		query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND EXISTS (SELECT 1 FROM check_milks WHERE check_milks.lactation_id = lactations.id AND check_milks.check_date BETWEEN ? AND ?))", bdFrom.UTC(), bdTo.UTC()).Preload("Lactation").Preload("Lactation.CheckMilks")
	} else if bodyData.ControlMilkingDateFrom != nil && *bodyData.ControlMilkingDateFrom != "" {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.ControlMilkingDateFrom)
		if err != nil {
			return nil, err
		}
		query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND EXISTS (SELECT 1 FROM check_milks WHERE check_milks.lactation_id = lactations.id AND check_milks.check_date >= ?))", bdFrom.UTC()).Preload("Lactation").Preload("Lactation.CheckMilks")
	} else if bodyData.ControlMilkingDateTo != nil && *bodyData.ControlMilkingDateTo != "" {
		bdTo, err := time.Parse(time.DateOnly, *bodyData.ControlMilkingDateTo)
		if err != nil {
			return nil, err
		}
		query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND EXISTS (SELECT 1 FROM check_milks WHERE check_milks.lactation_id = lactations.id AND check_milks.check_date <= ?))", bdTo.UTC()).Preload("Lactation").Preload("Lactation.CheckMilks")
	}

	// ====================================================================================================
	// ==================================== Filter by calving date ========================================
	// ====================================================================================================
	if bodyData.CalvingDateFrom != nil && bodyData.CalvingDateTo != nil &&
		*bodyData.CalvingDateFrom != "" && *bodyData.CalvingDateTo != "" {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.CalvingDateFrom)
		if err != nil {
			return nil, err
		}
		bdTo, err := time.Parse(time.DateOnly, *bodyData.CalvingDateTo)
		if err != nil {
			return nil, err
		}
		query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.calving_date BETWEEN ? AND ?)", bdFrom.UTC(), bdTo.UTC()).Preload("Lactation")
	} else if bodyData.CalvingDateFrom != nil && *bodyData.CalvingDateFrom != "" {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.CalvingDateFrom)
		if err != nil {
			return nil, err
		}
		query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.calving_date >= ?)", bdFrom.UTC()).Preload("Lactation")
	} else if bodyData.CalvingDateTo != nil && *bodyData.CalvingDateTo != "" {
		bdTo, err := time.Parse(time.DateOnly, *bodyData.CalvingDateTo)
		if err != nil {
			return nil, err
		}
		query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.calving_date <= ?)", bdTo.UTC()).Preload("Lactation")
	}

	// ====================================================================================================
	// ======================= Filter by birth parameters (e/g stillborn, twins ...)=======================
	// ====================================================================================================
	if bodyData.IsStillBorn != nil && *bodyData.IsStillBorn { // stillborn means, that 0 cows are born
		query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.calving_count = ?)", 0).Preload("Lactation")
	}
	if bodyData.IsStillBorn != nil && !*bodyData.IsStillBorn { // stillborn means, that 0 cows are born
		query = query.Where("NOT EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.calving_count = ?)", 0).Preload("Lactation")
	}

	if bodyData.IsTwins != nil && *bodyData.IsTwins { // twins means, that 2 cows are born
		query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.calving_count = ?)", 2).Preload("Lactation")
	}
	if bodyData.IsTwins != nil && !*bodyData.IsTwins { // twins means, that 2 cows are born
		query = query.Where("NOT EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.calving_count = ?)", 2).Preload("Lactation")
	}

	if bodyData.IsAborted != nil && *bodyData.IsAborted { // abort is marked by flag for some reason
		query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.abort = ?)", true).Preload("Lactation")
	}
	if bodyData.IsAborted != nil && !*bodyData.IsAborted { // abort is marked by flag for some reason
		query = query.Where("NOT EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.abort = ?)", true).Preload("Lactation")
	}

	if bodyData.ExteriorFrom != nil && bodyData.ExteriorTo != nil {
		query = query.Where("EXISTS(SELECT 1 FROM exteriors WHERE exteriors.cow_id = cows.id AND exteriors.rating BETWEEN ? AND ?)", bodyData.ExteriorFrom, bodyData.ExteriorTo).Preload("Exterior")
	} else if bodyData.ExteriorFrom != nil {
		query = query.Where("EXISTS(SELECT 1 FROM exteriors WHERE exteriors.cow_id = cows.id AND exteriors.rating >= ?)", bodyData.ExteriorFrom).Preload("Exterior")
	} else if bodyData.ExteriorTo != nil {
		query = query.Where("EXISTS(SELECT 1 FROM exteriors WHERE exteriors.cow_id = cows.id AND exteriors.rating <= ?)", bodyData.ExteriorTo).Preload("Exterior")
	}

	// ====================================================================================================
	// ===================================  Filter by birinkg date ========================================
	// ====================================================================================================
	if bodyData.BirkingDateFrom != nil && bodyData.BirkingDateTo != nil &&
		*bodyData.BirkingDateFrom != "" && *bodyData.BirkingDateTo != "" {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.BirkingDateFrom)
		if err != nil {
			return nil, err
		}
		bdTo, err := time.Parse(time.DateOnly, *bodyData.BirkingDateTo)
		if err != nil {
			return nil, err
		}
		query = query.Where("birking_date BETWEEN ? AND ?", bdFrom, bdTo)
	} else if bodyData.BirkingDateFrom != nil && *bodyData.BirkingDateFrom != "" {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.BirkingDateFrom)
		if err != nil {
			return nil, err
		}
		query = query.Where("birking_date >= ?", bdFrom)
	} else if bodyData.BirkingDateTo != nil && *bodyData.BirkingDateTo != "" {
		bdTo, err := time.Parse(time.DateOnly, *bodyData.BirkingDateTo)
		if err != nil {
			return nil, err
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
			return nil, err
		}
		bdTo, err := time.Parse(time.DateOnly, *bodyData.InseminationDateTo)
		if err != nil {
			return nil, err
		}
		query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.insemenation_date BETWEEN ? and ?)", bdFrom.UTC(), bdTo.UTC()).Preload("Lactation")
	} else if bodyData.InseminationDateFrom != nil {
		bdFrom, err := time.Parse(time.DateOnly, *bodyData.InseminationDateFrom)
		if err != nil {
			return nil, err
		}
		query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.insemenation_date >= ?)", bdFrom.UTC()).Preload("Lactation")
	} else if bodyData.InseminationDateTo != nil {
		bdTo, err := time.Parse(time.DateOnly, *bodyData.InseminationDateTo)
		if err != nil {
			return nil, err
		}
		query = query.Where("EXISTS (SELECT 1 FROM lactations WHERE lactations.cow_id = cows.id AND lactations.insemenation_date <= ?)", bdTo.UTC()).Preload("Lactation")
	}

	// ====================================================================================================
	return query, nil
}
