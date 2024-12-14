<template>
<div>
    <div class="doc-title">Загрузка документов</div>
    <form class="download-doc-form" action="https://genmilk.ru/api/load/document" method="post" enctype="multipart/form-data">
        <input type="file" id="Document" name="Document">
        <input type="text" id="CowID" name="CowID" required="" :value="$route.params.id">
        <button type="submit">Загрузить</button>
    </form>

</div>
</template>

<script>
export default {
    
}
</script>

<style scoped>
.doc-title {
    font-size: 130%;
    color: rgb(37, 0, 132);
    padding-bottom: 10px;
}
.download-doc-form {
    margin: 20px 0;
}

#CowID {
    display: none;
}
</style>