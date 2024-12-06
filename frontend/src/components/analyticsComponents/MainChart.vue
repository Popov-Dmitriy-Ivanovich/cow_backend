<template>
    <div class="analytics-chart">
        <apexchart 
        id="analytics" 
        width="700" 
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
                    id: 'analytics'
                },
                xaxis: {
                    categories: []
                },
                yaxis: {
                    axisBorder: {
                        show: false
                    },
                    axisTicks: {
                        show: false,
                    },
                    labels: {
                        show: false,
                    },
                },
                dataLabels: {
                    enabled: false
                },
                colors: ['#6e5add']
            },
            series: [],
        }
    },
    async created() {
        let response = await fetch(`/api/analitics/genotyped/years`);
        let result = await response.json();
        let obj = { data: []};
        for (let key in result) {
            if(key == -1) {
                this.options.xaxis.categories.push('Все года');
                obj.data.push(1);
            } else {
                this.options.xaxis.categories.push(key);
                obj.data.push(1);
            }
        }
        this.series.push(obj);
    },
    methods: {
        clickHandler(event, chartContext, config){
            let year = this.options.xaxis.categories[config.dataPointIndex];
            if(year == 'Все года') year = 40000;
            else year = Number(year);
            this.$router.push(`/analytics/${year}`)
        },
    }
}
</script>

<style scope>
.analytics-chart {
    margin-top: 120px;
}
</style>