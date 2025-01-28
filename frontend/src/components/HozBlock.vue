<template>
<div class="hoz">
    <div class="hozblock-title">Хозяйства</div>
    <div class="hoz-blocks">
        <div v-for="farm in farms" :key="farm[0]">
            <HozItem v-bind:farm_item="farm"/>
        </div> 
    </div>
    <div class="show-hoz" @click="$router.push('/hoz')">Показать все</div>
</div>
</template>

<script>
import HozItem from '@/components/HozItem.vue';

export default {
    components: {
        HozItem,
    },
    data() {
        return {
            farms: [],
        }
    },
    async created() {
        const response = await fetch('/api/farms/hoz');
        const res_farms = await response.json();
        for (let i = 0; i < 3; i ++) {
            this.farms.push(res_farms[i]);
        }
    }
}
</script>

<style scoped>
.hoz {
    height: 600px;
    text-align: center;
    font-family: Open Sans, sans-serif;
    color: rgb(10, 113, 75);
}

.hozblock-title {
    font-size: 190%;
    padding: 70px 0 10px 0;
}

.hoz-blocks {
    display: flex;
    justify-content: space-around;
    margin: 20px;
}

.show-hoz {
    cursor: pointer;
    transition: 0.3s;
}

.show-hoz:hover {
    color:rgb(49, 201, 145);
}
</style>