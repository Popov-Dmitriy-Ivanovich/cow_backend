<template>
    <ComboBox v-bind:start_value="options" @valueHasSelected="HasSelected" v-bind:clear="clearHoz" v-bind:valueFromOutside="valueFromOutside"></ComboBox>
</template>
    
<script>
import ComboBox from '@/components/ComboBox.vue';

export default {
    props: {
        clearHoz: {
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
        }
    },
    methods: {
        HasSelected(newValue) {
            this.$emit('sendToMain', newValue);
        }
    },
    async mounted() {
        this.options = [];
        const response = await fetch('/api/farms?parrent_id=null',
            {
                headers: {
                    'Authorization': localStorage.getItem('jwt')
                }
            }
        );
        const hozs = await response.json();
        console.log(hozs, hozs.length, 'hozs');
        for (let i = 0; i < hozs.length; i++) {
            let hoz = {name: hozs[i].Name, id: hozs[i].ID};
            this.options.push(hoz);
        }
    }
}
    
</script>
    
<style>
 
</style>