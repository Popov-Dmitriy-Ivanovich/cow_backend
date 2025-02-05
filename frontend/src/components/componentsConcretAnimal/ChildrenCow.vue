<template>
    <div class="child-title">Потомки</div>
    <div v-if="!isLoading">
        <div class="n-children">Количество потомков: {{ nChildren }}</div>
        <!-- <div class="parent-table">
            <table class="child-table">
                <thead>
                    <tr class="child-header">
                        <th>Идентификационный номер</th>
                        <th>Номер РСХН</th>
                        <th>Кличка</th>
                        <th>Дата рождения</th>
                        <th>Порода</th>
                        <th class="hoz">Хозяйство</th>
                    </tr>
                </thead>
                <tbody class="child-tablebody">
                    <tr v-for="child in children" :key="child.ID" class="child-line" @click="clickChild(child.Child.ID)">
                        <td>{{child.Child.IdentificationNumber || 'Нет информации'}}</td>
                        <td>{{child.Child.RSHNNumber || 'Нет информации'}}</td>
                        <td>{{child.Child.Name || 'Нет информации'}}</td>
                        <td>{{dateConverter(child.Child.BirthDate) || 'Нет информации'}}</td>
                        <td> {{ child.breed_name || 'Нет информации'}}</td>
                        <td class="hoz">{{ child.hoz_name || 'Нет информации'}}</td>
                    </tr>
                </tbody>
            </table>
        </div> -->
    </div>
    <div v-if="isLoading">Идёт загрузка...</div>
</template>
    
<script>
export default {
    data() {
        return {
            children: [],
            farm: {},
            namefarm: '',
            isLoading: false,
        }
    },
    async created() {
        let mass_route = this.$route.path.split('/');
        let cow_id = mass_route[2];
        this.isLoading = true;
        this.fetchInfo(cow_id);
        this.isLoading = false;
    },
    methods: {
        clickChild(id) {
            if(id) this.$router.push(`/animals/${id}`);
        },
        async fetchInfo(param) {
            let response = await fetch(`/api/cows/${param}/children`);
            let result = await response.json();
            if(result) {
                this.children = result;
            }
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
        nChildren() {
            return this.children.length || 0;
        }
    },
    watch: {
        $route(new_val) {
            this.fetchInfo(new_val.params.id);
        },
    }
}
</script>

<style scoped>
.child-title {
    font-size: 130%;
    color: rgb(10, 113, 75);
    padding-bottom: 30px;
    width: max-content;
}

.n-children {
    padding-bottom: 20px;
}

.parent-table {
    width: 50vw;
    overflow-x: auto;
}

.child-table {
    margin-bottom: 30px;
    text-align: left;

}

th {
    font-weight: normal;
}
th, td {
    border: none;
    padding: 0 4px;
}

.child-header {
    color: grey;
}

.child-header th {
    padding-right: 30px;
    padding-bottom: 15px;
}

.child-tablebody {
    text-align: left;
}

.parent-table::-webkit-scrollbar {
    height: 12px;
}

.parent-table::-webkit-scrollbar-track {
    background: rgb(241, 241, 241);
}

.parent-table::-webkit-scrollbar-thumb {
    background-color: rgb(183, 183, 183);
    border-radius: 20px;
    border: 3px solid rgb(241, 241, 241);
}

.child-line {
    cursor: pointer;
    transition: 0.3s;
}

.child-line:hover {
    color: rgb(58, 107, 84);
}

.hoz {
    width: 130px;
}
</style>