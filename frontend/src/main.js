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
            isCows: true,
            isBulls: false,
            isChild: false,
            showFilters: {},
            currentHoz: {},
        }
    },
    getters: {
        FILTERS(state) {
            return state.filters;
        },
        FILTERS_2(state) {
            return state.filters_2;
        },
        ISCOWS(state) {
            return state.isCows;
        },
        ISBULLS(state) {
            return state.isBulls;
        },
        ISCHILD(state) {
            return state.isChild;
        },
        SHOWFILTERS(state) {
            return state.showFilters;
        },
        CURRENTHOZ(state) {
            return state.currentHoz;
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
        },
        SET_ISCOWS(state, payload) {
            state.isCows = payload;
        },
        SET_ISBULLS(state, payload) {
            state.isBulls = payload;
        },
        SET_ISCHILD(state, payload) {
            state.isChild = payload;
        },
        SET_SHOWFILTERS(state, payload) {
            state.showFilters = payload;
        },
        SET_CURRENTHOZ(state, payload) {
            state.currentHoz = payload;
        }
    }
})

const app = createApp(App);
app.use(store);
app.use(router).mount('#app');
app.use(VueApexCharts);


