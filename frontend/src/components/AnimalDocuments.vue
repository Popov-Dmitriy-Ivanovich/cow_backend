<template>
    <div class="doc-title">Документы</div>
    <div class="doc-info">
        <DownloadDocument @changeUspeh="Uspeh"/>
        <hr class="doc-sep">
        <DocumentCow v-bind:uspeh="uspeh"/>
    </div>
</template>

<script>
import DocumentCow from './componentsConcretAnimal/DocumentCow.vue';
import DownloadDocument from '@/components/componentsConcretAnimal/DownloadDocument.vue'

export default {
    data() {
        return {
            uspeh: false,
        }
    },
    components: {
        DocumentCow, DownloadDocument
    },
    methods: {
        Uspeh(data) {
            this.uspeh = data;
        }
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

.doc-sep {
    border: 1px solid rgb(224, 224, 224);
    margin-bottom: 40px;
}
</style>