package cows_filter

import "cow_backend/filters"

var ALL_FILTERS = []filters.Filter{
	ByAbort{},
	ByAnyIllneses{},
	ByBirkingDate{},
	ByBreed{},
	ByBrithDate{},
	ByCalvingDate{},
	ByControlMilkingDate{},
	ByCreatedAt{},
	ByDeath{},
	ByDepartDate{},
	ByExterior{},
	ByHoz{},
	ByIllDate{},
	ByInbrindingCoeffByFamily{},
	ByInbrindingCoeffByGenotype{},
	ByIsGenotyped{},
	ByInsemenationDate{},
	BySearchString{},
	BySex{},
	ByStillBorn{},
	ByTwins{},
	ByMonogeneticIllnesses{},
	OrderBy{},
	ByRegion{},
	AliveInYear{},
	LiveInDistrict{},
	LiveInHolding{},
	LiveInRegion{},
}
