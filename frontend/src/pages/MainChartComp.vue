<template>
    <div>
        <!-- <div class="analytics-title">Статистика для сравнительного анализа хозяйств и регионов</div> -->
        <!-- <MainChart/> -->
        <div class="prev-chart" @click="toPrev" v-if="clickToPrev">🠔 Назад</div>
        <apexchart 
        id="analit_click" 
        width="1000" 
        type="bar" 
        :options="optionsClick" 
        :series="seriesClick" 
        ref="analit"
        @dataPointSelection="clickHandler"
        ></apexchart>
        <div v-if="!clickToPrev" class="description">
            График показывает разницу между индексным значением лучшей коровы (лучшая из лучших) и худшей коровы (худшая из худших) для данного хозяйства. Аналогично, выводится разница для коровы, замыкающей первую пятерку (худшей из лучших), и коровы, возглавляющей последнюю пятерку (лучшая из худших).
            Чем ниже получившаяся разница - тем более однородное стадо
        </div>
        <div v-else class="description">
            Вертикальная ось - количество голов КРС<br>
            Горизонтальная ось - селекционный индекс разбитый на 3 диапазона по 33% (худшие, средние, лучшие)
        </div>
    </div>
</template>

<script>
// import MainChart from '@/components/analyticsComponents/MainChart.vue';

export default {
    components: {
        // MainChart
    },
    data() {
        return {
            optionsClick: {
                chart: {
                    id: 'analit_click',
                    stacked: false,
                    zoom: {
                        enabled: false,
                    },
                },
                xaxis: {
                    categories: [],
                    dataLabels: {
                        enabled: true,
                        style: {
                            fontSize: '10px',
                        }
                    },
                    labels: {
                        style: {
                            fontSize: '10px',
                        },
                        hideOverlappingLabels: true,
                        trim: true,
                    }
                },

                colors: ['#78DABC','#6e5add','#75a2e7'],
                title: {
                    text: 'Разница индексных значений между 5 лучшими и 5 худшими коровами',
                    align: 'center',
                    style: {
                        fontSize:  '15px',
                    },
                },
                // tooltip: {
                //     enabled: false,
                // }
            },
            seriesClick: [],
            
            newX: [],
            result: [],
            clickToPrev: false,
        }
    },
    async mounted() {
        this.seriesClick = [];
        let response = await fetch('/api/analitics/total/23/regionalStatistics/');
        this.result = await response.json();
        this.chooseChart();
    },
    methods: {
        toPrev() {
            this.$router.back();
        },
        async clickHandler(event, chartContext, config) {
            let hoz = {};
            for (let i = 0; i < this.result.length; i++) {
                if(this.result[i].Farm) {
                    if(this.result[i].Farm.Name == this.newX[config.dataPointIndex]) {
                        hoz = this.result[i];
                    }
                } else {
                    if (this.newX[config.dataPointIndex] == 'Весь регион') {
                        hoz = this.result[i];
                    }
                }
            }
            if (this.$route.query.hoz) {
                let currAnim = [];
                if (config.dataPointIndex === 0) currAnim = hoz.MinCowIds;
                else if (config.dataPointIndex === 1) currAnim = hoz.AvgCowIds;
                else if (config.dataPointIndex == 2) currAnim = hoz.MaxCowIds;
                this.$store.commit('SET_CURRENTANIMALS', currAnim);
                this.$router.push('/animals');
            } else {
                await this.$router.push({query: {hoz: hoz.ID.toString()}});
            }
        },
        chooseChart() {
            let q = this.$route.query;
            if (q.hoz) {
                for (let i = 0; i < this.result.length; i++) {
                    if (this.result[i].ID == q.hoz) {
                        this.clickToPrev = true;
                        this.seriesClick = [{data: []}];
                        let currentHoz = this.result[i];
                        this.seriesClick[0].data.push(
                            currentHoz.MinCount, currentHoz.AvgCount, currentHoz.MaxCount
                        );
                        
                        this.$refs.analit.updateOptions({
                            xaxis: {
                                labels: {
                                    show: false,
                                }
                            }
                        });

                        let newtitle;
                        if (currentHoz.Farm) {
                            newtitle = ['Число голов в ', currentHoz.Farm.Name];
                        } else {
                            newtitle = 'Число голов во всем регионе';
                        }
                        this.$refs.analit.updateOptions({
                            title: {
                                text: newtitle
                            }
                        });
                        return
                    }
                }
            } else {
                this.clickToPrev = false;
                this.seriesClick = [];
                this.newX = []; 
                let newY = {name: 'Разница между лучшей из лучших и худшей из худших', data: []};
                // let index;
                for (let i = 0; i < this.result.length; i++) {
                    if (this.result[i].Farm) {
                        this.newX.push(this.result[i].Farm.Name);
                        newY.data.push(this.result[i].MaxIndex);
                    } else {
                        // index = i;
                    }
                }
                // this.newX.push('Весь регион');
                // newY.data.push(this.result[index].MaxIndex);
                this.seriesClick.push(newY);
                this.seriesClick.push( {name: 'Разница между худшей из лучших и лучшей из худших',data: [245.57, 273.36, 292.52, 307.85, 375.19, 406.50, 418.41]});

                this.$refs.analit.updateOptions({
                    xaxis: {
                        categories: this.newX,
                    }
                });

                this.$refs.analit.updateOptions({
                    title: {
                        text: 'Разница индексных значений между 5 лучшими и 5 худшими коровами'
                    }
                });
            }
        },
        clickToAnimals(event, chartContext, config) {
            console.log(config);
        }
    },
    watch: {
        $route() {
            this.chooseChart();
        }
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
    justify-self: flex-start;
    align-self: flex-start;
}

.prev-chart:hover {
    color: rgb(63, 205, 120);
    padding-left: 10px;
    width: max-content;
}

.description {
    font-family: Open Sans, sans-serif;
    margin-top: 10px;
    text-align: center;
    color: rgb(90, 90, 90);
}
</style>