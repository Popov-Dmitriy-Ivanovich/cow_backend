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
import MainChartComp from '@/pages/MainChartComp.vue';

const routes = [
    {
        path: '/',
        component: MainPage
    },
    {
        path: '/login',
        component: LoginPage,
        meta: {noAuth: true}
    },
    {
        path: '/help',
        component: HelpPage
    },
    {
        path: '/registration',
        component: RegistrationPage,
        meta: {noAuth: true}
    },
    {
        path: '/animals',
        component: TestPage,
        meta: {requiresAuth: true}
    },
    {
        path: '/animals/:id',
        component: ConcretAnimalPage,
        meta: {requiresAuth: true}
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
        component: AnaliticPage,
        meta: {requiresAuth: true},
        children: [
            {
                path: '',
                component: MainChartComp,
            },
            {
                path: ':id',
                component: AnaliticYear
            },
            {
                path: ':id/:region',
                component: AnaliticRegion
            },
            {
                path: ':id/:region/:district',
                component: AnaliticDistrict
            }
        ]
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

router.beforeEach((to) => {
    if(to.meta.requiresAuth && !(localStorage.getItem('jwt'))) return '/';
    if(to.meta.noAuth && (localStorage.getItem('jwt'))) return '/';
})

export default router;