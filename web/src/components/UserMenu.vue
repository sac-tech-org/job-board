<script setup lang="ts">
import { useRouter } from 'vue-router';
import { useUserStore } from '@/stores/user';

import UserMenuItem from '@/components/UserMenuItem.vue';
import { toRefs } from 'vue';

interface Props {
  username: string;
}

const props = defineProps<Props>();
const { username } = toRefs(props)

const router = useRouter();
const userStore = useUserStore();

const emit = defineEmits(['closeMenu']);

function clickAndClose(fn: () => any) {
  fn();
  emit('closeMenu');
}

function goToProfile() {
  router.push('/user/' + username.value);
}

</script>

<template>
  <div class="absolute w-24 flex flex-col divide-y right-0 min-w-40 p-1 rounded-sm bg-slate-100">
    <UserMenuItem @click="clickAndClose(goToProfile)" icon="fa-regular fa-address-card" text="Profile" />
    <UserMenuItem @click="clickAndClose(userStore.logout)" icon="fa-solid fa-arrow-right-from-bracket" text="Log Out" />
  </div>
</template>
