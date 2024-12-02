<template>
    <div>
        <div class="id-title">Родословная</div>
        <div class="idnum-flex">
            <div class="item-block">
                <div class="id-min-title">Отец</div>
                <div @click="clickFather" class="link-parent">{{ father.Name }}</div>
            </div>
            <div class="item-block">
                <div class="id-min-title">Мать</div>
                <div @click="clickMother" class="link-parent">{{ mother.Name }}</div>
            </div>
        </div>
        <div>
            <div class="item-block">
                <div class="id-min-title">Коэффициент инбридинга по родословной</div>
                <div>{{ cow_info.InbrindingCoeffByFamily }}</div>
            </div>
            <div class="item-block">
                <div class="id-min-title">Коэффициент инбридинга по генотипу</div>
                <div>{{ genetic.InbrindingCoeffByGenotype }}</div>
            </div>
        </div>
    </div>
</template>
    
<script>
export default {
    data() {
        return {
            cow_info: {},
            mother: {},
            father: {},
            genetic: {},
        }
    },
    async created() {
        let mass_route = this.$route.path.split('/');
        let cow_id = mass_route[2];
        let response = await fetch(`/api/cows/${cow_id}`);
        let result = await response.json();
        this.cow_info = result;
        this.mother = this.cow_info.Mother;
        this.father = this.cow_info.Father;
        this.genetic = this.cow_info.Genetic;
        console.log(this.cow_info);
    },
    methods: {
        clickFather() {
            if (this.cow_info.Father.ID) {
                this.$router.push(`/animals/${this.cow_info.Father.ID}`)
            }
        },
        clickMother() {
            if (this.cow_info.Mother.ID) {
                this.$router.push(`/animals/${this.cow_info.Mother.ID}`)
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
}

.id-min-title {
    color: grey;
    margin-bottom: 7px;
}

.link-parent {
    cursor: pointer;
}
</style>