<template>
<div class="cm-title">Контрольные доения</div>
<div v-if="!isLoading">
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
                    <th class="cm-milkday">Дойный день</th>
                    <th class="cm-milk">Удой, кг</th>
                    <th class="cm-zhir">Жир, %</th>
                    <th class="cm-belok">Белок, %</th>
                    <!-- <th>Жир, кг</th>
                    <th>Белок, кг</th>
                    <th>Количество соматических клеток</th> -->
                </tr>
            </thead>
            <tbody class="cm-tablebody">
                <tr v-for="milking in cow_info" :key="milking.CheckDate">
                    <td>{{ milking.LactationNumber }}</td>
                    <td>{{dateConverter(milking.CheckDate)}}</td>
                    <td>{{ milking.MilkingDays || 'Нет информации'}}</td>
                    <td v-if="milking.Milk">{{ milking.Milk.toFixed(1) }}</td><td v-else>Нет информации</td>    
                    <td v-if="milking.Fat">{{ milking.Fat.toFixed(1) }}</td><td v-else>Нет информации</td>
                    <td v-if="milking.Protein">{{ milking.Protein.toFixed(1) }}</td><td v-else>Нет информации</td>
                    <!-- <td>{{ milking.FatRegard || 'Нет информации'}}</td>
                    <td>{{ milking.ProteinRegard || 'Нет информации'}}</td>
                    <td>{{ milking.SomaticNucCount || 'Нет информации'}}</td> -->
                </tr>
            </tbody>
        </table>
    </div>

    <div v-if="isChart">
        <div>
            <div class="nlact">
                <div class="chart-flex">
                    <div class="chart-nlact" >Номер лактации: </div>
                    <div class="check-lact">
                        <select v-model="check_lact" class="select-param">
                            <option v-for="nlact in count_lactations" :value="nlact" :key="nlact">Лактация {{ nlact }}</option>
                        </select>
                    </div>
                </div>

            </div>

            <div class="chart-flex">
                <div class="chart-param">Показатель: </div>
                <select v-model="param_milking" class="select-param">
                    <option value="Milk">Удой, кг</option>
                    <option value="FatRegard">Жир, кг</option>
                    <option value="ProteinRegard">Белок, кг</option>
                    <option value="Fat">Жир, %</option>
                    <option value="Protein">Белок, %</option>
                    <option value="DryMatter">Сухое вещество, %</option>
                </select>
            </div>

        </div>
        <apexchart ref="linechart" id="ControlMilking" width="600" type="line" :options="options" :series="series"></apexchart>
    </div>
</div>
<div v-if="isLoading">Идёт загрузка...</div>
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
            check_lact: 1,
            isLact: false,

            isLoading: false,
        }
    },
    async created() { 
        this.isLoading = true;
        let mass_route = this.$route.path.split('/');
        let cow_id = mass_route[2];
        let response = await fetch(`/api/cows/${cow_id}/checkMilks`);
        let result = await response.json();
        this.cow_info = result;
        
        for (let i = 0; i < this.cow_info.length; i++) {
            this.cow_info[i].FatRegard = ((this.cow_info[i].Fat * this.cow_info[i].Milk)/100).toFixed(1);
            this.cow_info[i].ProteinRegard = ((this.cow_info[i].Protein * this.cow_info[i].Milk)/100).toFixed(1);
            if (this.cow_info[i].LactationNumber == this.check_lact) {
                this.options.xaxis.categories.push(this.dateConverter(this.cow_info[i].CheckDate));
            }
        }

        try {
            this.count_lactations.push(this.cow_info[0]['LactationNumber']);
            for (let i = 1; i < this.cow_info.length; i++) {
                if (this.cow_info[i]['LactationNumber'] !== this.cow_info[i-1]['LactationNumber']) {
                    this.count_lactations.push(this.cow_info[i]['LactationNumber']);
                }
            }
            this.check_lact = this.count_lactations[0];
        } catch(err) {
            console.log(err);
        }

        this.addParam(this.cow_info, this.series, this.param_milking, this.check_lact)
        this.isLoading = false;
    },
    methods: {
        addInArr(obj, arr, param, nlact) {
            for (let i=0; i<obj.length; i++) {
                if(nlact === obj[i].LactationNumber) {
                    arr.push(obj[i][param]);
                }
            }
            console.log('параметры',arr);
        },
        addParam(obj, arr, param, nlact) {
            let serie = {
                name: `Лактация ${nlact}`,
                data: []
            };
            for (let j = 0; j < obj.length; j++) {
                if (obj[j].LactationNumber === nlact) {
                    let num = Math.round(obj[j][param] * 100)/100;
                    serie.data.push(num);
                }
            }
            arr.push(serie);
        },
        dateConverter(date) {
            let arr = date.split('-');
            let result = '';
            result += arr[2]; result += '.';
            result += arr[1]; result += '.';
            result += arr[0];
            return result;
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
            this.addParam(this.cow_info, this.series, this.param_milking, new_value);
            let newX = [];
            for (let i = 0; i < this.cow_info.length; i++) {
                if (this.cow_info[i].LactationNumber == new_value) {
                    newX.push(this.dateConverter(this.cow_info[i].CheckDate));
                }
            }
            this.$refs.linechart.updateOptions({
                xaxis: {
                    categories: newX,
                }
            });
            console.log(this.options.xaxis.categories);
        }

    }
}
</script>

<style scoped>
.cm-title {
    font-size: 130%;
    color: rgb(10, 113, 75);
    padding-bottom: 10px;
}

.parent-table {
    width: 49vw;
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

td {
    width: auto;
    min-width: 100px;
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
    border: 1px solid rgb(10, 113, 75);
    background-color: white;
    color: rgb(10, 113, 75);
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
    background-color: rgb(214, 239, 233);
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
    color: rgb(78, 145, 116)
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
    border: 1px solid rgb(31, 174, 122);
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