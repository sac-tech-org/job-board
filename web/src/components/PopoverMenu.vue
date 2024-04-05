<script setup lang="ts">
interface Props {
  open: boolean;
}

const { open } = defineProps<Props>();
</script>

<template>
  <dialog :open class="relative">
    <slot></slot>
  </dialog>
</template>
