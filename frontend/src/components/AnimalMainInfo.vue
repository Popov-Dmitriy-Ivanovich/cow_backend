<template>
<div class="main-info" v-if="!isLoading">
    <div class="cowname">{{ cow_info.Name || 'Нет информации'}}</div>
    <div class="pol"> | {{ cow_info.SexName  || 'Нет информации' }}</div>
    <div class="cow-microinfo">
        <div class="bdate">Дата рождения: {{ cow_info.BirthDate  || 'Нет информации'}}</div> 
        <div class="rshn">Номер РСХН: {{ cow_info.RSHNNumber  || 'Нет информации'}}</div>
        <div class="pol">| {{ status }}</div>
    </div> 
</div>
<div class="main-info" v-else>Идёт загрузка...</div>
</template>

<script>
export default {
    data() {
        return {
            cow_info: {},
            status:"",

            isLoading: false,
        }
    },
    async created() {
        this.isLoading = true;
        let mass_route = this.$route.path.split('/');
        let cow_id = mass_route[2];
        await this.fetchInfo(cow_id);
        this.isLoading = false;
    },
    methods: {
        getPol(id_pol) {
            if (id_pol === 4) return 'Корова';
            if (id_pol === 3) return 'Бык';
            if (id_pol === 2) return 'Тёлка';
            if (id_pol === 1) return 'Бычок';
        },
        dateConverter(date) {
            let arr = date.split('-');
            let result = '';
            result += arr[2]; result += '.';
            result += arr[1]; result += '.';
            result += arr[0];
            return result;
        },
        async fetchInfo(param) {
            let response = await fetch(`/api/cows/${param}`);
            let result = await response.json();
            this.cow_info = result;
            if (this.cow_info.BirthDate) this.cow_info.BirthDate = this.dateConverter(this.cow_info.BirthDate);
            if (this.cow_info.DepartDate) {
                if(this.cow_info.SexId === 2 || this.cow_info.SexId === 4) this.status = 'Выбыла'
                else this.status = 'Выбыл'
            } else {
                if(this.cow_info.SexId === 2 || this.cow_info.SexId === 4) this.status = 'Не выбыла'
                else this.status = 'Не выбыл'
            }
        }
    },
    watch: {
        $route(new_val) {
            this.fetchInfo(new_val.params.id);
        }
    }
}
</script>

<style scoped>
.main-info {
    width: 56vw;
    min-width: 800px;
    background-color: white;
    border-radius: 10px;
    padding: 30px 40px;
    margin-bottom: 30px;
    box-shadow: rgba(100, 100, 111, 0.1) 0px 7px 29px 0px;
    font-size: 120%; 
    position: sticky;
    top: 100px;
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

.bdate {
    margin-right: 25px;
}

.rshn {
    margin-right: 10px;
}
</style>