import { ref } from 'vue';
import router from '@/router';
import { defineStore } from 'pinia';
import { type User } from '../types';
import Session from 'supertokens-web-js/recipe/session';
import ThirdPartyEmailPassword from 'supertokens-web-js/recipe/thirdpartyemailpassword';

interface Errors {
  errorMessage?: string;
  emailError?: string;
  passwordError?: string;
}

export const useUserStore = defineStore('user', () => {
  const user = ref<User | null>(null);
  const hasError = ref(false);
  const errors = ref<Errors>({ errorMessage: 'Something went wrong' });
  const loggedIn = ref(false);

  function clearErrors() {
    errors.value = {};
    hasError.value = false;
  }

  async function login(email: string, password: string) {
    const response = await ThirdPartyEmailPassword.emailPasswordSignIn({
      formFields: [
        { id: 'email', value: email },
        { id: 'password', value: password },
      ],
    });

    if (response.status === 'WRONG_CREDENTIALS_ERROR') {
      errors.value.errorMessage = 'Invalid credentials';
      hasError.value = true;
      return;
    }

    if (response.status === 'FIELD_ERROR') {
      response.formFields.forEach((i) => {
        if (i.id === 'email') {
          errors.value.emailError = i.error;
        } else if (i.id === 'password') {
          errors.value.passwordError = i.error;
        }
      });
      return;
    }

    loggedIn.value = true;
    clearErrors();
    router.push('/');
  }

  async function logout() {
    await Session.signOut();
    user.value = null;
    loggedIn.value = false;
    clearErrors();
  }

  async function selectSocicalProvider(provider: 'GitHub' | 'Google') {
    if (provider === 'GitHub') {
      await handleGithubSelect();
    } else if (provider === 'Google') {
      await handleGoogleSelect();
    }
  }

  async function signUp(email: string, password: string) {
    const response = await ThirdPartyEmailPassword.emailPasswordSignUp({
      formFields: [
        {
          id: 'email',
          value: email,
        },
        {
          id: 'password',
          value: password,
        },
      ],
    });

    if (response.status === 'FIELD_ERROR') {
      response.formFields.forEach((item) => {
        if (item.id === 'email') {
          // this means that something was wrong with the entered email.
          // probably that it's not a valid email (from a syntax point of view)
          errors.value.emailError = item.error;
        } else if (item.id === 'password') {
          // this means that something was wrong with the entered password.
          // probably it doesn't meet the password validation criteria on the backend.
          errors.value.passwordError = item.error;
        }
      });
      return;
    }
  }

  return {
    clearErrors,
    errors,
    hasError,
    loggedIn,
    login,
    logout,
    selectSocicalProvider,
    signUp,
    user,
  };
});

async function handleGithubSelect() {
  const authURL = await ThirdPartyEmailPassword.getAuthorisationURLWithQueryParamsAndSetState({
    thirdPartyId: 'github',
    frontendRedirectURI: 'http://localhost:5173/auth/callback/github',
  });

  router.replace(authURL);
}

async function handleGoogleSelect() {
  const authURL = await ThirdPartyEmailPassword.getAuthorisationURLWithQueryParamsAndSetState({
    thirdPartyId: 'google',
    frontendRedirectURI: 'http://localhost:5173/auth/callback/google',
  });

  router.replace(authURL);
}
