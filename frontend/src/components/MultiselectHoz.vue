<template>
    <ComboBox v-bind:start_value="options" @valueHasSelected="HasSelected" v-bind:clear="clearHoz"></ComboBox>
</template>
    
<script>
import ComboBox from './ComboBox.vue';

export default {
    props: {
        clearHoz: {
            type: Boolean,
        }
    },
    components: {
       ComboBox
    },
    data() {
         return {
            options: [],
        }
    },
    methods: {
        HasSelected(newValue) {
            this.$emit('sendToMain', newValue);
        }
    },
    async created() {
        this.options = [];
        const response = await fetch('https://genmilk.ru:9050/api/farms?ID_FARMER=1');
        const hozs = await response.json();
        for (let i = 0; i < hozs.length; i++) {
            let hoz = {name: hozs[i][1], id: hozs[i][0]};
            this.options.push(hoz);
        }
    }
}
    
</script>
    
<style>
 
</style>