<script setup lang="ts">
import { useRouter } from 'vue-router';
import { useUserStore } from '@/stores/user';

import UserMenuItem from '@/components/UserMenuItem.vue';

const router = useRouter();
const userStore = useUserStore();

function goToProfile() {
  router.push('/user/1');
}
</script>

<template>
  <div class="w-24 flex flex-col divide-y">
    <UserMenuItem @click="goToProfile" icon="fa-regular fa-address-card" text="Profile" />
    <UserMenuItem
      @click="userStore.logout"
      icon="fa-solid fa-arrow-right-from-bracket"
      text="Log Out"
    />
  </div>
</template>
