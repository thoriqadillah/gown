<script setup lang="ts">
import { ref } from 'vue';

const activator = ref(false)
const loading = ref(false)
const loaded = ref(false)

// TODO: do a fetch of the link
const input = ref<HTMLInputElement>()
const onFile = ref(false)
const onURL = ref(true)
const slide = ref('slide-right')
function fetch() {
  loading.value = true

  setTimeout(() => {
    loading.value = false
    loaded.value = true
  }, 2000)

  onFile.value = true
  onURL.value = false
  input.value?.blur()
}

// TODO: implement download
function download() {
  onURL.value = true
  onFile.value = false
  activator.value = false
}

function next() {
  onFile.value = true
  onURL.value = false
  slide.value = 'slide-left'
}

function prev() {
  onURL.value = true
  onFile.value = false
  slide.value = 'slide-right'
}

</script>

<template>
  <v-dialog v-model="activator" activator="parent" width="40%" transition="dialog-top-transition" persistent>
    <v-card>
      <v-text-field v-if="onURL || !loaded" :loading="loading" color="primary" type="input" hint="Click enter to fetch the file data from the URL you want to download" class="tw-p-3" density="compact" variant="outlined" label="URL" append-inner-icon="mdi-link" single-line v-on:keyup.enter="fetch" ref="input"/>
      <div v-else-if="onFile || !loaded" class="tw-flex tw-items-center">
        <div class="tw-basis-5/6">
          <v-text-field color="primary" label="File name" append-inner-icon="mdi-file-document-edit" class="tw-px-3 tw-pt-3 -tw-mb-4" single-line v-on:keyup.enter="fetch" density="compact" variant="outlined" ref="input"/>
          <v-text-field color="primary" label="Save location" append-inner-icon="mdi-folder" type="input" hint="Save location" class="tw-p-3" single-line v-on:keyup.enter="fetch" density="compact" variant="outlined" ref="input"/>
        </div>
        <div class="tw-basis-1/6 tw-text-center tw-pr-2">
          <v-icon icon="mdi-file"></v-icon>
          <p class="text-body-1 tw-mt-5">40 MB</p>
        </div>
      </div>
      
      <div class="tw-flex tw-justify-between">
        <div class="tw-flex tw-flex-row-reverse">
          <v-card-actions>
            <v-btn density="compact" :disabled="onFile || !loaded" @click="next()" icon="mdi-arrow-right"></v-btn>
          </v-card-actions>
          <v-card-actions>
            <v-btn density="compact" :disabled="onURL || !loaded" @click="prev()" icon="mdi-arrow-left"></v-btn>
          </v-card-actions>
        </div>
        <div class="tw-flex tw-flex-row-reverse">
          <v-card-actions>
            <v-btn color="primary" block @click="activator = false" :disabled="!loaded">Download</v-btn>
          </v-card-actions>
          <v-card-actions>
            <v-btn variant="text" block @click="activator = false">Cancel</v-btn>
          </v-card-actions>
        </div>
      </div>
    </v-card>
  </v-dialog>
</template>