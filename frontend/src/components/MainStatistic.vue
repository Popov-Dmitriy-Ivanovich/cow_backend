<template>
<div class="statistic">
    <div class="statistic-title">Статистика по генотипированию</div>
    <div v-for="item in stat" :key="item.name" class="statis-text">
        <div>{{ item.name }}</div>
        <div>{{ item.value }} животных</div>
    </div>
</div>
</template>

<script>
export default {
    data() {
        return {
            stat: [],
        }
    },
    async created() {
        this.stat = [];
        const response = await fetch('/api/analitics/genotyped/40000/regions');
        const result = await response.json();
        for( let key in result) {
            let item = {name: key, value: result[key].Genotyped}
            this.stat.push(item);
        }
    }
}
</script>

<style scoped>
.statistic {
    width: 100%;
    background-color: white;
    margin: 50px 0;
    font-family: Open Sans, sans-serif;
    padding: 40px 0;
    display: flex;
    flex-direction: column;
    align-items: center;
}

.statistic-title {
    font-size: 190%;
    padding: 0 0 20px 0;
    color: rgb(37, 0, 132);
}

.statis-text {
    width: 40%;
    display: flex;
    justify-content: space-between;
}

.statis-text div {
    padding: 7px 0;
}
</style>