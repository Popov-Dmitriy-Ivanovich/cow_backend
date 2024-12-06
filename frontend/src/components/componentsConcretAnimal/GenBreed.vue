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
                <div>{{ status || 'Нет информации' }}</div>
            </div>
            <div class="item-block">
                <div class="id-min-title">Дата генотипирования</div>
                <div>{{ gendate || 'Нет информации' }}</div>
            </div>
            <div class="item-block">
                <div class="id-min-title">Порода</div>
                <div>{{ cow_info.BreedName || 'Нет информации' }}</div>
            </div>
            <div class="item-block">
                <div class="id-min-title">Метод происхождения</div>
                <div>{{ cow_info.BirthMethod || 'Нет информации' }}</div>
            </div>
        </div>
    </div>
</template>
    
<script>
export default {
    props: {
        cow_info: {
            type: Object,
        },
        genetic: {
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
    computed: {
        status() {
            let a;
            if(this.genetic.ResultDate) a = 'Да';
            else a = 'Нет';
            return a
        },
        gendate() {
            let a;
            if(this.genetic.ResultDate) a = this.dateConverter(this.genetic.ResultDate);
            return a
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