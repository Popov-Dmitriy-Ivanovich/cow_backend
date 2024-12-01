<template>
<div>
    <div class="cm-title">Контрольные доения</div>
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
        <table class="cm-table">
            <thead>
                <tr class="cm-header">
                    <th>Номер лактации</th>
                    <th class="cm-date">Дата</th>
                    <th class="cm-nprob">Номер пробы</th>
                    <th class="cm-milkday">Дойные дни</th>
                    <th class="cm-milk">Удой, л</th>
                    <th class="cm-dry">Сухое вещество</th>
                    <th class="cm-zhir">Жир, %</th>
                    <th class="cm-belok">Белок, %</th>
                </tr>
            </thead>
            <tbody class="cm-tablebody">
                <tr v-for="milking in cow_info" :key="milking.CheckDate">
                    <td>{{ milking.LactationId }}</td>
                    <td>{{ milking.CheckDate }}</td>
                    <td>{{ milking.ProbeNumber }}</td>
                    <td>{{ milking.MilkingDays }}</td>
                    <td>{{ milking.Milk }}</td>
                    <td>{{ milking.DryMatter }}</td>
                    <td>{{ milking.Fat }}</td>
                    <td>{{ milking.Protein }}</td>
                </tr>
            </tbody>
        </table>
    </div>

    <div v-if="isChart">
        <div>
            <div class="nlact">
                <div class="chart-nlact" @click="isLact=!isLact">> Номер лактации </div>
                <div v-if="isLact" class="check-lact">
                    <div v-for="nlact in count_lactations" :key="nlact">
                        <label><input type="checkbox" :value="nlact" v-model="check_lact">Лактация {{ nlact }}</label>
                    </div>
                </div>
            </div>

            <div class="chart-flex">
                <div class="chart-param">Показатель: </div>
                <select v-model="param_milking" class="select-param">
                    <option disabled value="">Выберите параметр</option>
                    <option value="Milk">Удой</option>
                    <option value="Fat">Жир</option>
                    <option value="Protein">Белок</option>
                    <option value="FatRegard">Жир, %</option>
                    <option value="ProteinRegard">Белок, %</option>
                </select>
            </div>

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
            count_lactations:[],
            param_milking: 'Milk',
            check_lact: [1],
            isLact: false,
        }
    },
    async created() {
        let mass_route = this.$route.path.split('/');
        let cow_id = mass_route[2];
        let response = await fetch(`/api/cows/${cow_id}/checkMilks`);
        let result = await response.json();
        this.cow_info = result;
        
        for (let i = 0; i < this.cow_info.length; i++) {
            this.cow_info[i].FatRegard = ((this.cow_info[i].Fat / this.cow_info[i].Milk)*100).toFixed(2);
            this.cow_info[i].ProteinRegard = ((this.cow_info[i].Protein / this.cow_info[i].Milk)*100).toFixed(2);
        }

        try {
            this.count_lactations.push(this.cow_info[0]['LactationId']);
            for (let i = 1; i < this.cow_info.length; i++) {
                if (this.cow_info[i]['LactationId'] !== this.cow_info[i-1]['LactationId']) {
                    this.count_lactations.push(this.cow_info[i]['LactationId']);
                }
            }
        } catch(err) {
            console.log(err);
        }

        this.addParam(this.cow_info, this.series, this.param_milking, this.check_lact)

        console.log(this.count_lactations);
        console.log(this.cow_info);
    },
    methods: {
        addInArr(obj, arr, param, nlact) {
            for (let i=0; i<obj.length; i++) {
                for (let j=0; j<nlact.length; j++) {
                    if(nlact[j] === obj[i].LactationId) {
                        arr.push(obj[i][param]);
                    }
                }
            }
            console.log('параметры',arr);
        },
        addParam(obj, arr, param, nlact) {
            for (let i = 0; i < nlact.length; i++) {
                let serie = {
                    name: `Лактация ${nlact[i]}`,
                    data: []
                };
                for (let j = 0; j < obj.length; j++) {
                    if (obj[j].LactationId === nlact[i]) {
                        serie.data.push(obj[j][param]);
                    }
                }
                arr.push(serie);
            }
        }
    },
    watch: {
        param_milking(new_value) {
            this.series = [];
            this.addParam(this.cow_info, this.series, new_value, this.check_lact)
            console.log(this.series, 'данные для столбцов');
        },
        check_lact(new_value) {
            this.series = [];
            this.addParam(this.cow_info, this.series, this.param_milking, new_value)
        }

    }
}
</script>

<style scoped>
.cm-title {
    font-size: 130%;
    color: rgb(37, 0, 132);
    padding-bottom: 10px;
}

.parent-table {
    width: 50vw;
    overflow-x: auto;
}

.cm-table {
    margin-bottom: 30px;
    /* text-align: center; */
    border-spacing: 10px;
    font-size: 85%;
}

th {
    font-weight: normal;
}

.cm-header {
    color: grey;
}

.cm-header th {
    padding-right: 10px;
    padding-bottom: 15px;
}

.cm-tablebody {
    text-align: center;
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

.chart-nlact {
    width: max-content;
    padding: 6px 4px;
    margin-bottom: 3px;
    font-size: 110%;
    cursor: pointer;
    transition: 0.3s;
}

.chart-nlact:hover {
    color: rgb(92, 78, 145)
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

.check-lact {
    display: flex;
    flex-direction: column;
    max-width: 300px;
    flex-wrap: wrap;
    background-color: white;
    padding: 20px 15px;
    border-radius: 10px;
    /* box-shadow: rgba(100, 100, 111, 0.1) 0px 7px 29px 0px; */
    z-index: 20;
}

.check-lact div {
    padding: 5px 0;
}

.nlact {
    display: inline-block;
}
</style>