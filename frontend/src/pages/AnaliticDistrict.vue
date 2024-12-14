<template>
    <div>
        <div class="prev-chart" @click="toPrev">🠔 Назад</div>
        <ChartDistrict/>
    </div>
</template>

<script>
import ChartDistrict from '@/components/analyticsComponents/ChartDistrict.vue';

export default {
    components: {
        ChartDistrict
    },
    methods: {
        toPrev() {
            this.$router.push(`/analytics/${this.$route.params.id}/${this.$route.params.region}`);
        }
    }
}
</script>

<style scoped>
.prev-chart {
    font-family: Open Sans, sans-serif;
    margin-top: 30px;
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