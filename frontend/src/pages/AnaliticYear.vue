<template>
    <div>
        <div class="prev-chart" @click="toPrev">🠔 Назад</div>
        <!-- <ChartYear/> -->
        <div class="row">
            <apexchart id="analit_click" width="500" type="bar" :options="optionsClick" :series="seriesClick" ref="analit" @dataPointSelection="clickHandler"></apexchart>
            <apexchart id="analit" width="500" type="bar" :options="options" :series="series"></apexchart>
        </div>
        <div class="row">
            <apexchart id="analit1" width="500" type="bar" :options="options1" :series="series1"></apexchart>
            <apexchart id="analit2" width="500" type="bar" :options="options2" :series="series2"></apexchart>
        </div>
        <div class="row">
            <apexchart id="analit3" width="500" type="bar" :options="options3" :series="series3"></apexchart>
            <apexchart id="analit4" width="500" type="bar" :options="options4" :series="series4"></apexchart>
        </div>
    </div>
</template>

<script>
// import ChartYear from '@/components/analyticsComponents/ChartYear.vue';

export default {
    components: {
        // ChartYear
    },
    methods: {
        toPrev() {
            this.$router.push(`/analytics`);
        },
    },
}
</script>

<style scoped>
.prev-chart {
    font-family: Open Sans, sans-serif;
    margin-top: 30px;
    margin-left: 20px;
    color:rgb(10, 113, 75);
    padding-bottom: 20px;
    cursor: pointer;
    transition: 0.3s;
}

.prev-chart:hover {
    color: rgb(63, 205, 132);
    padding-left: 10px;
    width: max-content;
}

</style>