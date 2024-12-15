<template>
<div>
    <table class="lac-table">
        <thead>
            <tr class="lac-header">
                <th>Группа события</th>
                <th>Событие</th>
                <th>Подвид события</th>
                <th>Дата</th>
                <th>Дней с начала лактации</th>
            </tr>
        </thead>
        <tbody class="lac-tablebody">
            <tr v-for="lact in cow_info" :key="lact.Number">
                <td v-if="lact.EventType">{{ lact.EventType.Name || 'Нет информации'}}</td><td v-else>Нет информации</td>
                <td v-if="lact.EventType1">{{ lact.EventType1.Name || 'Нет информации'}}</td><td v-else>Нет информации</td>
                <td v-if="lact.EventType2">{{ lact.EventType2.Name || 'Нет информации'}}</td><td v-else>Нет информации</td>
                <td v-if="lact.Date">{{ dateConverter(lact.Date) || 'Нет информации'}}</td><td v-else>Нет информации</td>
                <td v-if="lact.DaysFromLactation">{{ lact.DaysFromLactation || 'Нет информации'}}</td><td v-else>Нет информации</td>
            </tr>
        </tbody>
    </table>
</div>
</template>

<script>
export default {
    data() {
        return {
            cow_info: {},
        }
    },
    async mounted() {
        let response = await fetch(`/api/cows/${this.$route.params.id}/health`);
        let result = await response.json();
        this.cow_info = result;
        console.log(this.cow_info);
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
    padding-right: 7px;
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