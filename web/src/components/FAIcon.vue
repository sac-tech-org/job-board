<script setup lang="ts">
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

type IconProps = InstanceType<typeof FontAwesomeIcon>['$props'];

interface props extends /* @vue-ignore */ IconProps {
  hoverGrow?: boolean;
}

const { hoverGrow } = defineProps<props>();
</script>

<template>
  <FontAwesomeIcon v-bind="$attrs as any" :class="{ 'hover:scale-125': hoverGrow }" />
</template>
