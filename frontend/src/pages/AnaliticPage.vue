<template>
    <div class="analytics-flex">
        <DAnimalFilters class="analytics-filters" @applyFilters="fetchAnalyticsFilters" v-bind:fromAnal="forFilters"/>
        <div class="chart-block">
            <div class="analytics-title">Статистика для сравнительного анализа хозяйств и регионов</div>
            <!-- <select v-model="opt" class="filter-input">
                <option :value="''">генотипирование</option>
                <option :value="'lact'">средние показатели удоя</option>
                <option :value="'ill'">моногенные заболевания</option>
            </select> -->
            <!-- <router-view></router-view> -->
            <div class="row">
                <router-view></router-view>
                <apexchart id="analit" width="500" type="bar" :options="options" :series="series"></apexchart>
            </div>
            <div class="row">
                <apexchart id="analit1" width="500" type="bar" :options="options1" :series="series1"></apexchart>
                <apexchart id="analit2" width="500" type="bar" :options="options2" :series="series2"></apexchart>
            </div>
            <div class="row">
                <apexchart id="analit3" width="500" type="bar" :options="options3" :series="series3"></apexchart>
                <apexchart id="analit4" width="500" type="bar" :options="options4" :series="series4"></apexchart>
            </div>
        </div>

    </div>
</template>

<script>
import DAnimalFilters from '@/components/testpage/DAnimalFilters.vue';

export default {
    components: {
        DAnimalFilters
    },
    data() {
        return{
            opt: '',
            forFilters: true,

            options: {
                chart: {
                    id: 'analit',
                    stacked: true,
                    zoom: {
                        enabled: false,
                    }
                },
                xaxis: {
                //     categories: [['ООО Агрокомплекс',' Павловский'],'АО Рассвет', 'АО Родина',  ['АО фирма Агрокомплекс', 'им. Н.И. Ткачева'], 'ОАО Племзавод Воля', 
                // 'АО Виктория – Агро', ['ФГБОУ ВО Кубанский', 'ГАУ УОХ Краснодарское'], ['ООО Агрокомплекс', 'Новокубанский', 'ОСП Ленинский путь']],
                categories: ['ООО Агрокомплекс Павловский','АО Рассвет', 'АО Родина',  'АО фирма Агрокомплекс им. Н.И. Ткачева', 'ОАО Племзавод Воля', 
                'АО Виктория – Агро', 'ФГБОУ ВО Кубанский ГАУ УОХ Краснодарское', 'ООО Агрокомплекс Новокубанский ОСП Ленинский путь'],
                    dataLabels: {
                        enabled: true,
                        style: {
                            fontSize: '5px',
                        }
                    },
                    labels: {
                        style: {
                            fontSize: '8px',
                        },
                        hideOverlappingLabels: true,
                        trim: true,
                    }
                },

                colors: ['#63d9cb','#6e5add','#75a2e7'],
                title: {
                    text: 'Средний удой на одну голову за 2024 г',
                    align: 'center',
                    style: {
                        fontSize:  '12px',
                    },
                },
                tooltip: {
                    enabled: false,
                }
            },
            series: [
                {
                    data: [8315,10883,15734,9277,11033,11264,15378,11810],
                }
            ],

            options1: {
                chart: {
                    id: 'analit1',
                    stacked: true,
                    zoom: {
                        enabled: false,
                    }
                },
                xaxis: {
                //     categories: [['ООО Агрокомплекс',' Павловский'],'АО Рассвет', 'АО Родина',  ['АО фирма Агрокомплекс', 'им. Н.И. Ткачева'], 'ОАО Племзавод Воля', 
                // 'АО Виктория – Агро', ['ФГБОУ ВО Кубанский', 'ГАУ УОХ Краснодарское'], ['ООО Агрокомплекс', 'Новокубанский', 'ОСП Ленинский путь']],
                categories: ['ООО Агрокомплекс Павловский','АО Рассвет', 'АО Родина',  'АО фирма Агрокомплекс им. Н.И. Ткачева', 'ОАО Племзавод Воля', 
                'АО Виктория – Агро', 'ФГБОУ ВО Кубанский ГАУ УОХ Краснодарское', 'ООО Агрокомплекс Новокубанский ОСП Ленинский путь'],
                    dataLabels: {
                        enabled: true,
                        style: {
                            fontSize: '5px',
                        }
                    },
                    labels: {
                        style: {
                            fontSize: '8px',
                        },
                        hideOverlappingLabels: true,
                        trim: true,
                    }
                },

                colors: ['#63d9cb','#6e5add','#75a2e7'],
                title: {
                    text: 'Выход телят',
                    align: 'center',
                    style: {
                        fontSize:  '12px',
                    },
                },
                tooltip: {
                    enabled: false,
                }
            },
            series1: [
                {
                    data: [77,60,85,93,62,99,72,43],
                }
            ],

            options2: {
                chart: {
                    id: 'analit2',
                    stacked: true,
                    zoom: {
                        enabled: false,
                    }
                },
                xaxis: {
                //     categories: [['ООО Агрокомплекс',' Павловский'],'АО Рассвет', 'АО Родина',  ['АО фирма Агрокомплекс', 'им. Н.И. Ткачева'], 'ОАО Племзавод Воля', 
                // 'АО Виктория – Агро', ['ФГБОУ ВО Кубанский', 'ГАУ УОХ Краснодарское'], ['ООО Агрокомплекс', 'Новокубанский', 'ОСП Ленинский путь']],
                categories: ['ООО Агрокомплекс Павловский','АО Рассвет', 'АО Родина',  'АО фирма Агрокомплекс им. Н.И. Ткачева', 'ОАО Племзавод Воля', 
                'АО Виктория – Агро', 'ФГБОУ ВО Кубанский ГАУ УОХ Краснодарское', 'ООО Агрокомплекс Новокубанский ОСП Ленинский путь'],
                    labels: {
                        style: {
                            fontSize: '8px',
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
                colors: ['#63d9cb','#6e5add','#75a2e7'],
                title: {
                    text: 'Валовый объём молока',
                    align: 'center',
                    style: {
                        fontSize:  '12px',
                    },
                },
                tooltip: {
                    enabled: false,
                }
            },
            series2: [
                {
                    data: [120000,430000,280000,355000,175000,490000,315000,260000],
                }
            ],

            options3: {
                chart: {
                    id: 'analit3',
                    stacked: true,
                    zoom: {
                        enabled: false,
                    }
                },
                xaxis: {
                //     categories: [['ООО Агрокомплекс',' Павловский'],'АО Рассвет', 'АО Родина',  ['АО фирма Агрокомплекс', 'им. Н.И. Ткачева'], 'ОАО Племзавод Воля', 
                // 'АО Виктория – Агро', ['ФГБОУ ВО Кубанский', 'ГАУ УОХ Краснодарское'], ['ООО Агрокомплекс', 'Новокубанский', 'ОСП Ленинский путь']],
                categories: ['ООО Агрокомплекс Павловский','АО Рассвет', 'АО Родина',  'АО фирма Агрокомплекс им. Н.И. Ткачева', 'ОАО Племзавод Воля', 
                'АО Виктория – Агро', 'ФГБОУ ВО Кубанский ГАУ УОХ Краснодарское', 'ООО Агрокомплекс Новокубанский ОСП Ленинский путь'],
                    labels: {
                        style: {
                            fontSize: '8px',
                        },
                        hideOverlappingLabels: true,
                        trim: true,
                    }
                },
                colors: ['#63d9cb','#6e5add','#75a2e7'],
                title: {
                    text: 'Показатель 1 от Плотникова',
                    align: 'center',
                    style: {
                        fontSize:  '12px',
                    },
                },
                tooltip: {
                    enabled: false,
                }
            },
            series3: [
                {
                    data: [10,14,12,8,20,22,6,25],
                }
            ],

            options4: {
                chart: {
                    id: 'analit4',
                    stacked: true,
                    zoom: {
                        enabled: false,
                    }
                },
                xaxis: {
                //     categories: [['ООО Агрокомплекс',' Павловский'],'АО Рассвет', 'АО Родина',  ['АО фирма Агрокомплекс', 'им. Н.И. Ткачева'], 'ОАО Племзавод Воля', 
                // 'АО Виктория – Агро', ['ФГБОУ ВО Кубанский', 'ГАУ УОХ Краснодарское'], ['ООО Агрокомплекс', 'Новокубанский', 'ОСП Ленинский путь']],
                categories: ['ООО Агрокомплекс Павловский','АО Рассвет', 'АО Родина',  'АО фирма Агрокомплекс им. Н.И. Ткачева', 'ОАО Племзавод Воля', 
                'АО Виктория – Агро', 'ФГБОУ ВО Кубанский ГАУ УОХ Краснодарское', 'ООО Агрокомплекс Новокубанский ОСП Ленинский путь'],
                    labels: {
                        style: {
                            fontSize: '8px',
                        },
                        hideOverlappingLabels: true,
                        trim: true,
                    }
                },
                colors: ['#63d9cb','#6e5add','#75a2e7'],
                title: {
                    text: 'Показатель 2 от Плотникова',
                    align: 'center',
                    style: {
                        fontSize:  '12px',
                    },
                },
                tooltip: {
                    enabled: false,
                }
            },
            series4: [
                {
                    data: [190,203,175,230,250,180,212,214],
                }
            ],
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
        // splitString(str) {
        //     let arr = srt.split(' ');
        //     if (arr.length > 2) {

        //     }
        // }
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
    justify-content: space-around;
    width: 90vw;
}

.analytics-filters {
    margin-top: 130px;
}

.analytics-title {
    font-size: 200%;
    font-family: Open Sans, sans-serif;
    color: rgb(10, 113, 75);
    margin-top: 120px;
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
    width: 880px;
    margin-top: 30px;
}
</style>