<template>
<div>
    <multiselect v-model="value" 
    :options="options" 
    label="name" 
    track-by="name" 
    placeholder="Выберите регион"
    ><template v-slot:noResult>Ничего не найдено</template></multiselect>
</div>
</template>

<script>
import Multiselect from 'vue-multiselect';

export default {
    components: {
        Multiselect,
    },
    data() {
        return {
            value: null,
            options: [],
        }
    },
    methods: {
        
    },
    async created() {
        this.options = [];
        const response = await fetch('https://genmilk.ru:9050/api/regions');
        const regions = await response.json();
        for (let i = 0; i < regions.length; i++) {
            let region = {name: regions[i][1]};
            this.options.push(region);
        }
    }
}

</script>

<style>
    .multiselect__tags {
        height: 20px;
        display: block;
        border-radius: 10px;
        border: 3px solid rgb(195, 200, 212);
        background: #fff;
        font-size: 14px;
        padding: 5px 5px;
        transition: 0.3s;
    }

    .multiselect__tags:hover {
        border: 3px solid rgb(101, 102, 170);
    }

    .multiselect__input {
        outline: none;
        border: none;
    }

    .multiselect__input::placeholder {
        opacity: 1;
        color: gray;
        font-size: 14px;
    }

    .multiselect__input:focus {
        border: none;
    }

    .multiselect__content {
        list-style: none;
    }

    .multiselect__content-wrapper {
        padding: 10px 10px;
        border: 1px solid rgb(230, 230, 230);
        border-radius: 10px;
        overflow: scroll;
        position: absolute;
        background-color: white;
        width: 276px;
        font-size: 90%;
    }

    .multiselect__element {
        cursor: pointer;
        margin: 7px 0;
    }

    .multiselect__content-wrapper::-webkit-scrollbar {
        width: 12px;
    }

    .multiselect__content-wrapper::-webkit-scrollbar-track {
        background: rgb(241, 241, 241);
    }

    .multiselect__content-wrapper::-webkit-scrollbar-thumb {
        background-color: rgb(183, 183, 183);
        border-radius: 20px;
        border: 3px solid rgb(241, 241, 241);
    }

    .multiselect__content-wrapper::-webkit-scrollbar:horizontal {
        height: 0;
    }
</style>