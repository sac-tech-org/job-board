import { createWebHistory, createRouter } from 'vue-router';

import AuthView from '@/views/AuthView.vue';
import AuthCallbackView from '@/views/AuthCallbackView.vue';
import MainView from '@/views/MainView.vue';
import UserProfile from '@/views/UserProfile.vue';

const routes = [
  { path: '/', name: 'home', component: MainView },
  { path: '/auth', name: 'auth', component: AuthView },
  { path: '/auth/callback/:provider', name: 'authCallback', component: AuthCallbackView },
  { path: '/user/:id', name: 'userProfile', component: UserProfile },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
