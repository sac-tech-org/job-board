<script setup lang="ts">
import FAIcon from '@/components/FAIcon.vue';

interface Props {
  icon: string;
  text: string;
}

const { icon, text } = defineProps<Props>();
</script>

<template>
  <div class="cursor-pointer flex justify-between items-center">
    <FAIcon :icon="icon" size="lg" />
    <span class="text-md">{{ text }}</span>
  </div>
</template>
