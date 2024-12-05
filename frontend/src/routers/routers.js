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
import AnaliticYear from '@/pages/AnaliticYear.vue'

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
    // {
    //     path: '/test',
    //     component: TestPage
    // }
    {
        path: '/analytics',
        component: AnaliticPage
    },
    {
        path: '/analytics/:id',
        component: AnaliticYear
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