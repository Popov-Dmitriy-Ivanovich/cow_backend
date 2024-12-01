<template>
    <div>
        <div class="id-title">Идентификаторы</div>
        <div class="idnum-flex">
            <div class="item-block">
                <div class="id-min-title">Инвентарный номер</div>
                <div>{{ cow_info.NINV }}</div>
            </div>
            <div class="item-block">
                <div class="id-min-title">ID Селекса</div>
                <div>{{ cow_info.NSELEKS }}</div>
            </div>
            <!-- <div class="item-block">
                <div class="id-min-title">UN</div>
                <div>test</div>
            </div> -->
            <div class="item-block">
                <div class="id-min-title">Хозяйство</div>
                <div>{{ cow_info.NAMEHOZ }}</div>
            </div>
            <div class="item-block">
                <div class="id-min-title">Ферма</div>
                <div>{{ cow_info.NAMEFARM }}</div>
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
    margin: 0 35px 30px 0;
}

.id-min-title {
    color: grey;
    margin-bottom: 7px;
}
</style>