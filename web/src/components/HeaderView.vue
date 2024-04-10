<script setup lang="ts">
import { onMounted, ref, toRefs, watch } from 'vue';
import { useUserStore } from '@/stores/user'

import FAIcon from '@/components/FAIcon.vue';
import DropdownMenu from '@/components/DropdownMenu.vue';
import UserMenu from '@/components/UserMenu.vue';

const userStore = useUserStore()
const { getUser } = userStore;
const { loggedIn, user } = toRefs(userStore)

watch(loggedIn, (val) => {
  if (val) getUser()
})

const menuOpen = ref(false);

onMounted(() => {
  if (loggedIn.value) {
    getUser()
  }
})
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
            <UserMenu @close-menu="menuOpen = false" :username="user?.username || 'me'" />
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
