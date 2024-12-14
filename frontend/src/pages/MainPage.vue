<template>
    <MainPageBackground/>
    <ParticipantsBlock/>
    <HozBlock/>
    <PurposesProect/>
    <NewsBlock/>
    <MainStatistic/>
</template>

<script>
import MainPageBackground from '@/components/MainPageBackground.vue'
import ParticipantsBlock from '@/components/ParticipantsBlock.vue'
import HozBlock from '@/components/HozBlock.vue'
import NewsBlock from '@/components/NewsBlock.vue'
import MainStatistic from '@/components/MainStatistic.vue';
import PurposesProect from '@/components/PurposesProect.vue'

export default {
    components: {
        MainPageBackground, ParticipantsBlock, 
        HozBlock,
        NewsBlock, MainStatistic, PurposesProect
    }
}
</script>

<style scoped>

</style>