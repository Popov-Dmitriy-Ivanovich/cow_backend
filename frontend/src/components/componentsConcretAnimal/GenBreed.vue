<template>
    <div>
        <div class="id-title">Генотип и порода</div>
        <div class="idnum-flex">
            <div class="item-block">
                <div class="id-min-title">№ образца</div>
                <div>{{ genetic.ProbeNumber || 'Нет информации' }}</div>
            </div>
            <div class="item-block">
                <div class="id-min-title">Статус генотипирования</div>
                <div>{{ genetic.status || 'Нет информации' }}</div>
            </div>
            <div class="item-block">
                <div class="id-min-title">Дата генотипирования</div>
                <div>{{ gendate || 'Нет информации' }}</div>
            </div>
            <div class="item-block">
                <div class="id-min-title">Порода</div>
                <div>{{ cow_info.BreedName || 'Нет информации' }}</div>
            </div>
        </div>
    </div>
</template>
    
<script>
export default {
    data() {
        return {
            genetic: {},
            gendate: '',
        }
    },
    props: {
        cow_info: {
            type: Object,
        }
    },
    methods: {
        true_false(val) {
            if(val) return 'Да';
            else return 'Нет';
        },
        dateConverter(date) {
            let arr = date.split('-');
            let result = '';
            result += arr[2]; result += '.';
            result += arr[1]; result += '.';
            result += arr[0];
            return result;
        }
    },
    async created() {
        let mass_route = this.$route.path.split('/');
        let cow_id = mass_route[2];
        let response1 = await fetch(`/api/cows/${cow_id}/genetic`);
        let result1 = await response1.json();
        this.genetic = result1 || {};
        console.log(this.genetic);
        if(this.genetic.ResultDate) {
            this.genetic.status = 'Да';
            this.gendate = this.dateConverter(this.genetic.ResultDate);
        } 
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