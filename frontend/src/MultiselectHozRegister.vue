<template>
    <ComboBox v-bind:start_value="options" @valueHasSelected="HasSelected"></ComboBox>
</template>

<script>
import ComboBox from './components/ComboBox.vue';

export default {
    components: {
        ComboBox
    },
    props: {
        regionId: {
            type: Number,
        }
    },
    data() {
        return {
            options: [],
        }
    },
    watch: {
        async regionId(new_region){
            this.options = [];
            let response = await fetch(`/api/regions/${new_region}/getFarms`);
            let result = await response.json();
            if (result.farms.length) {
                for (let i = 0; i < result.farms.length; i++) {
                    let obj = {name: result.farms[i].Name, id: result.farms[i].HozNumber};
                    this.options.push(obj);
                }
            } else {
                this.options = [];
            }
        }
    },
    methods: {
        HasSelected(newValue) {
            this.$emit('sendToMain', newValue);
        }
    }
}
</script>

<style scoped>

</style>