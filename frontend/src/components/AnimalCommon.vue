<template>
    <div class="common-title">Общая информация</div>
    <div class="general-info" v-if="!isLoading">
        <ID v-bind:cow_info="cow_info"/>
        <hr class="com-sep">
        <GenBreed v-bind:cow_info="cow_info" v-bind:genetic="genetic"/>
        <hr class="com-sep">
        <ParentsCow v-bind:father="father" v-bind:mother="mother" v-bind:genetic="genetic" v-bind:coeff-by-family="koeff"></ParentsCow>
        <hr class="com-sep">
        <MovementCow v-bind:cow_info="cow_info"></MovementCow>
    </div>
    <div class="general-info" v-if="isLoading">Идёт загрузка...</div>
</template>

<script>
import ID from '@/components/componentsConcretAnimal/ID.vue';
import GenBreed from '@/components/componentsConcretAnimal/GenBreed.vue';
import ParentsCow from '@/components/componentsConcretAnimal/ParentsCow.vue';
import MovementCow from './componentsConcretAnimal/MovementCow.vue';

export default {
    components: {
        ID, GenBreed, ParentsCow, MovementCow
    },
    data() {
        return {
            cow_info: {},
            father: {},
            mother: {},
            genetic: {},
            koeff: 0,

            isLoading: false,
        }
    },
    async mounted() {
        let mass_route = this.$route.path.split('/');
        let cow_id = mass_route[2];
        await this.fetchInfo(cow_id);
    },
    methods: {
        async fetchInfo(param) {
            this.isLoading = true;
            let response = await fetch(`/api/cows/${param}`);
            let result = await response.json();
            this.cow_info = result;
            let response1 = await fetch(`/api/cows/${param}/genetic`);
            let result1 = await response1.json();
            if(result1) {
                this.genetic = result1;
            } else {
                this.genetic = {}
            }
            if (this.cow_info.Mother) {
                this.mother = this.cow_info.Mother;
            } else {
                this.mother = {};
            }
            if(this.cow_info.Father) {
                this.father = this.cow_info.Father;
            } else {
                this.father = {};
            }
            this.koeff = this.cow_info.InbrindingCoeffByFamily;
            this.isLoading = false;
        }
    },
    watch: {
        async $route(new_val) {
            await this.fetchInfo(new_val.params.id);
        }
    }
}
</script>

<style scoped>
.common-title {
    font-size: 190%;
    margin-bottom: 30px;
}

.general-info {
    background-color: white;
    width: 56vw;
    min-width: 800px;
    height: max-content;
    min-height: 500px;
    border-radius: 10px;
    box-shadow: rgba(100, 100, 111, 0.1) 0px 7px 29px 0px;
    padding: 30px 40px;
}

.com-sep {
    border: 1px solid rgb(224, 224, 224);
    margin-bottom: 40px;
}
</style>