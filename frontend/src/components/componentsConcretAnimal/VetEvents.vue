<template>
    <div>
        <div class="vet-title">Ветеринарные мероприятия</div>
        <div class="parent-table">
            <table class="vet-table">
                <thead>
                    <tr class="vet-header">
                        <th>Дата</th>
                        <th>Название</th>
                        <th>Дойные дни</th>
                        <th>Примечание 1</th>
                        <th>Примечание 2</th>
                    </tr>
                </thead>
                <tbody class="vet-tablebody">
                    <tr>
                        <td>test</td>
                        <td>test</td>
                        <td>test</td>
                        <td>test</td>
                        <td>test</td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>
    
<script>
    
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