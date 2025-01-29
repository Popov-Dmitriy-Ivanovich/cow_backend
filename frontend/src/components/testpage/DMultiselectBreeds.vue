<template>
    <ComboBox v-bind:start_value="options" @valueHasSelected="HasSelected" v-bind:clear="clearBreed" id="breed" v-bind:valueFromOutside="value"></ComboBox>
</template>
    
<script>
import ComboBox from '@/components/ComboBox.vue';
    
export default {
    props: {
        clearBreed: {
            type: Boolean,
        },
        valueFromOutside: {
            type: Number,
        }
    },
    components: {
        ComboBox
    },
    data() {
        return {
            options: [],
            value: null,
        }
    },
    methods: {
        HasSelected(newValue) {
            this.$emit('sendToMain', newValue);
        },
        async getBreeds() {
            const response = await fetch('/api/breeds');
            const breeds = await response.json();
            this.options = [];
            for (let i = 0; i < breeds.length; i++) {
                let breed = {name: breeds[i].Name, id: breeds[i].ID};
                this.options.push(breed); 
            }
        }
    },
    async mounted() {
        await this.getBreeds();
    },
    watch: {
        async valueFromOutside(newVal) {
            await this.getBreeds();
            this.value = newVal;
        }
    }
}
    
</script>
    
<style>

</style>