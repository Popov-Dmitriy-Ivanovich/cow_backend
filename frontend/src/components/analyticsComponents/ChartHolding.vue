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
                    id: 'analytics-holding',
                    stacked: true,
                    
                },
                xaxis: {
                    categories: []
                },
                colors: ['#6e5add','#75a2e7']
            },
            series: [],
        }
    },
    async created() {
        let mass_route = this.$route.path.split('/');
        let year_id = mass_route[2];
        let hold_id = mass_route[5];
        let response = await fetch(`/api/analitics/genotyped/${year_id}/byHold/${hold_id}/hoz`);
        let result = await response.json();
        let genyear_serie = {name: 'Генотипированных', data: []};
        let allyear_serie = {name: 'Всего', data: []};
        for (let key in result) {
            this.options.xaxis.categories.push(key);
            genyear_serie.data.push(result[key].Genotyped);
            allyear_serie.data.push(result[key].Alive);
        }
        this.series.push(allyear_serie);
        this.series.push(genyear_serie);
        
    }
}
</script>

<style scoped>
.year-chart {
    margin-top: 10px;
}
</style>