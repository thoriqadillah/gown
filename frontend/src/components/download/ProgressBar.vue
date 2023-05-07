<script setup lang="ts">
import { download } from '../../../wailsjs/go/models';

const props = defineProps<{
  totalpart: number
  chunks: download.Chunk[]
}>()

</script>

<template>
  <div class="progressWrapper tw-flex tw-justify-between tw-mt-1">
    <div v-for="part in props.totalpart" :class="`tw-w-full ` + `basis-1/${props.totalpart}`" >
      <div class="tw-h-0.5 tw-bg-green-500 -tw-mt-1" :style="{ width: `${props.chunks[part-1].progressbar}%` }"></div>
    </div> 
  </div>
</template>