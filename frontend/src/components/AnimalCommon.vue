<template>
    <div class="common-title">Общая информация</div>
    <div class="general-info">
        <ID v-bind:cow_info="cow_info"/>
        <hr class="com-sep">
        <GenBreed v-bind:cow_info="cow_info"/>
        <hr class="com-sep">
        <ParentsCow v-bind:father="father" v-bind:mother="mother" v-bind:genetic="genetic" v-bind:coeff-by-family="koeff"></ParentsCow>
    </div>
</template>

<script>
import ID from '@/components/componentsConcretAnimal/ID.vue';
import GenBreed from '@/components/componentsConcretAnimal/GenBreed.vue';
import ParentsCow from '@/components/componentsConcretAnimal/ParentsCow.vue';

export default {
    components: {
        ID, GenBreed, ParentsCow
    },
    data() {
        return {
            cow_info: {},
            father: {},
            mother: {},
            genetic: {},
            koeff: 0,
        }
    },
    async mounted() {
        let mass_route = this.$route.path.split('/');
        let cow_id = mass_route[2];
        let response = await fetch(`/api/cows/${cow_id}`);
        let result = await response.json();
        this.cow_info = result;
        this.mother = this.cow_info.Mother;
        this.father = this.cow_info.Father;
        this.genetic = this.cow_info.Genetic;
        this.koeff = this.cow_info.InbrindingCoeffByFamily;
        console.log(this.cow_info, '32');
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
    width: 80%;
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