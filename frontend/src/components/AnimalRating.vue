<template>
    <div class="rating-title">Оценки</div>
    <div class="rating-info">
        <RatingCow/>
    </div>
</template>

<script>
import RatingCow from '@/components/componentsConcretAnimal/RatingCow.vue'

export default {
    components: {
        RatingCow
    }
}
</script>

<style scoped>
.rating-title {
    font-size: 190%;
    margin-bottom: 30px;
}

.rating-info {
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