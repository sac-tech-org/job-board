<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import Session from "supertokens-web-js/recipe/session";
import ThirdPartyEmailPassword from 'supertokens-web-js/recipe/thirdpartyemailpassword';

import AuthForm from '@/components/AuthForm.vue';
import AuthProviderButton from '@/components/AuthProviderButton.vue';

const router = useRouter();

const isSignIn = ref(true);
const email = ref('');
const password = ref('');

const error = ref(false);
const errorMessage = ref('Something went wrong');
const emailError = ref('');
const passwordError = ref('');

const providers: { name: string, color: string, handler: () => Promise<void> }[] = [
  { name: 'GitHub', color: '#fff', handler: handleGithubSelect },
  { name: 'Google', color: '#fff', handler: handleGoogleSelect },
  // { name: 'Apple', color: '#fff', handler: handleAppleSelect },
  // { name: 'LinkedIn', color: '#fff', handler: handleLinkedInSelect },
]

async function checkForSession() {
  if (await Session.doesSessionExist()) {
    router.push('/');
  }
}

function handleSubmit() {
  // e.preventDefault();

  console.log('handleSubmit', email.value, password.value)

  if (isSignIn.value) {
    signIn();
  } else {
    signUp();
  }
}

async function handleAppleSelect() {
  const authURL = await ThirdPartyEmailPassword.getAuthorisationURLWithQueryParamsAndSetState({
    thirdPartyId: 'apple',
    frontendRedirectURI: 'http://localhost:5173/auth/callback/apple',
    redirectURIOnProviderDashboard: 'http://localhost:5173/auth/callback/apple',
  })

  window.location.assign(authURL);
}

async function handleGithubSelect() {
  const authURL = await ThirdPartyEmailPassword.getAuthorisationURLWithQueryParamsAndSetState({
    thirdPartyId: 'github',
    frontendRedirectURI: 'http://localhost:5173/auth/callback/github',
  })

  window.location.assign(authURL);
}

async function handleGoogleSelect() {
  const authURL = await ThirdPartyEmailPassword.getAuthorisationURLWithQueryParamsAndSetState({
    thirdPartyId: 'google',
    frontendRedirectURI: 'http://localhost:5173/auth/callback/google',
  })

  window.location.assign(authURL);
}

async function handleLinkedInSelect() {
  const authURL = await ThirdPartyEmailPassword.getAuthorisationURLWithQueryParamsAndSetState({
    thirdPartyId: 'linkedin',
    frontendRedirectURI: 'http://localhost:5173/auth/callback/linkedin',
  })

  window.location.assign(authURL);
}

async function handleProviderSelected(provider: string) {
  if (provider === 'apple') {
    handleAppleSelect();
  } else if (provider === 'GitHub') {
    handleGithubSelect();
  } else if (provider === 'Google') {
    handleGoogleSelect();
  } else if (provider === 'linkedin') {
    handleLinkedInSelect();
  }
}

function goToSignUp() {
  isSignIn.value = false;
}

function goToSignIn() {
  isSignIn.value = true;
}

async function signIn() {
  const response = await ThirdPartyEmailPassword.emailPasswordSignIn({
    formFields: [
      {
        id: 'email',
        value: email.value,
      },
      {
        id: 'password',
        value: password.value,
      },
    ],
  });

  if (response.status === 'WRONG_CREDENTIALS_ERROR') {
    errorMessage.value = 'Invalid credentials';
    error.value = true;
    return;
  }

  if (response.status === 'FIELD_ERROR') {
    response.formFields.forEach(i => {
      if (i.id === 'email') {
        emailError.value = i.error;
      } else if (i.id === 'password') {
        passwordError.value = i.error;
      }
    });
    return;
  }

  // router.push('/');
}

async function signOut() {
  await Session.signOut();
  router.push('/');
}

async function signUp() {
  const response = await ThirdPartyEmailPassword.emailPasswordSignUp({
    formFields: [
      {
        id: "email",
        value: email.value,
      },
      {
        id: "password",
        value: password.value,
      },
    ],
  });

  if (response.status === "FIELD_ERROR") {
    response.formFields.forEach((item) => {
      if (item.id === "email") {
        // this means that something was wrong with the entered email.
        // probably that it's not a valid email (from a syntax point of view)
        emailError.value = item.error;
      } else if (item.id === "password") {
        // this means that something was wrong with the entered password.
        // probably it doesn't meet the password validation criteria on the backend.
        passwordError.value = item.error;
      }
    });
    return;
  }
}

function validateEmail(email: string): boolean {
  return email.includes('@');
}
</script>

<template>
  <div class="flex flex-col justify-between size-full">
    <div class="my-6 mx-auto w-[420px] text-center shadow-lg bg-white rounded-xl">
      <div v-if="error" class="error-container">
        <div class="error-message">{{ errorMessage }}</div>
      </div>
      <div class="m-auto w-3/4 pt-10">
        <div class="text-2xl tracking-wider font-extrabold mb-0.5 text-black">{{ isSignIn ? 'Sign In' : 'Sign Up' }}
        </div>
        <div class="text-sm font-light tracking-wide text-gray-500">
          <span>
            {{ isSignIn ? 'Not yet registered?' : 'Already have an account?' }}
            <span class="cursor-pointer text-blue-500 hover:text-blue-700"
              @click="isSignIn ? goToSignUp() : goToSignIn()">
              {{ isSignIn ? 'Sign Up' : 'Sign In' }}
            </span>
          </span>
        </div>

        <div class="divider-container">
          <div class="divider" />
        </div>

        <AuthProviderButton v-for="p in providers" :key="p.name" :provider="p"
          @selected="handleProviderSelected(p.name)" />

        <div class="divider-container">
          <div class="divider" />
          <div class="divider-text">or</div>
          <div class="divider" />
        </div>

        <AuthForm v-model:email="email" v-model:password="password" :emailError="emailError"
          :passwordError="passwordError" :isSignIn="isSignIn" @submitForm="handleSubmit()" />

      </div>
      <div v-if="isSignIn">
        <router-link :to="{ path: `/auth/reset-password` }"> Forgot Password? </router-link>
      </div>
      <div style="margin-bottom: 10px" />
    </div>
  </div>

</template>

<style scoped>
.forgot-password-link {
  padding-left: 3px;
  padding-right: 3px;
  cursor: pointer;
  line-height: 26px;
  font-size: 14px;
  font-weight: 300;
  letter-spacing: 0.4px;
  color: rgb(101, 101, 101);
  margin-top: 10px;
}

.error-container {
  display: flex;
  position: absolute;
  width: calc(100% - 24px);
  background-color: #ffcdd2;
  justify-content: center;
  align-items: center;
  padding-top: 2px;
  padding-bottom: 2px;
  margin-left: 12px;
  margin-right: 12px;
  margin-top: 4px;
  border-radius: 6px;
  box-sizing: border-box;
  border-width: 1px;
  border: 1px solid #ff1744;
}

.divider {
  height: 0.3px;
  width: 100%;
  background-color: rgb(221, 221, 221);
}

.divider-container {
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  margin-top: 1em;
  margin-bottom: 1em;
}

.divider-text {
  margin-left: 12px;
  margin-right: 12px;
  color: rgb(101, 101, 101);
  flex: 1 1 0%;
}

@keyframes slideTop {
  0% {
    transform: translateY(-5px);
  }

  100% {
    transform: translateY(0px);
  }
}
</style>
