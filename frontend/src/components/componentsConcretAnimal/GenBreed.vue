<template>
    <div>
        <div class="id-title">Генотип и порода</div>
        <div class="idnum-flex">
            <div class="item-block">
                <div class="id-min-title">№ образца</div>
                <div>{{ cow_info.GEN_N }}</div>
            </div>
            <div class="item-block">
                <div class="id-min-title">Статус генотипирования</div>
                <div>{{ true_false(cow_info.GEN_STAT) }}</div>
            </div>
            <div class="item-block">
                <div class="id-min-title">Дата генотипирования</div>
                <div>{{  cow_info.D_GEN }}</div>
            </div>
            <div class="item-block">
                <div class="id-min-title">Порода</div>
                <div>{{ cow_info.BREED }}</div>
            </div>
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
    methods: {
        true_false(val) {
            if(val) return 'Да';
            else return 'Нет';
        }
    },
    async created() {
        let mass_route = this.$route.path.split('/');
        let cow_id = mass_route[2];
        let response = await fetch(`https://genmilk.ru:9050/api/cow_common?ID_COW=${cow_id}`);
        let result = await response.json();
        this.cow_info = result;
    },
}
</script>

<style scoped>
.id-title {
    font-size: 130%;
    color: rgb(37, 0, 132);
    padding-bottom: 30px;
    width: max-content;
}

.idnum-flex {
    display: flex;
    flex-wrap: wrap;
}

.item-block {
    width: max-content;
    margin: 0 25px 30px 0;
}

.id-min-title {
    color: grey;
    margin-bottom: 7px;
}
</style>