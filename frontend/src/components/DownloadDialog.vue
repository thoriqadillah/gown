<script setup lang="ts">
import { useDownloads } from '../store/downloads';
import { useSettings } from '../store/setting';
import { ref, watch } from 'vue';
import { parseSize } from '../utils/parser';
import { download } from '../../wailsjs/go/models';
import Downloader from '../services/downloader'
import Dialog from '../services/download-dialog'

const downloads = useDownloads()
const settings = useSettings()

const result = ref(new download.Download())
const input = ref()
const loading = ref(false)
const activator = ref(false)
const loaded = ref(false)
const onFile = ref(false)
const onURL = ref(true)

let url = ref('')
const urlErr = ref('')
const urlHasError = ref(false)
const savelocationErr = ref('')
const savelocationHasError = ref(false)
const filenameErr = ref('')
const filenameHasError = ref(false)

const dialog = new Dialog({ activator, loaded, loading, onFile, onURL })
const downloader = new Downloader()

watch([url, () => result.value.name, settings], () => {
  urlHasError.value = url.value.trim().length === 0
  urlErr.value = url.value.trim().length === 0 ? 'This field is required' : ''
  
  savelocationHasError.value = settings.saveLocation.trim().length === 0
  savelocationErr.value = settings.saveLocation.trim().length === 0 ? 'This field is required' : ''

  if (result.value.name.trim().length === 0) {
    filenameErr.value = 'This field is required'
    filenameHasError.value = true
  } else if (result.value.name.trim().length > 256) {
    filenameErr.value = 'File name must be below 256 characters'
    filenameHasError.value = true
  } else {
    filenameErr.value = ''
    filenameHasError.value = false
  }
})

async function fetch() {
  try {
    input.value.blur()
    dialog.loading()
    
    result.value = await downloader.fetch(url.value) as download.Download
    
    dialog.done()
    dialog.next()
  } catch (error) {
    input.value.blur()
    dialog.done()
    urlHasError.value = true
    urlErr.value = error as string
  }
}

async function execute() {
  dialog.close()
  downloader.download(result.value)
  downloads.add(result.value)
  url = ref('')
  urlErr.value = ''
  urlHasError.value = false
  filenameErr.value = ''
  filenameHasError.value = false
  savelocationErr.value = ''
  savelocationHasError.value = false
}
</script>

<template>
  <v-dialog v-model="activator" activator="parent" max-width="450px" transition="dialog-top-transition" persistent>
    <v-card>
      <v-text-field v-if="onURL || !loaded" :error="urlHasError" :error-messages="urlErr" v-model="url" :loading="loading" :autofocus="activator" color="primary" type="input" hint="Click enter to fetch the file data from the URL you want to download" class="tw-p-3" density="compact" variant="outlined" label="URL" append-inner-icon="mdi-link" append-icon="mdi-magnify" @click:append="fetch()" single-line v-on:keyup.enter="fetch()" ref="input"/>
      <div v-else-if="onFile || !loaded" class="tw-flex tw-items-center">
        <div class="tw-basis-9/12">
          <v-text-field color="primary" :error="filenameHasError" :error-messages="filenameErr" v-model="result.name" label="File name" append-inner-icon="mdi-file-document-edit" type="input" hint="File name" class="tw-px-3 tw-pt-3 -tw-mb-2" single-line density="compact" variant="outlined" />
          <v-text-field color="primary" :error="savelocationHasError" :error-messages="savelocationErr" v-model="settings.saveLocation" label="Save location" append-inner-icon="mdi-folder" type="input" hint="Save location" class="tw-p-3" single-line density="compact" variant="outlined" />
        </div>
        <div class="tw-basis-3/12 tw-text-center tw-pr-2 -tw-mt-5">
          <v-icon :icon="result.type.icon" :color="result.type.color"></v-icon>
          <p class="text-body-1 tw-mt-5">{{ parseSize(result.size) }}</p>
        </div>
      </div>
      
      <div class="tw-flex tw-justify-between">
        <div class="tw-flex tw-flex-row-reverse">
          <v-card-actions>
            <v-btn density="compact" :disabled="onFile || !loaded" @click="dialog.next()" icon="mdi-arrow-right"></v-btn>
          </v-card-actions>
          <v-card-actions>
            <v-btn density="compact" :disabled="onURL || !loaded" @click="dialog.prev()" icon="mdi-arrow-left"></v-btn>
          </v-card-actions>
        </div>
        <div class="tw-flex tw-flex-row-reverse">
          <v-card-actions>
            <v-btn color="primary" block @click="execute()" :disabled="(onURL || !loaded) || urlHasError || filenameHasError || savelocationHasError">Download</v-btn>
          </v-card-actions>
          <v-card-actions>
            <v-btn variant="text" block @click="dialog.close()">Cancel</v-btn>
          </v-card-actions>
        </div>
      </div>
    </v-card>
  </v-dialog>
</template>