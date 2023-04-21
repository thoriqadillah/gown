<script setup lang="ts">
import { ref, watch } from 'vue';
import { download } from '../../wailsjs/go/models';
import Dialog from '../services/download-dialog'
import Downloader from '../services/downloader'
import { useDownloads } from '../store/downloads';
import { useSettings } from '../store/setting';

const downloads = useDownloads()
const settings = useSettings()

const result = ref(new download.Download())
const input = ref()
const loading = ref(false)
const activator = ref(false)
const loaded = ref(false)
const onFile = ref(false)
const onURL = ref(true)
const size = ref('')
const url = ref('')
const urlErr = ref('')
const urlHasError = ref(false)
const savelocation = ref('')
const savelocationErr = ref('')
const savelocationHasError = ref(false)
const filename = ref('')
const filenameErr = ref('')
const filenameHasError = ref(false)

const dialog = new Dialog({ activator, loaded, loading, onFile, onURL })
const downloader = new Downloader()

watch(url, (newval, oldval) => {
  urlHasError.value = url.value.trim().length === 0
  urlErr.value = url.value.trim().length === 0 ? 'This field is required' : ''
})
watch(filename, (newval, oldval) => {
  if (filename.value.trim().length === 0) {
    filenameErr.value = 'This field is required'
    filenameHasError.value = true
  } else if (filename.value.trim().length > 256) {
    filenameErr.value = 'File name must be below 256 characters'
    filenameHasError.value = true
  } else {
    filenameErr.value = ''
    filenameHasError.value = false
  }
})
watch(savelocation, (newval, oldval) => {
  savelocationHasError.value = savelocation.value.trim().length === 0
  savelocationErr.value = savelocation.value.trim().length === 0 ? 'This field is required' : ''
})


async function fetch() {
  try {
    input.value.blur()
    dialog.loading()
    
    result.value = await downloader.fetch(url.value) as download.Download
    size.value = downloads.parseSize(result.value.size)
    filename.value = result.value.name
    savelocation.value = settings.value.saveLocation
    
    dialog.done()
    dialog.next()
  } catch (error) {
    input.value.blur()
    dialog.done()
    urlHasError.value = true
    urlErr.value = error as string
  }
}

// TODO: implement download
async function execute() {
  downloader.download(result.value)
  downloads.add(result.value)
  dialog.close()
}

</script>

<template>
  <v-dialog v-model="activator" activator="parent" max-width="450px" transition="dialog-top-transition" persistent>
    <v-card>
      <v-text-field v-if="onURL || !loaded" :error="urlHasError" :error-messages="urlErr" v-model="url" :loading="loading" :autofocus="activator" color="primary" type="input" hint="Click enter to fetch the file data from the URL you want to download" class="tw-p-3" density="compact" variant="outlined" label="URL" append-inner-icon="mdi-link" append-icon="mdi-magnify" @click:append="fetch()" single-line v-on:keyup.enter="fetch()" ref="input"/>
      <div v-else-if="onFile || !loaded" class="tw-flex tw-items-center">
        <div class="tw-basis-9/12">
          <v-text-field color="primary" :error="filenameHasError" :error-messages="filenameErr" v-model="filename" label="File name" append-inner-icon="mdi-file-document-edit" type="input" hint="File name" class="tw-px-3 tw-pt-3 -tw-mb-2" single-line density="compact" variant="outlined" />
          <v-text-field color="primary" :error="savelocationHasError" :error-messages="savelocationErr" v-model="savelocation" label="Save location" append-inner-icon="mdi-folder" type="input" hint="Save location" class="tw-p-3" single-line density="compact" variant="outlined" />
        </div>
        <div class="tw-basis-3/12 tw-text-center tw-pr-2 -tw-mt-5">
          <v-icon icon="mdi-file"></v-icon>
          <p class="text-body-1 tw-mt-5">{{ size }}</p>
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