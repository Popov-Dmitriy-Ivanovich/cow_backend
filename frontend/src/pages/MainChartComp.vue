<template>
    <div>
        <!-- <div class="analytics-title">Статистика для сравнительного анализа хозяйств и регионов</div> -->
        <MainChart/>
    </div>
</template>

<script>
import MainChart from '@/components/analyticsComponents/MainChart.vue';

export default {
    components: {
        MainChart
    },
}
</script>

<style scoped>

</style>