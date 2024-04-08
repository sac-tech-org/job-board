<script setup lang="ts">
import { ref } from 'vue';
import { storeToRefs } from 'pinia';
import { useUserStore } from '@/stores/user';

import AuthForm from '@/components/AuthForm.vue';
import AuthProviderButton from '@/components/AuthProviderButton.vue';
import DividerContainer from '@/components/DividerContainer.vue';
import DividerLine from '@/components/DividerLine.vue';

const userStore = useUserStore();
const isSignIn = ref(true);
const email = ref('');
const password = ref('');

const { errors, hasError } = storeToRefs(userStore);

const providers: { name: 'GitHub' | 'Google', color: string, }[] = [
  { name: 'GitHub', color: '#fff' },
  { name: 'Google', color: '#fff' },
]

async function handleSubmit() {
  userStore.clearErrors();
  if (isSignIn.value) {
    await userStore.login(email.value, password.value);
  } else {
    await userStore.signUp(email.value, password.value);
  }
}

async function handleProviderSelected(provider: 'GitHub' | 'Google') {
  await userStore.selectSocicalProvider(provider);
}

function goToSignUp() {
  userStore.clearErrors();
  isSignIn.value = false;
}

function goToSignIn() {
  userStore.clearErrors();
  isSignIn.value = true;
}

</script>

<template>
  <div class="flex flex-col justify-between size-full">
    <div class="my-6 mx-auto w-[420px] text-center shadow-lg bg-white rounded-xl">
      <div v-if="hasError && errors.errorMessage"
        class="flex justify-center items-center w-[calc(100%-24px)] bg-red-300 py-0.5 mx-auto mt-1 rounded-md border border-solid border-red-600">
        <div class="text-neutral-50">{{ errors.errorMessage }}</div>
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

        <DividerContainer>
          <DividerLine />
        </DividerContainer>

        <AuthProviderButton v-for="p in providers" :key="p.name" :provider="p"
          @selected="handleProviderSelected(p.name)" />

        <DividerContainer>
          <DividerLine />
          <span class="ml-3 mr-3 text-slate-500 flex-1">or</span>
          <DividerLine />
        </DividerContainer>

        <AuthForm v-model:email="email" v-model:password="password" :emailError="errors.emailError"
          :passwordError="errors.passwordError" :isSignIn="isSignIn" @submitForm="handleSubmit()" />
      </div>

      <RouterLink v-if="isSignIn" class="pl-1 pr-1 cursor-pointer text-sm font-light tracking-wide text-slate-500 mt-2"
        :to="{ path: `/auth/reset-password` }"> Forgot
        Password?
      </RouterLink>
      <div class="mb-2" />
    </div>
  </div>
</template>
