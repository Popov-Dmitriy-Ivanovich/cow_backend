<template>
    <ComboBox v-bind:start_value="options" @valueHasSelected="HasSelected" v-bind:clear="clearHoz" v-bind:valueFromOutside="value"></ComboBox>
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
            value: null,
        }
    },
    methods: {
        HasSelected(newValue) {
            this.$emit('sendToMain', newValue);
        },
        async getHoz() {
            const response = await fetch('/api/farms?parrent_id=null',
                {
                    headers: {
                        'Authorization': this.getJwt()
                    }
                }
            );
            const hozs = await response.json();
            this.options = [];
            for (let i = 0; i < hozs.length; i++) {
                let hoz = {name: hozs[i].Name, id: hozs[i].ID};
                this.options.push(hoz);
            }
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
    async mounted() {
       await this.getHoz();
    },
    watch: {
        async valueFromOutside(newVal) {
            await this.getHoz();
            this.value = newVal;
        }
    }
}
    
</script>
    
<style>
 
</style>