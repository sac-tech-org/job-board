import router from '@/router';
import { defineStore } from 'pinia';
import { type User } from '../types';
import { computed, ref, watch } from 'vue';
import Session from 'supertokens-web-js/recipe/session';
import EmailVerifcation from 'supertokens-web-js/recipe/emailverification';
import {
  useAPI,
  UserGetCurrentConfig,
  UserPutConfig,
  type UserGetResponse,
  type UserPutRequest,
} from '@/utils/api';
import ThirdPartyEmailPassword from 'supertokens-web-js/recipe/thirdpartyemailpassword';

interface Errors {
  errorMessage?: string;
  emailError?: string;
  passwordError?: string;
}

export const useUserStore = defineStore('user', () => {
  const user = ref<User | undefined>(undefined);
  const errors = ref<Errors>({});
  const loading = ref(false);
  const loggedIn = ref(false);

  const hasError = computed(() => {
    return !!errors.value.errorMessage || !!errors.value.emailError || !!errors.value.passwordError;
  });

  async function checkSession() {
    console.log('checking login status');
    const exists = await Session.doesSessionExist();
    console.log('session exists', exists);
    if (exists) {
      loggedIn.value = true;
    }
  }

  function clearErrors() {
    errors.value = {};
  }

  async function getUser() {
    const api = useAPI<UserGetResponse, UserGetResponse>(UserGetCurrentConfig);
    const { data, error } = api;

    watch(data, (val) => {
      user.value = val?.data;
    });
    watch(error, (val) => {
      errors.value.errorMessage = val?.error;
    });

    await withLoading(() => api.execute({}));
  }

  async function updateUser(request: UserPutRequest) {
    const api = useAPI<UserGetResponse, UserGetResponse>(UserPutConfig);
    const { data, error } = api;

    watch(data, (val) => {
      user.value = val?.data;
    });
    watch(error, (val) => {
      errors.value.errorMessage = val?.error;
    });

    await withLoading(() => api.execute(request));
  }

  async function login(email: string, password: string) {
    const response = await withLoading(() =>
      ThirdPartyEmailPassword.emailPasswordSignIn({
        formFields: [
          { id: 'email', value: email },
          { id: 'password', value: password },
        ],
      }),
    );

    if (response.status === 'WRONG_CREDENTIALS_ERROR') {
      errors.value.errorMessage = 'Invalid credentials';
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
    user.value = undefined;
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

  async function sendVerificationEmail() {
    const response = await withLoading(EmailVerifcation.sendVerificationEmail);
    if (response.status === 'EMAIL_ALREADY_VERIFIED_ERROR') {
      errors.value.errorMessage = 'Email already verified';
    }
  }

  async function signUp(
    email: string,
    firstName: string,
    lastName: string,
    password: string,
    username: string,
  ) {
    console.log(email, firstName, lastName, password, username);
    const response = await withLoading(() =>
      ThirdPartyEmailPassword.emailPasswordSignUp({
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
        options: {
          preAPIHook: async (input) => {
            // We add our own userContext object to the request body so the backend can use the
            // values to check for duplicate usernames, and then create a database record for the user.
            let requestInit = input.requestInit;

            const body = {
              ...JSON.parse(requestInit.body as string),
              userContext: {
                firstName,
                lastName,
                username,
              },
            };

            requestInit = {
              ...requestInit,
              body: JSON.stringify(body),
            };

            return {
              url: input.url,
              requestInit,
            };
          },
        },
        userContext: {
          firstName,
          lastName,
          username,
        },
      }),
    );

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
    if (response.status === 'SIGN_UP_NOT_ALLOWED') {
      errors.value.errorMessage = 'Sign up not allowed';
    }

    await EmailVerifcation.sendVerificationEmail();

    router.push('/');
  }

  async function verifyEmail(): Promise<boolean> {
    // const response = await EmailVerifcation.verifyEmail();
    const response = await withLoading(EmailVerifcation.verifyEmail);
    if (response.status === 'EMAIL_VERIFICATION_INVALID_TOKEN_ERROR') {
      errors.value.errorMessage = 'Invalid token';
      return false;
    }

    return true;
  }

  async function withLoading<T>(fn: () => Promise<T>) {
    loading.value = true;
    try {
      return await fn();
    } finally {
      loading.value = false;
    }
  }

  return {
    checkSession,
    clearErrors,
    errors,
    getUser,
    hasError,
    loading,
    loggedIn,
    login,
    logout,
    selectSocicalProvider,
    sendVerificationEmail,
    signUp,
    updateUser,
    user,
    verifyEmail,
  };
});

async function handleGithubSelect() {
  const authURL = await ThirdPartyEmailPassword.getAuthorisationURLWithQueryParamsAndSetState({
    thirdPartyId: 'github',
    frontendRedirectURI: 'http://localhost:5173/auth/callback/github',
  });

  window.location.assign(authURL);
}

async function handleGoogleSelect() {
  const authURL = await ThirdPartyEmailPassword.getAuthorisationURLWithQueryParamsAndSetState({
    thirdPartyId: 'google',
    frontendRedirectURI: 'http://localhost:5173/auth/callback/google',
  });

  window.location.assign(authURL);
}
