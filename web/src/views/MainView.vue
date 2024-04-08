<script setup lang="ts">
import { onMounted, toRefs } from 'vue';
import { useAPI, type HealthCheckResponse } from '@/utils/api'

import HeroSection from '@/components/HeroSection.vue';
import ListView from '@/views/ListView.vue';

const api = useAPI<HealthCheckResponse>();
const { data, error, loading } = toRefs(api);

onMounted(async () => {
  await api.execute();
})

onMounted
</script>

<template>
  <main class="w-full min-h-0 px-7 pb-5 flex flex-col grow font-mono text-sm">
    <HeroSection />
    <ListView />
    <div>
      <div class="text-left">Loading... {{ loading }}</div>
      <div class="text-left">Error: {{ error }}</div>
      <div class="text-left">Data: {{ data }}</div>
    </div>
  </main>
</template>
