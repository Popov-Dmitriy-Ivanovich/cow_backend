<template>
<div class="part-page">
    <div @click="$router.push('/')" class="link-to-main">🠔 Главная</div>
    <div class="part-title">Участники</div>
    <div class="list-part">
        <div v-for="participant in all_participants" :key="participant[0]">
            <ParticipantItem v-bind:participant="participant" class="part-item"></ParticipantItem>
        </div>
    </div>

</div>
</template>

<script>
import ParticipantItem from '@/components/ParticipantItem.vue';

export default {
    data() {
        return {
            all_participants: [],
        }
    },
    components: {
        ParticipantItem
    },
    async created() {
        const response = await fetch('/api/partners');
        const result = await response.json();
        this.all_participants = result;
    }
}
</script>

<style scoped>
.part-page {
    margin-top: 110px;
    width: 100%;
    font-family: Open Sans, sans-serif;
    color: rgb(37, 0, 132);
    display: flex;
    flex-direction: column;
    align-items: center;
}

.link-to-main {
    align-self: flex-start;
    cursor: pointer;
}

.part-title {
    font-size: 210%;
    padding-bottom: 60px;
}

.list-part {
    width: 100%;
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
}

.part-item {
    margin: 15px
}
</style>