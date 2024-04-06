import { createWebHistory, createRouter } from 'vue-router';

import AuthView from '@/views/AuthView.vue';
import AuthCallbackView from '@/views/AuthCallbackView.vue';
import MainView from '@/views/MainView.vue';
import UserProfile from '@/views/UserProfile.vue';

const routes = [
  { path: '/', component: MainView },
  { path: '/auth', component: AuthView },
  { path: '/auth/callback', component: AuthCallbackView },
  { path: '/user/:id', component: UserProfile },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
