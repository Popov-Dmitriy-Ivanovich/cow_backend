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
            options: [
                {
                    name: 'Федеральный чиновник',
                    id: 3
                },
                {
                    name: 'Региональный чиновник',
                    id: 3
                },
                {
                    name: 'Фермер',
                    id: 3
                },
            ],
        }
    },
    methods: {
        HasSelected(newValue) {
            this.$emit('sendToMain', newValue);
        }
    },
}
</script>

<style scoped>

</style>