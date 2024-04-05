<script setup lang="ts">
import { useRouter } from 'vue-router';

const router = useRouter();

function goHome() {
  router.push('/');
}
</script>

<template>
  <div class="w-96 size-96">
    hello
    <button @click="goHome">Go Home</button>
  </div>
</template>
