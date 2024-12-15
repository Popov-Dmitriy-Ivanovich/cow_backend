<template>
<div class="hoz-page">
    <div @click="$router.push('/')" class="link-to-main">🠔 Главная</div>
    <div class="hozpage-title">Хозяйства</div>
    <div class="list-hoz">
        <div v-for="farm in farms" :key="farm[0]">
            <HozItem v-bind:farm_item="farm" class="hoz-item"/>
        </div>
        
    </div>
</div>
</template>

<script>
import HozItem from '@/components/HozItem.vue';

export default {
    components: {
        HozItem,
    },
    data () {
        return {
            farms: [],
        }
    },
    async created() {
        const response = await fetch ('/api/farms/hoz');
        const res_farms = await response.json();
        this.farms = res_farms;
    }
}
</script>

<style>
.hoz-page {
    margin-top: 110px;
    width: 100%;
    font-family: Open Sans, sans-serif;
    color: rgb(37, 0, 132);
    display: flex;
    flex-direction: column;
    align-items: center;
}

.link-to-main {
    align-self: flex-start;
    cursor: pointer;
    margin: 10px 20px;
    transition: 0.3s;
}

.link-to-main:hover {
    color: rgb(103, 63, 205);
    padding-left: 10px;
    width: max-content;
}

.hozpage-title {
    font-size: 210%;
    padding-bottom: 60px;
}

.list-hoz {
    width: 100%;
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
}

.hoz-item {
    margin: 15px
}
</style>