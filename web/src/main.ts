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
    ThirdPartyEmailPassword.init({
      override: {
        functions: (o) => {
          return {
            ...o,
            emailPasswordSignUp: async function (input) {
              const user = await o.emailPasswordSignUp(input);
              if (user.status === 'OK') {
                const { error, execute } = api;
                const req: UserPostRequest = {
                  body: {
                    firstName: input.userContext.firstName || '',
                    id: user.user.id,
                    lastName: input.userContext.lastName || '',
                    username: input.userContext.username || '',
                  },
                };
                await execute(req);
                if (error.value) {
                  console.error('Error creating user', error.value);
                }
              }

              return user;
            },
            thirdPartySignInAndUp: async function (input) {
              const user = await o.thirdPartySignInAndUp(input);

              if (user.status === 'OK') {
                const { error, execute } = api;
                const req: UserPostRequest = {
                  body: {
                    firstName: input.userContext?.firstName || '',
                    id: user.user.id,
                    lastName: input.userContext?.lastName || '',
                    username: input.userContext?.username || '',
                  },
                };
                await execute(req);
                if (error.value) {
                  console.error('Error creating user', error.value);
                }
              }

              return user;
            },
          };
        },
      },
    }),
    EmailVerifcation.init(),
  ],
});

const pinia = createPinia();
const store = useUserStore(pinia);
await store.checkSession();

createApp(App).use(pinia).use(router).mount('#app');
