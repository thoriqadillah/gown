<script setup lang="ts">
import { useDownloads } from '../../store/downloads';
import { useSettings } from '../../store/setting';
import { ref, watch } from 'vue';
import { parseSize } from '../../utils/parser';
import { download } from '../../../wailsjs/go/models';
import Downloader from '../../services/downloader'
import Dialog from '../../services/download-dialog'

const downloads = useDownloads()
const settings = useSettings()

const result = ref<download.Download>()
const input = ref()
const loading = ref(false)
const activator = ref(false)
const loaded = ref(false)
const onFile = ref(false)
const onURL = ref(true)

let url = ref('')

const urlHasError = ref(false)
const savelocationHasError = ref(false)
const filenameHasError = ref(false)

const validate = ref({
  required: (v: string) => !!v.trim() || "This field cannot be empty",
  max: (v: string) => v.trim().length < 256 || "File name must be below 256 characters"
})

watch([url, () => result.value?.name, settings], () => {
  urlHasError.value = url.value.trim().length === 0
  savelocationHasError.value = settings.saveLocation.trim().length === 0
  filenameHasError.value = result.value!.name.trim().length === 0 || result.value!.name.trim().length > 256
})

const dialog = new Dialog({ activator, loaded, loading, onFile, onURL })
const downloader = Downloader.service()

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
    // TODO: add alert to error message
  }
}

async function execute() {
  dialog.close()
  downloader.download(result.value!)
  downloads.add(result.value!)
}
</script>

<template>
  <v-dialog v-model="activator" activator="parent" max-width="450px" transition="dialog-top-transition" persistent>
    <v-card>
      <v-text-field v-if="onURL || !loaded" :rules="[validate.required]" v-model="url" :loading="loading" :autofocus="activator" color="primary" type="input" hint="Click enter to fetch the file data from the URL you want to download" class="tw-p-3" density="compact" variant="outlined" label="URL" append-inner-icon="mdi-link" append-icon="mdi-magnify" @click:append="fetch()" single-line v-on:keyup.enter="fetch()" ref="input"/>
      <div v-else-if="onFile || !loaded" class="tw-flex tw-items-center">
        <div class="tw-basis-9/12">
          <v-text-field color="primary" :rules="[validate.required, validate.max]" v-model="result!.name" label="File name" append-inner-icon="mdi-file-document-edit" type="input" hint="File name" class="tw-px-3 tw-pt-3 -tw-mb-2" single-line density="compact" variant="outlined" />
          <v-text-field color="primary" :rules="[validate.required]" v-model="settings.saveLocation" label="Save location" append-inner-icon="mdi-folder" type="input" hint="Save location" class="tw-p-3" single-line density="compact" variant="outlined" />
        </div>
        <div class="tw-basis-3/12 tw-text-center tw-pr-2 -tw-mt-5">
          <v-icon :icon="result!.type.icon" :color="result!.type.color"></v-icon>
          <p class="text-body-1 tw-mt-5">{{ parseSize(result!.size) }}</p>
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