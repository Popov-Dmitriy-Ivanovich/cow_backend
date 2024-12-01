<template>
    <div>
        <div class="child-title">Потомки</div>
        <div class="n-children">Количество потомков: </div>
        <div class="parent-table">
            <table class="child-table">
                <thead>
                    <tr class="child-header">
                        <th>UN</th>
                        <th>Уровень</th>
                        <th>Дата рождения</th>
                        <th>Кличка</th>
                        <th>Селекционный индекс</th>
                        <th>Молочный индекс</th>
                    </tr>
                </thead>
                <tbody class="child-tablebody">
                    <tr>
                        <td>test</td>
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
.child-title {
    font-size: 130%;
    color: rgb(37, 0, 132);
    padding-bottom: 30px;
    width: max-content;
}

.n-children {
    padding-bottom: 20px;
}

.parent-table {
    width: 50vw;
    overflow-x: auto;
}

.child-table {
    margin-bottom: 30px;
    text-align: left;

}

th {
    font-weight: normal;
}

.child-header {
    color: grey;
}

.child-header th {
    padding-right: 30px;
    padding-bottom: 15px;
}

.child-tablebody {
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