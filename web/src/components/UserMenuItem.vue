<script setup lang="ts">
import FAIcon from '@/components/FAIcon.vue';

interface Props {
  icon: string;
  text: string;
}

const { icon, text } = defineProps<Props>();
</script>

<template>
  <div class="cursor-pointer flex justify-between">
    <FAIcon :icon="icon" size="lg" />
    <span class="text-xs">{{ text }}</span>
  </div>
</template>
