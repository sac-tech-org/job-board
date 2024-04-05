import { createWebHistory, createRouter } from 'vue-router';

import MainView from '@/views/MainView.vue';
import UserProfile from '@/views/UserProfile.vue';

const routes = [
  { path: '/', component: MainView },
  { path: '/user/:id', component: UserProfile },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
