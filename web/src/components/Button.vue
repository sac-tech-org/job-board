<!-- eslint-disable vue/multi-word-component-names -->
<script setup lang="ts">
import { computed, toRefs } from 'vue';

interface Props {
  bgColor?: string
  textColor?: string
}

const props = defineProps<Props>()
const { bgColor, textColor } = toRefs(props)

// set the default colors if there's no value in props
const defaultClasses: { [key: string]: boolean } = {
  'bg-green-500': !bgColor.value,
  'text-zinc-50': !textColor.value,
  'hover:bg-green-500/90': !bgColor.value
}

const classes = computed(() => {
  const out = { ...defaultClasses }

  // then override the default colors if there's a value in props
  if (bgColor.value) {
    delete out['bg-green-500']
    delete out['hover:bg-green-500/90']

    out[`bg-${bgColor.value}`] = true
    out[`hover:bg-${bgColor.value}/90`] = true
  }

  if (textColor.value) {
    delete out['text-zinc-50']
    out[`text-${textColor.value}`] = true
  }

  return out
})
</script>

<template>
  <button
    class="inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 h-10 px-4 py-2"
    :class="classes">
    <slot></slot>
  </button>
</template>
