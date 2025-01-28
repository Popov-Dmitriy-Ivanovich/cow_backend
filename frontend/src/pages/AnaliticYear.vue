<template>
    <div>
        <div class="prev-chart" @click="toPrev">🠔 Назад</div>
        <ChartYear/>
    </div>
</template>

<script>
import ChartYear from '@/components/analyticsComponents/ChartYear.vue';

export default {
    components: {
        ChartYear
    },
    methods: {
        toPrev() {
            this.$router.push(`/analytics`);
        },
    },
}
</script>

<style scoped>
.prev-chart {
    font-family: Open Sans, sans-serif;
    margin-top: 30px;
    margin-left: 20px;
    color:rgb(10, 113, 75);
    padding-bottom: 20px;
    cursor: pointer;
    transition: 0.3s;
}

.prev-chart:hover {
    color: rgb(63, 205, 132);
    padding-left: 10px;
    width: max-content;
}

</style>