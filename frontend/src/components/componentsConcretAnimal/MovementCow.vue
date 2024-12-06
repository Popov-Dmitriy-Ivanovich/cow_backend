<template>
    <div class="id-title">Перемещения</div>
    <div class="mov-flex">
        <div class="item-block">
            <div class="id-min-title">Хозяйство предыдущего прибывания</div>
            <div>{{ hozPrev || 'Нет информации' }}</div>
        </div>
        <div class="item-block">
            <div class="id-min-title">Хозяйство рождения</div>
            <div>{{ hozBirth || 'Нет информации' }}</div>
        </div>
    </div>


</template>

<script>
export default {
    props: {
        cow_info: {
            type: Object,
        }
    },
    data() {
        return {
            hozBirth: null,
            hozPrev: null,
        }
    },
    async create() {
        if(this.cow_info.BirthHozId) {
            let response = await fetch(`/api/farms/${this.cow_info.BirthHozId}`);
            let result = await response.json();
            this.hozBirth = result;
        }
        if (this.cow_info.PreviousHozId) {
            let response1 = await fetch(`/api/farm/${this.cow_info.PreviousHozId}`);
            let result1 = await response1.json();
            this.hozPrev = result1;
        }
    }
}
</script>

<style scoped>
.id-title {
    font-size: 130%;
    color: rgb(37, 0, 132);
    padding-bottom: 30px;
    width: max-content;
}

.mov-flex {
    display: flex;
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