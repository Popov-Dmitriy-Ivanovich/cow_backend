<template>
<div>
    <div class="animal-line" @click="$router.push(`/animals/${animal_item.ID}`)" :class="{'is-approved': animal_item.APPROVED==1}">
        <div class="animal-rshn">{{ animal_item.NRSHN }}</div>
        <div class="animal-inv">{{ animal_item.NINV }}</div>
        <div class="animal-name">{{ animal_item.NAME }}</div>
        <div class="animal-hoz">{{ animal_item.NAMEHOZ }}</div>
        <div class="animal-bdate">{{ animal_item.D_BIRTH }}</div>
        <div class="animal-genfact">{{ isGen(animal_item.IS_GEN) }}</div>
        <div v-if="filters.D_OUT_BEGIN || filters.D_OUT_END" class="animal-dateout">{{ animal_item.D_OUT }}</div>
        <div v-if="filters.ID_BREED" class="animal-breed">{{ animal_item.N_BREED }}</div>
        <div v-if="filters.D_GEN_BEGIN || filters.D_GEN_END" class="animal-dategen">{{ animal_item.D_GEN }}</div>
        <div v-if="filters.EXTERIOR" class="animal-exterior">{{ animal_item.EXTER }}</div>
        <div v-if="filters.D_OSEM_BEGIN || filters.D_OSEM_END" class="animal-dateosem">{{ animal_item.D_OSEMEN }}</div>
        <div v-if="filters.D_OTEL_BEGIN || filters.D_OTEL_END" class="animal-dateotel">{{ animal_item.D_OTEL }}</div>
        <div v-if="filters.D_BIRKING_BEGIN || filters.D_BIRKING_END" class="animal-datebirk">{{ animal_item.D_BIRKING }}</div>
        <div v-if="filters.K_INBR_ROD_BEGIN || filters.K_INBR_ROD_END" class="animal-krod">{{ animal_item.K_INBR }}</div>
        <div v-if="filters.K_INBR_FEN_BEGIN || filters.K_INBR_FEN_END" class="animal-kfen">{{ animal_item.INBRID_FENOTYPE }}</div>
    </div>
</div>
</template>

<script>
export default {
    props: {
        animal_item: {
            type: Object,
            Required: true
        },
        filters: {
            type: Object,
        }
    },
    methods: {
        isGen(a) {
            if(a) return 'Да';
            else return 'Нет';
        }
    }
}
</script>

<style scoped>
.animal-line {
    display: flex;
    padding: 10px 0;
    transition: 0.3s;
    cursor: pointer;
    width: max-content;
    color: grey;
}

.is-approved {
    color: black;
}

.animal-line:hover {
    background-color: rgb(235, 233, 245);
}

.animal-line div {
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
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
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
</style>