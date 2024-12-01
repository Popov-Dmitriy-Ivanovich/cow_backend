<template>
<div>
    <div class="cm-title">Ежедневные доения</div>
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
                    <th>Дата</th>
                    <th>Удой полный</th>
                    <th>Жир полный</th>
                    <th>Белок полный</th>
                </tr>
            </thead>
            <tbody class="cm-tablebody">
                <tr v-for="milking in cow_info" :key="milking.LactationId">
                    <td>{{ milking.LactationId }}</td>
                    <td>{{ milking.Date }}</td>
                    <td>{{ milking.Milk }}</td>
                    <td>{{ milking.Fat }}</td>
                    <td>{{ milking.Protein }}</td>
                </tr>
            </tbody>
        </table>
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
        }
    },
    async created() {
        let mass_route = this.$route.path.split('/');
        let cow_id = mass_route[2];
        let response = await fetch(`/api/cows/${cow_id}/checkMilks`);
        let result = await response.json();
        this.cow_info = result;
    }
}
</script>

<style scoped>

</style>