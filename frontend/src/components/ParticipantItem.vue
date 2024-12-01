<template>
<div class="participant">
    <div class="partic-logo">
        <img  width="100%" :src="logo">
    </div>
    <div class="partic-description"> {{ participant[5] }}</div>
</div>
</template>

<script>
export default {
    props: {
        participant: {
            type: Array,
            Required: true
        }
    },
    computed: {
        logo() {
            return "data:image/jpg;base64," + this.participant[6];
        }
    }
}
</script>

<style scoped>
    .participant {
        width: 300px;
        height: 400px;
        background-color: rgb(255, 255, 255);
        border-radius: 30px;
        display: flex;
        flex-direction: column;
        align-items: center;
        box-shadow: rgba(100, 100, 111, 0.2) 0px 7px 29px 0px;
    }

    .partic-logo, .partic-description {
        width: 80%;
        border-radius: 20px;
    }

    .partic-logo {
        height: min-content;
        margin: 15px 0;
    }

    .partic-description {
        height: 55%;
        font-size: 80%;
    }
</style>