import { createRouter, createWebHistory } from 'vue-router';

import Authorization from '../views/Authorization.vue';
import Home from '../views/Home.vue';
import Transactions from '../views/Transactions.vue';

const routes = [
    {
        path: '/login',
        name: 'Login',
        component: Authorization,
    },
    {
        path: '/',
        name: 'Home',
        component: Home,
    },
    {
        path: '/transactions',
        name: 'Transaction',
        component: Transactions,
    },
];

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes,
});

export default router;

