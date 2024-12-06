<template>
    <div>
        <div class="child-title">Потомки</div>
        <div class="n-children">Количество потомков: {{ nChildren }}</div>
        <div class="parent-table">
            <table class="child-table">
                <thead>
                    <tr class="child-header">
                        <th>Сэлекс</th>
                        <th>Хозяйство</th>
                        <th>Дата рождения</th>
                        <th>Кличка</th>
                        <!-- <th>Селекционный индекс</th>
                        <th>Молочный индекс</th> -->
                    </tr>
                </thead>
                <tbody class="child-tablebody">
                    <tr v-for="child in children" :key="child.ID" class="child-line" @click="clickChild(child.ID)">
                        <td>{{child.SelecsNumber}}</td>
                        <td></td>
                        <td>{{child.BirthDate}}</td>
                        <td>{{child.Name}}</td>
                        <!-- <td>test</td>
                        <td>test</td> -->
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>
    
<script>
export default {
    data() {
        return {
            children: [],
        }
    },
    async created() {
        let mass_route = this.$route.path.split('/');
        let cow_id = mass_route[2];
        this.fetchInfo(cow_id);
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
    },
    computed: {
        nChildren() {
            return this.children.length || 0;
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
.child-title {
    font-size: 130%;
    color: rgb(37, 0, 132);
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
    color: rgb(74, 58, 107);
}
</style>