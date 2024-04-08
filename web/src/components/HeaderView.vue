<script setup lang="ts">
import { ref, toRefs } from 'vue';

import FAIcon from '@/components/FAIcon.vue';
import DropdownMenu from '@/components/DropdownMenu.vue';
import UserMenu from '@/components/UserMenu.vue';

interface Props {
  loggedIn: boolean;
}

const menuOpen = ref(false);
const props = defineProps<Props>();
const { loggedIn } = toRefs(props);
</script>

<template>
  <header class="w-full px-1 py-2 border-b border-slate-300 sticky top-0">
    <div class="h-10 flex justify-between items-center">
      <RouterLink to="/">
        <h1 class="text-xl md:text-4xl max-h-full">Sac Tech Job Board</h1>
      </RouterLink>
      <div class="flex justify-between gap-2">
        <div v-if="loggedIn">
          <button @click="menuOpen = !menuOpen">
            <FAIcon icon="fa-regular fa-circle-user" size="2xl" />
          </button>
          <DropdownMenu :open="menuOpen">
            <UserMenu @close-menu="menuOpen = false" />
          </DropdownMenu>
        </div>
        <button v-else @click="$router.push('/auth')">
          <span class="pr-2 text-lg">Login</span>
          <FAIcon icon="fa-solid fa-arrow-right-to-bracket" size="xl" />
        </button>
      </div>
    </div>
  </header>
</template>
