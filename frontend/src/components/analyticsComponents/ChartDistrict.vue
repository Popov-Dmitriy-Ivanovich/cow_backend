<template>
    <div class="year-chart">
        <apexchart v-if="!err"
        id="analytics" 
        width="850" 
        type="bar" 
        :options="options" 
        :series="series"
        ref="chartdistr"
        ></apexchart>
        <div v-if="err">Недостаточный уровень доступа</div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            options: {
                chart: {
                    id: 'analytics-district',
                    stacked: true,
                    
                },
                xaxis: {
                    categories: []
                },
                colors: ['#6e5add','#75a2e7','#63d9cb'],
                title: {
                    text: 'Хозяйства',
                    align: 'center',
                    style: {
                        fontSize:  '24px',
                    },
                }
            },
            series: [],
            common_info: {},
            newX: [],
            err: false,
        }
    },
    async created() {
        if (this.changeOpt === '' || this.changeOpt == 'ill') {
            await this.fetchData();
        } else if(this.changeOpt == 'lact') {
            await this.fetchDataLact();
        }
    },
    methods: {
        async fetchData() {
            this.err = false;
            this.series = [];
            this.options.xaxis.categories = []; 
            let mass_route = this.$route.path.split('/');
            let year_id = mass_route[2];
            let district_id = mass_route[4];
            let response = await fetch(`/api/analitics/genotyped/${year_id}/byDistrict/${district_id}/hoz`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8',
                    'Authorization': this.getJwt(),
                },
                body: JSON.stringify(this.changeFilters),
            });
            let result = await response.json();
            if(result.error) {
                this.err = true;
                return 0;
            }
            this.common_info = result;
            let genyear_serie = {name: 'Генотипированных', data: []};
            let allyear_serie = {name: 'Всего', data: []};
            let illyear_serie = {name: 'Больных', data: []};
            this.newX = [];
            for (let key in result) {
                this.newX.push(key);
                genyear_serie.data.push(result[key].Genotyped);
                allyear_serie.data.push(result[key].Alive);
                if (this.changeOpt == 'ill') {
                    illyear_serie.data.push(result[key].Ill);
                }
            }
            this.series.push(allyear_serie);
            this.series.push(genyear_serie);
            if (this.changeOpt == 'ill') {
                this.series.push(illyear_serie);
            }

            this.$refs.chartdistr.updateOptions({
                xaxis: {
                    categories: this.newX,
                }
            });

            await this.getTitle();
        },
        async fetchDataLact() {
            this.err = false;
            this.series = [];
            this.options.xaxis.categories = []; 
            let mass_route = this.$route.path.split('/');
            let year_id = mass_route[2];
            let district_id = mass_route[4];
            let response = await fetch(`/api/analitics/checkMilks/${year_id}/byDistrict/${district_id}/byHoz`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8',
                    'Authorization': this.getJwt(),
                },
                body: JSON.stringify(this.changeFilters),
            });
            let result = await response.json();
            if(result.error) {
                this.err = true;
                return 0;
            }
            this.common_info = result;
            let milk_serie = {name: 'Удой', data: []};
            let fat_serie = {name: 'Жир', data: []};
            let protein_serie = {name: 'Белок', data: []};
            this.newX = [];
            for (let key in result) {
                this.newX.push(key);
                milk_serie.data.push(Math.round(result[key].Milk*100)/100);
                fat_serie.data.push(Math.round(result[key].Fat*100)/100);
                protein_serie.data.push(Math.round(result[key].Protein*100)/100);
            }

            this.series.push(milk_serie);
            this.series.push(fat_serie);
            this.series.push(protein_serie);

            this.$refs.chartdistr.updateOptions({
                xaxis: {
                    categories: this.newX,
                }
            });

            await this.getTitle();
        },
        getJwt() {
            let arr = document.cookie.split(';');
            for (let i = 0; i < arr.length; i++) {
                if (arr[i].split('=')[0] == 'jwt') {
                    return arr[i].split('=')[1];
                }
            }
            return null;
        },
        async getTitle() {
            let mass_route = this.$route.path.split('/');
            let year_id = mass_route[2];
            let region_id = mass_route[3];
            let district_id = mass_route[4];

            let chartTitle;

            if (this.changeOpt === '') {
                chartTitle = 'Генотипирование за ';
            } else if (this.changeOpt == 'ill') {
                chartTitle = 'Моногенные заболевания за ';
            } else if (this.changeOpt == 'lact') {
                chartTitle = 'Срадние показатели удоя за ';
            }

            if (year_id == 40000) {
                chartTitle += 'все года. ';
            } else {
                chartTitle += `${year_id} год. `
            }

            let res_region = await fetch(`/api/regions/${region_id}`);
            let region = await res_region.json();
            if (region) {
                chartTitle += region.Name;
            }
            chartTitle += ', ';

            let res_districts = await fetch(`/api/districts`);
            let districts = await res_districts.json();
            for (let i = 0; i < districts.length; i++) {
                if (districts[i].ID == district_id) {
                    chartTitle = chartTitle + districts[i].Name + ' район';
                }
            }

            this.$refs.chartdistr.updateOptions({
                xaxis: {
                    title: {
                        text: chartTitle,
                    }
                }
            })
        },
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

<style scoped>
.year-chart {
    margin-top: 10px;
}
</style>