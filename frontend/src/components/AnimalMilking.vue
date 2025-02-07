<template>
    <div class="milking-title">Контрольные дойки и лактации</div>
    <div class="milking-info">
        <ControlMilking/>
        <hr class="milking-sep">
        <LactCow/>
    </div>
</template>

<script>
import ControlMilking from '@/components/componentsConcretAnimal/ControlMilking.vue';
import LactCow from '@/components/componentsConcretAnimal/LactCow.vue'

export default {
    components: {
        ControlMilking, LactCow
    }
}
</script>

<style scoped>
.milking-title {
    font-size: 190%;
    margin-bottom: 30px;
}

.milking-info {
    background-color: white;
    width: 56vw;
    min-width: 800px;
    height: max-content;
    min-height: 500px;
    border-radius: 10px;
    box-shadow: rgba(100, 100, 111, 0.1) 0px 7px 29px 0px;
    padding: 30px 40px;
}

.milking-sep {
    border: 1px solid rgb(224, 224, 224);
    margin-bottom: 40px;
}
</style>