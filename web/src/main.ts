import '@/assets/main.css';
import App from '@/App.vue';
import router from '@/router';
import { createApp } from 'vue';
import { createPinia } from 'pinia';
import { useUserStore } from '@/stores/user';
import SuperTokens from 'supertokens-web-js';
import Session from 'supertokens-web-js/recipe/session';
import EmailVerifcation from 'supertokens-web-js/recipe/emailverification';
import ThirdPartyEmailPassword from 'supertokens-web-js/recipe/thirdpartyemailpassword';

SuperTokens.init({
  appInfo: {
    apiDomain: 'https://api.jobs.sac-tech.com',
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
store.checkSession();

createApp(App).use(pinia).use(router).mount('#app');
