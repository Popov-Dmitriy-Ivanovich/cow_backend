<template>
    <ComboBox v-bind:start_value="options" @valueHasSelected="HasSelected" v-bind:clear="clearHoz"></ComboBox>
</template>
    
<script>
import ComboBox from '@/components/ComboBox.vue';

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
        const response = await fetch('/api/farms?parrent_id=null');
        const hozs = await response.json();
        for (let i = 0; i < hozs.length; i++) {
            let hoz = {name: hozs[i].Name, id: hozs[i].ID};
            this.options.push(hoz);
        }
    }
}
    
</script>
    
<style>
 
</style>