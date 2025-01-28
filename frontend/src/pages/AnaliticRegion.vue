<template>
    <div>
        <div class="prev-chart" @click="toPrev">🠔 Назад</div>
        <ChartRegion/>
    </div>
</template>

<script>
import ChartRegion from '@/components/analyticsComponents/ChartRegion.vue';

export default {
    components: {
        ChartRegion
    },
    methods: {
        toPrev() {
            this.$router.push(`/analytics/${this.$route.params.id}`);
        }
    }
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
    color: rgb(63, 205, 120);
    padding-left: 10px;
    width: max-content;
}

</style>