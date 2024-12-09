<template>
    <div class="analytics-flex">
        <DAnimalFilters class="analytics-filters"/>
        <div>
            <div class="analytics-title">Статистика для сравнительного анализа хозяйств и регионов</div>
            <MainChart/>
        </div>

    </div>
</template>

<script>
import MainChart from '@/components/analyticsComponents/MainChart.vue';
import DAnimalFilters from '@/components/testpage/DAnimalFilters.vue';

export default {
    components: {
        MainChart, DAnimalFilters
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
</style>