<template>
    <div class="doc-title">Документы</div>
    <div class="doc-info">
        <DocumentCow/>
    </div>
</template>

<script>
import DocumentCow from './componentsConcretAnimal/DocumentCow.vue';

export default {
    components: {
        DocumentCow
    }
}
</script>

<style scoped>
.doc-title {
    font-size: 190%;
    margin-bottom: 30px;
}

.doc-info {
    background-color: white;
    width: 56vw;
    min-width: 800px;
    height: max-content;
    min-height: 500px;
    border-radius: 10px;
    box-shadow: rgba(100, 100, 111, 0.1) 0px 7px 29px 0px;
    padding: 30px 40px;
}
</style>