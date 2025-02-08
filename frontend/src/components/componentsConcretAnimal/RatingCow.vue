<template>
    <div class="rating-columns" v-if="!isLoading">
        <!-- <div>
                <div class="rat-title">Оценка КРС по хозяйству</div>
            <div class="rating-item">
                <div class="rating-param">Общая индексная оценка:</div>
                <div v-if="ratings_hoz">{{ ratings_hoz.GeneralValue || 'Нет информации'}}</div>
            </div>
            <div class="rating-item">
                <div class="rating-param">EBV по среднему удою за 305 дней:</div>
                <div>{{ round(ratings_hoz.EbvMilk) || 'Нет информации'}}</div>
            </div>
            <div class="rating-item">
                <div class="rating-param">EBV по среднему жиру за 305 дней:</div>
                <div>{{ round(ratings_hoz.EbvFat) || 'Нет информации'}}</div>
            </div>
            <div class="rating-item">
                <div class="rating-param">EBV по среднему белку за 305 дней:</div>
                <div>{{ round(ratings_hoz.EbvProtein) || 'Нет информации'}}</div>
            </div>
            <div class="rating-item">
                <div class="rating-param">EBV по средней кратности осеменения:</div>
                <div>{{ round(ratings_hoz.EbvInsemenation) || 'Нет информации'}}</div>
            </div>
            <div class="rating-item">
                <div class="rating-param">EBV по средней длительности сервис-периода:</div>
                <div>{{ round(ratings_hoz.EvbService) || 'Нет информации'}}</div>
            </div>
        </div> -->

        <div>
            <div class="rat-title">Оценка КРС по региону</div>
            <!-- <div class="rating-item">
                <div class="rating-param">Общая индексная оценка:</div>
                <div>{{ round(ratings_reg.GeneralValue) || 'Нет информации'}}</div>
            </div>
            <div class="rating-item">
                <div class="rating-param">EBV по среднему удою за 305 дней:</div>
                <div>{{ round(ratings_reg.EbvMilk) || 'Нет информации'}}</div>
            </div>
            <div class="rating-item">
                <div class="rating-param">EBV по среднему жиру за 305 дней:</div>
                <div>{{ round(ratings_reg.EbvFat) || 'Нет информации'}}</div>
            </div>
            <div class="rating-item">
                <div class="rating-param">EBV по среднему белку за 305 дней:</div>
                <div>{{ round(ratings_reg.EbvProtein) || 'Нет информации'}}</div>
            </div> -->
            <div class="table-chart">
                <table>
                    <thead>
                        <tr class="lac-header">
                            <th></th>
                            <th>Оценка</th>
                            <th>Достоверность</th>
                        </tr>
                    </thead>
                    <tbody class="lac-tablebody">
                        <tr>
                            <td>Селекционный индекс:</td>
                            <td>{{ round(ratings_reg.GeneralValue) || '-'}}</td>
                            <td></td>
                        </tr>
                        <tr>
                            <td>EBV по среднему удою за 305 дней, кг:</td>
                            <td>{{ round(ratings_reg.EbvMilk) || '-'}}</td>
                            <td>{{ round(ratings_reg.EbvMilkReliability) || '-'}}</td>
                        </tr>
                        <tr>
                            <td>EBV по среднему жиру за 305 дней, кг:</td>
                            <td>{{ round(ratings_reg.EbvFat) || '-'}}</td>
                            <td>{{ round(ratings_reg.EbvFatReliability) || '-'}}</td>
                        </tr>
                        <tr>
                            <td>EBV по среднему белку за 305 дней, кг:</td>
                            <td>{{ round(ratings_reg.EbvProtein) || '-'}}</td>
                            <td>{{ round(ratings_reg.EbvProteinReliability) || '-'}}</td>
                        </tr>
                        <tr>
                            <td>EBV по содержанию жира, %:</td>
                            <td>{{ round(ratings_reg.EbvFatPercents) || '-'}}</td>
                            <td>{{ round(ratings_reg.EbvFatPercentsReliability) || '-'}}</td>
                        </tr>
                        <tr>
                            <td>EBV по содержанию белка, %:</td>
                            <td>{{ round(ratings_reg.EbvProteinPercents) || '-'}}</td>
                            <td>{{ round(ratings_reg.EbvProteinPercentsReliability) || '-'}}</td>
                        </tr>
                        <tr>
                            <td>EBV по кратности осеменения, ед:</td>
                            <td>{{ round(ratings_reg.EbvInsemenation) || '-'}}</td>
                            <td>{{ round(ratings_reg.EbvInsemenationReliability) || '-'}}</td>
                        </tr>
                        <tr>
                            <td>EBV по сервис-периоду, дни:</td>
                            <td>{{ round(ratings_reg.EbvService) || '-'}}</td>
                            <td>{{ round(ratings_reg.EbvServiceReliability) || '-'}}</td>
                        </tr>
                        <tr>
                            <td>EBV по маститу, число случаев:</td>
                            <td>{{ round(ratings_reg.EbvMastit) || '-'}}</td>
                            <td>{{ round(ratings_reg.EbvMastitReliability) || '-'}}</td>
                        </tr>
                        <tr>
                            <td>EBV по продуктивному долголетию, число лактаций:</td>
                            <td>{{ round(ratings_reg.EbvProductiveLongevity) || '-'}}</td>
                            <td>{{ round(ratings_reg.EbvProductiveLongevityReliability) || '-'}}</td>
                        </tr>
                        <tr>
                            <td>EBV по соматическим клеткам, тыс. соматических клеток:</td>
                            <td>{{ round(ratings_reg.EbvSomaticNucs) || '-'}}</td>
                            <td>{{ round(ratings_reg.EbvSomaticNucsReliability) || '-'}}</td>
                        </tr>
                    </tbody>
                </table>

                <apexchart id="rating" width="400" height="632" type="bar" :options="options" :series="series" ref="rating"></apexchart>
            </div>
            

            <!-- <div class="rating-item">
                <div class="rating-param">EBV по средней кратности осеменения:</div>
                <div>{{ round(ratings_reg.EbvInsemenation) || 'Нет информации'}}</div>
            </div>
            <div class="rating-item">
                <div class="rating-param">EBV по средней длительности сервис-периода:</div>
                <div>{{ round(ratings_reg.EvbService) || 'Нет информации'}}</div>
            </div> -->
        </div>
    </div>
    <div v-if="isLoading">Идёт загрузка...</div>
</template>

<script>
export default {
    data() {
        return{ 
            ratings_hoz: {},
            ratings_reg: {},
            percents: {},

            options: {
                chart: {
                    id: 'rating',
                },
                xaxis: {
                    stepSize: 10,
                    categories: [],
                    min: -50,
                    max: 50,
                },
                plotOptions: {
                    bar: {
                        horizontal: true,
                    },
                },
                yaxis: {
                    show: false,
                },
                tooltip: {
                    enabled: false,
                },
                dataLabels: {
                    enabled: false,
                },
                title: {
                    text: ['Показатель животного относительно', 'среднего по региону, %'],
                    align: 'center',
                    offsetY: -4,
                    style: {
                        fontWeight: 'normal',
                        color: 'grey',
                    }
                }
            },
            series: [
                {
                    data: [
                        {x: '/', y: [0, 10]},
                        {x: '//', y: -7},
                        {x: '///', y: [0, 7]},
                        {x: '////', y: [0, 12]},
                    ]
                }
            ],

            isLoading: false,
        }
    },
    async created() {
        this.isLoading = true;
        let mass_route = this.$route.path.split('/');
        let cow_id = mass_route[2];
        let response = await fetch(`/api/cows/${cow_id}/grades`)
        let result = await response.json();
        console.log(result);
        if(result.ByHoz) {
            this.ratings_hoz = result.ByHoz;
        }
        if(result.ByRegion) {
            this.ratings_reg = result.ByRegion;
        }
        if(result.PercentsRegion) {
            this.percents = result.PercentsRegion;
            this.getPercents();
        }
        this.isLoading = false;
    },
    methods: {
        round(num) {
            return Math.round(num*100)/100;
        },
        getPercents() {
            this.series = [];
            let oneSerie = {
                data: [],
            }
            oneSerie.data.push({
                x: '1',
                y: this.round(Number(this.percents.GeneralValue))
            })
            oneSerie.data.push({
                x: '2',
                y: this.round(Number(this.percents.EbvMilk))
            })
            oneSerie.data.push({
                x: '3',
                y: this.round(Number(this.percents.EbvFat))
            })
            oneSerie.data.push({
                x: '4',
                y: this.round(Number(this.percents.EbvProtein))
            })
            oneSerie.data.push({
                x: '5',
                y: this.round(Number(this.percents.EbvFatPercents))
            })
            oneSerie.data.push({
                x: '6',
                y: this.round(Number(this.percents.EbvProteinPercents))
            })
            oneSerie.data.push({
                x: '7',
                y: this.round(Number(this.percents.EbvInsemenation))
            })
            oneSerie.data.push({
                x: '8',
                y: this.round(Number(this.percents.EbvService))
            })
            oneSerie.data.push({
                x: '9',
                y: this.round(Number(this.percents.EbvMastit))
            })
            oneSerie.data.push({
                x: '9',
                y: this.round(Number(this.percents.EbvProductiveLongevity))
            })
            oneSerie.data.push({
                x: '9',
                y: this.round(Number(this.percents.EbvSomaticNucs))
            })

            this.series = [oneSerie];
        }
    }
}
</script>

<style scoped>
.rat-title {
    font-size: 130%;
    color: rgb(10, 113, 75);
    padding: 0 0 30px 0;
    width: max-content;
}

.sub-title {
    color: red;
}

.rating-item {
    display: flex;
    justify-content: space-between;
    align-items: end;
    margin-right: 30px;
    border-bottom: 1px solid rgb(237, 248, 243);
    width: 370px;
}

.rating-param {
    margin: 5px 0;
    color: rgb(74, 74, 74);
    width: 250px;
}

.rating-columns {
    display: flex;
    font-size: 90%;
}

.lac-header {
    color: grey;
}

.lac-tablebody {
    text-align: left;
}

table {
    border-collapse: collapse;
    margin-top: 8px;
}

th {
    font-weight: normal;
    text-align: center;
    padding-bottom: 35px;
}

td {
    width: auto;
    min-width: 80px;
    text-align: center;
    height: 49px;
}

.table-chart {
    display: flex;
    align-items: flex-start;
}
</style>