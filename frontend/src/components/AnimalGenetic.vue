<template>
    <div class="genetic-title">Генетика</div>
    <div class="genetic-info">
        <GenData/>
        <hr class="gen-sep">
        <GenIllness/>
    </div>
</template>

<script>
import GenData from '@/components/componentsConcretAnimal/GenData.vue';
import GenIllness from './componentsConcretAnimal/GenIllness.vue';

export default {
    components: {
        GenData, GenIllness
    }
}
</script>

<style scoped>
.genetic-title {
    font-size: 190%;
    margin-bottom: 30px;
}

.genetic-info {
    background-color: white;
    width: 56vw;
    min-width: 800px;
    height: max-content;
    min-height: 500px;
    border-radius: 10px;
    box-shadow: rgba(100, 100, 111, 0.1) 0px 7px 29px 0px;
    padding: 30px 40px;
}

.gen-sep {
    border: 1px solid rgb(224, 224, 224);
    margin-bottom: 40px;
}
</style>