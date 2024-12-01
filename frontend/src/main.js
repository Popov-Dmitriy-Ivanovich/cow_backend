import { createApp } from 'vue'
import App from './App.vue'
import router from './routers/routers'
import VueApexCharts from "vue3-apexcharts";

const app = createApp(App);

app.use(router).mount('#app');
app.use(VueApexCharts);