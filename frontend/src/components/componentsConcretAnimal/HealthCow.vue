<template>
<div class="parent-table" v-if="!isLoading">
    <table class="lac-table">
        <thead>
            <tr class="lac-header">
                <th>Дата</th>
                <th>Группа события</th>
                <th>Событие</th>
                <th>Разновидность события</th>
                <th>Дней с начала лактации</th>
            </tr>
        </thead>
        <tbody class="lac-tablebody">
            <tr v-for="lact in cow_info" :key="lact.Number">
                <td v-if="lact.Date">{{ dateConverter(lact.Date) || 'Нет информации'}}</td><td v-else>Нет информации</td>
                <td v-if="lact.EventType">{{ lact.EventType.Name || 'Нет информации'}}</td><td v-else>Нет информации</td>
                <td v-if="lact.EventType1">{{ lact.EventType1.Name || 'Нет информации'}}</td><td v-else>Нет информации</td>
                <td v-if="lact.EventType2">{{ lact.EventType2.Name || 'Нет информации'}}</td><td v-else>Нет информации</td>
                <td v-if="lact.DaysFromLactation">{{ lact.DaysFromLactation || 'Нет информации'}}</td><td v-else>Нет информации</td>
            </tr>
        </tbody>
    </table>
</div>
<div class="parent-table" v-else>Идёт загрузка...</div>
</template>

<script>
export default {
    data() {
        return {
            cow_info: {},
            isLoading: false,
        }
    },
    async mounted() {
        this.isLoading = true;
        let response = await fetch(`/api/cows/${this.$route.params.id}/health`);
        let result = await response.json();
        this.cow_info = result;
        this.isLoading = false;
    },
    methods: {
        dateConverter(date) {
            let arr = date.split('-');
            let result = '';
            result += arr[2]; result += '.';
            result += arr[1]; result += '.';
            result += arr[0];
            return result;
        },
    }
}
</script>

<style scoped>
.parent-table {
    width: 49vw;
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
    padding-right: 15px;
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
</style>