<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { useUserStore } from '@/stores/user'
import { type UserPutRequest } from '@/utils/api'
import UserProfileForm from '@/components/UserProfileForm.vue'

const userStore = useUserStore()

const email = ref('')
const firstName = ref('')
const lastName = ref('')
const username = ref('')

userStore.$subscribe((_, state) => {
  email.value = state.user?.email.address || ''
  firstName.value = state.user?.firstName || ''
  lastName.value = state.user?.lastName || ''
  username.value = state.user?.username || ''
})

async function updateUser() {
  const req: UserPutRequest = {
    body: {
      firstName: firstName.value,
      lastName: lastName.value,
      username: username.value
    },
    params: {
      path: {
        username: username.value
      }
    }
  }
  console.log('req:', req)
  await userStore.updateUser(req)
}

onMounted(() => {
  userStore.getUser()
})

</script>

<template>
  <section class="w-full p-1 md:p-6">
    <header class="flex justify-start items-center gap-2 border-b">
      <div class="rounded bg-cyan-600 md:size-12"></div>
      <hgroup>
        <h1 class="text-2xl font-bold">{{ username || '~~~' }}</h1>
        <p class="pl-1">tagline/title</p>
      </hgroup>
    </header>

    <div class="flex flex-col">
      <div class="font-bold">Form Values:</div>
      <span>email: <span>{{ email }}</span></span>
      <span>first name: <span>{{ firstName }}</span></span>
      <span>last name: <span>{{ lastName }}</span></span>
    </div>
    <UserProfileForm @submitForm="updateUser" v-model:firstName="firstName" v-model:lastName="lastName"
      v-model:username="username" />
  </section>
</template>
