<template>
    <div class="analytics-flex">
        <DAnimalFilters class="analytics-filters" @applyFilters="fetchAnalyticsFilters"/>
        <div class="chart-block">
            <div class="analytics-title">Статистика для сравнительного анализа хозяйств и регионов</div>
            <select v-model="opt" class="filter-input">
                <option :value="''">генотипирование</option>
                <option :value="'lact'">средние показатели удоя</option>
                <option :value="'ill'">моногенные заболевания</option>
            </select>
            <router-view></router-view>
        </div>

    </div>
</template>

<script>
import DAnimalFilters from '@/components/testpage/DAnimalFilters.vue';

export default {
    components: {
        DAnimalFilters
    },
    data() {
        return{
            opt: '',
        }
    },
    methods: {
        async fetchAnalyticsFilters(filters){
            this.$store.commit('SET_FILTERS', filters);
            console.log(this.$store.getters.FILTERS, 'in store');
        }
    },
    watch: {
        opt(new_val) {
            this.$store.commit('SET_OPTION', new_val);
        }
    }
}
</script>

<style scoped>
.analytics-flex {
    display: flex;
    justify-content: space-around;
    width: 90vw;
}

.analytics-filters {
    margin-top: 130px;
}

.analytics-title {
    font-size: 200%;
    font-family: Open Sans, sans-serif;
    color: rgb(37, 0, 132);
    margin-top: 120px;
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
</style>