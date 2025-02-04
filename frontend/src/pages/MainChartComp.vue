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
                    stacked: true,
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
                    text: 'Лучший селекционный индекс',
                    align: 'center',
                    style: {
                        fontSize:  '15px',
                    },
                },
                tooltip: {
                    enabled: false,
                }
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
                                categories: [
                                    'Минимальный индекс ' + currentHoz.MinIndex,
                                    'Средний индекс ' + currentHoz.AvgIndex,
                                    'Максимальный индекс ' + currentHoz.MaxIndex,
                                ],
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
                let newY = {data: []};
                let index;
                for (let i = 0; i < this.result.length; i++) {
                    if (this.result[i].Farm) {
                        this.newX.push(this.result[i].Farm.Name);
                        newY.data.push(this.result[i].MaxIndex);
                    } else {
                        index = i;
                    }
                }
                this.newX.push('Весь регион');
                newY.data.push(this.result[index].MaxIndex);
                
                this.seriesClick.push(newY);
                this.$refs.analit.updateOptions({
                    xaxis: {
                        categories: this.newX,
                    }
                });

                this.$refs.analit.updateOptions({
                    title: {
                        text: 'Лучший селекционный индекс'
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
</style>