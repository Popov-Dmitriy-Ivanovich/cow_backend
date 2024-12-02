<template>
    <div>
        <div class="id-title">Генотип и порода</div>
        <div class="idnum-flex">
            <div class="item-block">
                <div class="id-min-title">№ образца</div>
                <div>{{ genetic.ProbeNumber }}</div>
            </div>
            <div class="item-block">
                <div class="id-min-title">Статус генотипирования</div>
                <div>{{ genetic.status }}</div>
            </div>
            <div class="item-block">
                <div class="id-min-title">Дата генотипирования</div>
                <div>{{ genetic.ResultDate }}</div>
            </div>
            <div class="item-block">
                <div class="id-min-title">Порода</div>
                <div>{{ cow_info.BreedName }}</div>
            </div>
        </div>
    </div>
</template>
    
<script>
export default {
    data() {
        return {
            cow_info: {},
            genetic: {},
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
        let response = await fetch(`/api/cows/${cow_id}`);
        let result = await response.json();
        this.cow_info = result;
        let response1 = await fetch(`/api/cows/${cow_id}/genetic`);
        let result1 = await response1.json();
        this.genetic = result1;
        if(this.genetic.ResultDate) this.genetic.status = 'Да';
        else this.genetic.status = 'Нет'
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