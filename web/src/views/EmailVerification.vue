<script setup lang="ts">
import { useRouter } from 'vue-router';
import { useUserStore } from '@/stores/user';
import { computed, onBeforeUnmount, onMounted, ref, toRefs } from 'vue';

import Button from '@/components/Button.vue';
import LoadingOverlay from '@/components/LoadingOverlay.vue';

const router = useRouter();
const userStore = useUserStore();

const verified = ref(false);
const countdownTimer = ref(5)
const alreadyVerified = ref(false);
const { currentRoute } = router
const { errors, hasError, loading, loggedIn } = toRefs(userStore);

const token = computed(() => currentRoute.value.query.token);

function startTimer() {
  const interval = setInterval(() => {
    countdownTimer.value--;
    if (countdownTimer.value === 0) {
      clearInterval(interval);
      router.push({ name: 'home' });
    }
  }, 1000);
}

onMounted(async () => {
  if (loggedIn.value) {
    alreadyVerified.value = await userStore.checkEmailVerification();
    if (alreadyVerified.value) {
      startTimer();
      return
    }
  }

  if (token.value) {
    verified.value = await userStore.verifyEmail();
  }
});

onBeforeUnmount(() => {
  userStore.clearErrors()
});

</script>

<template>
  <main class="w-full min-h-0 grow relative">
    <div class="flex flex-col items-center justify-center h-full">
      <header class="text-xl md:text-3xl text-left w-full font-bold p-2">Email Verification</header>
      <section class="flex flex-col gap-4 grow w-full p-5">
        <template v-if="loading">
          <LoadingOverlay />
        </template>
        <template v-else>
          <div v-if="hasError && errors.errorMessage" class="text-red-500">{{ errors.errorMessage }}</div>

          <div v-if="loggedIn">
            <div v-if="alreadyVerified" class="text-center">Your email is already verified. Redirecting in {{
              countdownTimer }}...</div>
            <div v-else>
              <div>Check your email for an email verification link.</div>
              <div>Need a new verification email? Click the button to send a new one.</div>
              <Button @click="userStore.sendVerificationEmail">Send
                Verification
                Email</Button>
            </div>
          </div>

          <div v-else>
            <div class="flex flex-col gap-2 mb-4">
              <div v-if="!token">Check your email for a verification link.</div>
              <div :class="{ 'text-green-500': verified }">
                {{ verified ?
                  'Email verified! You can now login.' :
                  'Need a new verification email? Login to send a new one.' }}
              </div>
            </div>
            <Button @click="router.push({ name: 'auth' })">Login</Button>
          </div>

        </template>
      </section>
    </div>
  </main>
</template>
