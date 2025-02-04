<template>
    <div class="analytics-flex">
        <div class="chart-block">
            <div class="analytics-title">Аналитика</div>
            <div class="prev-chart" @click="toPrev">🠔 Назад</div>
            <div class="row">
                <apexchart id="analit1" width="570" type="bar" :options="options1" :series="series1"></apexchart>
                <apexchart id="analit2" width="570" type="bar" :options="options2" :series="series2"></apexchart>
            </div>
            <div class="row">
                <apexchart id="analit3" width="570" type="bar" :options="options3" :series="series3"></apexchart>
                <apexchart id="analit4" width="570" type="bar" :options="options4" :series="series4"></apexchart>
            </div>
        </div>

    </div>
</template>

<script>
export default {
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
                    categories: ['АО "Родина"', 'ФГБОУ ВО Кубанский ГАУ (учхоз Краснодарское)', 'АО "Рассвет"', 'АО "Виктория-Агро"', 
                    'АО ф. «Агрокомплекс» им. Н.И.Ткачева', 'ОАО "Племзавод "Воля"', 'АО «Агрокомплекс Павловский»'],
                    labels: {
                        style: {
                            fontSize: '12px',
                        },
                        hideOverlappingLabels: true,
                        trim: true,
                    }
                },
                dataLabels: {
                    enabled: false,
                },
                colors: ['#78DABC','#6e5add','#75a2e7','#AB80EB'],
                title: {
                    text: 'Число голов в хозяйствах проекта',
                    align: 'center',
                    style: {
                        fontSize:  '15px',
                    },
                },
            },
            series1: [
                {
                    name: '2021',
                    data: [6268,1050,3405,709,20014,916,2310],
                },
                {
                    name: '2022',
                    data: [8615, 1050, 3618, 710, 21090, 909, 2370]
                },
                {
                    name: '2023',
                    data: [8860,1050,3522,711,22429,0,0]
                },
                {
                    name: '2024',
                    data: [11168,1100,3518,736,22722,0,0]
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
                    categories: ['АО "Родина"', 'ФГБОУ ВО Кубанский ГАУ (учхоз Краснодарское)', 'АО "Рассвет"', 'АО "Виктория-Агро"', 
                    'АО ф. «Агрокомплекс» им. Н.И.Ткачева', 'ОАО "Племзавод "Воля"', 'АО «Агрокомплекс Павловский»'],
                    labels: {
                        style: {
                            fontSize: '12px',
                        },
                        hideOverlappingLabels: true,
                        trim: true,
                    }
                },
                dataLabels: {
                    enabled: false,
                },
                colors: ['#78DABC','#6e5add','#75a2e7','#AB80EB'],
                title: {
                    text: 'Удой на корову с начала года',
                    align: 'center',
                    style: {
                        fontSize:  '15px',
                    },
                },
            },
            series2: [
                {
                    name: '2021',
                    data: [14021,13690,10739,9541,8500,8360,9778],
                },
                {
                    name: '2022',
                    data: [14372,14049,10754,9536,9110,9122,9737],
                },
                {
                    name: '2023',
                    data: [14313,14493,11327,11259,9341,0,0],
                },
                {
                    name: '2024',
                    data: [15734,15378,10883,11264,9277,0,0],
                },
            ],

            options3: {
                chart: {
                    id: 'analit3',
                    zoom: {
                        enabled: false,
                    }
                },
                xaxis: {
                    categories: ['АО "Родина"', 'ФГБОУ ВО Кубанский ГАУ (учхоз Краснодарское)', 'АО "Рассвет"', 'АО "Виктория-Агро"', 
                    'АО ф. «Агрокомплекс» им. Н.И.Ткачева', 'ОАО "Племзавод "Воля"', 'АО «Агрокомплекс Павловский»'],
                    labels: {
                        style: {
                            fontSize: '12px',
                        },
                        hideOverlappingLabels: true,
                        trim: true,
                    }
                },
                yaxis: {
                    stepSize: 0.1,
                    min: 3.4,
                    max: 3.9,
                },
                dataLabels: {
                    enabled: false,
                },
                colors: ['#78DABC','#6e5add','#75a2e7','#AB80EB'],
                title: {
                    text: 'Жир, %',
                    align: 'center',
                    style: {
                        fontSize:  '15px',
                    },
                },
            },
            series3: [
            {
                    name: '2021',
                    data: [3.90,3.53,3.85,3.75,3.55,3.70,3.63],
                },
                {
                    name: '2022',
                    data: [3.91,3.54,3.76,3.94,3.60,3.93,3.48],
                },
                {
                    name: '2023',
                    data: [3.93,3.45,3.78,4.30,3.61,0.00,0.00],
                },
                {
                    name: '2024',
                    data: [3.40,3.25,3.20,3.38,3.17,3.38,3.30],
                },
            ],

            options4: {
                chart: {
                    id: 'analit3',
                    zoom: {
                        enabled: false,
                    }
                },
                xaxis: {
                    categories: ['АО "Родина"', 'ФГБОУ ВО Кубанский ГАУ (учхоз Краснодарское)', 'АО "Рассвет"', 'АО "Виктория-Агро"', 
                    'АО ф. «Агрокомплекс» им. Н.И.Ткачева', 'ОАО "Племзавод "Воля"', 'АО «Агрокомплекс Павловский»'],
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
                    enabled: false,
                },
                colors: ['#78DABC','#6e5add','#75a2e7','#AB80EB'],
                title: {
                    text: 'Белок, %',
                    align: 'center',
                    style: {
                        fontSize:  '15px',
                    },
                },
            },
            series4: [
                {
                    name: '2021',
                    data: [3.31,3.24,3.27,3.24,3.13,3.18,3.26],
                },
                {
                    name: '2022',
                    data: [3.40,3.25,3.20,3.38,3.17,3.38,3.30],
                },
                {
                    name: '2023',
                    data: [3.44,3.35,3.19,3.50,3.18,0.00,0.00],
                },
                {
                    name: '2024',
                    data: [3.41,3.38,3.25,3.45,3.18,0.00,0.00],
                },
            ],
        }
    },
    methods: {
        toPrev() {
            this.$router.push(`/analytics`);
        }
    },
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

.axis {
    padding: 10px;
}

.row {
    display: flex;
    margin-top: 30px;
    width: 75vw;
}

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