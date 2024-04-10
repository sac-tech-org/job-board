<script setup lang="ts">
import FormGroup from '@/components/FormGroup.vue';

const firstName = defineModel<string>('firstName', { required: true });
const lastName = defineModel<string>('lastName', { required: true });
const username = defineModel<string>('username', { required: true });

const emit = defineEmits(['submitForm']);
</script>
<template>
  <form class="flex flex-col gap-4 p-3" @submit.prevent="emit('submitForm')">
    <FormGroup label="First Name" name="firstName" :placeholder="firstName || 'First Name'" type="text"
      v-model:val="firstName" required />
    <FormGroup label="Last Name" name="lastName" :placeholder="lastName || 'Last Name'" type="text"
      v-model:val="lastName" required />
    <FormGroup label="Username" name="username" :placeholder="username || 'username'" type="text" v-model:val="username"
      required />

    <button type="submit" class="bg-blue-500 text-white rounded max-w-32 p-2">Save</button>
  </form>
</template>
