<script setup lang="ts">
import { defineProps } from 'vue';
import { useDownloads } from '../store/downloads';
import { EventsOn } from '../../wailsjs/runtime/runtime';
import { download } from '../../wailsjs/go/models';
import { useDateFormat } from '@vueuse/shared';
import { ref, watch } from 'vue';

const downloads = useDownloads()
const props = defineProps<{
  list: download.Download[]
}>()

const editIcon = ref<HTMLElement[]>([])
const filenames = ref<HTMLElement[]>() 
const hovered = ref<HTMLElement>()

//FIXME: fix editable newly downloaded file
watch(filenames, () => {
  for (let i = 0; i < filenames.value!.length; i++) {
    filenames.value![i].addEventListener('mouseover', () => {
      editIcon.value[i].style.opacity = '100'
      hovered.value = filenames.value![i]
    });
    filenames.value![i].addEventListener('mouseleave', () => {
      editIcon.value[i].style.opacity = '0'
    });
  }
})

const click = ref(0)
const onEdit = ref(false)
const selectedID = ref('')
const newFilename = ref()
const selected = ref()

function editFilename() {
  click.value += 1
  if (click.value == 2) {
    onEdit.value = true
    selectedID.value = hovered.value!.id
    selected.value = hovered.value
    newFilename.value = selected.value.innerText
    
    click.value = 0
  }
}

async function doneEditing() {
  try {
    onEdit.value = false
    for (const el of downloads.list) {
      if (el.id == selectedID.value) {
        await downloads.updateName(el.name, newFilename.value, el.id)
        el.name = newFilename.value
        break
      }
    }
  
    await downloads.updateData(downloads.list)
  } catch (error) {
    console.log(error);
    // TODO: print this to log
  }
}

// TODO: fix the progress bar not responding
// EventsOn("transfered", async (...data) => {
//   console.log(data);
  
//   let prog = data[2]
//   const progressBar = document.getElementById(`progressBar-${data[0]}-${data[1]}`) as HTMLElement
//   progressBar.style.opacity = '50'
//   progressBar.style.width = prog + '%'

//   for (const el of downloads.list) {
//     if (el.id == data[0]) {
//       el.timeElapsed = downloads.parseElapsedTime(downloads.toDownload.date)
//       el.progress += data[3]
//       break
//     }
//   }
// })

EventsOn("downloaded", async (...data) => {
  const [id, combined] = data
  
  for (const el of downloads.list) {
    if (el.id != id) continue

    if (!combined) {
      el.status = {
        name: 'Combining',
        icon: 'mdi-file-arrow-left-right-outline',
        color: 'info'
      }

      // TODO: delete this
      el.timeElapsed = downloads.parseElapsedTime(downloads.toDownload.date)
      el.progress += downloads.toDownload.size
      break
    } 

    el.status = {
      name: 'Success',
      icon: 'mdi-check-circle-outline',
      color: 'success'
    }
    for (let i = 0; i < el.metadata.totalpart; i++) {
      const progressBar = document.getElementById(`progressBar-${data[0]}-${i}`) as HTMLElement
      progressBar.style.opacity = '0'
    } 
    break
  }

  await downloads.updateData(downloads.list)
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
          <th class="tw-text tw-cursor-pointer-left">
            <span class="tw-text-sm">Progress</span>
          </th>
          <th class="tw-text tw-cursor-pointer-left">
            <span class="tw-text-sm">Status</span>
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
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in props.list" :key="item.name">
          <td color="primary" class="tw-rounded-sm namecol" :id="item.id" ref="filenames" @click="editFilename()">
            <div class="tw-flex tw-justify-between tw-mt-1">
              <div class="tw-overflow-x-hidden tw-w-max tw-flex">
                <v-icon :icon="item.type.icon" :color="item.type.color" class="tw-opacity-70 tw-mr-2"></v-icon>
                <input v-if="selectedID == item.id && onEdit" type="text" v-model="newFilename" autofocus :size="item.name.length-13" class="tw-text-sm border tw-pb-1" @keyup.enter="doneEditing()" @keyup.esc="onEdit = false">
                <span v-else class="tw-text-sm tw-inline">{{ item.name }}</span>
              </div>
              <span ref="editIcon" class="tw-opacity-0">
                <v-icon icon="mdi-square-edit-outline" class="tw-text-sm tw-opacity-50 tw-mx-3 edit-icon tw-mb-1" @click="onEdit = true"></v-icon>
              </span>
            </div>
            <div class="progressWrapper tw-flex tw-justify-between" :id="item.id">
              <div v-for="part in item.metadata.totalpart" :class="`tw-w-full ` + `basis-1/${item.metadata.totalpart}`" >
                <div class="tw-h-0.5 tw-bg-green-500 tw-opacity-0 tw-mt-1 tw-w-1 tw-rounded-lg" :id="`progressBar-${item.id}-${part-1}`"></div>
              </div> 
            </div>
          </td>
          <td class="tw-text-sm tw-text-center">{{  ((item.progress/item.size) * 100).toFixed(0) + '%' }}</td>
          <td class="tw-text-sm tw-text-center">
            <v-tooltip :text="item.status.name" location="top">
              <template v-slot:activator="{ props }">
                <v-icon v-bind="props" :icon="item.status.icon" :color="item.status.color" class="tw-opacity-90 tw-ml-2"></v-icon>
              </template>
            </v-tooltip>
          </td>
          <td class="tw-text-sm tw-text-left tw-w-32">{{ item.timeElapsed }}</td>
          <td class="tw-text-sm tw-text-left tw-w-20">{{ downloads.parseSize(item.size) }}</td>
          <td class="tw-text-sm tw-text-left tw-w-32">{{ useDateFormat(item.date, 'MMMM DD, YYYY HH:mm').value }}</td>
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