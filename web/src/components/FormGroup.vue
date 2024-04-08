<script setup lang="ts">
import { toRefs } from 'vue';

interface Props {
  autocomplete?: string;
  label: string;
  name: string;
  placeholder: string;
  required: boolean;
  type: string;
}

const props = defineProps<Props>();
const val = defineModel<string>('val', { required: true });
const { autocomplete, label, name, placeholder, required } = toRefs(props);
</script>

<template>
  <div class="flex flex-col gap-1 max-w-64">
    <label :for="name" class="font-semibold">{{ label }}</label>
    <input :name="name" :placeholder="placeholder" :autocomplete="autocomplete" :required="required" :type="type"
      class="rounded" v-model="val" />
  </div>
</template>
