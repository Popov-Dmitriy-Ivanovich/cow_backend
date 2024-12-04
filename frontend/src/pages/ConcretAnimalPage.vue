<template>
    <div class="concret-animal">
        <div class="animal-sideblock">
            <div @click="$router.push('/animals')" class="link-to-prev">🠔 Животные</div>
            <AnimalSection @changeSection="parentChangeSection"/>
        </div>
        <div class="animal-info">
            <AnimalMainInfo/>
            <AnimalCommon v-if="section.common"/>
            <AnimalRating v-if="section.rating"/>
            <AnimalGenetic v-if="section.genetic"/>
            <AnimalExterior v-if="section.exterior"/>
            <AnimalHealth v-if="section.health"/>
            <AnimalMilking v-if="section.milking"/>
            <AnimalDocuments v-if="section.documents"/>
            <AnimalLineage v-if="section.lineage"/>
        </div>
    </div>
</template>

<script>
import AnimalSection from '@/components/AnimalSections.vue';
import AnimalCommon from '@/components/AnimalCommon.vue';
import AnimalRating from '@/components/AnimalRating.vue';
import AnimalGenetic from '@/components/AnimalGenetic.vue';
import AnimalExterior from '@/components/AnimalExterior.vue';
import AnimalHealth from '@/components/AnimalHealth.vue';
import AnimalMilking from '@/components/AnimalMilking.vue';
import AnimalDocuments from '@/components/AnimalDocuments.vue';
import AnimalLineage from '@/components/AnimalLineage.vue';
import AnimalMainInfo from '@/components/AnimalMainInfo.vue';

export default {
    components: {
        AnimalSection, AnimalCommon, AnimalRating, AnimalGenetic, AnimalExterior, AnimalHealth, AnimalMilking, 
        AnimalDocuments, AnimalLineage, AnimalMainInfo
    },
    data() {
        return {
            section: {
                common: true,
                rating: false,
                genetic: false,
                exterior: false,
                health: false,
                milking: false,
                documents: false,
                lineage: false,
            }
        }
    },
    methods: {
        parentChangeSection(section_name){
            for (let key in this.section) {
                if(key == section_name) {
                    this.section[key] = true;
                } else {
                    this.section[key] = false;
                }
            }
        }
    }
}
</script>

<style scoped>
.concret-animal {
    font-family: Open Sans, sans-serif;
    margin-top: 140px;
    display: flex;
}

.animal-sideblock {
    padding-left: 90px;
}

.link-to-prev {
    color:rgb(37, 0, 132);
    padding-bottom: 20px;
    cursor: pointer;
    transition: 0.3s;
}

.link-to-prev:hover {
    color: rgb(103, 63, 205);
    padding-left: 10px;
    width: max-content;
}

.animal-info {
    margin: 0 100px;
    width: 100%;
}
</style>