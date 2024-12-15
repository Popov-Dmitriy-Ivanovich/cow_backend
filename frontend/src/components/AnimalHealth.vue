<template>
    <div class="health-title">Ветеринарные мероприятия</div>
    <div class="health-info">
        <HealthCow/>
    </div>
</template>

<script>
import HealthCow from './componentsConcretAnimal/HealthCow.vue';



export default {
    components: {
        HealthCow
    }
}
</script>

<style scoped>
.health-title {
    font-size: 190%;
    margin-bottom: 30px;
}

.health-info {
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