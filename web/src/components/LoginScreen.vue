<script setup lang="ts">
import { useUserStore } from '@/stores/user';

const userStore = useUserStore();
</script>

<template>
  <div>
    <button @click="userStore.login">Login</button>
  </div>
</template>
