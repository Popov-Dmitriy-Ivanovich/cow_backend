<template>
    <div class="year-chart">
        <apexchart 
        id="analytics" 
        width="850" 
        type="bar" 
        :options="options" 
        :series="series"
        @dataPointSelection="clickHandler"
        ></apexchart>
    </div>
</template>

<script>
export default {
    data() {
        return {
            options: {
                chart: {
                    id: 'analytics-region',
                    stacked: true,
                    
                },
                xaxis: {
                    categories: []
                },
                colors: ['#6e5add','#75a2e7']
            },
            series: [],
            common_info: {},
        }
    },
    async created() {
        let mass_route = this.$route.path.split('/');
        let year_id = mass_route[2];
        let region_id = mass_route[3];
        let response = await fetch(`/api/analitics/genotyped/${year_id}/byRegion/${region_id}/districts`);
        let result = await response.json();
        this.common_info = result;
        let genyear_serie = {name: 'Генотипированных', data: []};
        let allyear_serie = {name: 'Всего', data: []};
        for (let key in result) {
            this.options.xaxis.categories.push(key);
            genyear_serie.data.push(result[key].Genotyped);
            allyear_serie.data.push(result[key].Alive);
        }
        this.series.push(allyear_serie);
        this.series.push(genyear_serie);
    },
    methods: {
        clickHandler(event, chartContext, config) {
            let nameReg = this.options.xaxis.categories[config.dataPointIndex];
            let reg_id = this.$route.params.region;
            let dist_id = this.common_info[nameReg].DistrictID;
            let year = this.$route.params.id;
            this.$router.push(`/analytics/${year}/${reg_id}/${dist_id}`);
        }
    }
}
</script>

<style scoped>
.year-chart {
    margin-top: 10px;
}
</style>