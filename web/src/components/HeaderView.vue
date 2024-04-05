<script setup lang="ts">
import { ref } from 'vue';

import FAIcon from '@/components/FAIcon.vue';
import PopoverMenu from '@/components/PopoverMenu.vue';
import UserMenu from '@/components/UserMenu.vue';

interface Props {
  loggedIn: boolean;
}

const menuOpen = ref(false);
const { loggedIn } = defineProps<Props>();
const emit = defineEmits(['openLoginModal']);
</script>

<template>
  <header class="w-dvw px-5 py-3 border-b border-slate-300 sticky top-0">
    <div class="h-10 flex justify-between items-center">
      <h1 class="text-4xl max-h-full">Sac Tech Job Board</h1>
      <div class="flex justify-between gap-2">
        <button v-if="loggedIn" @click="menuOpen = !menuOpen">
          <FAIcon icon="fa-regular fa-circle-user" size="2xl" />
        </button>
        <button v-else @click="emit('openLoginModal')">
          <span class="pr-2 text-lg">Login</span>
          <FAIcon icon="fa-solid fa-arrow-right-to-bracket" size="xl" />
        </button>
      </div>
    </div>
  </header>

  <PopoverMenu :open="menuOpen">
    <UserMenu />
  </PopoverMenu>
</template>
