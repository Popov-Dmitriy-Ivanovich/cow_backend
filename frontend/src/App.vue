<template>
    <div class="app">
        <router-view></router-view>
        <MainFooter class="foot"/>
    </div>
    <TopBar class="bar"/>
</template>

<script>
import TopBar from './components/TopBar.vue';
import MainFooter from '@/components/MainFooter.vue'

export default {
    components: {
        TopBar, MainFooter
    }
}
</script>

<style>
* {
  margin: 0;
  padding: 0;
}

body {
    background-color: rgb(248, 247, 252);
}

html {
  scroll-behavior: smooth;
}

body::-webkit-scrollbar {
  width: 12px;
}

body::-webkit-scrollbar-track {
  background: rgb(241, 241, 241);
}

body::-webkit-scrollbar-thumb {
  background-color: rgb(183, 183, 183);
  border-radius: 20px;
  border: 3px solid rgb(241, 241, 241);
}

.bar {
    position: fixed;
    top: 0;
}

input[type="number"]::-webkit-outer-spin-button,
input[type="number"]::-webkit-inner-spin-button {
    -webkit-appearance: none;
}

input[type='number'],
input[type="number"]:hover,
input[type="number"]:focus {
    appearance: none;
    -moz-appearance: textfield;
}
</style>
