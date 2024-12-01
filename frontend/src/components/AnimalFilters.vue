<template>
<div class="filters">
    <div class="filter-title">Фильтры</div>
    <div class="filter-category">
        <div>Название хозяйства/фермы</div>
        <MultiselectHoz @sendToMain="setIdHoz" v-bind:clearHoz="clearHoz"/>
    </div>
    <div class="filter-category">
        <div>Дата рождения</div>
        <label class="range">От: <input type='date' class="filter-input filter-date" v-model="filters.D_BIRTH_BEGIN"></label><br>
        <label class="range">До: <input type='date' class="filter-input filter-date" v-model="filters.D_BIRTH_END"></label>
    </div>
    <div class="filter-category">
        <div>Дата выбытия</div>
        <label class="range">От: <input type='date' class="filter-input filter-date" v-model="filters.D_OUT_BEGIN"></label><br>
        <label class="range">До: <input type='date' class="filter-input filter-date" v-model="filters.D_OUT_END"></label>
    </div>
    <div class="filter-category">
        <div>Животное мертво</div>
        <select class="filter-input" v-model="filters.IS_DEATH">
            <option :value="''">не важно</option>
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
        <label class="range">От: <input type='date' class="filter-input filter-date" v-model="filters.D_GEN_BEGIN"></label><br>
        <label class="range">До: <input type='date' class="filter-input filter-date" v-model="filters.D_GEN_END"></label>
    </div>
    <div class="filter-category">
        <div>Контрольная дойка</div>
        <label class="range">От: <input type='date' class="filter-input filter-date" v-model="filters.D_MILKING_BEGIN"></label><br>
        <label class="range">До: <input type='date' class="filter-input filter-date" v-model="filters.D_MILKING_END"></label>
    </div>
    <div class="filter-category">
        <div>Оценка экстерьера</div>
        <input type='number' class="filter-input filter-num" v-model="filters.EXTERIOR">
    </div>
    <div class="filter-category">
        <div>Осеменение</div>
        <label class="range">От: <input type='date' class="filter-input filter-date" v-model="filters.D_OSEM_BEGIN"></label><br>
        <label class="range">До: <input type='date' class="filter-input filter-date" v-model="filters.D_OSEM_END"></label>
    </div>
    <div class="filter-category">
        <div>Отел</div>
        <label class="range">От: <input type='date' class="filter-input filter-date" v-model="filters.D_OTEL_BEGIN"></label><br>
        <label class="range">До: <input type='date' class="filter-input filter-date" v-model="filters.D_OTEL_END"></label>
        <div>Двойня</div>
        <select class="filter-input" v-model="filters.IS_TWINS">
            <option :value="''">не важно</option>
            <option :value="true">да</option>
            <option :value="false">нет</option>
        </select><br>
        <div>Мертворожденный</div>
        <select class="filter-input" v-model="filters.IS_STILLBORN">
            <option :value="''">не важно</option>
            <option :value="true">да</option>
            <option :value="false">нет</option>
        </select><br>
        <div>Аборт</div>
        <select class="filter-input" v-model="filters.IS_ABORTED">
            <option :value="''">не важно</option>
            <option :value="true">да</option>
            <option :value="false">нет</option>
        </select><br>
    </div>
    <div class="filter-category">
        <div>Перебирковка</div>
        <label class="range">От: <input type='date' class="filter-input filter-date" v-model="filters.D_BIRKING_BEGIN"></label><br>
        <label class="range">До: <input type='date' class="filter-input filter-date" v-model="filters.D_BIRKING_END"></label>
    </div>
    <div class="filter-category">
        <div>Значения коэффициента инбридинга по родословной</div>
        <label class="range">От: <input type='number' class="filter-input filter-num" v-model="filters.K_INBR_ROD_BEGIN"></label><br>
        <label class="range">До: <input type='number' class="filter-input filter-num" v-model="filters.K_INBR_ROD_END"></label>
        <div>Значения коэффициента инбридинга по генотипу</div>
        <label class="range">От: <input type='number' class="filter-input filter-num" v-model="filters.K_INBR_FEN_BEGIN"></label><br>
        <label class="range">До: <input type='number' class="filter-input filter-num" v-model="filters.K_INBR_FEN_END"></label>
    </div>
    <div class="filter-category category-last">
        <div>Наличие моногенных заболеваний</div>
        <select class="filter-input" v-model="filters.IS_ILL">
            <option :value="''">не важно</option>
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
import MultiselectBreeds from './MultiselectBreeds.vue';
import MultiselectHoz from './MultiselectHoz.vue';
import MultiselectIllness from './MultiselectIllness.vue';

export default {
    components: {
        MultiselectHoz, MultiselectBreeds, MultiselectIllness
    },
    data() {
        return {
            filters: {
                ID_HOZ: '',
                D_BIRTH_BEGIN: '',
                D_BIRTH_END: '',
                D_OUT_BEGIN: '',
                D_OUT_END: '',
                IS_DEATH: '',
                ID_BREED: '',
                D_GEN_BEGIN: '',
                D_GEN_END: '',
                D_MILKING_BEGIN: '',
                D_MILKING_END: '',
                EXTERIOR: '',
                D_OSEM_BEGIN: '',
                D_OSEM_END: '',
                D_OTEL_BEGIN: '',
                D_OTEL_END: '',
                IS_TWINS: '',
                IS_STILLBORN: '',
                IS_ABORTED: '',
                D_BIRKING_BEGIN: '',
                D_BIRKING_END: '',
                K_INBR_ROD_BEGIN: '',
                K_INBR_ROD_END: '',
                K_INBR_GEN_BEGIN: '',
                K_INBR_GEN_END: '',
                MONOG_ILLNESS: '',
                IS_ILL: '',
            },

            clearBreed: false,
            clearHoz: false,
            clearIllness: false,
        }
    },
    methods: {
        fetchFilters(){
            let send_filters = this.filters;
            this.$emit('applyFilters', send_filters);
            window.scrollTo(0,0);
        },
        clearFilters() {
            for(let key in this.filters) {
                this.filters[key] = '';
            }
            this.clearBreed = !this.clearBreed;
            this.clearHoz = !this.clearHoz;
            this.clearIllness = !this.clearIllness;
            this.$emit('applyFilters', this.filters);
            window.scrollTo(0,0);
        },
        setIdHoz(hozid) {
            this.filters.ID_HOZ = hozid;
        },
        setIdBreed(breedid) {
            this.filters.ID_BREED = breedid;
            
        },
        setIdIllness(illid) {
            this.filters.MONOG_ILLNESS = illid;
        }
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