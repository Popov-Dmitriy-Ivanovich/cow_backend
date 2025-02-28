<template>
    <div class="participants">
        <div class="part-text">
            Участники
        </div>
        <div class="part-text-low">Для реализации проекта образован консорциум из нескольких компаний, 
            включая крупные агропромышленные холдинги, образовательные учреждения и генетические лаборатории.</div>
        <div class="part-blocks">
            <div v-for="participant in participants" :key="participant[0]" class="part-block">
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
        if(result.length > 2) {
            for (let i = 0; i < 3; i++) {
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
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
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

    .part-text-low {
        color: black;
        width: 50%;
        margin-bottom: 30px;
    }

    .part-block {
        margin: 0 50px;
    }
</style>