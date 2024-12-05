<template>
    <div class="year-chart">
        <apexchart 
        id="analytics" 
        width="700" 
        type="bar" 
        :options="options" 
        :series="series"
        ></apexchart>
    </div>
</template>

<script>
export default {
    data() {
        return {
            options: {
                chart: {
                    id: 'analytics-years',
                    stacked: true
                },
                xaxis: {
                    categories: []
                },
            },
            series: [],
        }
    },
    async created() {
        let mass_route = this.$route.path.split('/');
        let year_id = mass_route[2];
        let response = await fetch(`/api/analitics/genotyped/${year_id}/regions`);
        let result = await response.json();
        let genyear_serie = {name: 'Генотипированных', data: []};
        let allyear_serie = {name: 'Всего', data: []};
        for (let key in result) {
            this.options.xaxis.categories.push(key);
            genyear_serie.data.push(result[key].Genotyped);
            allyear_serie.data.push(result[key].Alive);
        }
        this.series.push(genyear_serie);
        this.series.push(allyear_serie);
    }
}
</script>

<style scoped>
.year-chart {
    margin-top: 120px;
}
</style>