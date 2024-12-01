<template>
    <div class="common-title">Общая информация</div>
    <div class="general-info">
        <ID/>
        <hr class="com-sep">
        <GenBreed/>
        <hr class="com-sep">
        <ParentsCow/>
    </div>
</template>

<script>
import ID from '@/components/componentsConcretAnimal/ID.vue';
import GenBreed from '@/components/componentsConcretAnimal/GenBreed.vue';
import ParentsCow from '@/components/componentsConcretAnimal/ParentsCow.vue';

export default {
    components: {
        ID, GenBreed, ParentsCow
    }
}
</script>

<style scoped>
.common-title {
    font-size: 190%;
    margin-bottom: 30px;
}

.general-info {
    background-color: white;
    width: 80%;
    height: max-content;
    min-height: 500px;
    border-radius: 10px;
    box-shadow: rgba(100, 100, 111, 0.1) 0px 7px 29px 0px;
    padding: 30px 40px;
}

.com-sep {
    border: 1px solid rgb(224, 224, 224);
    margin-bottom: 40px;
}
</style>