<script setup lang="ts">
import { onMounted } from "vue";
import { useRouter } from "vue-router";
import ThirdPartyEmailPassword from "supertokens-web-js/recipe/thirdpartyemailpassword";

const router = useRouter();

onMounted(async () => {
  try {
    const response = await ThirdPartyEmailPassword.thirdPartySignInAndUp({});
    if (response.status !== "OK") {
      return router.push('/auth?error=signing-in');
    }

    router.push('/');
  } catch (_) {
    router.push('/auth?error=signing-in');
  }
});
</script>

<template>
  <div class="fill">
    <div class="spinner">
      <svg version="1.1" viewBox="25 25 50 50">
        <circle cx="50" cy="50" r="20" fill="none" strokeWidth="20" stroke="rgb(255, 155, 51)" strokeLinecap="round"
          strokeDashoffset="0" strokeDasharray="200, 200">
          <animateTransform attributeName="transform" attributeType="XML" type="rotate" from="0 50 50" to="360 50 50"
            dur="4s" repeatCount="indefinite" />
          <animate attributeName="stroke-dashoffset" values="0;-30;-124" dur="2s" repeatCount="indefinite" />
          <animate attributeName="stroke-dasharray" values="0,200;110,200;110,200" dur="2s" repeatCount="indefinite" />
        </circle>
      </svg>
    </div>
  </div>
</template>

<style scoped>
.fill {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 100%;
}

.spinner {
  width: 80px;
  height: auto;
  padding-top: 20px;
  padding-bottom: 40px;
  margin: 0 auto;
}
</style>
