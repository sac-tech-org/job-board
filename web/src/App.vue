<script setup lang="ts">
import { ref } from 'vue';
import { useUserStore } from '@/stores/user';

import AuthView from '@/views/AuthView.vue';
import HeaderView from '@/components/HeaderView.vue';
import ModalView from '@/components/ModalView.vue';

const modalOpen = ref(false);

const userStore = useUserStore();

userStore.$subscribe((_, state) => {
  if (!state.loggedIn) {
    closeLoginModal();
  }
});

function closeLoginModal() {
  modalOpen.value = false;
}

function openLoginModal() {
  modalOpen.value = true;
}
</script>

<template>
  <HeaderView :loggedIn="userStore.loggedIn" @open-login-modal="openLoginModal" />
  <RouterView />

  <ModalView :open="modalOpen" @modalClosed="closeLoginModal" title="Login" class="h-3/5 w-1/2 top-[10%]">
    <AuthView />
  </ModalView>
</template>
