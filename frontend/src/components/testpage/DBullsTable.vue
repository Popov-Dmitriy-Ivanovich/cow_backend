<template>
    <div class="parent-table">
        <div class="animals-table">
            <div class="table-header">
                <div class="animal-rshn">Номер РСХН</div>
                <div class="animal-inv">Инвентарный номер</div>
                <div class="animal-name">Кличка</div>
                <div class="animal-hoz">Хозяйство</div>
                <div class="animal-bdate">Дата рождения</div>
                <div class="animal-genfact">Факт генотипирования</div>
                <div v-if="filters.D_OUT_BEGIN || filters.D_OUT_END" class="animal-dateout">Дата выбытия</div>
                <div v-if="filters.ID_BREED" class="animal-breed">Порода</div>
                <div v-if="filters.D_GEN_BEGIN || filters.D_GEN_END" class="animal-dategen">Дата генотипирования</div>
                <div v-if="filters.D_MILKING_BEGIN || filters.D_MILKING_END" class="animal-datemilking">Контрольная дойка</div>
                <div v-if="filters.EXTERIOR" class="animal-exterior">Оценка экстерьера</div>
                <div v-if="filters.D_OSEM_BEGIN || filters.D_OSEM_END" class="animal-dateosem">Дата осеменения</div>
                <div v-if="filters.D_OTEL_BEGIN || filters.D_OTEL_END" class="animal-dateotel">Дата отела</div>
                <div v-if="filters.D_BIRKING_BEGIN || filters.D_BIRKING_END" class="animal-datebirk">Дата перебирковки</div>
                <div v-if="filters.K_INBR_ROD_BEGIN || filters.K_INBR_ROD_END" class="animal-krod">Коэффициент инбридинга по родословной</div>
                <div v-if="filters.K_INBR_FEN_BEGIN || filters.K_INBR_FEN_END" class="animal-kfen">Коэффициент инбридинга по фенотипу</div>
            </div>
            <div v-if="!isSearch&!search_error">
                <div v-for="animal in animals" :key="animal[0]">
                    <DConcretAnimalLine v-bind:animal_item="animal" v-bind:filters="filters"/>
                </div>
            </div>
            <div v-if="isSearch&!search_error">
                <div v-for="search_animal in search_result" :key="search_animal[0]">
                    <DConcretAnimalLine v-bind:animal_item="search_animal" v-bind:filters="filters"/>
                </div>
            </div>
            <div v-if="search_error || errorr" class="search-error">
                Ничего не найдено
            </div>
        </div>
        <NumberPages v-bind:current-page="cp" v-bind:total-pages="tp" @changePage="changePage"/>
    </div>

</template>

<script>
import DConcretAnimalLine from './DConcretAnimalLine.vue';
import NumberPages from '@/components/NumberPages.vue';

export default {
    components: {
        DConcretAnimalLine, NumberPages
    },
    data () {
        return {
            animals: [],
            
            errorr: false,
        }
    },
    props: {
        isSearch: {
            Type: Boolean
        },
        search_result: {
            type: Array
        },
        search_error: {
            type: Boolean
        },
        cp: {
            type: Number,
            required: true
        },
        tp: {
            type: Number,
            required: true
        },
        filters: {
            type: Object,
        }
    },
    methods: {
        changePage(new_value) {
            this.$emit('defPages', new_value, this.tp);
        }
    },
    async mounted() {
        this.errorr = false;
        let search_params = this.filters;
        search_params.sex = [3];
        search_params.pageNumber = 1;
        const response = await fetch('/api/cows/filter', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json;charset=utf-8'
            },
            body: JSON.stringify(search_params),
        });
        const res_animals = await response.json();

        this.animals = res_animals.LST;
        if(res_animals.LST.length == 0) this.errorr = true;
        //Передаю текущую первую страницу и кол-во страниц наверх
        this.$emit('defPages', search_params.pageNumber, Math.ceil(res_animals.N/50));
    },
    watch: {
        async cp(newValue) {

            if(!this.isSearch) {
                let search_params = {};
                search_params.sex = [3];
                search_params.pageNumber = newValue;
                const response = await fetch('/api/cows/filter', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json;charset=utf-8'
                    },
                    body: JSON.stringify(search_params),
                });
                const res_animals = await response.json();
                this.animals = res_animals.LST;
            } else {
                this.$emit('changePageButSearch', newValue);
            }

        }
    },
}
</script>

<style scoped>
.parent-table {
    display: flex;
    flex-direction: column;
    align-items: center;
}

.animals-table {
    width: 930px;
    height: max-content;
    min-height: 500px;
    background-color: white;
    box-shadow: rgba(100, 100, 111, 0.1) 0px 7px 29px 0px;
    margin: 0 50px 0 10px;
    font-family: Open Sans, sans-serif;
    padding: 20px 30px;
    overflow-x: auto;
    font-size: 90%;
}

.animals-table::-webkit-scrollbar {
    height: 12px;
}

.animals-table::-webkit-scrollbar-track {
    background: rgb(241, 241, 241);
}

.animals-table::-webkit-scrollbar-thumb {
    background-color: rgb(183, 183, 183);
    border-radius: 20px;
    border: 3px solid rgb(241, 241, 241);
}

.table-header {
    display: flex;
    color: grey;
    padding: 10px 0;
    width: max-content;
}

.table-header div {
    padding: 0 10px;
}

.animal-selex {
    width: 140px;
}

.animal-inv {
    width: 115px;
}

.animal-rshn {
    width: 100px;
}

.animal-name {
    width: 140px;
    overflow-wrap: break-word;
}

.animal-hoz {
    width: 250px;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
}

.animal-bdate {
    width: 120px;
}

.animal-genfact {
    width: 120px;
}

.animal-dateout {
    width: 120px;
}

.animal-breed, .animal-dategen, .animal-datemilking, 
.animal-dateosem, .animal-dateotel, .animal-datebirk {
    width: 150px;
}

.animal-exterior {
    width: 100px;
}

.animal-krod, .animal-kfen {
    width: 100px;
}

.search-error {
    width: 100%;
    text-align: center;
    padding: 20px 0;
}
</style>