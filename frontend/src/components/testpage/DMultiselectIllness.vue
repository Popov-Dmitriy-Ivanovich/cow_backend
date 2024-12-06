<template>
    <ComboBox v-bind:start_value="options" @valueHasSelected="HasSelected" v-bind:clear="clearIll"></ComboBox>
</template>
    
<script>
import ComboBox from '@/components/ComboBox.vue';
    
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
        const response = await fetch('/api/monogenetic_illnesses');
        const illness = await response.json();
        let il = {name: 'Все', id: []};
        for (let i = 0; i < illness.length; i++) {
            il.id.push(illness[i].ID);
        }
        this.options.push(il);
        for (let i = 0; i < illness.length; i++) {
            let ill = {name: illness[i].Name, id: illness[i].ID};
            this.options.push(ill);
        }
    }
}
    
</script>
    
<style>
  
</style>