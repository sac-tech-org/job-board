import { defineStore } from 'pinia';
import { type User } from '../types';
import { ref } from 'vue';

export const useUserStore = defineStore('user', () => {
  const user = ref<User | null>(null);
  const loggedIn = ref(false);

  function login() {
    user.value = {
      id: 1,
      name: 'Eduardo',
      email: 'rusher2004@gmail.com',
    };

    loggedIn.value = true;
  }

  function logout() {
    user.value = null;
    loggedIn.value = false;
  }

  return { loggedIn, login, logout, user };
});
