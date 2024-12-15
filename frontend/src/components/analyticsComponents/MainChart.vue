<template>
    <div class="analytics-chart">
        <apexchart 
        id="analytics" 
        width="850px" 
        type="bar" 
        :options="options" 
        :series="series"
        @dataPointSelection="clickHandler"
        ref="mainchart"
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
                colors: ['#6e5add'],
                tooltip: {
                    enabled: false,
                }
            },
            series: [],
            newX: [],
        }
    },
    async mounted() {
        if (this.changeOpt === '' || this.changeOpt == 'ill') {
            await this.fetchData();
        } else if(this.changeOpt == 'lact') {
            await this.fetchDataLact();
        }
    },
    methods: {
        clickHandler(event, chartContext, config){
            console.log(config);
            console.log(this.options.xaxis.categories);
            let year = this.newX[config.dataPointIndex];
            if(year == 'Все года') year = 40000;
            else year = Number(year);
            this.$router.push({path:`/analytics/${year}`});
        },
        async fetchData() {
            this.series = [];
            this.options.xaxis.categories = []; 
            console.log(localStorage.getItem('jwt'));
            let response = await fetch(`/api/analitics/genotyped/years`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8',
                    'Authorization': localStorage.getItem('jwt'),
                },
                body: JSON.stringify(this.changeFilters),
            });
            let result = await response.json();
            let obj = { data: []};
            this.newX = [];
            for (let key in result) {
                this.newX.push(key);
                obj.data.push(1);
            }

            this.newX.push('Все года');
            obj.data.push(1);

            this.$refs.mainchart.updateOptions({
                xaxis: {
                    categories: this.newX,
                }
            });
            console.log(this.options.xaxis.categories);
            this.series.push(obj);
            console.log('функция 1');
        },
        async fetchDataLact() {
            this.series = [];
            this.options.xaxis.categories = []; 
            let response = await fetch(`/api/analitics/checkMilks/years`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8',
                    'Authorization': localStorage.getItem('jwt'),
                },
                body: JSON.stringify(this.changeFilters),
            });
            let result = await response.json();
            let obj = { data: []};
            this.newX = [];
            for (let key in result) {
                this.newX.push(key);
                obj.data.push(1);
            }

            this.$refs.mainchart.updateOptions({
                xaxis: {
                    categories: this.newX,
                }
            });

            this.series.push(obj);
            console.log('функция 2');
        }
    },
    watch: {
        async changeFilters() {
            if (this.changeOpt === '' || this.changeOpt == 'ill') {
                await this.fetchData();
            } else if(this.changeOpt == 'lact') {
                await this.fetchDataLact();
            }
        },
        async changeOpt() {
            if (this.changeOpt === '' || this.changeOpt == 'ill') {
                await this.fetchData();
            } else if(this.changeOpt == 'lact') {
                await this.fetchDataLact();
            }
        }
    },
    computed: {
        changeFilters(){
            let a = this.$store.state.filters;
            return a;
        },
        changeOpt(){
            let a = this.$store.state.option;
            return a;
        }
    }
}
</script>

<style scope>
.analytics-chart {
    margin-top: 30px;
}
</style>