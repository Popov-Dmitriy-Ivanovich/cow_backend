<template>
    <div class="lineage-title">Родословная</div>
    <div class="lineage-info">
        <LineageTree/>
        <hr class="lineage-sep">
        <ChildrenCow/>
    </div>
</template>

<script>
import ChildrenCow from '@/components/componentsConcretAnimal/ChildrenCow.vue';
import LineageTree from '@/components/componentsConcretAnimal/LineageTree.vue';

export default {
    components: {
        ChildrenCow, LineageTree
    }
}
</script>

<style scoped>
.lineage-title {
    font-size: 190%;
    margin-bottom: 30px;
}

.lineage-info {
    background-color: white;
    width: 800px;
    height: max-content;
    min-height: 500px;
    border-radius: 10px;
    box-shadow: rgba(100, 100, 111, 0.1) 0px 7px 29px 0px;
    padding: 30px 40px;
}

.lineage-sep {
    border: 1px solid rgb(224, 224, 224);
    margin-bottom: 40px;
}
</style>