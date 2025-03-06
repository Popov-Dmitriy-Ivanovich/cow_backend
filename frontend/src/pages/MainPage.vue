<template>
    <MainPageBackground/>
    <ParticipantsBlock/>
    <HozBlock/>
    <PurposesProect/>
    <NewsBlock/>
    <!-- <MainStatistic/> -->
</template>

<script>
import MainPageBackground from '@/components/MainPageBackground.vue'
import ParticipantsBlock from '@/components/ParticipantsBlock.vue'
import HozBlock from '@/components/HozBlock.vue'
import NewsBlock from '@/components/NewsBlock.vue'
// import MainStatistic from '@/components/MainStatistic.vue';
import PurposesProect from '@/components/PurposesProect.vue'

export default {
    components: {
        MainPageBackground, ParticipantsBlock, 
        HozBlock,
        NewsBlock, 
        // MainStatistic, 
        PurposesProect
    },
    // async created() {
    //     let response = await fetch('/api/user/whoami', {
    //         headers: {
    //             'Authorization': this.getJwt()
    //         },
    //     });
    //     let result = await response.json();
    //     console.log(result);
    // },
    // methods: {
    //     getJwt() {
    //         let arr = document.cookie.split(';');
    //         for (let i = 0; i < arr.length; i++) {
    //             if (arr[i].split('=')[0] == 'jwt') {
    //                 return arr[i].split('=')[1];
    //             }
    //         }
    //         return null;
    //     }
    // }
}
</script>

<style scoped>

</style>