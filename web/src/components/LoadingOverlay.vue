<template>
  <div class="absolute inset-0 bg-black bg-opacity-30 z-50 flex justify-center items-center">
    <div class="gradient flex justify-center items-center p-4 rounded-xl h-1/2 w-3/4 max-h-40 max-w-60 shadow-xl">
      <div class="text-center text-2xl">Loading...</div>
    </div>
  </div>
</template>

<style scoped>
.gradient {
  background: linear-gradient(-45deg, #22c55e, #14b8a6, #0ea5e9);
  background-size: 400% 400%;
  animation: gradient 4s ease-in-out infinite;
}

@keyframes gradient {
  0% {
    background-position: 0% 50%;
  }

  50% {
    background-position: 100% 50%;
  }

  100% {
    background-position: 0% 50%;
  }
}
</style>
