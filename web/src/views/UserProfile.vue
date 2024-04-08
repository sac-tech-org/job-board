<script setup lang="ts">
import { useUserStore } from '@/stores/user'
import UserProfileForm from '@/components/UserProfileForm.vue'
import { ref } from 'vue';

const userStore = useUserStore()

userStore.$subscribe((_, state) => {
  email.value = state.user?.email || ''
  firstName.value = state.user?.firstName || ''
  lastName.value = state.user?.lastName || ''
})

const email = ref('')
const firstName = ref('')
const lastName = ref('')

</script>


<template>
  <section class="w-full p-1 md:p-6">
    <header class="flex justify-start items-center gap-2 border-b">
      <div class="rounded bg-cyan-600 md:size-12"></div>
      <hgroup>
        <h1 class="text-2xl font-bold">{{ $route.params.username }}</h1>
        <p class="pl-1">tagline/title</p>
      </hgroup>
    </header>

    <div class="flex flex-col">
      <span>email: <span>{{ email }}</span></span>
      <span>first name: <span>{{ firstName }}</span></span>
      <span>last name: <span>{{ lastName }}</span></span>
    </div>
    <UserProfileForm v-model:email="email" v-model:firstName="firstName" v-model:lastName="lastName" />
  </section>
</template>
