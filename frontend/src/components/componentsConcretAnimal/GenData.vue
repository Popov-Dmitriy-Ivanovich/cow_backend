<template>
    <div>
        <div class="datagen-title">Данные о генотипировании</div>
        <div>Файл</div>
        <div class="datagen-download">    
            <div class="download-file"></div>
            <div class="download-btn">Скачать файл</div>
        </div>
        <table class="genfile-table">
                <thead>
                    <tr class="genfile-header">
                        <th>Дата загрузки</th>
                        <th>№ образца</th>
                        <th>Лаборатория</th>
                    </tr>
                </thead>
                <tbody class="genfile-tablebody">
                    <tr>
                        <td>test</td>
                        <td>test</td>
                        <td>test</td>
                    </tr>
                </tbody>
            </table>
    </div>
</template>
    
<script>
    
</script>
    
<style scoped>
.datagen-title {
    font-size: 130%;
    color: rgb(37, 0, 132);
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
    background-color: rgb(241, 240, 246);
    margin: 10px 20px 0 0;
}

.download-btn {
    color: rgb(37, 0, 132);
    cursor: pointer;
}

.genfile-table {
    margin: 30px 0;
    text-align: left;
    
}
    
th {
    font-weight: normal;
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
</style>