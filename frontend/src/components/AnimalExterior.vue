<template>
    <div class="exterior-title">Экстерьер</div>
    <div class="exterior-info">
        <ExteriorRating/>
        <hr class="exterior-sep">
        <VetEvents/>
    </div>
</template>

<script>
import VetEvents from '@/components/componentsConcretAnimal/VetEvents.vue';
import ExteriorRating from '@/components/componentsConcretAnimal/ExteriorRating.vue'

export default {
    components: {
        VetEvents, ExteriorRating
    }
}
</script>

<style scoped>
.exterior-title {
    font-size: 190%;
    margin-bottom: 30px;
}

.exterior-info {
    background-color: white;
    width: 80%;
    height: max-content;
    min-height: 500px;
    border-radius: 10px;
    box-shadow: rgba(100, 100, 111, 0.1) 0px 7px 29px 0px;
    padding: 30px 40px;
}

.exterior-sep {
    border: 1px solid rgb(224, 224, 224);
    margin-bottom: 40px;
}
</style>