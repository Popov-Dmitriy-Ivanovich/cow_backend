<template>
    <div>
        <div class="id-title">Родословная</div>
        <div class="idnum-flex">
            <div class="item-block animal-parent"  @click="clickFather" :class="{'par-hover': father.Name}">
                <div class="id-min-title">Отец</div>
                <div class="link-parent">{{ father.Name || 'Нет информации' }}</div>
            </div>
            <div class="item-block animal-parent" @click="clickMother" :class="{'par-hover': mother.Name}">
                <div class="id-min-title">Мать</div>
                <div class="link-parent">{{ mother.Name || 'Нет информации' }}</div>
            </div>
                <div class="item-block">
                <div class="id-min-title">Коэф. инбридинга по родословной</div>
                <div v-if="coeffByFamily || coeffByFamily === 0">{{ coeffByFamily }}</div>
                <div v-else>Нет информации</div>
            </div>
            <div class="item-block">
                <div class="id-min-title">Коэф. инбридинга по генотипу</div>
                <div v-if="genetic.InbrindingCoeffByGenotype || genetic.InbrindingCoeffByGenotype === 0">{{ genetic.InbrindingCoeffByGenotype }}</div>
                <div v-else>Нет информации</div>
            </div>
        </div>
    </div>
</template>
    
<script>
export default {
    props: {
        mother: {
            type: Object,
            required: true,
        },
        father: {
            type: Object,
            required: true,
        },
        genetic: {
            type: Object,
            required: true,
        },
        coeffByFamily: {
            type: Number,
        }
    },
    methods: {
        clickFather() {
            if (this.father.ID) {
                this.$router.push(`/animals/${this.father.ID}`)
            }
        },
        clickMother() {
            if (this.mother.ID) {
                this.$router.push(`/animals/${this.mother.ID}`)
            }
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
    padding: 7px;
}

.id-min-title {
    color: grey;
    margin-bottom: 7px;
}

.animal-parent {
    transition: 0.3s;
}

.par-hover {
    cursor: pointer;
}

.par-hover:hover {
    background-color: rgb(236, 232, 245);
    border-radius: 10px;
}

</style>