<template>
    <ComboBox v-bind:start_value="options" @valueHasSelected="HasSelected" v-bind:clear="clearBreed" id="regions"></ComboBox>
</template>

<script>
import ComboBox from './ComboBox.vue';

export default {
    components: {
        ComboBox
    },
    props: {
        clearBreed: {
            type: Boolean,
        }
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
        const response = await fetch('/api/regions');
        const regions = await response.json();
        for (let i = 0; i < regions.length; i++) {
            let region = {name: regions[i].Name, id: regions[i].ID};
            this.options.push(region); 
        }
    },
}
</script>

<style scoped>

</style>