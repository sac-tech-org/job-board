import './assets/main.css';
import App from './App.vue';
import router from './router';
import { createApp } from 'vue';
import { createPinia } from 'pinia';
import SuperTokens from 'supertokens-web-js';
import Session from 'supertokens-web-js/recipe/session';
import ThirdPartyEmailPassword from 'supertokens-web-js/recipe/thirdpartyemailpassword';

SuperTokens.init({
  appInfo: {
    apiDomain: 'http://localhost:8080',
    apiBasePath: '/auth',
    appName: 'Sac Tech Job Board',
  },
  recipeList: [Session.init(), ThirdPartyEmailPassword.init()],
});

const pinia = createPinia();

createApp(App).use(router).use(pinia).mount('#app');
