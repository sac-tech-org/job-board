<script setup lang="ts">
interface Props {
  provider: {
    color: string
    name: string
  }
}

const props = defineProps<Props>()
const emit = defineEmits(['selected'])
</script>

<template>
  <div class="providerContainer">
    <span>
      <button class="providerButton providerGithub" @click="emit('selected')">
        <div class="providerButtonLeft">
          <div class="providerButtonLogo">
            <div class="providerButtonLogoCenter">
              <svg xmlns="http://www.w3.org/2000/svg" width="18" height="17.556" viewBox="0 0 18 17.556">
                <path fill="#fff" fill-rule="evenodd"
                  d="M145.319 107.44a9 9 0 0 0-2.844 17.54c.45.082.614-.2.614-.434 0-.214-.008-.78-.012-1.531-2.5.544-3.032-1.206-3.032-1.206a2.384 2.384 0 0 0-1-1.317c-.817-.559.062-.547.062-.547a1.89 1.89 0 0 1 1.378.927 1.916 1.916 0 0 0 2.619.748 1.924 1.924 0 0 1 .571-1.2c-2-.227-4.1-1-4.1-4.448a3.479 3.479 0 0 1 .927-2.415 3.233 3.233 0 0 1 .088-2.382s.755-.242 2.475.923a8.535 8.535 0 0 1 4.506 0c1.718-1.165 2.472-.923 2.472-.923a3.234 3.234 0 0 1 .09 2.382 3.473 3.473 0 0 1 .925 2.415c0 3.458-2.1 4.218-4.11 4.441a2.149 2.149 0 0 1 .611 1.667c0 1.2-.011 2.174-.011 2.469 0 .24.162.52.619.433a9 9 0 0 0-2.851-17.539z"
                  transform="translate(-136.32 -107.44)"></path>
              </svg>
            </div>
          </div>
        </div>
        <div class="providerButtonText">Continue with {{ provider.name }}</div>
      </button>
    </span>
  </div>
</template>

<style scoped>
.providerButtonLeft {
  width: 40px;
}

.providerButtonLogo {
  height: 30px;
  display: flex;
  border-right: 1px solid rgba(255, 255, 255, 0.6);
}

.providerButtonLogoCenter {
  margin: auto;
}

.providerButtonText {
  margin: auto;
  text-align: center;
  justify-content: center;
  -webkit-box-pack: center;
  font-weight: inherit;
}

.providerContainer {
  padding-top: 9px;
  padding-bottom: 9px;
}

.providerButton {
  width: 100%;
  min-height: 34px;
  display: flex;
  flex-direction: row;
  padding: 2px 0px;
  transition: all 0.4s ease 0s;
  cursor: pointer;
  height: auto !important;
  border-radius: 6px;
  border-width: 1px;
  font-weight: 700;
  color: white;
}

.providerGithub {
  border-color: black;
  background-color: black;
}

.providerGithub:hover {
  filter: brightness(1.1);
}

.providerGoogle {
  border-color: rgb(234, 55, 33);
  background-color: rgb(234, 55, 33);
}

.providerGoogle:hover {
  filter: brightness(0.95);
}

.providerApple {
  border-color: rgb(7, 9, 60);
  background-color: rgb(1, 0, 48);
}

.providerApple:hover {
  filter: brightness(1.1);
}
</style>
