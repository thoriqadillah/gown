<script setup lang="ts">
import { defineProps, ref } from 'vue';
import { useDownloads } from '../store/downloads';
import { EventsOn } from '../../wailsjs/runtime/runtime';
import { watch } from 'vue';
import { download } from '../../wailsjs/go/models';
import { useDateFormat } from '@vueuse/shared';

const downloads = useDownloads()
const props = defineProps<{
  list: download.Download[]
}>()

const transfered = ref(0)
const progress = ref(0)
const progressWrapper = ref<HTMLElement[]>([])
const progressBar = ref()
const totalparts = ref(0)

watch(downloads.list, (newval, oldval) => {
  totalparts.value = downloads.toDownload.totalpart
})

EventsOn("transfered", (...data) => {
  transfered.value += data[1] / (1024*1024)
  progress.value = ((transfered.value / (downloads.toDownload.size / (1024*1024))) * 100)
  
  let prog = progress.value.toFixed(0)
  const progressBar = document.getElementById(`progressBar-${data[0]}`) as HTMLElement
  progressBar.style.display = 'block'
  progressBar.style.width = prog + '%'
  
  if (prog == '100') {
    progressBar.style.display = 'none'
    // refresh the status icon and time elapsed
  }
})

</script>

<template>
  <div class="tw-px-5 xl:tw-px-0 xl:tw-pl-5">
    <div class="tw-flex tw-gap-2 md:tw-gap-4 tw-mb-2">
      <v-btn variant="outlined" @click="downloads.setDefault()" prepend-icon="mdi-select-all">ALL</v-btn>
      <v-btn variant="outlined" @click="downloads.filterByImage()" prepend-icon="mdi-image" color="red-accent-2">IMAGE</v-btn>
      <v-btn variant="outlined" @click="downloads.filterByVideo()" prepend-icon="mdi-video" color="deep-orange-accent-2">VIDEO</v-btn>
      <v-btn variant="outlined" @click="downloads.filterByDocument()" prepend-icon="mdi-file-document" color="blue-accent-2">DOCUMENT</v-btn>
      <v-btn variant="outlined" @click="downloads.filterByCompressed()" prepend-icon="mdi-zip-box" color="yellow-accent-4">COMPRESSED</v-btn>
      <v-btn variant="outlined" @click="downloads.filterByMusic()" prepend-icon="mdi-music-box" color="purple-accent-2">MUSIC</v-btn>
    </div>
    
    <v-table density="compact">
      <thead>
        <tr>
          <th class="text-left tw-cursor-pointer" @click="downloads.sortByName()">
            <div class="tw-flex tw-justify-between tw-items-center">
              <span class="tw-text-sm">Name</span>
              <v-icon icon="mdi-arrow-up-down" class="tw-text-sm"></v-icon>
            </div>
          </th>
          <th class="text-left tw-cursor-pointer" @click="downloads.sortByTimeElapsed()">
            <div class="tw-flex tw-justify-between tw-items-center tw-w-max md:tw-w-full">
              <span class="tw-text-sm tw-mr-3 tw-w-32">Time Elapsed</span>
              <v-icon icon="mdi-arrow-up-down" class="tw-text-sm"></v-icon>
            </div>
          </th>
          <th class="text-left tw-cursor-pointer" @click="downloads.sortBySize">
            <div class="tw-flex tw-justify-between tw-items-center">
              <span class="tw-text-sm tw-mr-3 tw-w-20">Size</span>
              <v-icon icon="mdi-arrow-up-down" class="tw-text-sm"></v-icon>
            </div>
          </th>
          <th class="text-left tw-cursor-pointer" @click="downloads.sortByDate()">
            <div class="tw-flex tw-justify-between tw-items-center">
              <span class="tw-text-sm tw-w-32">Date</span>
              <v-icon icon="mdi-arrow-up-down" class="tw-text-sm"></v-icon>
            </div>
          </th>
          <th class="tw-text tw-cursor-pointer-left">
            <span class="tw-text-sm">Status</span>
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(item,i) in props.list" :key="item.name">
          <td color="primary" class="tw-rounded-sm" id="nameCol">
            <div class="tw-flex tw-justify-between tw-mt-1">
              <div class="tw-overflow-x-hidden tw-w-max">
                <v-icon :icon="item.type.icon" :color="item.type.color" class="tw-opacity-70 tw-mr-2"></v-icon>
                <span class="tw-text-sm">{{ item.name }}</span>
              </div>
            </div>
            <div ref="progressWrapper" class="progressWrapper tw-flex tw-justify-between">
              <!-- <div v-for="part in totalparts" :class="`tw-h-0.5 tw-bg-green-500 tw-opacity-50 tw-my-1 tw-w-1 tw-rounded-lg ` + `tw-basis-1/${totalparts}`" :id="`progressBar-${i}-${part-1}`" ref="progressBar"></div>  -->
              <div class="tw-h-0.5 tw-bg-green-500 tw-opacity-50 tw-my-1 tw-w-1 tw-hidden tw-rounded-lg" :id="`progressBar-${item.id}`" ref="progressBar"></div> 
            </div>
          </td>
          <td class="tw-text-sm tw-rounded-sm text-left tw-w-32">{{ item.timeElapsed == 0 ? '' : item.timeElapsed }}</td>
          <td class="tw-text-sm tw-rounded-sm text-left tw-w-20">{{ downloads.parseSize(item.size) }}</td>
          <td class="tw-text-sm tw-rounded-sm text-left tw-w-32">{{ useDateFormat(item.date, 'MMMM DD, YYYY HH:mm').value }}</td>
          <td class="tw-text-sm tw-text-left"><v-icon :icon="item.status.icon" :color="item.status.color" class="tw-opacity-90 tw-ml-2"></v-icon></td>
        </tr>
      </tbody>
    </v-table>
  </div>
</template>

<style>
.v-table .v-table__wrapper > table > tbody > tr > td {
  border-bottom: none !important;
}
.v-table .v-table__wrapper > table > tbody > tr:nth-child(even) {
  background: rgb(87, 83, 78, 0.05);
}
</style>