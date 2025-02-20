import { createRouter, createWebHistory } from 'vue-router';

// Import your components/pages
import Authorization from '../views/Authorization.vue';
import Home from '../views/Home.vue';

// Define routes
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
];

// Create the router instance
const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes,
});

export default router;

