import '@/assets/main.css';
import App from '@/App.vue';
import router from '@/router';
import { createApp } from 'vue';
import { createPinia } from 'pinia';
import { useUserStore } from '@/stores/user';
import SuperTokens from 'supertokens-web-js';
import Session from 'supertokens-web-js/recipe/session';
import EmailVerifcation from 'supertokens-web-js/recipe/emailverification';
import { UserPostConfig, useAPI, type UserPostRequest } from './utils/api';
import ThirdPartyEmailPassword from 'supertokens-web-js/recipe/thirdpartyemailpassword';

const api = useAPI<UserPostRequest, UserPostRequest>(UserPostConfig);

SuperTokens.init({
  appInfo: {
    apiDomain: 'http://localhost:8080',
    apiBasePath: '/auth',
    appName: 'Sac Tech Job Board',
  },
  recipeList: [
    Session.init(),
    ThirdPartyEmailPassword.init(),
    EmailVerifcation.init(),
  ],
});

const pinia = createPinia();
const store = useUserStore(pinia);
await store.checkSession();

createApp(App).use(pinia).use(router).mount('#app');
