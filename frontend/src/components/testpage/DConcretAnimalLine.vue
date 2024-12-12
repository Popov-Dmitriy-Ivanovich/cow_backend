<template>
    <div>
        <div class="animal-line" @click="$router.push(`/animals/${animal_item.ID}`)" :class="{'is-approved': animal_item.Approved==1}">
            <div class="animal-rshn">{{ animal_item.RSHNNumber }}</div>
            <div class="animal-inv">{{ animal_item.InventoryNumber }}</div>
            <div class="animal-name">{{ animal_item.Name }}</div>
            <div class="animal-hoz">{{ animal_item.FarmGroupName }}</div>
            <div class="animal-bdate">{{ dateConverter(bdate(animal_item.BirthDate)) }}</div>
            <div class="animal-genfact">{{ isGen(animal_item.Genotyped) }}</div>

            <div v-if="animal_item.DepartDate" class="animal-dateout">{{ dateConverter(animal_item.DepartDate) }}</div>
            <div v-else-if="filters.departDateFrom || filters.departDateTo" class="animal-dateout"> - </div>

            <div v-if="animal_item.IsDead===true || animal_item.IsDead===false" class="animal-dead">{{ isGen(animal_item.IsDead) }}</div>
            <div v-else-if="filters.isDead===true || filters.isDead===false" class="animal-dead"> - </div>

            <div v-if="animal_item.BreedName" class="animal-breed">{{ animal_item.BreedName }}</div>
            <div v-else-if="filters.breedId" class="animal-breed"> - </div>

            <div v-if="animal_item.CheckMilkDate" class="animal-contrmilk">{{ dateConverter(check_milk) }}</div>
            <div v-else-if="filters.controlMilkingDateFrom || filters.controlMilkingDateTo" class="animal-contrmilk"> - </div>

            <div v-if="animal_item.CreatedAt" class="animal-contrmilk">{{ dateConverter(animal_item.CreatedAt) }}</div>
            <div v-else-if="filters.createdAtFrom || filters.createdAtTo" class="animal-krod"> - </div>

            <div v-if="animal_item.Exterior" class="animal-exterior">{{ animal_item.Exterior }}</div>
            <div v-else-if="filters.exteriorFrom || filters.exteriorTo" class="animal-exterior"> - </div>

            <div v-if="animal_item.InsemenationDate" class="animal-dateosem">{{ dateConverter(insemination) }}</div>
            <div v-else-if="filters.insemenationDateFrom || filters.inseminationDateTo" class="animal-dateosem"> - </div>

            <div v-if="animal_item.CalvingDate" class="animal-dateotel">{{ dateConverter(calving) }}</div>
            <div v-else-if="filters.calvingDateFrom || filters.calvingDateTo" class="animal-dateotel"> - </div>

            <div v-if="animal_item.IsTwins===true || animal_item.IsTwins===false" class="animal-genfact">{{ isGen(animal_item.IsTwins) }}</div>
            <div v-else-if="filters.isTwins===true || filters.isTwins===false" class="animal-genfact"> - </div>

            <div v-if="animal_item.IsStillBorn===true || animal_item.IsStillBorn===false" class="animal-genfact">{{ isGen(animal_item.IsStillBorn) }}</div>
            <div v-else-if="filters.isStillBorn===true || filters.isStillBorn===false" class="animal-genfact"> - </div>

            <div v-if="animal_item.IsAborted===true || animal_item.IsAborted===false" class="animal-genfact">{{ isGen(animal_item.IsAborted) }}</div>
            <div v-else-if="filters.isAborted===true || filters.isAborted===false" class="animal-genfact"> - </div>

            <div v-if="animal_item.BirkingDate" class="animal-datebirk">{{ dateConverter(animal_item.BirkingDate) }}</div>
            <div v-else-if="filters.birkingDateFrom || filters.birkingDateTo" class="animal-datebirk"> - </div>

            <div v-if="animal_item.InbrindingCoeffByFamily" class="animal-krod">{{ animal_item.InbrindingCoeffByFamily }}</div>
            <div v-else-if="filters.inbrindingCoeffByFamilyFrom || filters.inbrindingCoeffByFamilyTo" class="animal-krod"> - </div>

            <div v-if="animal_item.InbrindingCoeffByGenotype" class="animal-kfen">{{ animal_item.InbrindingCoeffByGenotype }}</div>
            <div v-else-if="filters.inbrindingCoeffByGenotypeFrom || filters.inbrindingCoeffByGenotypeTo" class="animal-kfen"> - </div>

            <div v-if="animal_item.Events" class="animal-kfen">{{ dateConverter(cevent) }}</div>
            <div v-else-if="filters.illDateFrom || filters.illDateTo" class="animal-krod"> - </div>
        </div>
    </div>
</template>

<script>
export default {
    props: {
        animal_item: {
            type: Object,
            Required: true
        },
        filters: {
            type: Object,
        }
    },
    data() {
        return {
            check_milk: '',
            insemination: '',
            calving: '',
            cevent: '',
        }
    },
    methods: {
        isGen(a) {
            if(a) return 'Да';
            else return 'Нет';
        },
        bdate(birth) {
            return birth.split('T')[0];
        },
        dateConverter(date) {
            let arr = date.split('-');
            let result = '';
            result += arr[2]; result += '.';
            result += arr[1]; result += '.';
            result += arr[0];
            return result;
        }
    },
    mounted() {
        this.check_milk = '';
        if (this.animal_item.CheckMilkDate) {
            if (this.animal_item.CheckMilkDate.length) {
                this.check_milk = this.animal_item.CheckMilkDate[0];
            }
        }
        this.insemination = '';
        if (this.animal_item.InsemenationDate) {
            if (this.animal_item.InsemenationDate.length) {
                this.insemination = this.animal_item.InsemenationDate[0];
            }
        }
        this.calving = '';
        if (this.animal_item.CalvingDate) {
            if(this.animal_item.CalvingDate.length) {
                this.calving = this.animal_item.CalvingDate[0];
            }
        }
        this.cevent = '';
        if (this.animal_item.Events) {
            if (this.animal_item.Events.length) {
                if (this.animal_item.Events[0].Date) {
                    this.cevent = this.animal_item.Events[0].Date;
                }
            }
        }
    },
}
</script>

<style scoped>
.animal-line {
    display: flex;
    padding: 10px 0;
    transition: 0.3s;
    cursor: pointer;
    width: max-content;
    color: grey;
}

.is-approved {
    color: black;
}

.animal-line:hover {
    background-color: rgb(235, 233, 245);
}

.animal-line div {
    padding: 0 10px;
}

.animal-selex {
    width: 140px;
}

.animal-inv {
    width: 115px;
}

.animal-rshn {
    width: 120px;
}

.animal-name {
    width: 140px;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
}

.animal-hoz {
    width: 230px;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
}

.animal-bdate {
    width: 120px;
}

.animal-genfact {
    width: 120px;
}

.animal-dateout {
    width: 120px;
}

.animal-breed, .animal-dategen, .animal-datemilking, 
.animal-dateosem, .animal-dateotel, .animal-datebirk {
    width: 150px;
    height: 100%;
}

.animal-exterior {
    width: 100px;
}

.animal-krod, .animal-kfen {
    width: 100px;
}

.animal-dead {
    width: 100px;
}

.animal-contrmilk  {
    width: 130px;
}
</style>