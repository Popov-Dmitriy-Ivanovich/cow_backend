<template>
    <div class="id-title">Перемещения</div>
    <div class="mov-flex">
        <div class="item-block">
            <div class="id-min-title">Хозяйство предыдущего пребывания</div>
            <div v-if="hozPrev">{{ hozPrev.Name || 'Нет информации' }}</div>
            <div v-else>Нет информации</div>
        </div>
        <div class="item-block">
            <div class="id-min-title">Хозяйство рождения</div>
            <div v-if="hozBirth">{{ hozBirth.Name || 'Нет информации' }}</div>
            <div v-else>Нет информации</div>
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
    async created() {
        if(this.cow_info.BirthHozId) {
            let response = await fetch(`/api/farms/${this.cow_info.BirthHozId}`);
            let result = await response.json();
            this.hozBirth = result;
        }
        if (this.cow_info.PreviousHozId) {
            let response1 = await fetch(`/api/farms/${this.cow_info.PreviousHozId}`);
            let result1 = await response1.json();
            this.hozPrev = result1;
        }
    },
    watch: {
        async cow_info(new_val) {
            if(new_val.BirthHozId) {
                let response = await fetch(`/api/farms/${new_val.BirthHozId}`);
                let result = await response.json();
                this.hozBirth = result;
            }
            if (new_val.PreviousHozId) {
                let response1 = await fetch(`/api/farms/${new_val.PreviousHozId}`);
                let result1 = await response1.json();
                this.hozPrev = result1;
            }
        }
    }
}
</script>

<style scoped>
.id-title {
    font-size: 130%;
    color: rgb(10, 113, 75);
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