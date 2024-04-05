<script setup lang="ts">
import { ref, watch } from 'vue';
import { useUserStore } from '@/stores/user';

const userStore = useUserStore();

const email = ref('');

watch(email, (value) => {
  if (value) {
    if (!validateEmail(value)) {
      console.log('Invalid email');
    }
  }
});

function validateEmail(email: string): boolean {
  return email.includes('@');
}
</script>

<template>
  <div>
    <form>
      <input type="email" placeholder="Email" v-model="email" class="rounded" required />
    </form>
    <button @click="userStore.login">Login</button>
  </div>
</template>
