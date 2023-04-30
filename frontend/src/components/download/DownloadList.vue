<script setup lang="ts">
import { computed } from 'vue';
import { parseElapsedTime, parseSize } from '../../utils/parser';
import { useDateFormat } from '@vueuse/shared';
import { useDownloads } from '../../store/downloads';
import { EventsOn } from '../../../wailsjs/runtime/runtime';
import Actionable from './Actionable.vue';
import Deletable from './Deletable.vue';

const store = useDownloads()
const items = computed(() => store.filter(store.search))

EventsOn("transfered", async (...data) => {
  const [id, index, transfered, progress] = data

  const progressBar = document.getElementById(`progressBar-${id}-${index}`) as HTMLElement
  progressBar.style.width = transfered + '%'

  const target = store.list[store.indexOf(id)]
  target.timeElapsed = parseElapsedTime(target.date)
  target.progress += progress // FIXME: a float can't be saved into backend. Need int64
})

EventsOn("downloaded", async (...data) => {
  const [id, combined] = data
  
  const target = store.list[store.indexOf(id)]
  target.progress = 100
  target.status = {
    name: !combined ? 'Combining' : 'Success',
    icon: !combined ? 'mdi-file-arrow-left-right-outline' : 'mdi-check-circle-outline',
    color: !combined ? 'info' : 'success'
  }
  
  await store.updateData(store.list)
})
</script>

<template>
  <div class="tw-px-5 xl:tw-px-0 xl:tw-pl-5">
    <div class="tw-flex tw-gap-2 md:tw-gap-4 tw-mb-2">
      <v-btn variant="outlined" @click="store.setDefault()" prepend-icon="mdi-select-all">ALL</v-btn>
      <v-btn variant="outlined" @click="store.filterByImage()" prepend-icon="mdi-image" color="red-accent-2">IMAGE</v-btn>
      <v-btn variant="outlined" @click="store.filterByVideo()" prepend-icon="mdi-video" color="deep-orange-accent-2">VIDEO</v-btn>
      <v-btn variant="outlined" @click="store.filterByDocument()" prepend-icon="mdi-file-document" color="blue-accent-2">DOCUMENT</v-btn>
      <v-btn variant="outlined" @click="store.filterByCompressed()" prepend-icon="mdi-zip-box" color="yellow-accent-4">COMPRESSED</v-btn>
      <v-btn variant="outlined" @click="store.filterByMusic()" prepend-icon="mdi-music-box" color="purple-accent-2">MUSIC</v-btn>
    </div>
    
    <v-table density="compact">
      <thead>
        <tr>
          <th class="text-left tw-cursor-pointer" @click="store.sortByName()">
            <div class="tw-flex tw-justify-between tw-items-center">
              <span class="tw-text-sm">Name</span>
              <v-icon icon="mdi-arrow-up-down" class="tw-text-sm"></v-icon>
            </div>
          </th>
          <th class="text-center">
            <span class="tw-text-sm">Progress</span>
          </th>
          <th class="text-center">
            <span class="tw-text-sm">Status</span>
          </th>
          <th class="text-left tw-cursor-pointer" @click="store.sortByTimeElapsed()">
            <div class="tw-flex tw-justify-between tw-items-center tw-w-max md:tw-w-full">
              <span class="tw-text-sm tw-mr-3 tw-w-32">Time Elapsed</span>
              <v-icon icon="mdi-arrow-up-down" class="tw-text-sm"></v-icon>
            </div>
          </th>
          <th class="text-left tw-cursor-pointer" @click="store.sortBySize">
            <div class="tw-flex tw-justify-between tw-items-center">
              <span class="tw-text-sm tw-mr-3 tw-w-20">Size</span>
              <v-icon icon="mdi-arrow-up-down" class="tw-text-sm"></v-icon>
            </div>
          </th>
          <th class="text-left tw-cursor-pointer" @click="store.sortByDate()">
            <div class="tw-flex tw-justify-between tw-items-center">
              <span class="tw-text-sm tw-w-32">Date</span>
              <v-icon icon="mdi-arrow-up-down" class="tw-text-sm"></v-icon>
            </div>
          </th>
          <th class="text-center">
            <div class="tw-flex tw-justify-between tw-items-center">
              <span class="tw-text-sm tw-w-10">Action</span>
            </div>
          </th>
        </tr>
      </thead>
      <tbody class="tw-relative">
        <tr v-for="item in items" :key="item.name">
          <td color="primary" class="tw-rounded-sm bordered name-col">
            <div v-if="item.progress != 100" class="progressWrapper tw-flex tw-justify-between tw-absolute tw-w-full tw-left-0 tw-right-0 tw-px-5 tw-mt-1" :id="item.id">
              <div v-for="part in item.metadata.totalpart" :class="`tw-w-full ` + `basis-1/${item.metadata.totalpart}`" >
                <div class="tw-h-9 tw-bg-green-500 tw-opacity-10 tw-w-0 -tw-mt-1.5" :id="`progressBar-${item.id}-${part-1}`"></div>
              </div> 
            </div>
            <div class="tw-flex tw-justify-between tw-mt-1 tw-mr-3 tw-items-center group">
              <div class="tw-overflow-x-hidden tw-w-max tw-flex">
                <v-icon :icon="item.type.icon" :color="item.type.color" class="tw-opacity-70 tw-mr-2"></v-icon>
                <!-- TODO: add mark if not resumable -->
                <span class="tw-text-sm tw-inline">{{ item.name }}</span>
              </div>
              <Actionable :active="item.progress != 100" :filename="item.name" :id="item.id" :statusname="item.status.name"/>
            </div>
          </td>
          <td class="tw-text-sm bordered tw-text-center">{{ item.progress.toFixed(0) + '%' }}</td>
          <td class="tw-text-sm tw-text-center bordered">
            <v-tooltip :text="item.status.name" location="bottom">
              <template v-slot:activator="{ props }">
                <v-icon v-bind="props" :icon="item.status.icon" :color="item.status.color" class="tw-opacity-90"></v-icon>
              </template>
            </v-tooltip>
          </td>
          <td class="tw-text-sm tw-text-left bordered">{{ item.timeElapsed }}</td>
          <td class="tw-text-sm tw-text-left bordered">{{ parseSize(item.size) }}</td>
          <td class="tw-text-sm tw-text-left bordered">{{ useDateFormat(item.date, 'MMMM DD, YYYY HH:mm').value }}</td>
          <td class="tw-text-sm tw-text-center bordered tw-w-10">
            <Deletable :filename="item.name" :id="item.id"/>
          </td>
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
.bordered {
  border-left: 1px solid rgb(87, 83, 78);
  border-left: 1px solid rgb(87, 83, 78, 0.05);
  border-right: 1px solid rgba(87, 83, 78, 0.05);
  border-right: 1px solid rgba(87, 83, 78, 0.05);
  -webkit-background-clip: padding-box; /* for Safari */
  background-clip: padding-box; /* for IE9+, Firefox 4+, Opera, Chrome */
}
.name-col:hover .icon-name {
  opacity: 100;
}
</style>