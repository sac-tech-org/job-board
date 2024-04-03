<script setup lang="ts">
import { ref } from 'vue';

import JobView from './JobView.vue';
import ListItem from './ListItem.vue';
import ModalView from './ModalView.vue';

import { type Job } from '../types'

const jobs = [
  { id: 1, title: 'Job 1', description: 'This is a job description.' },
  { id: 2, title: 'Job 2', description: 'This is a job description.' },
  { id: 3, title: 'Job 3', description: 'This is a job description.' },
  { id: 4, title: 'Job 4', description: 'This is a job description.' },
  { id: 5, title: 'Job 5', description: 'This is a job description.' },
  { id: 6, title: 'Job 6', description: 'This is a job description.' },
  { id: 7, title: 'Job 7', description: 'This is a job description.' },
  { id: 8, title: 'Job 8', description: 'This is a job description.' },
  { id: 9, title: 'Job 9', description: 'This is a job description.' },
  { id: 10, title: 'Job 10', description: "'Under The Skin' asks: what's beneath the surface - of people we admire, of the ideas that define our time, of the history we are told. Speaking with guests from the world of academia, popular culture and the arts, they'll help us to see the ulterior truth behind our constructed reality. And have a laugh." },
  { id: 11, title: 'Job 11', description: 'This is a job description.' },
  { id: 12, title: 'Job 12', description: 'This is a job description.' },
  { id: 13, title: 'Job 13', description: 'This is a job description.' },
  { id: 14, title: 'Job 14', description: 'This is a job description.' },
  { id: 15, title: 'Job 15', description: 'This is a job description.' },
  { id: 16, title: 'Job 16', description: 'This is a job description.' },
  { id: 17, title: 'Job 17', description: 'This is a job description.' },
  { id: 18, title: 'Job 18', description: 'This is a job description.' },
  { id: 19, title: 'Job 19', description: 'This is a job description.' },
  { id: 20, title: 'Job 20', description: 'This is a job description.' },
  // { id: 21, title: 'Job 21', description: 'This is a job description.' },
  // { id: 22, title: 'Job 22', description: 'This is a job description.' },
  // { id: 23, title: 'Job 23', description: 'This is a job description.' },
  // { id: 24, title: 'Job 24', description: 'This is a job description.' },
  // { id: 25, title: 'Job 25', description: 'This is a job description.' },
  // { id: 26, title: 'Job 26', description: 'This is a job description.' },
  // { id: 27, title: 'Job 27', description: 'This is a job description.' },
  // { id: 28, title: 'Job 28', description: 'This is a job description.' },
  // { id: 29, title: 'Job 29', description: 'This is a job description.' },
  // { id: 30, title: 'Job 30', description: 'This is a job description.' },
  // { id: 31, title: 'Job 31', description: 'This is a job description.' },
  // { id: 32, title: 'Job 32', description: 'This is a job description.' },
  // { id: 33, title: 'Job 33', description: 'This is a job description.' },
  // { id: 34, title: 'Job 34', description: 'This is a job description.' },
  // { id: 35, title: 'Job 35', description: 'This is a job description.' },
  // { id: 36, title: 'Job 36', description: 'This is a job description.' },
  // { id: 37, title: 'Job 37', description: 'This is a job description.' },
  // { id: 38, title: 'Job 38', description: 'This is a job description.' },
  // { id: 39, title: 'Job 39', description: 'This is a job description.' },
  // { id: 40, title: 'Job 40', description: 'This is a job description.' },
  // { id: 41, title: 'Job 41', description: 'This is a job description.' },
  // { id: 42, title: 'Job 42', description: 'This is a job description.' },
  // { id: 43, title: 'Job 43', description: 'This is a job description.' },
  // { id: 44, title: 'Job 44', description: 'This is a job description.' },
  // { id: 45, title: 'Job 45', description: 'This is a job description.' },
  // { id: 46, title: 'Job 46', description: 'This is a job description.' },
  // { id: 47, title: 'Job 47', description: 'This is a job description.' },
  // { id: 48, title: 'Job 48', description: 'This is a job description.' },
  // { id: 49, title: 'Job 49', description: 'This is a job description.' },
  // { id: 50, title: 'Job 50', description: 'This is a job description.' },
];

const modalOpen = ref(false)
const selectedJob = ref<Job | null>(null)

function onSelected(id: number) {
  console.log('selected', id)
  const job = jobs.find(j => j.id === id);

  selectedJob.value = job || null;
  if (selectedJob.value) {
    modalOpen.value = true;
  }
}

function modalClosed() {
  selectedJob.value = null;
  modalOpen.value = false;
}

</script>

<template>
  <ModalView :open="modalOpen" @modalClosed="modalClosed" :title="'Job Details'" class="h-5/6 w-1/2">
    <JobView v-if="selectedJob" :job="selectedJob" />
  </ModalView>

  <ul role="list" class="divide-y divide-gray-100 max-w-fit">
    <ListItem v-for="j in jobs" :job="j" :key="j.id" :disableClick="selectedJob !== null && selectedJob.id !== j.id"
      @selected="onSelected(j.id)" />
  </ul>
</template>
