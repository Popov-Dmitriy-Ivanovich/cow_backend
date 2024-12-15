<template>
    <div class="question">
        <div class="quest-title" @click="showAnswer = !showAnswer">
            <slot name="title"></slot>
        </div>
        <div class="quest-answer" v-if="showAnswer">
            <slot name="answer"></slot>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            showAnswer: false,
        }
    }
}
</script>

<style scoped>
.question {
    margin: 20px 25px;
}

.quest-title {
    font-size: 110%;
    color: rgb(37, 0, 132);
    margin-bottom: 15px;
    cursor: pointer;
}
</style>