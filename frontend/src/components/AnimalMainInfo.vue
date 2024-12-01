<template>
<div class="main-info">
    <div class="cowname">{{ cow_info.NAME }}</div>
    <div class="pol"> | {{ getPol(cow_info.POL) }}</div>
    <div class="cow-microinfo">
        <div>Дата рождения: {{ cow_info.D_BIRTH}}</div> 
        <div>Номер РСХН: {{ cow_info.NRSHN }}</div>
    </div>

</div>
</template>

<script>
export default {
    data() {
        return {
            cow_info: {},
        }
    },
    async created() {
        let mass_route = this.$route.path.split('/');
        let cow_id = mass_route[2];
        let response = await fetch(`https://genmilk.ru:9050/api/cow_common?ID_COW=${cow_id}`);
        let result = await response.json();
        this.cow_info = result;
    },
    methods: {
        getPol(id_pol) {
            if (id_pol === 4) return 'Корова';
            if (id_pol === 3) return 'Бык';
            if (id_pol === 2) return 'Тёлка';
            if (id_pol === 1) return 'Бычок';
        }
    }
}
</script>

<style scoped>
.main-info {
    width: 80%;
    background-color: white;
    border-radius: 10px;
    padding: 30px 40px;
    margin-bottom: 30px;
    box-shadow: rgba(100, 100, 111, 0.1) 0px 7px 29px 0px;
    font-size: 120%; 
    position: sticky;
    top: 80px;
    z-index: 30;
}

.cowname {
    color: rgb(37, 0, 132);
    margin: 0 10px 5px 0;
    display: inline-block;
}

.pol {
    font-size: 90%;
    color: gray;
    display: inline-block;
}

.cow-microinfo {
    display: flex;
}

.cow-microinfo div {
    margin-right: 30px;
}
</style>