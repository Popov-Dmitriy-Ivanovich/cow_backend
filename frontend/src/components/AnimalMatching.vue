<template>
    <div class="match-title">Подбор пар</div>
    <div class="general-info">
        <MatchingParams/>
    </div>
</template>

<script>
import MatchingParams from './componentsConcretAnimal/MatchingParams.vue';

export default {
    components: {
        MatchingParams
    }
}
</script>

<style scoped>
.match-title {
    font-size: 190%;
    margin-bottom: 30px;
}

.general-info {
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