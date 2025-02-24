<template>
    <div class="animals">
        <Teleport to="body">
            <div class="modal-back" v-if="isApproveDelete" @click="isApproveDelete = false">
                <div class="modal" @click.stop="">
                    <div>Вы действительно хотите удалить всех отфильтрованных животных?</div>
                    <div class="delete-buttons">
                        <div class="confirm-delete" @click="deleteFiltered">Удалить</div>
                        <div class="reject-delete" @click="isApproveDelete = false">Отмена</div>
                    </div>
                </div>
            </div>
        </Teleport>
        <table>
            <tbody>
                <tr>
                    <td>
                        <div class="flex-logo">
                            <div class="animal-title">Животные</div>
                            <div class="cows-bulls">
                                <button class="cow-btn" :class="{'current-animal-btn': isCows}" @click="cowsClick">Коровы</button>
                                <button class="bull-btn" :class="{'current-animal-btn': isBulls}" @click="bullsClick">Быки</button>
                                <button class="cow-btn" :class="{'current-animal-btn': isChild}" @click="childClick">Молодняк</button>
                            </div>
                        </div>
                    </td>
                    <td class="ceil-search">
                        <div class="search-block">
                            <div>
                                <!-- <div class="search-text">Поиск по кличке, инвентарному номеру, Сэлекс или идентификационному номеру РСХН</div> -->
                                <input class="search-animals" 
                                type="text" 
                                placeholder="Кличка, инвентарный номер, Сэлекс, РСХН" 
                                @keyup.enter="searchCowsOrBulls" 
                                
                                id="search-animals"
                                autocomplete="off">
                            </div>
                            <button class="search-btn" @click="searchCowsOrBulls">Найти</button>
                        </div>
                    </td>
                </tr>
                <tr>
                    <td>

                    </td>
                    <td class="ceil-line">
                        <div class="count-animals">
                            <div class="nanimals">
                                Количество животных: {{ number_of_animals }}
                            </div>
                            <div class="animals-on-page">
                                <div>На странице:</div>
                                <input type="number" class="input-nanimals" v-model="animal_filters.entitiesOnPage" @keyup.enter="searchCowsOrBulls">
                            </div>
                        </div>
                        <div class="sort">
                            <div class="save-btns">
                                <div>Сортировать по: </div>
                                <select v-model="sort" v-on:change="searchCowsOrBulls"  class="filter-input">
                                    <!-- <option :value="null">-нет-</option> -->
                                    <option :value="null">РСХН</option>
                                    <option :value="'Name'">Кличка</option>
                                    <option :value="'InventoryNumber'">Инв. номер</option>
                                    <option :value="'BirthDate'">Дата рождения</option>
                                </select>
                                <select class="filter-input" v-on:change="searchCowsOrBulls" v-model="order">
                                    <option :value="false">по возрастанию</option>
                                    <option :value="true">по убыванию</option>
                                </select>
                            </div>

                            <div class="save-btns">
                                <button class="delete-animals" @click="isApproveDelete = true">Удалить</button>
                                <button class="save-table" @click="saveCSV">CSV</button>
                                <button class="save-table" @click="saveXLS">XLS</button>
                            </div>
                        </div>
                    </td>
                </tr>
                <tr>
                    <td class="ceil-filter">
                        <DAnimalFilters @applyFilters="findAnimals"/>
                    </td>
                    <td class="ceil-animal">
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
                        @changeN="changeN"
                        v-bind:isLoading="isLoading"
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
                        @changeN="changeN"
                        v-bind:isLoading="isLoading"
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
                        @changeN="changeN"
                        v-bind:isLoading="isLoading"
                        />
                    </td>
                </tr>
            </tbody>
        </table>
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
            animal_filters: {entitiesOnPage: 50},

            isLoading: false,
            sort: null,
            order: false,

            isApproveDelete: false,
            number_of_animals: 0,
        }
    },
    methods: {
        async searchCowsOrBulls() {
            try {
                this.isLoading = true;
                console.log(this.isLoading);
                this.current_page = 1;

                this.searching_animal = [];

                let search_params = this.current_filters;
                search_params.searchQuery = document.getElementById('search-animals').value;
                if(this.isCows) search_params.sex = [4];
                if(this.isBulls) search_params.sex = [3];
                if(this.isChild) search_params.sex = [1,2];
                search_params.entitiesOnPage = this.animal_filters.entitiesOnPage;
                if(this.sort) {
                    search_params.orderByDesc = this.order;
                    search_params.orderBy = this.sort;
                } else {
                    search_params.orderBy = 'RSHN';
                    search_params.orderByDesc = this.order;
                }

                this.current_filters = search_params;
                this.animal_filters = search_params;

                console.log(JSON.stringify(search_params), 'параметры для отправки');

                const response = await fetch('/api/cows/filter', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json;charset=utf-8',
                        'Authorization': this.getJwt()
                    },
                    body: JSON.stringify(search_params),
                });
                let result = await response.json();
                this.searching_animal = result.LST;
                this.number_of_animals = result.N;
                this.search = true;

                this.total_pages = Math.ceil(result.N/search_params.entitiesOnPage);
                
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
            this.isLoading = false;
        },
        async findAnimals(filters){
            try {
                this.isLoading = true;
                this.current_page = 1;
                this.search = true;
                this.searching_animal = [];
                let search_params = filters;
                search_params.searchQuery = document.getElementById('search-animals').value;

                if(this.isCows) search_params.sex = [4];
                if(this.isBulls) search_params.sex = [3];
                if(this.isChild) search_params.sex = [1,2];
                console.log(this.animal_filters);
                search_params.entitiesOnPage = this.animal_filters.entitiesOnPage;
                if(this.sort) {
                    search_params.orderByDesc = this.order;
                    search_params.orderBy = this.sort;
                } else {
                    search_params.orderBy = 'RSHN';
                    search_params.orderByDesc = this.order;
                }
                this.current_filters = search_params;

                let response = await fetch('/api/cows/filter', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json;charset=utf-8',
                        'Authorization': this.getJwt()
                    },
                    body: JSON.stringify(search_params),
                });
                let result = await response.json();

                this.total_pages = Math.ceil(result.N/search_params.entitiesOnPage);

                if(this.isCows) this.search_error_cows = false;
                if(this.isBulls) this.search_error_bulls = false;
                if(this.isChild) this.search_error_child = false;

                if(result.LST.length == 0 || result.N == 0) {
                    if(this.isCows) this.search_error_cows = true;
                    if(this.isBulls) this.search_error_bulls = true;
                    if(this.isChild) this.search_error_child = true;
                }

                this.searching_animal = result.LST;
                this.number_of_animals = result.N;
                Object.assign(this.animal_filters, search_params);
            } catch(err) {
                if(this.isCows) this.search_error_cows = true;
                if(this.isBulls) this.search_error_bulls = true;
                if(this.isChild) this.search_error_child = true;
            }
            this.isLoading = false;
            // this.animal_filters = search_params;
        },
        async changePage(newVal) {
            this.isLoading = true;
            this.current_page = newVal;
            this.current_filters.pageNumber = newVal;
            console.log(this.current_filters);

            let response = await fetch('/api/cows/filter', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json;charset=utf-8',
                        'Authorization': this.getJwt()
                    },
                    body: JSON.stringify(this.current_filters),
                });
            let result = await response.json();
            this.searching_animal = result.LST;
            this.isLoading = false;
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
            // document.getElementById('search-animals').value = '';
        },
        cowsClick() {
            console.log(this.search_error_bulls, this.search_error_child, this.search_error_cows);
            this.searching_animal = [];
            this.isCows = true;
            this.isChild = false;
            this.isBulls = false;
            this.search = false;
            this.search_error_cows = false;
            // document.getElementById('search-animals').value = '';
        },
        childClick() {
            console.log(this.search_error_bulls, this.search_error_child, this.search_error_cows);
            this.searching_animal = [];
            this.isCows = false;
            this.isChild = true;
            this.isBulls = false;
            this.search = false;
            this.search_error_child = false;
            // document.getElementById('search-animals').value = '';
        },
        changeN(newN) {
            this.number_of_animals = newN;
        },
        getJwt() {
            let arr = document.cookie.split(';');
            for (let i = 0; i < arr.length; i++) {
                if (arr[i].split('=')[0] == 'jwt') {
                    return arr[i].split('=')[1];
                }
            }
            return null;
        },
        setCurrentFilters() {
            if (this.isCows) {
                this.current_filters.sex = [4];
            }
            if (this.isBulls) {
                this.current_filters.sex = [3];
            }
            if (this.isChild) {
                this.current_filters.sex = [1,2];
            }
            this.current_filters.pageNumber = 1;
            this.current_filters.entitiesOnPage = 50;
            if(!this.current_filters.orderBy) {
                this.current_filters.orderBy = 'RSHN';
                this.current_filters.orderByDesc = false;
            }
        },
        async saveCSV() {
            if (!Object.keys(this.current_filters).length) {
                this.setCurrentFilters();
            }
            let response = await fetch('/api/cows/filterCSV', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8',
                    'Authorization': this.getJwt()
                },
                body: JSON.stringify(this.current_filters),
            });
            let result = await response.json();
            console.log(result);
            let filename = result.LST;
            let url = '/api/static/csv/' + filename;
            let link = document.createElement('a');
            link.href = url;
            link.setAttribute('download', filename);
            document.body.appendChild(link);
            link.click();
            link.remove();
        },
        async saveXLS() {
            if (!Object.keys(this.current_filters).length) {
                this.setCurrentFilters();
            }
            let response = await fetch('/api/cows/filterExcel', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8',
                    'Authorization': this.getJwt()
                },
                body: JSON.stringify(this.current_filters),
            });
            let result = await response.json();
            console.log(result);
            let filename = result.LST;
            let url = '/api/static/excel/' + filename;
            let link = document.createElement('a');
            link.href = url;
            link.setAttribute('download', filename);
            document.body.appendChild(link);
            link.click();
            link.remove();
        },
        async deleteFiltered() {
            if (!Object.keys(this.current_filters).length) {
                this.setCurrentFilters();
            }
            let response = await fetch('/api/cows/delByFilters', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8',
                    'Authorization': this.getJwt()
                },
                body: JSON.stringify(this.current_filters),
            });
            let result = await response.json();
            if (result.N) {
                this.isApproveDelete = false;
                await this.findAnimals(this.current_filters);
            }
            
        }
    }
}
</script>

<style scoped>
.animals {
    margin-top: 130px;
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
    align-items: center;
}

.flex-logo {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
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
    padding: 30px 40px 30px 40px;
    border-radius: 10px;
    box-shadow: rgba(100, 100, 111, 0.1) 0px 7px 29px 0px;
    display: flex;
    align-items: center;
}

.search-animals {
    width: 770px;
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

.sort {
    font-family: Open Sans, sans-serif;
    display: flex;
    align-items: center;
    margin-bottom: 10px;
    justify-content: space-between;
    width: 100%;
}

.filter-input {
    height: 30px;
    width: 200px;
    padding: 0 10px;
    font-size: 14px;
    box-sizing: border-box;
    outline: none;
    border: 1px solid rgb(101, 102, 170);
    border-radius: 10px;
    transition: 0.3s;
    margin-left: 15px;
}

.save-btns {
    display: flex;
    align-items: center;
}

.save-table {
    background-color: white;
    border: 1px solid rgb(101, 102, 170);
    color: rgb(101, 102, 170);
    padding: 0 7px;
    height: 30px;
    border-radius: 10px;
    width: 80px;
    cursor: pointer;
    margin: 0 5px;
    transition: 0.3s;
}

.save-table:hover {
    background-color: rgb(101, 102, 170);
    color: white;
}

.delete-animals {
    background-color: white;
    border: 1px solid rgb(204, 99, 99);
    color: rgb(204, 99, 99);
    padding: 0 7px;
    height: 30px;
    border-radius: 10px;
    width: 80px;
    cursor: pointer;
    margin: 0 5px;
    transition: 0.3s;
}

.delete-animals:hover {
    background-color: rgb(204, 99, 99);
    color: white;
}

.modal-back {
    background-color: rgba(0,0,0,0.2);
    position: fixed;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    z-index: 500;
    display: flex;
    justify-content: center;
    align-items: center;
}

.modal {
    background-color: white;
    padding: 30px 40px;
    border-radius: 10px;
    font-family: Open Sans, sans-serif;
    max-width: 700px;
}

.delete-buttons {
    margin-top: 40px;
    display: flex;
    width: 100%;
    justify-content: center;
}

.confirm-delete, .reject-delete {
    padding: 8px 10px;
    border-radius: 5px;
    cursor: pointer;
    border: none;
    transition: 0.3s;
}

.reject-delete {
    margin-left: 20px;
    color: black;
}

.reject-delete:hover {
    background-color: rgb(243, 240, 246);
}

.confirm-delete {
    color: red;
}

.confirm-delete:hover {
    background-color: rgb(255, 225, 225);
}

.count-animals {
    width: 100%;
    display: flex;
    justify-content: space-between;
    align-items: start;
}

.nanimals {
    font-family: Open Sans, sans-serif;
    font-size: 120%;
    color: rgb(66, 66, 66);
    padding-right: 40px;
    padding-bottom: 15px;
}

.ceil-animal {
    vertical-align: top;
}

.ceil-line {
    padding-top: 20px;
}

.ceil-filter {
    padding-right: 20px;
}

.ceil-search {
    vertical-align: bottom;
}

.animals-on-page {
    display: flex;
    align-items: center;
    font-family: Open Sans, sans-serif;
}

.input-nanimals {
    padding: 0 10px;
    font-size: 16px;
    box-sizing: border-box;
    outline: none;
    border: 1px solid rgb(101, 102, 170);
    border-radius: 10px;
    height: 30px;
    width: 80px;
    margin: 0 6px;
    text-align: center;
}
</style>