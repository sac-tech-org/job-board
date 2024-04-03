<script setup lang="ts">
import { type Job } from '@/types';

interface Props {
  job: Job;
}

const { job } = defineProps<Props>()

</script>

<template>
  <div class="flex flex-col gap-y-4">
    <h1 class="text-lg font-semibold">{{ job.title }}</h1>
    <p class="text-sm leading-6 text-gray-500">Company Name</p>
    <p class="text-sm leading-6 text-gray-500">Location</p>
    <p class="text-sm leading-6 text-gray-500">Salary</p>
    <p class="text-sm leading-6 text-gray-500">Job Description</p>
  </div>
</template>
