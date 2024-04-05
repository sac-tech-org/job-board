<script setup lang="ts">
import { computed } from 'vue';

interface Props {
  job: {
    id: number;
    title: string;
    description: string;
  };
  disableClick: boolean;
}

const props = defineProps<Props>();
const emit = defineEmits(['selected']);

const disabled = computed(() => props.disableClick);

function handleClick() {
  if (!disabled.value) {
    emit('selected');
  }
}
</script>

<template>
  <div
    class="flex justify-between gap-x-6 px-3 py-5 w-full cursor-pointer hover:shadow-lg"
    @click="handleClick"
  >
    <div class="flex min-w-0 gap-x-4">
      <img
        class="size-12 flex-none rounded-full"
        src="@/assets/trash-panda.svg?url"
        alt="Trash Panda"
      />
      <div class="min-w-0 flex-auto">
        <p class="text-sm font-semibold leading-6">{{ job.title }}</p>
        <p class="mt-1 truncate text-xs leading-5">{{ job.description }}</p>
      </div>
    </div>
    <div class="hidden shrink-0 sm:flex sm:flex-col sm:items-end">
      <p class="text-sm leading-6">Company Name</p>
      <p class="mt-1 text-xs leading-5">Something small</p>
    </div>
  </div>
</template>
