<script setup lang="ts">
import { onMounted, ref, toRefs } from 'vue';
import { useRouter } from 'vue-router';
import { useUserStore } from '@/stores/user';

const router = useRouter();
const userStore = useUserStore();

const verified = ref(false);
const { currentRoute } = toRefs(router);
const { errors, hasError, loading } = toRefs(userStore);

onMounted(async () => {
  if (currentRoute.value.query.token) {
    verified.value = await userStore.verifyEmail();
  }
});
</script>

<template>
  <div>
    <div v-if="loading" class="text-center">Loading...</div>
    <div v-else-if="hasError && errors.errorMessage" class="text-center text-red-500">{{ errors.errorMessage }}</div>
    <div v-else-if="verified" class="text-center text-green-500">Email verified! You can now login.</div>
  </div>
</template>
