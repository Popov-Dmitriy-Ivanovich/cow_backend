<template>
    <div class="animals">
        <div class="flex-top">
            <div class="flex-logo">
                <div class="animal-title">Животные</div>
                <div class="cows-bulls">
                    <button class="cow-btn" :class="{'current-animal-btn': isCows}" @click="cowsClick">Коровы</button>
                    <button class="bull-btn" :class="{'current-animal-btn': isBulls}" @click="bullsClick">Быки</button>
                    <button class="cow-btn" :class="{'current-animal-btn': isChild}" @click="childClick">Молодняк</button>
                </div>
            </div>
            <div class="search-block">
                <div>
                    <div class="search-text">Поиск по кличке, инвентарному номеру или идентификационному номеру РСХН</div>
                    <input class="search-animals" 
                    type="text" 
                    placeholder="Введите значение" 
                    @keyup.enter="searchCowsOrBulls" 
                    
                    id="search-animals"
                    autocomplete="off">
                </div>
                <button class="search-btn" @click="searchCowsOrBulls">Найти</button>
            </div>
        </div>
        <div class="filters-and-table">
            <DAnimalFilters @applyFilters="findAnimals"/>
             <!-- <AnimalFilters2 @applyFilters="findAnimals"/> -->
            <DCowsTable 
            v-if="isCows" 
            v-bind:isSearch="search" 
            v-bind:search_result="searching_animal" 
            v-bind:search_error="search_error_cows"
            v-bind:cp="current_page"
            v-bind:tp="total_pages"
            v-bind:filters="animal_filters"
            @defPages="setPages"
            @changePageButSearch="changePage"
            /> 
            
            <DBullsTable 
            v-if="isBulls" 
            v-bind:isSearch="search" 
            v-bind:search_result="searching_animal" 
            v-bind:search_error="search_error_bulls"
            v-bind:cp="current_page"
            v-bind:tp="total_pages"
            v-bind:filters="animal_filters"
            @defPages="setPages"
            @changePageButSearch="changePage"
            />

            <DChildTable 
            v-if="isChild" 
            v-bind:isSearch="search" 
            v-bind:search_result="searching_animal" 
            v-bind:search_error="search_error_child"
            v-bind:cp="current_page"
            v-bind:tp="total_pages"
            v-bind:filters="animal_filters"
            @defPages="setPages"
            @changePageButSearch="changePage"
            />
        </div>
    </div>
</template>

<script>
import DCowsTable from '@/components/testpage/DCowsTable.vue';
import DBullsTable from '@/components/testpage/DBullsTable.vue';
import DChildTable from '@/components/testpage/DChildTable.vue';
import DAnimalFilters from '@/components/testpage/DAnimalFilters.vue';

export default {
    components: {
        DCowsTable, DBullsTable, DChildTable, DAnimalFilters,
    },
    data() {
        return {
            isCows: true,
            isBulls: false,
            isChild: false,
            searching_animal: [],
            search: false,
            search_error_cows: false,
            search_error_bulls: false,
            search_error_child: false,
            current_filters: {},
            current_page: 1,
            total_pages: 1,
            animal_filters: {},
        }
    },
    methods: {
        async searchCowsOrBulls() {
            try {
                this.current_page = 1;

                this.searching_animal = [];

                let search_params = this.current_filters;
                search_params.searchQuery = document.getElementById('search-animals').value;
                if(this.isCows) search_params.sex = [4];
                if(this.isBulls) search_params.sex = [3];
                if(this.isChild) search_params.sex = [1,2];

                this.current_filters = search_params;

                console.log(JSON.stringify(search_params), 'параметры для отправки');

                const response = await fetch('/api/cows/filter', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json;charset=utf-8'
                    },
                    body: JSON.stringify(search_params),
                });
                let result = await response.json();
                this.searching_animal = result.LST;
                this.search = true;

                console.log(result);

                this.total_pages = Math.ceil(result.N/50);
                
                if(this.isCows) this.search_error_cows = false;
                if(this.isBulls) this.search_error_bulls = false;
                if(this.isChild) this.search_error_child = false;
                if (result.N == 0) {
                    if(this.isCows) this.search_error_cows = true;
                    if(this.isBulls) this.search_error_bulls = true;
                    if(this.isChild) this.search_error_child = true;
                }
            } catch(err) {
                if(this.isCows) this.search_error_cows = true;
                if(this.isBulls) this.search_error_bulls = true;
                if(this.isChild) this.search_error_child = true;
            }
        },
        async findAnimals(filters){
            try {
                this.current_page = 1;
                this.search = true;
                this.searching_animal = [];
                let search_params = filters;
                search_params.searchQuery = document.getElementById('search-animals').value;

                if(this.isCows) search_params.sex = [4];
                if(this.isBulls) search_params.sex = [3];
                if(this.isChild) search_params.sex = [1,2];


                this.current_filters = search_params;
                this.animal_filters = filters;

                let response = await fetch('/api/cows/filter', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json;charset=utf-8'
                    },
                    body: JSON.stringify(search_params),
                });
                let result = await response.json();

                this.total_pages = Math.ceil(result.N/50);

                if(this.isCows) this.search_error_cows = false;
                if(this.isBulls) this.search_error_bulls = false;
                if(this.isChild) this.search_error_child = false;

                console.log(this.search_error_bulls, this.search_error_child, this.search_error_cows);

                if(result.LST.length == 0 || result.N == 0) {
                    if(this.isCows) this.search_error_cows = true;
                    if(this.isBulls) this.search_error_bulls = true;
                    if(this.isChild) this.search_error_child = true;
                }

                this.searching_animal = result.LST;
            } catch(err) {
                if(this.isCows) this.search_error_cows = true;
                if(this.isBulls) this.search_error_bulls = true;
                if(this.isChild) this.search_error_child = true;
            }
        },
        async changePage(newVal) {
            this.current_page = newVal;
            this.current_filters.pageNumber = newVal;

            let response = await fetch('/api/cows/filter', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json;charset=utf-8'
                    },
                    body: JSON.stringify(this.current_filters),
                });
            let result = await response.json();
            this.searching_animal = result.LST;
        },
        setPages(npage, total) {
            this.current_page = npage;
            this.total_pages = total;
        },
        bullsClick() {
            console.log(this.search_error_bulls, this.search_error_child, this.search_error_cows);
            this.searching_animal = [];
            this.isCows = false;
            this.isChild = false;
            this.isBulls = true;
            this.search = false;
            this.search_error_bulls = false;
            document.getElementById('search-animals').value = '';
        },
        cowsClick() {
            console.log(this.search_error_bulls, this.search_error_child, this.search_error_cows);
            this.searching_animal = [];
            this.isCows = true;
            this.isChild = false;
            this.isBulls = false;
            this.search = false;
            this.search_error_cows = false;
            document.getElementById('search-animals').value = '';
        },
        childClick() {
            console.log(this.search_error_bulls, this.search_error_child, this.search_error_cows);
            this.searching_animal = [];
            this.isCows = false;
            this.isChild = true;
            this.isBulls = false;
            this.isSearch = false;
            this.search_error_child = false;
            document.getElementById('search-animals').value = '';
        }
    }
}
</script>

<style scoped>
.animals {
    margin-top: 110px;
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
}

.filters-and-table {
    display: flex;
    align-items: flex-start;
    justify-content: center;
    width: 100%;
    margin: 50px;
}

.flex-top {
    display: flex;
}

.flex-logo {
    padding: 10px 50px 0 0; 
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    margin: 0 20px;
    min-width: max-content;
}

.animal-title {
    font-size: 180%;
    font-family: Open Sans, sans-serif;
    color:rgb(37, 0, 132);
}

.cow-btn, .bull-btn {
    border: 1px solid white;
    background-color: white;
    color: rgb(37, 0, 132);
    padding: 15px 0;
    margin-top: 30px;
    width: 100px;
    border-radius: 10px;
    transition: 0.3s;
    cursor: pointer;
    margin-right: 15px;
}

.cow-btn:hover, .bull-btn:hover {
    border: 1px solid rgb(37, 0, 132);
}

.current-animal-btn {
    border: 1px solid rgb(37, 0, 132);
    background-color: rgb(232, 233, 248);
}

.search-block {
    background-color: white;
    padding: 10px 40px 20px 40px;
    border-radius: 5px;
    box-shadow: rgba(100, 100, 111, 0.1) 0px 7px 29px 0px;
    display: flex;
    align-items: flex-end;
}

.search-animals {
    width: 550px;
    height: 30px;
    padding: 0 15px;
    font-size: 110%;
    outline: none;
    border-radius: 10px;
    border: 2px solid white;
    background-color: rgb(248, 247, 252);
    transition: 0.3s;
}

.search-animals:focus {
    border: 2px solid rgb(122, 123, 193);
}

.search-animals::placeholder {
    opacity: 0.5;
    color:rgb(79, 60, 126);
}

.search-btn {
    background-color: white;
    border: 1px solid rgb(101, 102, 170);
    color: rgb(101, 102, 170);
    padding: 0 7px;
    height: 30px;
    border-radius: 10px;
    width: 70px;
    cursor: pointer;
    margin: 0 15px;
    transition: 0.3s;
}

.search-btn:hover {
    background-color: rgb(101, 102, 170);
    color: white;
}

.search-text {
    width: 460px;
    font-family: Open Sans, sans-serif;
    font-size: 90%;
    padding: 18px 7px;
    color:black;
}
</style>