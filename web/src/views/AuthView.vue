<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import Session from "supertokens-web-js/recipe/session";
import ThirdPartyEmailPassword from 'supertokens-web-js/recipe/thirdpartyemailpassword';

const router = useRouter();

const isSignIn = ref(true);
const email = ref('');
const password = ref('');

const error = ref(false);
const errorMessage = ref('Something went wrong');
const emailError = ref('');
const passwordError = ref('');

async function checkForSession() {
  if (await Session.doesSessionExist()) {
    router.push('/');
  }
}

function handleSubmit(e: Event) {
  e.preventDefault();

  if (isSignIn.value) {
    signIn();
  } else {
    signUp();
  }
}

async function handleAppleSelect() {
  const authURL = await ThirdPartyEmailPassword.getAuthorisationURLWithQueryParamsAndSetState({
    thirdPartyId: 'apple',
    frontendRedirectURI: 'http://localhost:5173/auth/callback/apple',
    redirectURIOnProviderDashboard: 'http://localhost:5173/auth/callback/apple',
  })

  window.location.assign(authURL);
}

async function handleGithubSelect() {
  const authURL = await ThirdPartyEmailPassword.getAuthorisationURLWithQueryParamsAndSetState({
    thirdPartyId: 'github',
    frontendRedirectURI: 'http://localhost:5173/auth/callback/github',
  })

  window.location.assign(authURL);
}

async function handleGoogleSelect() {
  const authURL = await ThirdPartyEmailPassword.getAuthorisationURLWithQueryParamsAndSetState({
    thirdPartyId: 'google',
    frontendRedirectURI: 'http://localhost:5173/auth/callback/google',
  })

  window.location.assign(authURL);
}

async function handleLinkedInSelect() {
  const authURL = await ThirdPartyEmailPassword.getAuthorisationURLWithQueryParamsAndSetState({
    thirdPartyId: 'linkedin',
    frontendRedirectURI: 'http://localhost:5173/auth/callback/linkedin',
  })

  window.location.assign(authURL);
}

function goToSignUp() {
  isSignIn.value = false;
}

function goToSignIn() {
  isSignIn.value = true;
}

async function signIn() {
  const response = await ThirdPartyEmailPassword.emailPasswordSignIn({
    formFields: [
      {
        id: 'email',
        value: email.value,
      },
      {
        id: 'password',
        value: password.value,
      },
    ],
  });

  if (response.status === 'WRONG_CREDENTIALS_ERROR') {
    errorMessage.value = 'Invalid credentials';
    error.value = true;
    return;
  }

  if (response.status === 'FIELD_ERROR') {
    response.formFields.forEach(i => {
      if (i.id === 'email') {
        emailError.value = i.error;
      } else if (i.id === 'password') {
        passwordError.value = i.error;
      }
    });
    return;
  }

  // router.push('/');
}

async function signUp() {
  const response = await ThirdPartyEmailPassword.emailPasswordSignUp({
    formFields: [
      {
        id: "email",
        value: email.value,
      },
      {
        id: "password",
        value: password.value,
      },
    ],
  });

  if (response.status === "FIELD_ERROR") {
    response.formFields.forEach((item) => {
      if (item.id === "email") {
        // this means that something was wrong with the entered email.
        // probably that it's not a valid email (from a syntax point of view)
        emailError.value = item.error;
      } else if (item.id === "password") {
        // this means that something was wrong with the entered password.
        // probably it doesn't meet the password validation criteria on the backend.
        passwordError.value = item.error;
      }
    });
    return;
  }
}

function validateEmail(email: string): boolean {
  return email.includes('@');
}
</script>

<template>
  <div class="auth-container">
    <div class="auth-form-container">
      <div v-if="error" class="error-container">
        <div class="error-message">{{ errorMessage }}</div>
      </div>
      <div class="auth-form-content-container">
        <div class="form-title" v-if="isSignIn">Sign In</div>
        <div class="form-title" v-else>Sign Up</div>
        <div class="sign-in-up-text-container">
          <span v-if="isSignIn">Not yet registered?
            <span class="clickable-text" v-on:click="goToSignUp">Sign Up</span></span>
          <span v-else>Already have an account?
            <span class="clickable-text" v-on:click="goToSignIn">Sign In</span></span>
        </div>
        <div class="divider-container">
          <div class="divider" />
        </div>
        <div class="providerContainer">
          <span>
            <button class="providerButton providerGithub" @click="handleGithubSelect">
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
              <div class="providerButtonText">Continue with Github</div>
            </button>
          </span>
        </div>

        <div class="providerContainer">
          <span>
            <button class="providerButton providerGoogle" @click="handleGoogleSelect">
              <div class="providerButtonLeft">
                <div class="providerButtonLogo">
                  <div class="providerButtonLogoCenter">
                    <svg xmlns="http://www.w3.org/2000/svg" width="18.001" height="18" viewBox="0 0 18.001 18">
                      <g id="Group_9292" transform="translate(-534 -389)">
                        <path id="Path_85803"
                          d="M3.989 144.285l-.627 2.339-2.29.048a9.016 9.016 0 0 1-.066-8.4l2.039.374.893 2.027a5.371 5.371 0 0 0 .05 3.616z"
                          transform="translate(534 255.593)" style="fill: rgb(255, 255, 255)"></path>
                        <path id="Path_85804"
                          d="M270.273 208.176a9 9 0 0 1-3.208 8.7l-2.568-.131-.363-2.269a5.364 5.364 0 0 0 2.308-2.739h-4.813v-3.56h8.645z"
                          transform="translate(281.57 188.143)" style="fill: rgb(255, 255, 255)"></path>
                        <path id="Path_85805"
                          d="M44.07 314.549a9 9 0 0 1-13.561-2.749l2.917-2.387a5.353 5.353 0 0 0 7.713 2.741z"
                          transform="translate(504.564 90.469)" style="fill: rgb(255, 255, 255)"></path>
                        <path id="Path_85806"
                          d="M42.362 2.072l-2.915 2.387a5.352 5.352 0 0 0-7.89 2.8l-2.932-2.4a9 9 0 0 1 13.737-2.787z"
                          transform="translate(506.383 389)" style="fill: rgb(255, 255, 255)"></path>
                      </g>
                    </svg>
                  </div>
                </div>
              </div>
              <div class="providerButtonText">Continue with Google</div>
            </button>
          </span>
        </div>

        <div class="providerContainer">
          <span>
            <button class="providerButton providerApple" @click="handleAppleSelect">
              <div class="providerButtonLeft">
                <div class="providerButtonLogo">
                  <div class="providerButtonLogoCenter">
                    <svg xmlns="http://www.w3.org/2000/svg" width="15.614" height="18.737" viewBox="0 0 15.614 18.737">
                      <g id="iconfinder_logo_brand_brands_logos_apple_ios_2993701" transform="translate(-2)">
                        <path id="Path_91415"
                          d="M14.494 11.075a4.29 4.29 0 0 1 2.372-3.836A4.888 4.888 0 0 0 13.713 6a4.783 4.783 0 0 0-1.83.406 4.357 4.357 0 0 1-1.684.375 4.357 4.357 0 0 1-1.684-.375A4.783 4.783 0 0 0 6.685 6C6.206 6 2 6.153 2 11.465c0 4.344 2.964 8.587 4.685 8.587a5.18 5.18 0 0 0 2.194-.464 3.2 3.2 0 0 1 1.32-.316 3.2 3.2 0 0 1 1.32.317 5.187 5.187 0 0 0 2.194.463c1.206 0 2.922-2.085 3.9-4.826a4.209 4.209 0 0 1-3.119-4.151z"
                          transform="translate(0 -1.316)" style="fill: rgb(255, 255, 255)"></path>
                        <path id="XMLID_1339_" d="M12 4.684A4.734 4.734 0 0 0 15.906 0 4.734 4.734 0 0 0 12 4.684z"
                          transform="translate(-2.193)" style="fill: rgb(255, 255, 255)"></path>
                        <path id="Path_91416"
                          d="M6.685 6.2a4.783 4.783 0 0 1 1.83.406 4.357 4.357 0 0 0 1.684.375 4.357 4.357 0 0 0 1.684-.381 4.783 4.783 0 0 1 1.83-.406 4.953 4.953 0 0 1 3.014 1.126c.047-.026.091-.058.14-.082A4.888 4.888 0 0 0 13.713 6a4.783 4.783 0 0 0-1.83.406 4.357 4.357 0 0 1-1.684.375 4.357 4.357 0 0 1-1.684-.375A4.783 4.783 0 0 0 6.685 6C6.206 6 2 6.153 2 11.465v.107C2.053 6.352 6.208 6.2 6.685 6.2z"
                          transform="translate(0 -1.316)" style="fill: rgb(255, 255, 255); opacity: 0.1"></path>
                        <path id="Path_91417"
                          d="M13.713 21.368a5.187 5.187 0 0 1-2.194-.463 3.2 3.2 0 0 0-1.32-.317 3.2 3.2 0 0 0-1.32.316 5.18 5.18 0 0 1-2.194.464c-1.707 0-4.633-4.174-4.681-8.48v.088c0 4.344 2.964 8.587 4.685 8.587a5.18 5.18 0 0 0 2.19-.463 3.2 3.2 0 0 1 1.32-.316 3.2 3.2 0 0 1 1.32.317 5.187 5.187 0 0 0 2.194.463c1.206 0 2.922-2.085 3.9-4.826-.023-.006-.043-.017-.066-.023-.991 2.654-2.655 4.653-3.834 4.653z"
                          transform="translate(0 -2.826)" style="fill: rgb(255, 255, 255); opacity: 0.2"></path>
                        <path id="Path_91418"
                          d="M15.888.4A4.621 4.621 0 0 1 12 4.544v.2A4.745 4.745 0 0 0 15.9.261c0 .039 0 .098-.012.139z"
                          transform="translate(-2.193 -.057)" style="fill: rgb(255, 255, 255); opacity: 0.2"></path>
                        <path id="Path_91419"
                          d="M12.005 4.477c.009-.051.02-.192.032-.278q.012-.161.035-.317A4.491 4.491 0 0 1 15.9.2V0a4.738 4.738 0 0 0-3.895 4.477z"
                          transform="translate(-2.194)" style="fill: rgb(255, 255, 255); opacity: 0.1"></path>
                        <path id="Path_91420"
                          d="M14.494 9.759a4.29 4.29 0 0 1 2.372-3.836 4.888 4.888 0 0 0-3.153-1.239 4.783 4.783 0 0 0-1.83.406 4.357 4.357 0 0 1-1.684.375 4.357 4.357 0 0 1-1.684-.375 4.783 4.783 0 0 0-1.83-.406C6.206 4.684 2 4.838 2 10.15c0 4.344 2.964 8.587 4.685 8.587a5.18 5.18 0 0 0 2.194-.464 3.2 3.2 0 0 1 1.32-.316 3.2 3.2 0 0 1 1.32.317 5.187 5.187 0 0 0 2.194.463c1.206 0 2.922-2.085 3.9-4.826a4.209 4.209 0 0 1-3.119-4.152zM13.713 0a4.734 4.734 0 0 0-3.9 4.684A4.734 4.734 0 0 0 13.713 0z"
                          style="fill: rgb(255, 255, 255)"></path>
                      </g>
                    </svg>
                  </div>
                </div>
              </div>
              <div class="providerButtonText">Continue with Apple</div>
            </button>
          </span>
        </div>

        <div class="divider-container">
          <div class="divider" />
          <div class="divider-text">or</div>
          <div class="divider" />
        </div>

        <form v-on:submit="handleSubmit" autocomplete="on" novalidate>
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
                <input autocomplete="current-password" class="input" type="password" name="password"
                  placeholder="Password" v-model="password" />
              </div>
            </div>
            <div v-if="passwordError" class="input-error">{{ `${passwordError}` }}</div>
          </div>

          <div class="input-section-container">
            <div v-if="isSignIn">
              <button type="submit" class="button">SIGN IN</button>
            </div>
            <div v-else>
              <button type="submit" class="button">SIGN UP</button>
            </div>
          </div>
        </form>
      </div>
      <div v-if="isSignIn">
        <router-link :to="{ path: `/auth/reset-password` }"> Forgot Password? </router-link>
      </div>
      <div style="margin-bottom: 10px" />
    </div>
  </div>

</template>

<style scoped>
.auth-container {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  width: 100%;
  height: 100%;
}

.auth-form-container {
  margin: 26px auto;
  width: 420px;
  text-align: center;
  border-radius: 8px;
  box-shadow: rgb(0, 0, 0, 0.16) 1px 1px 10px;
  background-color: white;
}

.form-title {
  font-size: 24px;
  line-height: 40px;
  letter-spacing: 0.58px;
  font-weight: 800;
  margin-bottom: 2px;
  color: rgb(34, 34, 34);
}

.sign-in-up-text-container {
  font-size: 14px;
  font-weight: 300;
  letter-spacing: 0.4px;
  color: rgb(101, 101, 101);
}

.auth-form-content-container {
  margin: auto;
  width: 76%;
  padding-top: 40px;
}

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

.input-section-container {
  display: flex;
  flex-direction: column;
  padding-top: 0px;
  padding-bottom: 34px;
}

.input-section-container.error {
  padding-bottom: 0px;
}

form {
  display: block;
  margin-top: 0em;
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

.forgot-password-link {
  padding-left: 3px;
  padding-right: 3px;
  cursor: pointer;
  line-height: 26px;
  font-size: 14px;
  font-weight: 300;
  letter-spacing: 0.4px;
  color: rgb(101, 101, 101);
  margin-top: 10px;
}

.error-container {
  display: flex;
  position: absolute;
  width: calc(100% - 24px);
  background-color: #ffcdd2;
  justify-content: center;
  align-items: center;
  padding-top: 2px;
  padding-bottom: 2px;
  margin-left: 12px;
  margin-right: 12px;
  margin-top: 4px;
  border-radius: 6px;
  box-sizing: border-box;
  border-width: 1px;
  border: 1px solid #ff1744;
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

@keyframes slideTop {
  0% {
    transform: translateY(-5px);
  }

  100% {
    transform: translateY(0px);
  }
}
</style>