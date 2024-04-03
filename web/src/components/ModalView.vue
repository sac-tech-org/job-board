<script setup lang="ts">
import FAIcon from './FAIcon.vue';

interface Props {
  open: boolean;
  title: string;
}

const { open, title } = defineProps<Props>();
const emit = defineEmits(["modalClosed"]);

</script>

<template>
  <dialog :open class="flex-auto py-2 px-4 rounded-lg shadow-lg">
    <header class="flex justify-between">
      <h2 class="text-lg font-semibold">{{ title }}</h2>
      <button @click="emit('modalClosed')">
        <FAIcon :icon="'fa-regular fa-circle-xmark'" :size="'xl'" hover-grow />
      </button>
    </header>
    <hr class="my-2" />
    <slot></slot>
  </dialog>
</template>
