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
        },
        getJwt() {
            let arr = document.cookie.split(';');
            for (let i = 0; i < arr.length; i++) {
                if (arr[i].split('=')[0] == 'jwt') {
                    return arr[i].split('=')[1];
                }
            }
            return null;
        }
    },
    async created() {
        this.options = [];
        const response = await fetch('/api/farms?parrent_id=null',
            {
                headers: {
                    'Authorization': this.getJwt()
                }
            }
        );
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