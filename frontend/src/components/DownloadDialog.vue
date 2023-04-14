<script setup lang="ts">
import { ref } from 'vue';
import { Fetch } from '../../wailsjs/go/main/App'

const activator = ref(false)
const loading = ref(false)
const loaded = ref(false)
const url = ref('')
const filename = ref('')
const size = ref()
const saveLocation = ref('')

// TODO: do a fetch of the link
const input = ref<HTMLInputElement>()
const onFile = ref(false)
const onURL = ref(true)
function fetch() {
  input.value?.blur()
  loading.value = true

  Fetch(url.value).then(res => {
    const MB = 1024 * 1024
    
    filename.value = res.Filename
    size.value = (res.Size / MB).toFixed(2)
    saveLocation.value = res.SaveLocation

    loading.value = false
    loaded.value = true
    onFile.value = true
    onURL.value = false
  })
}

// TODO: implement download
function download() {
  onURL.value = true
  onFile.value = false
  activator.value = false
}

function cancel() {
  onURL.value = true
  onFile.value = false
  activator.value = false
}

function next() {
  onFile.value = true
  onURL.value = false
}

function prev() {
  onURL.value = true
  onFile.value = false
}

</script>

<template>
  <v-dialog v-model="activator" activator="parent" max-width="450px" transition="dialog-top-transition" persistent>
    <v-card>
      <v-text-field v-if="onURL || !loaded" v-model="url" :loading="loading" color="primary" type="input" hint="Click enter to fetch the file data from the URL you want to download" class="tw-p-3" density="compact" variant="outlined" label="URL" append-inner-icon="mdi-link" single-line v-on:keyup.enter="fetch" ref="input"/>
      <div v-else-if="onFile || !loaded" class="tw-flex tw-items-center">
        <div class="tw-basis-9/12">
          <v-text-field color="primary" v-model="filename" label="File name" append-inner-icon="mdi-file-document-edit" class="tw-px-3 tw-pt-3 -tw-mb-4" single-line v-on:keyup.enter="fetch" density="compact" variant="outlined" ref="input"/>
          <v-text-field color="primary" v-model="saveLocation" label="Save location" append-inner-icon="mdi-folder" type="input" hint="Save location" class="tw-p-3" single-line v-on:keyup.enter="fetch" density="compact" variant="outlined" ref="input"/>
        </div>
        <div class="tw-basis-3/12 tw-text-center tw-pr-2 -tw-mt-5">
          <v-icon icon="mdi-file"></v-icon>
          <p class="text-body-1 tw-mt-5">{{ size }} MB</p>
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
            <v-btn color="primary" block @click="download()" :disabled="!loaded">Download</v-btn>
          </v-card-actions>
          <v-card-actions>
            <v-btn variant="text" block @click="cancel()">Cancel</v-btn>
          </v-card-actions>
        </div>
      </div>
    </v-card>
  </v-dialog>
</template>