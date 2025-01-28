import { createApp } from 'vue'
import App from './App.vue'
import router from './routers/routers'
import VueApexCharts from "vue3-apexcharts";
import { createStore } from 'vuex';

const store = createStore({
    state(){
        return{
            filters: {},
            filters_2: {},
            option: '',
            isLogged: Boolean(localStorage.getItem('jwt')),
        }
    },
    getters: {
        FILTERS(state) {
            return state.filters;
        },
        FILTERS_2(state) {
            return state.filters_2;
        }
    },
    mutations: {
        SET_FILTERS(state, payload) {
            state.filters = payload;
        },
        SET_FILTERS_2(state, payload) {
            state.filters_2 = payload;
        },
        SET_OPTION(state, payload) {
            state.option = payload;
        },
        SET_ISLOGGED(state, payload) {
            state.isLoad = payload;
        }
    }
})

const app = createApp(App);
app.use(store);
app.use(router).mount('#app');
app.use(VueApexCharts);


