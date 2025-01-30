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
            <apexchart id="analit" width="950" type="bar" :options="options" :series="series" ref="analit"></apexchart>
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
                    categories: ['Гармаша','Кубанский ГАУ УОХ «Краснодарское»', 'Север Кубани',  'Агрокомплекс Павловский', 'Рассвет Юг', 
                    'Ревко', 'Крупское', 'Ткачева Родина', 'АО ВикторияАгро', 'АК Павловский айрширы'],
                    dataLabels: {
                        enabled: true,
                    }
                    // axisTicks: {
                    //     offsetX: -35,
                    // },
                    // group: {
                    //     groups: [
                    //         {title: 'Кубанский ГАУ УОХ «Краснодарское»', cols: 1},
                    //         {title: 'Север Кубани', cols: 1},
                    //         {title: 'Гармаша', cols: 1},
                    //         {title: 'Агрокомплекс Павловский', cols: 1},
                    //         {title: 'Рассвет Юг', cols: 1},
                    //         {title: 'Ревко', cols: 1},
                    //         {title: 'Крупское', cols: 1},
                    //         {title: 'Ткачева Родина', cols: 1},
                    //         {title: 'АО ВикторияАгро', cols: 1},
                    //         {title: 'АК Павловский айрширы', cols: 1},
                    //     ]
                    // }
                },
                yaxis: {
                    title: {
                        text: 'Удой, кг',
                        offsetX: -8,
                        style: {
                            fontSize: 15,
                            cssClass: 'axis',
                        }
                    }
                },
                colors: ['#63d9cb','#6e5add','#75a2e7'],
                // title: {
                //     text: 'Хозяйства',
                //     align: 'center',
                //     style: {
                //         fontSize:  '24px',
                //     },
                // }
                // tooltip: {
                //     enabled: false,
                // }
            },
            series: [
                {
                    name: '2022 г',
                    group: '2022',
                    data: [10, 10, 15, 16, 20, 10, 13, 15, 16, 20],
                },
                {
                    name: '2023 г',
                    group: '2023',
                    data: [10, 13, 15, 16, 20, 10, 10, 15, 16, 20],
                },
                {
                    name: '2024 г',
                    group: '2024',
                    data: [11, 14, 12, 20, 13, 10, 10, 15, 30, 10],
                },
                // {
                //     name: 'Жир за 2022 г, кг',
                //     group: '2022',
                //     data: [10, 10, 15, 30, 10, 11, 14, 12, 20, 13],
                // },
                // {
                //     name: 'Жир за 2023 г, кг',
                //     group: '2023',
                //     data: [10, 10, 15, 13, 16, 9, 7, 13, 16, 19],
                // },
                // {
                //     name: 'Жир за 2024 г, кг',
                //     group: '2024',
                //     data: [9, 7, 13, 16, 19, 10, 10, 15, 13, 16],
                // },
                // {
                //     name: 'Белок за 2022 г, кг',
                //     group: '2022',
                //     data: [10, 10, 15, 23, 4, 10, 10, 15, 5, 9],
                // },
                // {
                //     name: 'Белок за 2023 г, кг',
                //     group: '2023',
                //     data: [10, 10, 15, 5, 9, 10, 10, 15, 23, 4],
                // },
                // {
                //     name: 'Белок за 2024 г, кг',
                //     group: '2024',
                //     data: [11, 13, 14, 13, 19, 20, 17, 16, 9, 12],
                // }
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
</style>