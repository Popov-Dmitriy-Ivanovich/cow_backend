<template>
    <div class="participants">
        <div class="part-text">
            Участники
        </div>
        <div class="part-blocks">
            <div v-for="participant in participants" :key="participant[0]">
                <ParticipantItem v-bind:participant="participant"/>
            </div>
        </div>
        <div class="show-part" @click="$router.push('/participants')">Показать всех</div>
    </div>
</template>

<script>
import ParticipantItem from '@/components/ParticipantItem.vue';

export default {
    components: {
        ParticipantItem,
    },
    data() {
        return {
            participants: [],
        }
    },
    async created() {
        const response = await fetch('/api/partners');
        const result = await response.json();
        if(result.length > 3) {
            for (let i = 0; i < 4; i++) {
                this.participants.push(result[i]);
            }
        } else {
            for (let i = 0; i < result.length; i++) {
                this.participants.push(result[i]);
            }
        }

    }
}
</script>

<style scoped>
    .participants {
        height: 600px;
        text-align: center;
        font-family: Open Sans, sans-serif;
        color: rgb(37, 0, 132);
    }

    .part-text {
        font-size: 190%;
        padding: 40px 0;
    }

    .part-blocks {
        display: flex;
        justify-content: space-around;
        margin: 20px;
    }

    .show-part {
        cursor: pointer;
        transition: 0.3s;
    }

    .show-part:hover {
        color:rgb(83, 101, 237);
    }
</style>