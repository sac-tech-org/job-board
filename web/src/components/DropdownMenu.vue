<script setup lang="ts">
interface Props {
  open: boolean;
}

const { open } = defineProps<Props>();
</script>

<template>
  <div class="absolute right-0 z-10 shadow-lg" v-bind:class="open ? 'block' : 'hidden'">
    <slot></slot>
  </div>
</template>
