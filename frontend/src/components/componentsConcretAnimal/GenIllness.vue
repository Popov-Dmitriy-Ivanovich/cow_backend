<template>
    <div>
        <div class="vet-title">Моногенные заболевания</div>
        <div class="parent-table">
            <table class="vet-table">
                <thead>
                    <tr class="vet-header">
                        <th>Патология</th>
                        <th>Статус</th>
                        <th>Описание</th>
                    </tr>
                </thead>
                <tbody class="vet-tablebody">
                    <tr v-for="(value, name) in cow_info" :key="name">
                        <td>{{ name }}</td>
                        <td>{{ value }}</td>
                        <td></td>
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
            cow_info: {},
        }
    },
    async created() {
        let mass_route = this.$route.path.split('/');
        let cow_id = mass_route[2];
        let response = await fetch(`https://genmilk.ru:9050/api/cow_health?ID_COW=${cow_id}`);
        let result = await response.json();
        this.cow_info = result;
    },
}    
</script>
    
<style scoped>
.vet-title {
    font-size: 130%;
    color: rgb(37, 0, 132);
    padding-bottom: 30px;
    width: max-content;
}
    
.parent-table {
    width: 50vw;
    overflow-x: auto;
}
    
.vet-table {
    margin-bottom: 30px;
    text-align: left;
    
}
    
th {
    font-weight: normal;
}
    
.vet-header {
    color: grey;
}
    
.vet-header th {
    padding-right: 30px;
    padding-bottom: 15px;
}
    
.vet-tablebody {
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
</style>