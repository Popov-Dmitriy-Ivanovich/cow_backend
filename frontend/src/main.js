import { createApp } from 'vue'
import App from './App.vue'
import router from './routers/routers'
import VueApexCharts from "vue3-apexcharts";
import { createStore } from 'vuex';

const store = createStore({
    state(){
        return{
            filters: {},
            option: '',
            isLogged: Boolean(localStorage.getItem('jwt')),
        }
    },
    getters: {
        FILTERS(state) {
            return state.filters;
        },
    },
    mutations: {
        SET_FILTERS(state, payload) {
            state.filters = payload;
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


