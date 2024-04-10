<script setup lang="ts">
import { toRefs } from 'vue'
interface Props {
  emailError?: string;
  passwordError?: string;
  isSignIn: boolean;
}

const props = defineProps<Props>();
const { emailError, passwordError, isSignIn } = toRefs(props)
const emit = defineEmits(['submitForm']);

const email = defineModel<string>('email', { required: true });
const firstName = defineModel<string>('firstName', { required: true });
const lastName = defineModel<string>('lastName', { required: true });
const password = defineModel<string>('password', { required: true });
const username = defineModel<string>('username', { required: true });
</script>

<template>
  <form @submit.prevent="emit('submitForm')" autocomplete="on" novalidate>
    <div class="input-section-container" v-bind:class="emailError ? 'error' : ''">
      <div class="input-label">Email</div>
      <div class="input-container">
        <div class="input-wrapper" v-bind:class="emailError ? 'error' : ''">
          <input autocomplete="email" class="input" type="email" name="email" placeholder="Email address"
            v-model="email" required />
        </div>
      </div>
      <div v-if="emailError" class="input-error">{{ `${emailError}` }}</div>
    </div>

    <div v-if="!isSignIn">
      <div class="input-section-container">
        <div class="input-label">First Name</div>
        <div class="input-container">
          <div class="input-wrapper">
            <input class="input" type="text" name="firstName" placeholder="First Name" v-model="firstName" required />
          </div>
        </div>
      </div>

      <div class="input-section-container">
        <div class="input-label">Last Name</div>
        <div class="input-container">
          <div class="input-wrapper">
            <input class="input" type="text" name="lastName" placeholder="Last Name" v-model="lastName" required />
          </div>
        </div>
      </div>

      <div class="input-section-container">
        <div class="input-label">Username</div>
        <div class="input-container">
          <div class="input-wrapper">
            <input class="input" type="text" name="username" placeholder="Username" v-model="username" required />
          </div>
        </div>
      </div>
    </div>

    <div class="input-section-container" v-bind:class="passwordError ? 'error' : ''">
      <div class="input-label">Password</div>
      <div class="input-container">
        <div class="input-wrapper" v-bind:class="passwordError ? 'error' : ''">
          <input autocomplete="current-password" class="input" type="password" name="password" placeholder="Password"
            v-model="password" required />
        </div>
      </div>
      <div v-if="passwordError" class="input-error">{{ `${passwordError}` }}</div>
    </div>

    <div class="input-section-container">
      <button type="submit" class="button">SIGN {{ isSignIn ? 'IN' : 'UP' }}</button>
    </div>
  </form>
</template>

<style scoped>
form {
  display: block;
  margin-top: 0em;
}

.input-section-container {
  display: flex;
  flex-direction: column;
  padding-top: 0px;
  padding-bottom: 24px;
}

.input-section-container.error {
  padding-bottom: 0px;
}

.input-label {
  text-align: left;
  font-weight: 600;
  font-size: 14px;
  line-height: 24px;
  color: rgb(34, 34, 34);
}

.input-container {
  margin-top: 6px;
}

.input-wrapper {
  box-sizing: border-box;
  width: 100%;
  display: flex;
  -webkit-box-align: center;
  align-items: center;
  background-color: rgb(250, 250, 250);
  height: 34px;
  border-radius: 6px;
  border: 1px solid rgb(224, 224, 224);
}

.input-wrapper.error {
  border-color: rgb(255, 23, 23);
}

.input {
  box-sizing: border-box;
  padding-left: 15px;
  filter: none;
  color: rgb(34, 34, 34);
  background-color: transparent;
  border-radius: 6px;
  font-size: 14px;
  border: none;
  padding-right: 25px;
  letter-spacing: 1.2px;
  flex: 9 1 75%;
  width: 75%;
  height: 32px;
}

.button {
  width: 100%;
  height: 34px;
  background-color: rgb(255, 155, 51);
  color: white;
  font-weight: 700;
  border-width: 1px;
  border-style: solid;
  border-color: rgb(238, 141, 35);
  border-radius: 6px;
  background-position: center center;
  transition: all 0.4s ease 0s;
  cursor: pointer;
}

.button:hover {
  filter: brightness(0.95);
}

.input-error {
  padding-top: 5px;
  padding-bottom: 5px;
  color: rgb(255, 23, 23);
  line-height: 24px;
  font-weight: 400;
  font-size: 14px;
  text-align: left;
  animation: slideTop 0.5s cubic-bezier(0.25, 0.46, 0.45, 0.94) 0s 1 normal both;
  max-width: 330px;
}
</style>
