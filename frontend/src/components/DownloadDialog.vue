<script setup lang="ts">
import { ref } from 'vue';
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
const url = ref('')
const size = ref('')

const dialog = new Dialog({ activator, loaded, loading, onFile, onURL })
const downloader = new Downloader()

async function fetch() {
  input.value.blur()
  dialog.loading()

  result.value = await downloader.fetch(url.value) as download.Download
  size.value = downloads.parseSize(result.value.size)

  dialog.done()
  dialog.next()
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
      <v-text-field v-if="onURL || !loaded" v-model="url" :loading="loading" :autofocus="activator" color="primary" type="input" hint="Click enter to fetch the file data from the URL you want to download" class="tw-p-3" density="compact" variant="outlined" label="URL" append-inner-icon="mdi-link" append-icon="mdi-magnify" @click:append="fetch()" single-line v-on:keyup.enter="fetch()" ref="input"/>
      <div v-else-if="onFile || !loaded" class="tw-flex tw-items-center">
        <div class="tw-basis-9/12">
          <v-text-field color="primary" v-model="result.name" label="File name" append-inner-icon="mdi-file-document-edit" class="tw-px-3 tw-pt-3 -tw-mb-4" single-line v-on:keyup.enter="fetch" density="compact" variant="outlined" ref="input"/>
          <v-text-field color="primary" v-model="settings.value.saveLocation" label="Save location" append-inner-icon="mdi-folder" type="input" hint="Save location" class="tw-p-3" single-line v-on:keyup.enter="fetch" density="compact" variant="outlined" ref="input"/>
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
            <v-btn color="primary" block @click="execute()" :disabled="onURL || !loaded">Download</v-btn>
          </v-card-actions>
          <v-card-actions>
            <v-btn variant="text" block @click="dialog.close()">Cancel</v-btn>
          </v-card-actions>
        </div>
      </div>
    </v-card>
  </v-dialog>
</template>