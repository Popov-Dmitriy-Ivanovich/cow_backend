<template>
    <div class="analytics-flex">
        <!-- <DAnimalFilters class="analytics-filters" @applyFilters="fetchAnalyticsFilters" v-bind:fromAnal="forFilters"/> -->
        
        <div class="chart-block">
            <div class="analytics-title">Аналитика</div>
            <!-- <select v-model="opt" class="filter-input">
                <option :value="''">генотипирование</option>
                <option :value="'lact'">средние показатели удоя</option>
                <option :value="'ill'">моногенные заболевания</option>
            </select> -->
            <!-- <router-view></router-view> -->
            <div class="analyt-btns">
                <button
                class="btn"
                :class="{'active-btn':isSelex}"
                @click="isRegion = false; isSelex = true"
                >Селекционный индекс</button>

                <button
                class="btn"
                :class="{'active-btn':isRegion}"
                @click="isRegion = true; isSelex = false"
                >Статистика по региону</button>
            </div>
            
            <div class="row" v-if="isSelex">
                <router-view></router-view>
            </div>
            <div class="row" v-if="isRegion">
                <apexchart id="analit1" width="500" type="bar" :options="options1" :series="series1" @dataPointSelection="clickHandler"></apexchart>
                <apexchart id="analit2" width="500" type="bar" :options="options2" :series="series2" @dataPointSelection="clickHandler"></apexchart>
            </div>
            <div class="row" v-if="isRegion">
                <apexchart id="analit3" width="500" type="bar" :options="options3" :series="series3" @dataPointSelection="clickHandler"></apexchart>
                <apexchart id="analit4" width="500" type="bar" :options="options4" :series="series4" @dataPointSelection="clickHandler"></apexchart>
            </div>
        </div>

    </div>
</template>

<script>
// import DAnimalFilters from '@/components/testpage/DAnimalFilters.vue';

export default {
    components: {
        // DAnimalFilters
    },
    data() {
        return{
            opt: '',
            forFilters: true,

            options1: {
                chart: {
                    id: 'analit1',
                    zoom: {
                        enabled: false,
                    }
                },
                xaxis: {
                //     categories: [['ООО Агрокомплекс',' Павловский'],'АО Рассвет', 'АО Родина',  ['АО фирма Агрокомплекс', 'им. Н.И. Ткачева'], 'ОАО Племзавод Воля', 
                // 'АО Виктория – Агро', ['ФГБОУ ВО Кубанский', 'ГАУ УОХ Краснодарское'], ['ООО Агрокомплекс', 'Новокубанский', 'ОСП Ленинский путь']],
                categories: ['2021', '2022', '2023', '2024'],
                    labels: {
                        style: {
                            fontSize: '12px',
                        },
                        hideOverlappingLabels: true,
                        trim: true,
                    }
                },
                dataLabels: {
                    enabled: true,
                    style: {
                        fontSize: '9px',
                    }
                },
                legend: {
                    show: false,
                },
                colors: ['#78DABC','#6e5add','#75a2e7'],
                title: {
                    text: ['Соотношение общего числа голов', 'к числу голов охваченных проектом'],
                    align: 'center',
                    style: {
                        fontSize:  '15px',
                    },
                },
            },
            series1: [
                {
                    name: 'Число всех голов',
                    data: [36712, 44039, 45569, 49435],
                },
                {
                    name: 'Число голов в проекте',
                    data: [0, 9732, 21422, 29612],
                }
            ],

            options2: {
                chart: {
                    id: 'analit2',
                    zoom: {
                        enabled: false,
                    }
                },
                xaxis: {
                //     categories: [['ООО Агрокомплекс',' Павловский'],'АО Рассвет', 'АО Родина',  ['АО фирма Агрокомплекс', 'им. Н.И. Ткачева'], 'ОАО Племзавод Воля', 
                // 'АО Виктория – Агро', ['ФГБОУ ВО Кубанский', 'ГАУ УОХ Краснодарское'], ['ООО Агрокомплекс', 'Новокубанский', 'ОСП Ленинский путь']],
                categories: ['2021', '2022', '2023', '2024'],
                    labels: {
                        style: {
                            fontSize: '12px',
                        },
                        hideOverlappingLabels: true,
                        trim: true,
                    }
                },
                dataLabels: {
                    enabled: true,
                },
                legend: {
                    show: false,
                },
                colors: ['#78DABC','#6e5add','#75a2e7'],
                title: {
                    text: 'Средний удой, кг',
                    align: 'center',
                    style: {
                        fontSize:  '15px',
                    },
                },
            },
            series2: [
                {
                    name: 'Средний удой',
                    data: [9733.84, 10353.29, 10665, 11742],
                }
            ],

            options3: {
                chart: {
                    id: 'analit3',
                    zoom: {
                        enabled: false,
                    }
                },
                xaxis: {
                //     categories: [['ООО Агрокомплекс',' Павловский'],'АО Рассвет', 'АО Родина',  ['АО фирма Агрокомплекс', 'им. Н.И. Ткачева'], 'ОАО Племзавод Воля', 
                // 'АО Виктория – Агро', ['ФГБОУ ВО Кубанский', 'ГАУ УОХ Краснодарское'], ['ООО Агрокомплекс', 'Новокубанский', 'ОСП Ленинский путь']],
                    categories: ['2021', '2022', '2023', '2024'],
                    labels: {
                        style: {
                            fontSize: '12px',
                        },
                        hideOverlappingLabels: true,
                        trim: true,
                    }
                },
                dataLabels: {
                    enabled: true,
                    offsetY: 730,
                },
                yaxis: {
                    stepSize: 0.1,
                    min: 3.4,
                    max: 3.9,
                },
                legend: {
                    show: false,
                },
                colors: ['#78DABC','#6e5add','#75a2e7'],
                title: {
                    text: 'Средний жир, %',
                    align: 'center',
                    style: {
                        fontSize:  '15px',
                    },
                },
            },
            series3: [
                {
                    name: 'Средний жир, %',
                    data: [3.80, 3.78, 3.73, 3.69],
                }
            ],

            options4: {
                chart: {
                    id: 'analit3',
                    zoom: {
                        enabled: false,
                    }
                },
                xaxis: {
                //     categories: [['ООО Агрокомплекс',' Павловский'],'АО Рассвет', 'АО Родина',  ['АО фирма Агрокомплекс', 'им. Н.И. Ткачева'], 'ОАО Племзавод Воля', 
                // 'АО Виктория – Агро', ['ФГБОУ ВО Кубанский', 'ГАУ УОХ Краснодарское'], ['ООО Агрокомплекс', 'Новокубанский', 'ОСП Ленинский путь']],
                    categories: ['2021', '2022', '2023', '2024'],
                    labels: {
                        style: {
                            fontSize: '12px',
                        },
                        hideOverlappingLabels: true,
                        trim: true,
                    }
                },
                yaxis: {
                    stepSize: 0.05,
                    min: 3.15,
                    max: 3.3,
                },
                dataLabels: {
                    enabled: true,
                    offsetY: 2270,
                },
                legend: {
                    show: false,
                },
                colors: ['#78DABC','#6e5add','#75a2e7'],
                title: {
                    text: 'Средний белок, %',
                    align: 'center',
                    style: {
                        fontSize:  '15px',
                    },
                },
            },
            series4: [
                {
                    name: 'Средний белок, %',
                    data: [3.2677, 3.27, 3.26, 3.27],
                }
            ],

            isRegion: false,
            isSelex: true,
        }
    },
    created() {
        this.$store.commit('SET_OPTION', this.opt);
    },
    methods: {
        async fetchAnalyticsFilters(filters){
            this.$store.commit('SET_FILTERS', filters);
            console.log(this.$store.getters.FILTERS, 'in store');
        },
        clickHandler() {
            this.$router.push('/analytics/general');
        }
    },
    watch: {
        opt(new_val) {
            this.$store.commit('SET_OPTION', new_val);
        }
    }
}
</script>

<style scoped>
.analytics-flex {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    margin-top: 30px;
}

.analytics-filters {
    margin-top: 130px;
}

.chart-block {
    display: flex;
    flex-direction: column;
    align-items: center;
    width: max-content;
}

.analytics-title {
    font-size: 200%;
    font-family: Open Sans, sans-serif;
    color: rgb(10, 113, 75);
    margin-top: 120px;
    margin-bottom: 20px;
}

.filter-input {
    margin: 10px 9px;
    height: 30px;
    width: 200px;
    padding: 0 10px;
    font-size: 14px;
    box-sizing: border-box;
    outline: none;
    border: 3px solid rgb(195, 212, 205);
    border-radius: 10px;
    transition: 0.3s;
}

.axis {
    padding: 10px;
}

.row {
    display: flex;
    margin-top: 30px;
    width: 65vw;
}

.analyt-btns {
    align-self: flex-start;
    margin-left: 30px;
}

.btn {
    border: 1px solid rgb(10, 113, 75);
    background-color: white;
    color: rgb(10, 113, 75);
    padding: 10px 0;
    margin: 20px 0;
    width: 200px;
    border-radius: 10px;
    transition: 0.3s;
    cursor: pointer;
    margin-right: 15px;
    transition: 0.3s;
}

.active-btn {
    background-color: rgb(214, 239, 233);
}
</style>