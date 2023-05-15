<script setup lang="ts">
import { ref, watch } from 'vue';
import { useSettings } from '../../store/setting';
import { DefaultSetting } from "../../../wailsjs/go/store/fileStore";

const setting = useSettings()
const activator = ref(false)
watch(activator, (val) => emit('dialog', val))

const emit = defineEmits<{
  (emit: 'dialog', val: boolean): void
}>()

const saveLocation = ref(setting.saveLocation)
const workers = ref(setting.concurrency)
const maxtries = ref(setting.maxtries)

</script>

<template>
  <v-dialog v-model="activator" width="450px" activator="parent" persistent>
    <v-card class="tw-py-3">
      <v-card-title>
        <p class="text-h6 tw-ml-1">Setting</p>
      </v-card-title>

      <div class="tw-mx-5 tw-mt-5 tw-mb-3">
        <p class="text-body-2 tw-mb-1">Download location</p>
        <v-text-field  v-model="setting.saveLocation" color="primary" type="input" hint="Make sure it's a valid folder" density="compact" variant="outlined" label="Save location" single-line/>  
      </div>
      
      <div class="tw-mx-5 tw-flex">
        <div class="tw-basis-3/4">
          <p class="text-body-1">Workers</p>
          <p class="text-body-2 tw-opacity-70">Amount of workers to perform concurrent download</p>
        </div>
        <div class="tw-basis-1/4 tw-mt-1">
          <v-text-field  v-model="setting.concurrency" color="primary" type="number" variant="outlined" label="Workers" single-line/>  
        </div>
      </div>

      <div class="tw-mx-5 tw-flex">
        <div class="tw-basis-3/4">
          <p class="text-body-1">Max retries</p>
          <p class="text-body-2 tw-opacity-70">Amount of retries when error happens while downloading</p>
        </div>
        <div class="tw-basis-1/4 tw-mt-1">
          <v-text-field  v-model="setting.maxtries" color="primary" type="number" variant="outlined" label="Workers" single-line/>  
        </div>
      </div>

      <div class="tw-mx-5 tw-mt-5 tw-opacity-70 tw-flex tw-items-center tw-gap-2">
        <v-icon>mdi-alert-circle-outline</v-icon>
        <div>
          <p class="text-caption">Changes will take effect only after being saved. You can always go back to default</p>
        </div>
      </div>

      <div class="tw-flex tw-justify-between">
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn variant="text" @click="setting.backDefault()">Set Default</v-btn>
        </v-card-actions>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn variant="text" @click="activator = false">Close</v-btn>
          <v-btn color="primary" variant="text" @click="() => { setting.update(); activator = false }">Save</v-btn>
        </v-card-actions>
      </div>
    </v-card>
  </v-dialog>
</template>