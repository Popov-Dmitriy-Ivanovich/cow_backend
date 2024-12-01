<template>
    <MainPageBackground/>
    <ParticipantsBlock/>
    <HozBlock/>
    <NewsBlock/>
    <MainStatistic/>
</template>

<script>
import MainPageBackground from '@/components/MainPageBackground.vue'
import ParticipantsBlock from '@/components/ParticipantsBlock.vue'
import HozBlock from '@/components/HozBlock.vue'
import NewsBlock from '@/components/NewsBlock.vue'
import MainStatistic from '@/components/MainStatistic.vue';

export default {
    components: {
        MainPageBackground, ParticipantsBlock, HozBlock, NewsBlock, MainStatistic
    }
}
</script>

<style scoped>

</style>