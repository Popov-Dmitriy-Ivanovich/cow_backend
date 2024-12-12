<template>
    <div class="analytics-flex">
        <div class="sub-block"></div>
        <div>
            <div class="prev-chart" @click="toPrev">🠔 Назад</div>
            <ChartYear/>
        </div>
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

.sub-block {
    width: 320px;
    background: none;
    height: 20px;
}
.analytics-flex {
    display: flex;
    justify-content: space-around;
    width: 90vw;
}
</style>