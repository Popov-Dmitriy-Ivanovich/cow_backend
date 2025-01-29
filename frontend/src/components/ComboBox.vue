<template>
<div class="combobox">
    <div class="combobox__flex">
        <input type="text" 
        readonly 
        @click="isVisible = !isVisible" 
        value="" 
        class="str-value" 
        placeholder="Выберите значение" 
        v-model="text1"
        :class="{'str-value__click': isVisible}">

        <div class="combobox__clear-value" @click="clearValue">🗙</div> 
    </div>
      
    <div class="combobox__search-select" v-if="isVisible"> 
        <form>
            <input 
            type="text" 
            placeholder="Поиск..." 
            @keyup="search" 
            class="search-field" 
            autocomplete="off" 
            v-model="text2"
            :class="{'search-field__click': isVisible}">
            <div v-for="elem in list" :key="elem" @click="clickValue(elem)" class="combobox__list">
                {{ elem.name }}
            </div>
        </form>    
    </div>

</div>
</template>

<script>
export default {
    props: {
        start_value: {
            type: Array,
        },
        clear: {
            type: Boolean,
        },
        valueFromOutside: {
            type: Number,
            default: null,
        }
    },
    data() {
        return {
            isVisible: false,
            list: [],
            selected_value: [],
            text1: "",
            text2: "",
        }
    },
    methods: {
        clickValue(new_value) {
            this.selected_value = [];
            this.text1 = new_value.name;
            this.selected_value.push(new_value.id);
            this.isVisible = false;
            this.$emit('valueHasSelected', this.selected_value[0]);
        },
        search() {
            let search_val = this.text2;
            this.list = [];
            for (let i = 0; i < this.start_value.length; i++) {
                if (this.start_value[i].name.toLowerCase().includes(search_val.toLowerCase())) {
                    this.list.push(this.start_value[i]);
                }
            }
        },
        clearValue(){
            this.text1="";
            this.selected_value = [];
            this.$emit('valueHasSelected', null);
            this.isVisible = false;
        }
    },
    created() {
        this.list = this.start_value;
    },
    watch: {
        clear() {
            this.text1 = "";
        },
        start_value(new_val) {
            this.list = new_val;
        },
        valueFromOutside(new_val) {
            if (new_val) {
                this.selected_value = [];
                this.selected_value.push(new_val);
                console.log(this.start_value.length);
                for (let val of this.start_value) {
                    if (val.id === new_val) {
                        this.text1 = val.name;
                    }
                }
                this.isVisible = false;
            }
        }
    }
}
</script>

<style scoped>
.combobox {
    width: max-content;
    margin-bottom: 10px;
}

.combobox__search-select {
    position: absolute;
    max-height: 200px;
    overflow: auto;
    background-color: white;
    z-index: 40;
    border: 1px solid rgb(205, 205, 205);
    border-radius: 0 0 10px 10px;
    width: 240px;
}

.combobox__flex {
    display: flex;
    align-items: center;
}

.combobox__clear-value {
    color: rgb(176, 176, 176);
    cursor: pointer;
    margin: 6px 5px 0 5px;
    transition: 0.3s;
}

.combobox__clear-value:hover {
    color: rgb(101, 170, 144);
}

.combobox__list {
    padding: 5px 3px;
    width: auto;
    font-size: 90%;
    cursor: pointer;
    transition: 0.3s;
}

.combobox__list:hover {
    background-color: rgb(238, 245, 241);
}

.str-value {
    border-radius: 10px;
    border: 3px solid rgb(203, 227, 219);
    background: #fff;
    font-size: 14px;
    padding: 5px 5px;
    transition: 0.3s;
    cursor: pointer;
    outline: none;
    margin-top: 10px;
    width: 230px;
}

.str-value:focus {
    border: 3px solid rgb(203, 227, 219);
}

.str-value:hover {
    border: 3px solid rgb(112, 189, 151);
}

.str-value__click {
    border-radius: 10px 10px 0 0;
}

.search-field {
    width: 230px;
    padding: 5px 5px;
    border: none;
    border-bottom: 1px solid rgb(205, 205, 205);
    border-top: none;
    outline: none;
}

.combobox__search-select::-webkit-scrollbar {
    width: 12px;
}

.combobox__search-select::-webkit-scrollbar-track {
    background: rgb(241, 241, 241);
}

.combobox__search-select::-webkit-scrollbar-thumb {
    background-color: rgb(183, 183, 183);
    border-radius: 20px;
    border: 3px solid rgb(241, 241, 241);
}

.combobox__search-select::-webkit-scrollbar:horizontal {
    height: 0;
}
</style>