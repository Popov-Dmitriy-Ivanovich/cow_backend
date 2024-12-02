<template>
<div>
    <div class="lac-title">Лактации</div>
    <button @click="isTable=true;isChart=false" 
    class="table-chart"
    :class="{'active-btn':isTable}">
    Таблица</button>
    <button 
    @click="isTable=false;isChart=true" 
    class="table-chart"
    :class="{'active-btn':isChart}">
    График</button>
    <div class="parent-table" v-if="isTable">
        <table class="lac-table">
            <thead>
                <tr class="lac-header">
                    <th>Номер лактации</th>
                    <th>Дата лактации</th>
                    <th>Кратность осеменения</th>
                    <th>Дата осеменения</th>
                    <th>Количество рожденных телят</th>
                    <th>Дата отела</th>
                    <th>Удой полный, л</th>
                    <th>Удой 305 дней, л</th>
                    <th>Жир полный, кг</th>
                    <th>Жир 305 дней, кг</th>
                    <th>Белок полный, кг</th>
                    <th>Белок 305 дней, кг</th>
                    <th>Количество дней дойки</th>
                </tr>
            </thead>
            <tbody class="lac-tablebody">
                <tr v-for="lact in cow_info" :key="lact.Number">
                    <td>{{ lact.Number }}</td>
                    <td>{{ lact.Date }}</td>
                    <td>{{ lact.InsemenationNum }}</td>
                    <td>{{ lact.InsemenationDate }}</td>
                    <td>{{ lact.CalvingCount }}</td>
                    <td>{{ lact.CalvingDate }}</td>
                    <td>{{ lact.MilkAll }}</td>
                    <td>{{ lact.Milk305 }}</td>
                    <td>{{ lact.FatAll }}</td>
                    <td>{{ lact.Fat305 }}</td>
                    <td>{{ lact.ProteinAll }}</td>
                    <td>{{ lact.Protein305 }}</td>
                    <td>{{ lact.Days }}</td>
                </tr>
            </tbody>
        </table>
    </div>

    <div v-if="isChart">
        <div class="chart-flex">
            <div class="chart-param">Показатель: </div>
            <select v-model="param_milking" class="select-param">
                <option value="MilkAll">Удой полный</option>
                <option value="Milk305">Удой 305 дней</option>
                <option value="FatAll">Жир полный</option>
                <option value="Fat305">Жир 305 дней</option>
                <option value="ProteinAll">Белок полный</option>
                <option value="Protein305">Белок 305 дней</option>
                <option value="MilkDaily">Удой среднесуточный</option>
                <option value="FatDaily">Жир среднесуточный</option>
                <option value="ProteinDaily">Белок среднесуточный</option>
            </select>
        </div>

        <apexchart id="ControlMilking" width="600" type="bar" :options="options" :series="series"></apexchart>
    </div>
</div>
</template>

<script>
export default {
    data() {
        return {
            cow_info: [],
            isTable: true,
            isChart: false,
            options: {
                chart: {
                    id: 'ControlMilking'
                },
                xaxis: {
                    categories: []
                }
            },
            series: [],
            param_milking: 'MilkAll',
        }
    },
    async created() {
        let mass_route = this.$route.path.split('/');
        let cow_id = mass_route[2];
        let response = await fetch(`/api/cows/${cow_id}/lactations`);
        let result = await response.json();
        this.cow_info = result;
        console.log(this.cow_info);

        let serie = {name:'Удой полный',data: []};
        for (let i = 0; i < this.cow_info.length; i++) {
            this.cow_info[i].MilkDaily = (this.cow_info[i].MilkAll / this.cow_info[i].Days).toFixed(2);
            this.cow_info[i].FatDaily = (this.cow_info[i].FatAll / this.cow_info[i].Days).toFixed(2);
            this.cow_info[i].ProteinDaily = (this.cow_info[i].ProteinAll / this.cow_info[i].Days).toFixed(2);

            this.options.xaxis.categories.push('Лактация ' + this.cow_info[i].Number);
            serie.data.push(this.cow_info[i].MilkAll);
        }
        this.series.push(serie);
    },
    methods: {
        addParam(obj, arr, param) {
            let serie = {
                name: `${param}`,
                data: []
            };
            for (let i = 0; i < obj.length; i++) {
                serie.data.push(obj[i][param]);
            }
            arr.push(serie);
        }
    },
    watch: {
        param_milking(new_value) {
            this.series = [];
            this.addParam(this.cow_info, this.series, new_value);
        }
    }
}
</script>

<style scoped>
.lac-title {
    font-size: 130%;
    color: rgb(37, 0, 132);
    padding-bottom: 30px;
    width: max-content;
}

.parent-table {
    width: 40vw;
    overflow-x: auto;
}

.lac-table {
    margin-bottom: 30px;
    text-align: left;
}

th {
    font-weight: normal;
}

td {
    width: auto;
    min-width: 130px;
}

.lac-header {
    color: grey;
}

.lac-header th {
    padding-right: 30px;
    padding-bottom: 15px;
}

.lac-tablebody {
    text-align: left;
}

.parent-table::-webkit-scrollbar {
    height: 12px;
}

.parent-table::-webkit-scrollbar-track {
    background: rgb(241, 241, 241);
}

.parent-table::-webkit-scrollbar-thumb {
    background-color: rgb(183, 183, 183);
    border-radius: 20px;
    border: 3px solid rgb(241, 241, 241);
}

.table-chart {
    border: 1px solid rgb(37, 0, 132);
    background-color: white;
    color: rgb(37, 0, 132);
    padding: 10px 0;
    margin: 20px 0;
    width: 100px;
    border-radius: 10px;
    transition: 0.3s;
    cursor: pointer;
    margin-right: 15px;
    transition: 0.3s;
}

.active-btn {
    background-color: rgb(219, 214, 239);
}

.chart-param {
    width: max-content;
    padding: 6px 4px;
    margin-bottom: 3px;
    font-size: 110%;
}

.chart-flex {
    display: flex;
    align-items: center;
    justify-content: space-between;
    width: 350px;
}

.select-param {
    padding: 5px;
    border: 1px solid rgb(37, 0, 132);
    border-radius: 3px;
}
</style>