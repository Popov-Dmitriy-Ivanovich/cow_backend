<template>
    <div class="filters">
        <div class="filter-title">Фильтры</div>
        <div class="filter-category">
            <div>Название хозяйства/фермы</div>
            <MultiselectHoz @sendToMain="setIdHoz" v-bind:clearHoz="clearHoz"/>
        </div>
        <div class="filter-category">
            <div>Дата рождения</div>
            <label class="range">От: <input type='date' class="filter-input filter-date" v-model="filters.birthDateFrom"></label><br>
            <label class="range">До: <input type='date' class="filter-input filter-date" v-model="filters.birthDateTo"></label>
        </div>
        <div class="filter-category">
            <div>Дата выбытия</div>
            <label class="range">От: <input type='date' class="filter-input filter-date" v-model="filters.departDateFrom"></label><br>
            <label class="range">До: <input type='date' class="filter-input filter-date" v-model="filters.departDateTo"></label>
        </div>
        <div class="filter-category">
            <div>Животное мертво</div>
            <select class="filter-input" v-model="filters.isDead">
                <option :value="null">не важно</option>
                <option :value="true">да</option>
                <option :value="false">нет</option>
            </select>
        </div>
        <div class="filter-category">
            <div>Порода</div>
            <MultiselectBreeds @sendToMain="setIdBreed" v-bind:clearBreed="clearBreed"/>
        </div>
        <div class="filter-category">
            <div>Генотипирование</div>
            <label class="range">От: <input type='date' class="filter-input filter-date" v-model="filters.genotypingDateFrom"></label><br>
            <label class="range">До: <input type='date' class="filter-input filter-date" v-model="filters.genotypingDateTo"></label>
        </div>
        <div class="filter-category">
            <div>Контрольная дойка</div>
            <label class="range">От: <input type='date' class="filter-input filter-date" v-model="filters.controlMilkingDateFrom"></label><br>
            <label class="range">До: <input type='date' class="filter-input filter-date" v-model="filters.controlMilkingDateTo"></label>
        </div>
        <div class="filter-category">
            <div>Оценка экстерьера</div>
            <input type='number' class="filter-input filter-num" v-model="filters.exterior">
        </div>
        <div class="filter-category">
            <div>Осеменение</div>
            <label class="range">От: <input type='date' class="filter-input filter-date" v-model="filters.inseminationDateFrom"></label><br>
            <label class="range">До: <input type='date' class="filter-input filter-date" v-model="filters.inseminationDateTo"></label>
        </div>
        <div class="filter-category">
            <div>Отел</div>
            <label class="range">От: <input type='date' class="filter-input filter-date" v-model="filters.calvingDateFrom"></label><br>
            <label class="range">До: <input type='date' class="filter-input filter-date" v-model="filters.calvingDateTo"></label>
            <div>Двойня</div>
            <select class="filter-input" v-model="filters.isTwins">
                <option :value="null">не важно</option>
                <option :value="true">да</option>
                <option :value="false">нет</option>
            </select><br>
            <div>Мертворожденный</div>
            <select class="filter-input" v-model="filters.isStillBorn">
                <option :value="null">не важно</option>
                <option :value="true">да</option>
                <option :value="false">нет</option>
            </select><br>
            <div>Аборт</div>
            <select class="filter-input" v-model="filters.isAborted">
                <option :value="null">не важно</option>
                <option :value="true">да</option>
                <option :value="false">нет</option>
            </select><br>
        </div>
        <div class="filter-category">
            <div>Перебирковка</div>
            <label class="range">От: <input type='date' class="filter-input filter-date" v-model="filters.birkingDateFrom"></label><br>
            <label class="range">До: <input type='date' class="filter-input filter-date" v-model="filters.birkingDateTo"></label>
        </div>
        <div class="filter-category">
            <div>Значения коэффициента инбридинга по родословной</div>
            <label class="range">От: <input type='number' class="filter-input filter-num" v-model="filters.inbrindingCoeffByFamilyFrom"></label><br>
            <label class="range">До: <input type='number' class="filter-input filter-num" v-model="filters.inbrindingCoeffByFamilyTo"></label>
            <div>Значения коэффициента инбридинга по генотипу</div>
            <label class="range">От: <input type='number' class="filter-input filter-num" v-model="filters.inbrindingCoeffByFenotypeFrom"></label><br>
            <label class="range">До: <input type='number' class="filter-input filter-num" v-model="filters.inbrindingCoeffByFenotypeTo"></label>
        </div>
        <div class="filter-category category-last">
            <div>Наличие моногенных заболеваний</div>
            <select class="filter-input" v-model="filters.isIll">
                <option :value="null">не важно</option>
                <option :value="true">есть</option>
                <option :value="false">отсутствует</option>
            </select><br>
            <div>Заболевание</div>
            <MultiselectIllness @sendToMain="setIdIllness" v-bind:clearIll="clearIllness" class="illness"/>
        </div>
        <div class="filters-buttons">
            <button class="filters-apply" @click="fetchFilters">Применить</button>
            <button class="clear-filters" @click="clearFilters">Сбросить фильтры</button>
        </div>  
    </div>
</template>

<script>
import MultiselectBreeds from '@/components/testpage/DMultiselectBreeds.vue';
import MultiselectHoz from '@/components/testpage/DMultiselectHoz.vue';
import MultiselectIllness from '@/components/testpage/DMultiselectIllness.vue';

export default {
    components: {
        MultiselectHoz, MultiselectBreeds, MultiselectIllness
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
                genotypingDateFrom: null,
                genotypingDateTo: null,
                controlMilkingDateFrom: null,
                controlMilkingDateTo: null,
                exterior: null,
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
                inbrindingCoeffByFenotypeFrom: null,
                inbrindingCoeffByFenotypeTo: null,
                monogeneticIllneses: null,
                isIll: null,
            },

            clearBreed: false,
            clearHoz: false,
            clearIllness: false,
        }
    },
    methods: {
        fetchFilters(){
            let send_filters = this.filters;
            console.log(send_filters, 'filers');
            console.log(JSON.stringify(send_filters))
            this.$emit('applyFilters', send_filters);
            window.scrollTo(0,0);
        },
        clearFilters() {
            for(let key in this.filters) {
                this.filters[key] = null;
            }
            this.clearBreed = !this.clearBreed;
            this.clearHoz = !this.clearHoz;
            this.clearIllness = !this.clearIllness;
            console.log(this.filters, 'сбросить');
            this.$emit('applyFilters', this.filters);
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
    }
}
</script>

<style scoped>
.filters {
    background-color: white;
    height: 500px;
    min-height: max-content;
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
    color: rgb(37, 0, 132);
}

.filter-category {
    margin-bottom: 15px;
    width: 100%;
    border-bottom: 1px solid rgb(218, 217, 230);
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
    border: 3px solid rgb(195, 200, 212);
    border-radius: 10px;
    transition: 0.3s;
}

.filter-input:hover {
    border: 3px solid rgb(101, 102, 170);
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
    background-color: rgb(101, 82, 183);
    border: 1px solid rgb(101, 82, 183);
}

.clear-filters {
    background: none;
    border: 1px solid white;
    margin: 5px 20px 20px 20px;
    padding: 7px 0;
    color: rgb(101, 82, 183);
}

.filters-apply:hover {
    background-color: white;
    color: rgb(101, 82, 183);
}

.clear-filters:hover {
    background-color: rgb(231, 228, 245);
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
    border-top: 1px solid rgb(218, 217, 230);
}

/* .illness {
    position: relative;
    z-index: 300;
} */
</style>