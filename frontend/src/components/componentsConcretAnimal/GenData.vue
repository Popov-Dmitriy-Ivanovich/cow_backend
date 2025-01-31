<template>
    <div class="datagen-title">Данные о генотипировании</div>
    <div v-if="!isLoading">
        <div>Файл</div>
        <div class="datagen-download">    
            <div class="download-file">{{ cow_info.GtcFilePath || 'Нет информации'}} </div>
            <div class="download-btn" v-if="cow_info.GtcFilePath"><a :href="`/api/static/gtc/${cow_info.GtcFilePath}`" :download="cow_info.GtcFilePath" class="download-gtc">Скачать файл</a></div>
        </div>
        <table class="genfile-table">
                <thead>
                    <tr class="genfile-header">
                        <th>Факт генотипирования</th>
                        <th>№ образца</th>
                        <th>Дата отбора образца</th>
                    </tr>
                </thead>
                <tbody class="genfile-tablebody">
                    <tr>
                        <td>{{true_false(cow_info.ResultDate) || 'Нет информации'}}</td>
                        <td>{{cow_info.ProbeNumber || 'Нет информации'}}</td>
                        <td>{{ blood_date || 'Нет информации'}}</td>
                    </tr>
                </tbody>
            </table>
    </div>
    <div v-if="isLoading">Идёт загрузка...</div>
</template>
    
<script>
export default {
    data() {
        return {
            cow_info:{},
            blood_date: null,
            isLoading: false,
        }
    },
    async created() {
        this.isLoading = true;
        let mass_route = this.$route.path.split('/');
        let cow_id = mass_route[2];
        let response = await fetch(`/api/cows/${cow_id}/genetic`);
        let result = await response.json();
        console.log(result);
        if (result) {
            this.cow_info = result;
            if (this.cow_info.BloodDate) this.blood_date = this.dateConverter(this.cow_info.BloodDate);
        }
        this.isLoading = false;
    },
    methods: {
        true_false(val) {
            if(val) return 'Да';
            else return 'Нет';
        },
        dateConverter(date) {
            let arr = date.split('-');
            let result = '';
            result += arr[2]; result += '.';
            result += arr[1]; result += '.';
            result += arr[0];
            return result;
        }
    }
}
</script>
    
<style scoped>
.datagen-title {
    font-size: 130%;
    color: rgb(10, 113, 75);
    padding-bottom: 30px;
    width: max-content;
}

.datagen-download {
    display: flex;
    align-items: flex-end;
}

.download-file {
    width: 70%;
    height: 25px;
    background-color: rgb(240, 246, 244);
    margin: 10px 20px 0 0;
    padding: 7px 0 0 7px;
}

.download-btn {
    color: rgb(10, 113, 75);
    cursor: pointer;
}

.genfile-table {
    margin: 30px 0;
    text-align: left;
    
}
    
th {
    font-weight: normal;
}

td {
    padding-right: 20px;
}
    
.genfile-header {
    color: grey;
}
    
.genfile-header th {
    padding-right: 30px;
    padding-bottom: 5px;
}
    
.genfile-tablebody {
    text-align: left;
}

.download-gtc {
    color: rgb(10, 113, 75);
    text-decoration: none;
}
</style>