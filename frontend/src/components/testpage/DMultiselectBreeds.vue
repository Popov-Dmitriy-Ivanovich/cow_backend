<template>
    <ComboBox v-bind:start_value="options" @valueHasSelected="HasSelected" v-bind:clear="clearBreed" id="breed"></ComboBox>
</template>
    
<script>
import ComboBox from '@/components/ComboBox.vue';
    
export default {
    props: {
        clearBreed: {
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
        const response = await fetch('/api/breeds');
        const breeds = await response.json();
        for (let i = 0; i < breeds.length; i++) {
            let breed = {name: breeds[i].Name, id: breeds[i].ID};
            this.options.push(breed);
        }
    },
}
    
</script>
    
<style>

</style>