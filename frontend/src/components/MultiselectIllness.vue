<template>
    <ComboBox v-bind:start_value="options" @valueHasSelected="HasSelected" v-bind:clear="clearIll"></ComboBox>
</template>
    
<script>
import ComboBox from './ComboBox.vue';
    
export default {
    props: {
        clearIll: {
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
        const response = await fetch('https://genmilk.ru:9050/api/genetic_diseases');
        const illness = await response.json();
        for (let i = 0; i < illness.length; i++) {
            let ill = {name: illness[i][1], id: illness[i][0]};
            this.options.push(ill);
        }
    }
}
    
</script>
    
<style>
  
</style>