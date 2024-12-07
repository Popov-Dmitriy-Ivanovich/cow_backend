<template>
<div>
    <div class="prev-chart" @click="toPrev">🠔 Назад</div>
    <ChartHolding/>
</div>
</template>

<script>
import ChartHolding from '@/components/analyticsComponents/ChartHolding.vue';

export default {
    components: {
        ChartHolding
    },
    methods: {
        toPrev() {
            this.$router.push(`/analytics/${this.$route.params.id}/${this.$route.params.region}/${this.$route.params.district}`);
        }
    }
}
</script>

<style scoped>
.prev-chart {
    font-family: Open Sans, sans-serif;
    margin-top: 130px;
    margin-left: 20px;
    color:rgb(37, 0, 132);
    padding-bottom: 20px;
    cursor: pointer;
    transition: 0.3s;
}

.prev-chart:hover {
    color: rgb(103, 63, 205);
    padding-left: 10px;
    width: max-content;
}
</style>