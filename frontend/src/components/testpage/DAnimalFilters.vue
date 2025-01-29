<template>
    <div class="filters">
        <div class="filter-title">Фильтры</div>
        <div class="filter-category">
            <div @click="showFilters.hozName = !showFilters.hozName" class="ill-list-title">> Название хозяйства/фермы</div>
            <MultiselectHoz 
            @sendToMain="setIdHoz" 
            v-bind:clearHoz="clearHoz" 
            v-bind:value-from-outside="hozIdFromOutside"
            v-show="showFilters.hozName"/>
        </div>
        <div class="filter-category">
            <div @click="showFilters.birthDate = !showFilters.birthDate" class="ill-list-title">> Дата рождения</div>
            <div v-if="showFilters.birthDate">
                <label class="range">От: <input type='date' class="filter-input filter-date" v-model="filters.birthDateFrom"></label><br>
                <label class="range">До: <input type='date' class="filter-input filter-date" v-model="filters.birthDateTo"></label>
            </div>
        </div>
        <div class="filter-category">
            <div @click="showFilters.departDate = !showFilters.departDate" class="ill-list-title">> Дата выбытия</div>
            <div v-if="showFilters.departDate">
                <label class="range">От: <input type='date' class="filter-input filter-date" v-model="filters.departDateFrom"></label><br>
                <label class="range">До: <input type='date' class="filter-input filter-date" v-model="filters.departDateTo"></label>
            </div>
        </div>
        <div class="filter-category" v-if="!fromAnal">
            <div @click="showFilters.isDead = !showFilters.isDead" class="ill-list-title">> Животное мертво</div>
            <select class="filter-input" v-model="filters.isDead" v-if="showFilters.isDead">
                <option :value="null">не важно</option>
                <option :value="true">да</option>
                <option :value="false">нет</option>
            </select>
        </div>
        <div class="filter-category">
            <div @click="showFilters.breed = !showFilters.breed" class="ill-list-title">> Порода</div>
            <MultiselectBreeds 
            @sendToMain="setIdBreed" 
            v-bind:clearBreed="clearBreed"
            v-bind:value-from-outside="breedIdFromOutside"
            v-show="showFilters.breed"/>
        </div>
        <div class="filter-category">
            <div @click="showFilters.dateKRS = !showFilters.dateKRS" class="ill-list-title">> Дата внесения данных о КРС</div>
            <div v-if="showFilters.dateKRS">
                <label class="range">От: <input type='date' class="filter-input filter-date" v-model="filters.createdAtFrom"></label><br>
                <label class="range">До: <input type='date' class="filter-input filter-date" v-model="filters.createdAtTo"></label>
            </div>
        </div>
        <div class="filter-category">
            <div @click="showFilters.geneticStatus = !showFilters.geneticStatus" class="ill-list-title">> Статус генотипирования</div>
            <select class="filter-input" v-model="filters.isGenotyped" v-if="showFilters.geneticStatus">
                <option :value="null">не важно</option>
                <option :value="true">да</option>
                <option :value="false">нет</option>
            </select>
        </div>
        <div class="filter-category">
            <div @click="showFilters.milking = !showFilters.milking" class="ill-list-title">> Контрольная дойка</div>
            <div v-if="showFilters.milking">
                <label class="range">От: <input type='date' class="filter-input filter-date" v-model="filters.controlMilkingDateFrom"></label><br>
                <label class="range">До: <input type='date' class="filter-input filter-date" v-model="filters.controlMilkingDateTo"></label>
            </div>
        </div>
        <div class="filter-category">
            <div @click="showFilters.exterior = !showFilters.exterior" class="ill-list-title">> Оценка экстерьера</div>
            <!-- <input type='number' class="filter-input filter-num" v-model="filters.exterior"> -->
            <select class="filter-input" v-model="exterior" v-if="showFilters.exterior">
                <option :value="null">Не важно</option>
                <option :value="'низкая'">Низкая (50-64)</option>
                <option :value="'средняя'">Средняя (65-74)</option>
                <option :value="'хорошая'">Хорошая (75-79)</option>
                <option :value="'хорошая+'">Хорошая+ (80-84)</option>
                <option :value="'очень хорошая'">Очень хорошая (85-89)</option>
                <option :value="'отличная'">Отличная (90-100)</option>
            </select>
        </div>
        <div class="filter-category">
            <div @click="showFilters.osem = !showFilters.osem" class="ill-list-title">> Осеменение</div>
            <div v-if="showFilters.osem">
                <label class="range">От: <input type='date' class="filter-input filter-date" v-model="filters.inseminationDateFrom"></label><br>
                <label class="range">До: <input type='date' class="filter-input filter-date" v-model="filters.inseminationDateTo"></label>
            </div>
        </div>
        <div class="filter-category">
            <div @click="showFilters.otel = !showFilters.otel" class="ill-list-title">> Отел</div>
            <div v-if="showFilters.otel">
                <label class="range">От: <input type='date' class="filter-input filter-date" v-model="filters.calvingDateFrom"></label><br>
                <label class="range">До: <input type='date' class="filter-input filter-date" v-model="filters.calvingDateTo"></label>
            </div>
        </div>
        <div class="filter-category">
            <div @click="showFilters.twin = !showFilters.twin" class="ill-list-title">> Двойня</div>
            <select class="filter-input" v-model="filters.isTwins" v-if="showFilters.twin">
                <option :value="null">не важно</option>
                <option :value="true">да</option>
                <option :value="false">нет</option>
            </select>
        </div>
        <div class="filter-category">
            <div @click="showFilters.birthdead = !showFilters.birthdead" class="ill-list-title">> Мертворожденный</div>
            <select class="filter-input" v-model="filters.isStillBorn" v-if="showFilters.birthdead">
                <option :value="null">не важно</option>
                <option :value="true">да</option>
                <option :value="false">нет</option>
            </select>
        </div>
        <div class="filter-category">
            <div @click="showFilters.abort = !showFilters.abort" class="ill-list-title">> Аборт</div>
            <select class="filter-input" v-model="filters.isAborted" v-if="showFilters.abort">
                <option :value="null">не важно</option>
                <option :value="true">да</option>
                <option :value="false">нет</option>
            </select>
        </div>
        <div class="filter-category">
            <div @click="showFilters.birking = !showFilters.birking" class="ill-list-title">> Перебирковка</div>
            <div v-if="showFilters.birking">
                <label class="range">От: <input type='date' class="filter-input filter-date" v-model="filters.birkingDateFrom"></label><br>
                <label class="range">До: <input type='date' class="filter-input filter-date" v-model="filters.birkingDateTo"></label>
            </div>
        </div>
        <div class="filter-category">
            <div @click="showFilters.inbrFamily = !showFilters.inbrFamily" class="ill-list-title">> Значения коэффициента инбридинга по родословной</div>
            <div v-if="showFilters.inbrFamily">
                <label class="range">От: <input type='number' class="filter-input filter-num" v-model="filters.inbrindingCoeffByFamilyFrom"></label><br>
                <label class="range">До: <input type='number' class="filter-input filter-num" v-model="filters.inbrindingCoeffByFamilyTo"></label>
            </div>
        </div>
        <div class="filter-category">
            <div @click="showFilters.inbrGen = !showFilters.inbrGen" class="ill-list-title">> Значения коэффициента инбридинга по генотипу</div>
            <div v-if="showFilters.inbrGen">
                <label class="range">От: <input type='number' class="filter-input filter-num" v-model="filters.inbrindingCoeffByGenotypeFrom"></label><br>
                <label class="range">До: <input type='number' class="filter-input filter-num" v-model="filters.inbrindingCoeffByGenotypeTo"></label>
            </div>
        </div>
        <div class="filter-category">
            <div @click="showFilters.ill = !showFilters.ill" class="ill-list-title">> Заболевание</div>
            <div v-if="showFilters.ill">
                <label class="range">От: <input type='date' class="filter-input filter-date" v-model="filters.illDateFrom"></label><br>
                <label class="range">До: <input type='date' class="filter-input filter-date" v-model="filters.illDateTo"></label>
            </div>
        </div>
        <div class="filter-category">
            <div @click="showFilters.availIll = !showFilters.availIll" class="ill-list-title">> Наличие моногенных заболеваний</div>
            <select class="filter-input" v-model="list_ill_parameters" v-if="showFilters.availIll">
                <option :value="0">не важно</option>
                <option :value="1">есть любое</option>
                <option :value="2">отсутствуют все</option>
                <option :value="3">есть из списка</option>
                <option :value="4">отсутствует(ют) из списка</option>
            </select>
        </div>    
        <div class="filter-category category-last">
            <div @click="isVisibleIll = !isVisibleIll" class="ill-list-title">> Список заболеваний</div>
            <!-- <MultiselectIllness @sendToMain="setIdIllness" v-bind:clearIll="clearIllness" class="illness"/> -->
            <div v-if="isVisibleIll">
                <div v-for="ill in options" :key="ill.id" class="ill-item">
                    <label><input type="checkbox" :value="ill.id" v-model="filters.monogeneticIllneses"> {{ ill.name }}</label>
                </div>
            </div> 
        </div>
        <div class="filters-buttons">
            <button class="filters-apply" @click.stop="fetchFilters">Применить</button>
            <button class="clear-filters" @click.stop="clearFilters">Сбросить фильтры</button>
        </div> 
    </div>
</template>

<script>
import MultiselectBreeds from '@/components/testpage/DMultiselectBreeds.vue';
import MultiselectHoz from '@/components/testpage/DMultiselectHoz.vue';
//import MultiselectIllness from '@/components/testpage/DMultiselectIllness.vue';

export default {
    components: {
        MultiselectHoz, MultiselectBreeds, //MultiselectIllness
    },
    props: {
        fromAnal: {
            type: Boolean,
        }
    },
    data() {
        return {
            filters: {
                hozId: null,
                birthDateFrom: null,
                birthDateTo: null,
                departDateFrom: null,
                departDateTo: null,
                isDead: null,
                breedId: null,
                createdAtFrom: null,
                createdAtTo: null,
                controlMilkingDateFrom: null,
                controlMilkingDateTo: null,
                exteriorFrom: null,
                exteriorTo: null,
                inseminationDateFrom: null,
                inseminationDateTo: null,
                calvingDateFrom: null,
                calvingDateTo: null,
                isTwins: null,
                isStillBorn: null,
                isAborted: null,
                birkingDateFrom: null,
                birkingDateTo: null,
                inbrindingCoeffByFamilyFrom: null,
                inbrindingCoeffByFamilyTo: null,
                inbrindingCoeffByGenotypeFrom: null,
                inbrindingCoeffByGenotypeTo: null,
                monogeneticIllneses: [],
                isIll: null,
                hasAnyIllnes: null,
                isGenotyped: null,
                illDateFrom: null,
                illDateTo: null,
            },

            clearBreed: false,
            clearHoz: false,
            clearIllness: false,

            hozIdFromOutside: null,
            breedIdFromOutside: null,

            exterior: null,
            options: [],

            checked_ill: [],
            list_ill_parameters: 0,

            isVisibleIll: false,

            showFilters: {
                hozName: false,
                birthDate: false,
                departDate: false,
                isDead: false,
                breed: false,
                dateKRS: false,
                geneticStatus: false,
                milking: false,
                exterior: false,
                osem: false,
                otel: false,
                twin: false,
                birthdead: false,
                abort: false,
                birking: false,
                inbrFamily: false,
                inbrGen: false,
                ill: false,
                availIll: false,
            }
        }
    },
    methods: {
        fetchFilters(){
            if (this.filters.inbrindingCoeffByFamilyFrom==='') this.filters.inbrindingCoeffByFamilyFrom = null;
            if (this.filters.inbrindingCoeffByFamilyTo==='') this.filters.inbrindingCoeffByFamilyTo = null;
            if (this.filters.inbrindingCoeffByGenotypeFrom==='') this.filters.inbrindingCoeffByGenotypeFrom = null;
            if (this.filters.inbrindingCoeffByGenotypeTo==='') this.filters.inbrindingCoeffByGenotypeTo = null;
            // let send_filters = this.filters;
            let send_filters = {};
            Object.assign(send_filters, this.filters);
            this.$emit('applyFilters', send_filters);
            window.scrollTo(0,0);
        },
        clearFilters() {
            let send_filters = this.filters;
            for(let key in this.filters) {
                this.filters[key] = null;
                if(key==='monogeneticIllneses') {
                    this.filters[key] = [];
                }
            }
            this.exterior = null;
            this.list_ill_parameters = 0,
            this.clearBreed = !this.clearBreed;
            this.clearHoz = !this.clearHoz;
            this.clearIllness = !this.clearIllness;
            this.isVisibleIll = false;
            this.$emit('applyFilters', send_filters);
            window.scrollTo(0,0);
        },
        setIdHoz(hozid) {
            this.filters.hozId = hozid;
        },
        setIdBreed(breedid) {
            if(breedid) this.filters.breedId = [breedid];
            else this.filters.breedId = breedid;
            
        },
        setIdIllness(illid) {
            if (illid) this.filters.monogeneticIllneses = [illid];
            else this.filters.monogeneticIllneses = illid;
        },
    },
    watch: {
        exterior(new_val) {
            if (new_val === 'низкая') {
                this.filters.exteriorFrom = 50;
                this.filters.exteriorTo = 64;
            } else if(new_val === 'средняя') {
                this.filters.exteriorFrom = 65;
                this.filters.exteriorTo = 74;
            } else if (new_val === 'хорошая') {
                this.filters.exteriorFrom = 75;
                this.filters.exteriorTo = 79;
            } else if (new_val ==='хорошая+') {
                this.filters.exteriorFrom = 80;
                this.filters.exteriorTo = 84;
            } else if (new_val === 'очень хорошая') {
                this.filters.exteriorFrom = 85;
                this.filters.exteriorTo = 89;
            } else if (new_val === 'отличная') {
                this.filters.exteriorFrom = 90;
                this.filters.exteriorTo = 100;
            } else {
                this.filters.exteriorFrom = null;
                this.filters.exteriorTo = null;
            }
        },
        list_ill_parameters(new_val) {
            if (new_val === 0) {
                this.filters.isIll = null;
                this.filters.hasAnyIllnes = null;
            } else if (new_val === 1) {
                this.filters.isIll = null;
                this.filters.hasAnyIllnes = true;
            } else if (new_val === 2) {
                this.filters.isIll = null;
                this.filters.hasAnyIllnes = false;
            } else if (new_val === 3) {
                this.filters.isIll = true;
                this.filters.hasAnyIllnes = null;
            } else if (new_val === 4) {
                this.filters.isIll = false;
                this.filters.hasAnyIllnes = null;
            }
        }
    },
    async created() {
        this.options = [];
        const response = await fetch('/api/monogenetic_illnesses');
        const illness = await response.json();
        for (let i = 0; i < illness.length; i++) {
            let ill = {name: illness[i].Name, id: illness[i].ID};
            this.options.push(ill);
        }
        if (!this.fromAnal && Object.keys(this.$store.getters.FILTERS_2).length !== 0) {
            for (let key in this.$store.getters.FILTERS_2) {
                this.filters[key] = this.$store.getters.FILTERS_2[key];
            }
            if (this.$store.getters.FILTERS_2.hozId) {
                this.hozIdFromOutside = this.$store.getters.FILTERS_2.hozId;
            }
            if (this.$store.getters.FILTERS_2.breedId) {
                this.breedIdFromOutside = this.$store.getters.FILTERS_2.breedId[0];
            }
            this.$emit('applyFilters', this.filters);
        }
    },
}
</script>

<style scoped>
.filters {
    background-color: white;
    height: max-content;
    min-height: 500px;
    width: 350px;
    margin-left: 20px;
    box-shadow: rgba(100, 100, 111, 0.1) 0px 7px 29px 0px;
    padding: 10px 20px;
    font-family: Open Sans, sans-serif;
    display: flex;
    flex-direction: column;

}

.filter-title {
    margin-bottom: 10px;
    font-size: 140%;
    color: rgb(10, 113, 75);
}

.filter-category {
    margin-bottom: 15px;
    width: 100%;
    border-bottom: 1px solid rgb(217, 230, 223);
}

.category-last {
    border: none;
}

.filter-input {
    margin: 10px 9px;
    height: 30px;
    width: 200px;
    padding: 0 10px;
    font-size: 14px;
    box-sizing: border-box;
    outline: none;
    border: 3px solid rgb(203, 227, 219);
    border-radius: 10px;
    transition: 0.3s;
}

.filter-input:hover {
    border: 3px solid rgb(112, 189, 151);
}

.filter-date {
    width: 150px;
}

.range {
    color: rgb(114, 99, 145);
}

.filter-num {
    width: 90px;
}

.filters-apply, .clear-filters {
    width: 200px;
    color: white;
    border-radius: 10px;
    transition: 0.3s;
    cursor: pointer;
    align-self: center;
}

.filters-apply {
    font-size: 100%;
    margin: 15px 20px;
    padding: 7px 0;
    background-color: rgb(82, 183, 129);
    border: 1px solid rgb(82, 183, 129);
}

.clear-filters {
    background: none;
    border: 1px solid white;
    margin: 5px 20px 20px 20px;
    padding: 7px 0;
    color: rgb(23, 138, 86);
}

.filters-apply:hover {
    background-color: white;
    color: rgb(23, 138, 86);
}

.clear-filters:hover {
    background-color: rgb(228, 245, 237);
}

.check {
    padding: 0 7px 7px 0;
}

.filters-buttons {
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
    position: sticky;
    bottom: 0;
    background-color: white;
    border-top: 1px solid rgb(217, 230, 225);
}

.ill-item {
    padding-top: 7px;
}

.ill-list-title {
    cursor: pointer;
    transition: 0.3s;
}

.ill-list-title:hover {
    color: grey;
}
</style>