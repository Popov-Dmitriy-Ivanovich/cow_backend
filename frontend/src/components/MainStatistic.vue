<template>
<div class="statistic">
    <div class="statistic-title">Статистика по генотипированию</div>
    <div v-for="item in stat" :key="item.name" class="statis-text">
        <div>{{ item.name }}</div>
        <div>{{ item.value }} животных (~{{ item.regard }}% от общего количества)</div>
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
        let obj = {};
        console.log(localStorage.getItem('jwt'));
        const response = await fetch('/api/analitics/genotyped/40000/byRegion/36/districts',
        {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json;charset=utf-8',
                'Authorization': localStorage.getItem('jwt')
            },
            body: JSON.stringify(obj),
        });
        const result = await response.json();
        if (result.error) {
            this.stat.push({name: 'Ошибка доступа', value: '-', regard: '-'});
        } else {
            for( let key in result) {
            let val = Math.round(((result[key].Genotyped / result[key].Alive)*100)*100)/100;
            let item = {name: key + ' район', value: result[key].Genotyped, regard:val};
            this.stat.push(item);
        }
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
    width: 50%;
    display: flex;
    justify-content: space-between;
}

.statis-text div {
    padding: 7px 0;
}
</style>