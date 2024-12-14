<template>
    <div class="parent-table">
        <div class="animals-table">
        <div class="table-header">
            <div class="animal-rshn">Номер РСХН</div>
            <div class="animal-inv">Инвентарный номер</div>
            <div class="animal-name">Кличка</div>
            <div class="animal-hoz">Хозяйство</div>
            <div class="animal-bdate">Дата рождения</div>
            <div class="animal-genfact">Факт генотипирования</div>

            <div v-if="filters.departDateFrom || filters.departDateTo" class="animal-dateout">Дата выбытия</div>
            <div v-if="filters.isDead===true || filters.isDead===false" class="animal-dead">Животное мертво</div>
            <div v-if="filters.breedId" class="animal-breed">Порода</div>
            <div v-if="filters.genotypingDateFrom || filters.genotypingDateTo" class="animal-dategen">Дата генотипирования</div>
            <div v-if="filters.controlMilkingDateFrom || filters.controlMilkingDateTo" class="animal-contrmilk">Дата контрольной дойки</div>
            <div v-if="filters.createdAtFrom || filters.createdAtTo" class="animal-krod">Дата внесения данных о КРС</div>
            <div v-if="filters.exteriorFrom || filters.exteriorTo" class="animal-exterior">Оценка экстерьера</div>            
            <div v-if="filters.inseminationDateFrom || filters.inseminationDateTo" class="animal-dateosem">Дата осеменения</div>
            <div v-if="filters.calvingDateFrom || filters.calvingDateTo" class="animal-dateotel">Дата отела</div>
            <div v-if="filters.isTwins===true || filters.isTwins===false" class="animal-genfact">Двойня</div>
            <div v-if="filters.isStillBorn===true || filters.isStillBorn===false" class="animal-genfact">Мертворождённый</div>
            <div v-if="filters.isAborted===true || filters.isAborted===false" class="animal-genfact">Аборт</div>
            <div v-if="filters.birkingDateFrom || filters.birkingDateTo" class="animal-datebirk">Дата перебирковки</div>
            <div v-if="(filters.inbrindingCoeffByFamilyFrom || filters.inbrindingCoeffByFamilyFrom===0) || (filters.inbrindingCoeffByFamilyTo || filters.inbrindingCoeffByFamilyTo===0)" class="animal-krod">Коэффициент инбридинга по родословной</div>
            <div v-if="(filters.inbrindingCoeffByGenotypeFrom || filters.inbrindingCoeffByGenotypeFrom===0) || (filters.inbrindingCoeffByGenotypeTo || filters.inbrindingCoeffByGenotypeTo===0)" class="animal-kfen">Коэффициент инбридинга по генотипу</div>
            <div v-if="filters.illDateFrom || filters.illDateTo" class="animal-krod">Дата заболевания</div>
            <div v-if="filters.monogeneticIllneses && filters.monogeneticIllneses.length || filters.isIll || filters.hasAnyIllnes" class="animal-krod">
                    HCD
                </div>
                <div v-if="filters.monogeneticIllneses && filters.monogeneticIllneses.length || filters.isIll || filters.hasAnyIllnes" class="animal-krod">
                    HH1
                </div>
                <div v-if="filters.monogeneticIllneses && filters.monogeneticIllneses.length || filters.isIll || filters.hasAnyIllnes" class="animal-krod">
                    HH3
                </div>
                <div v-if="filters.monogeneticIllneses && filters.monogeneticIllneses.length || filters.isIll || filters.hasAnyIllnes" class="animal-krod">
                    HH4
                </div>
                <div v-if="filters.monogeneticIllneses && filters.monogeneticIllneses.length || filters.isIll || filters.hasAnyIllnes" class="animal-krod">
                    HH5
                </div>
                <div v-if="filters.monogeneticIllneses && filters.monogeneticIllneses.length || filters.isIll || filters.hasAnyIllnes" class="animal-krod">
                    HH6
                </div>
                <div v-if="filters.monogeneticIllneses && filters.monogeneticIllneses.length || filters.isIll || filters.hasAnyIllnes" class="animal-krod">
                    BLAD
                </div>
                <div v-if="filters.monogeneticIllneses && filters.monogeneticIllneses.length || filters.isIll || filters.hasAnyIllnes" class="animal-krod">
                    CVM
                </div>
                <div v-if="filters.monogeneticIllneses && filters.monogeneticIllneses.length || filters.isIll || filters.hasAnyIllnes" class="animal-krod">
                    DUMPS
                </div>
                <div v-if="filters.monogeneticIllneses && filters.monogeneticIllneses.length || filters.isIll || filters.hasAnyIllnes" class="animal-krod">
                    BC
                </div>
                <div v-if="filters.monogeneticIllneses && filters.monogeneticIllneses.length || filters.isIll || filters.hasAnyIllnes" class="animal-krod">
                    MF
                </div>
                <div v-if="filters.monogeneticIllneses && filters.monogeneticIllneses.length || filters.isIll || filters.hasAnyIllnes" class="animal-krod">
                    FGFR2
                </div>
                <div v-if="filters.monogeneticIllneses && filters.monogeneticIllneses.length || filters.isIll || filters.hasAnyIllnes" class="animal-krod">
                    IH
                </div>
                <div v-if="filters.monogeneticIllneses && filters.monogeneticIllneses.length || filters.isIll || filters.hasAnyIllnes" class="animal-krod">
                    FXID
                </div>
        </div>
        <div v-if="!isSearch&!search_error">
            <div v-for="animal in animals" :key="animal[0]">
                <DConcretAnimalLine v-bind:animal_item="animal" v-bind:filters="filters"/>
            </div>
        </div>
        <div v-if="isSearch&!search_error">
            <div v-for="search_animal in search_result" :key="search_animal[0]">
                <DConcretAnimalLine v-bind:animal_item="search_animal" v-bind:filters="filters"/>
            </div>
        </div>
        <div v-if="search_error || errorr" class="search-error">
            Ничего не найдено
        </div>
        <div v-else-if="isLoading || isLoadingChild" class="search-error">Идёт загрузка...</div>
    </div>
    <NumberPages v-bind:current-page="cp" v-bind:total-pages="tp" @changePage="changePage"/>
</div>
</template>

<script>
import DConcretAnimalLine from './DConcretAnimalLine.vue';
import NumberPages from '@/components/NumberPages.vue';

export default {
    components: {
        DConcretAnimalLine, NumberPages
    },
    data () {
        return {
            animals: [],
            errorr: false,
            isLoadingChild: false,
        }
    },
    props: {
        isSearch: {
            Type: Boolean
        },
        search_result: {
            type: Array
        },
        search_error: {
            type: Boolean
        },
        cp: {
            type: Number,
            required: true
        },
        tp: {
            type: Number,
            required: true
        },
        filters: {
            type: Object,
        },
        isLoading: {
            type: Boolean,
        }
    },
    methods: {
        changePage(new_value) {
            this.$emit('defPages', new_value, this.tp);
        }
    },
    async mounted() {
        this.isLoadingChild = true;
        let search_params = this.filters;
        search_params.sex = [1,2];
        search_params.pageNumber = 1;
        search_params.entitiesOnPage = 25;
        if(!search_params.orderBy) {
            search_params.orderBy = 'Name';
            search_params.orderByDesc = false;
        }
        const response = await fetch('/api/cows/filter', {            
            method: 'POST',
            headers: {
                'Content-Type': 'application/json;charset=utf-8'
            },
            body: JSON.stringify(search_params),
        });
        const res_animals = await response.json();

        this.animals = res_animals.LST;
        if(res_animals.LST.length == 0) this.errorr = true;
        //Передаю текущую первую страницу и кол-во страниц наверх
        this.$emit('defPages', search_params.pageNumber, Math.ceil(res_animals.N/search_params.entitiesOnPage));
        this.isLoadingChild = false;
        console.log(res_animals);
    },
    watch: {
        async cp(newValue) {

            if(!this.isSearch) {
                let search_params = {};
                search_params.sex = [1,2];
                search_params.pageNumber = newValue;
                search_params.entitiesOnPage = 25;
                const response = await fetch('/api/cows/filter', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json;charset=utf-8'
                    },
                    body: JSON.stringify(search_params),
                });
                const res_animals = await response.json();
                this.animals = res_animals.LST;
                console.log(response);
            } else {
                this.$emit('changePageButSearch', newValue);
            }

        },
        search_result(newVal) {
            if(newVal.length == 0 && this.isSearch) this.errorr = true;
            else this.errorr = false;
        }
    },
}
</script>

<style scoped>
.parent-table {
    display: flex;
    flex-direction: column;
    align-items: center;
}

.animals-table {
    width: 930px;
    height: max-content;
    min-height: 500px;
    background-color: white;
    box-shadow: rgba(100, 100, 111, 0.1) 0px 7px 29px 0px;
    margin: 0 50px 0 10px;
    font-family: Open Sans, sans-serif;
    padding: 20px 30px;
    overflow-x: auto;
    font-size: 90%;
}

.animals-table::-webkit-scrollbar {
    height: 12px;
}

.animals-table::-webkit-scrollbar-track {
    background: rgb(241, 241, 241);
}

.animals-table::-webkit-scrollbar-thumb {
    background-color: rgb(183, 183, 183);
    border-radius: 20px;
    border: 3px solid rgb(241, 241, 241);
}

.table-header {
    display: flex;
    color: grey;
    padding: 10px 0;
    width: max-content;
}

.table-header div {
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
    overflow-wrap: break-word;
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

.search-error {
    width: 100%;
    text-align: center;
    padding: 20px 0;
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

.animal-ill {
    width: 150px;
    display: flex;
    justify-content: space-between;
}

.illflex {
    display: flex;
}
</style>