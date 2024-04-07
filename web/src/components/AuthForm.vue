<script setup lang="ts">
interface Props {
  emailError: string;
  passwordError: string;
  isSignIn: boolean;
}

const props = defineProps<Props>();
const emit = defineEmits(['submitForm']);

const email = defineModel<string>('email', { required: true });
const password = defineModel<string>('password', { required: true });
</script>

<template>
  <form @submit.prevent="emit('submitForm')" autocomplete="on" novalidate>
    <div class="input-section-container" v-bind:class="emailError ? 'error' : ''">
      <div class="input-label">Email</div>
      <div class="input-container">
        <div class="input-wrapper" v-bind:class="emailError ? 'error' : ''">
          <input autocomplete="email" class="input" type="email" name="email" placeholder="Email address"
            v-model="email" />
        </div>
      </div>
      <div v-if="emailError" class="input-error">{{ `${emailError}` }}</div>
    </div>

    <div class="input-section-container" v-bind:class="passwordError ? 'error' : ''">
      <div class="input-label">Password</div>
      <div class="input-container">
        <div class="input-wrapper" v-bind:class="passwordError ? 'error' : ''">
          <input autocomplete="current-password" class="input" type="password" name="password" placeholder="Password"
            v-model="password" />
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
  padding-bottom: 34px;
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
</style>
