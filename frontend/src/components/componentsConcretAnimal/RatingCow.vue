<template>
    <div v-if="!isLoading">
        <div class="select-year">
            <button @click="isReg=true;isHoz=false" 
            class="reg-hoz"
            :class="{'active-btn':isReg}">
            Оценка животного по региону</button>
            <button 
            @click="isReg=false;isHoz=true" 
            class="reg-hoz"
            :class="{'active-btn':isHoz}">
            Оценка животного по хозяйству</button>
            <div class="select-year">
                <div>Просмотр значений и индексов за </div> 
                <select class="select-param">
                    <option>2024 г.</option>
                </select>
            </div>
        </div>

        <div class="rating-columns">
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
                            <td>{{'-'}}</td>
                            <td></td>
                        </tr>
                        <tr>
                            <td>Племенная ценность по удою, кг:</td>
                            <td>{{'-'}}</td>
                            <td>{{'-'}}</td>
                        </tr>
                        <tr>
                            <td>Племенная ценность по количеству жира, кг:</td>
                            <td>{{'-'}}</td>
                            <td>{{'-'}}</td>
                        </tr>
                        <tr>
                            <td>Племенная ценность по количеству белка, кг:</td>
                            <td>{{'-'}}</td>
                            <td>{{'-'}}</td>
                        </tr>
                        <tr>
                            <td>Племенная ценность по доле жира, %:</td>
                            <td>{{'-'}}</td>
                            <td>{{'-'}}</td>
                        </tr>
                        <tr>
                            <td>Племенная ценность по доле белка, %:</td>
                            <td>{{'-'}}</td>
                            <td>{{'-'}}</td>
                        </tr>
                        <tr>
                            <td>Племенная ценность по продолжительности сервис-периода, дни:</td>
                            <td>{{'-'}}</td>
                            <td>{{'-'}}</td>
                        </tr>
                        <tr>
                            <td>Племенная ценность по кратности осеменения в 1 лактацию, ед.:</td>
                            <td>{{'-'}}</td>
                            <td>{{'-'}}</td>
                        </tr>
                        <tr>
                            <td>Племенная ценность по продуктивному долголетию, дни:</td>
                            <td>{{'-'}}</td>
                            <td>{{'-'}}</td>
                        </tr>
                        <tr>
                            <td>Племенная ценность по маститам, случаев:</td>
                            <td>{{'-'}}</td>
                            <td>{{'-'}}</td>
                        </tr>
                        <tr>
                            <td>Племенная ценность по количеству соматических клеток, ед:</td>
                            <td>{{'-'}}</td>
                            <td>{{'-'}}</td>
                        </tr>
                    </tbody>
                </table>

                <apexchart id="rating" width="400" height="632" type="bar" :options="options" :series="series" ref="rating"></apexchart>
            </div>
        </div>
            <div class="text">
                Используемая структура индекса:<br>
                I = 52.7×MF305 + 64.7× MP305 – 54.4 × SP – 0.9 × KO,<br>
                где: MF305 – Племенная ценность по количеству жира, MP305 – Племенная ценность по количеству белка, SP – Племенная ценность по продолжительности сервис-периода, KO – Племенная ценность по кратности осеменения в 1 лактацию.
            </div>
    </div>
    <div v-if="isLoading">Идёт загрузка...</div>
</template>

<script>
export default {
    data() {
        return{ 
            isReg: true,
            isHoz: false,
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
                        {x: '1', y: [0, 10]},
                        {x: '2', y: -7},
                        {x: '3', y: [0, 7]},
                        {x: '4', y: [0, 12]},
                        {x: '5', y: [0, 6]},
                        {x: '6', y: [0, 5]},
                        {x: '7', y: -10},
                        {x: '8', y: -11},
                        {x: '9', y: [0, 11]},
                        {x: '10', y: -13},
                        {x: '11', y: [0, 12]},
                    ]
                }
            ],

            isLoading: false,
        }
    },
    methods: {
        round(num) {
            return Math.round(num*100)/100;
        },
    }
}
</script>

<style scoped>
.rat-title {
    font-size: 130%;
    color: rgb(37, 0, 132);
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
    text-align: left;
    padding-bottom: 35px;
}

td {
    width: auto;
    min-width: 80px;
    text-align: left;
    height: 49px;
}

.table-chart {
    display: flex;
    align-items: flex-start;
}

.text {
    color: grey;
    font-size: 80%;
    /* text-align: center; */
}

.reg-hoz {
    border: 1px solid rgb(37, 0, 132);
    background-color: white;
    color: rgb(37, 0, 132);
    padding: 10px 0;
    margin: 20px 0;
    width: 220px;
    border-radius: 10px;
    transition: 0.3s;
    cursor: pointer;
    margin-right: 15px;
    transition: 0.3s;
}

.active-btn {
    background-color: rgb(219, 214, 239);
}

.select-year {
    display: flex;
    align-items: center;
    justify-content: space-between;
}

.select-param {
    padding: 5px;
    border: 1px solid rgb(37, 0, 132);
    border-radius: 3px;
    margin-left: 10px;
}
</style>