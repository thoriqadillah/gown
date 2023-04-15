<script setup lang="ts">
import { Download } from '../types/download';
import { defineProps } from 'vue';

const props = defineProps<{
  list: Download[]
}>()
</script>

<template>
  <div class="tw-px-5 xl:tw-px-0 xl:tw-pl-5">
    <div class="tw-flex tw-gap-2 md:tw-gap-4 tw-mb-2">
      <v-btn variant="outlined" prepend-icon="mdi-image" color="red-accent-2">IMAGE</v-btn>
      <v-btn variant="outlined" prepend-icon="mdi-video" color="deep-orange-accent-2">VIDEO</v-btn>
      <v-btn variant="outlined" prepend-icon="mdi-file-document" color="blue-accent-2">DOCS</v-btn>
      <v-btn variant="outlined" prepend-icon="mdi-zip-box" color="yellow-accent-4">COMPRESSED</v-btn>
      <v-btn variant="outlined" prepend-icon="mdi-music-box" color="purple-accent-2">MUSIC</v-btn>
    </div>

    <v-table density="compact">
      <thead>
        <tr>
          <th class="text-left tw-cursor-pointer">
            <div class="tw-flex tw-justify-between tw-items-center">
              <span class="tw-text-sm">Name</span>
              <v-icon icon="mdi-arrow-up-down" class="tw-text-sm"></v-icon>
            </div>
          </th>
          <th class="text-left tw-cursor-pointer">
            <div class="tw-flex tw-justify-between tw-items-center tw-w-max md:tw-w-full">
              <span class="tw-text-sm tw-mr-3">Time Elapsed</span>
              <v-icon icon="mdi-arrow-up-down" class="tw-text-sm"></v-icon>
            </div>
          </th>
          <th class="text-left tw-cursor-pointer">
            <div class="tw-flex tw-justify-between tw-items-center">
              <span class="tw-text-sm tw-mr-3">Size</span>
              <v-icon icon="mdi-arrow-up-down" class="tw-text-sm"></v-icon>
            </div>
          </th>
          <th class="text-left tw-cursor-pointer">
            <div class="tw-flex tw-justify-between tw-items-center">
              <span class="tw-text-sm">Date</span>
              <v-icon icon="mdi-arrow-up-down" class="tw-text-sm"></v-icon>
            </div>
          </th>
          <th class="tw-text tw-cursor-pointer-left">
            <span class="tw-text-sm">Status</span>
          </th>
        </tr>
      </thead>
      <tbody >
        <tr v-for="item in props.list" :key="item.name">
          <td class="tw-rounded-sm">
            <div class="tw-flex tw-justify-between">
              <div class="tw-w-max">
                <v-icon :icon="item.type.icon" :color="item.type.color" class="tw-opacity-70 tw-mr-2"></v-icon>
                <span class="tw-text-sm">{{ item.name }}</span>
              </div>
            </div>
          </td>
          <td class="tw-text-sm tw-rounded-sm text-left">{{ item.timeElapsed }}</td>
          <td class="tw-text-sm tw-rounded-sm text-left">{{ item.size }}</td>
          <div class="tw-w-max mt-2 tw-text-left">
            <td class="tw-text-sm tw-rounded-sm text-left">{{ item.date }}</td>
          </div>
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