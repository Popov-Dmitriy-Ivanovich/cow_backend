import { createRouter, createWebHistory } from 'vue-router';
import MainPage from '@/pages/MainPage.vue';
import LoginPage from '@/pages/LoginPage.vue';
import HelpPage from '@/pages/HelpPage.vue';
import RegistrationPage from '@/pages/RegistrationPage.vue';
// import AnimalsPage from '@/pages/AnimalsPage.vue';
import ConcretAnimalPage from '@/pages/ConcretAnimalPage.vue';
import ParticipantPage from '@/pages/ParticipantPage.vue';
import HozPage from '@/pages/HozPage.vue';
import TestPage from '@/pages/TestPage.vue';
import AnaliticPage from '@/pages/AnaliticPage.vue';
import AnaliticYear from '@/pages/AnaliticYear.vue';
import AnaliticRegion from '@/pages/AnaliticRegion.vue';
import AnaliticDistrict from '@/pages/AnaliticDistrict.vue';
import AnaliticHolding from '@/pages/AnaliticHolding.vue'

const routes = [
    {
        path: '/',
        component: MainPage
    },
    {
        path: '/login',
        component: LoginPage
    },
    {
        path: '/help',
        component: HelpPage
    },
    {
        path: '/registration',
        component: RegistrationPage
    },
    {
        path: '/animals',
        component: TestPage
    },
    {
        path: '/animals/:id',
        component: ConcretAnimalPage
    },
    {
        path: '/participants',
        component: ParticipantPage
    },
    {
        path: '/hoz',
        component: HozPage
    },

    {
        path: '/analytics',
        component: AnaliticPage
    },
    {
        path: '/analytics/:id',
        component: AnaliticYear
    },
    {
        path: '/analytics/:id/:region',
        component: AnaliticRegion
    },
    {
        path: '/analytics/:id/:region/:district',
        component: AnaliticDistrict
    },
    {
        path: '/analytics/:id/:region/:district/:hold',
        component: AnaliticHolding
    },
];

// const scrollBehavior = function () {
//     return { top: 0, left: 0 };
// };

const router = createRouter({
    routes,
    history: createWebHistory(process.env.BASE_URL),
    // scrollBehavior
});

export default router;