<script setup lang="ts">
import { ref } from 'vue'
import { useSettings } from '../../store/setting';

const setting = useSettings()
const dialog = ref(false)

const validation = ref({
  required: (v: string) => !!v.trim() || "This field cannot be empty",
  max: (v: string) => v.trim().length < 256 || "File name must be below 256 characters"
})

</script>

<template>
  <v-dialog activator="parent" v-model="dialog" persistent width="450px">
    <v-card>
      <v-card-title class="tw-mt-2">
        <span class="text-h5 tw-ml-1.5">Settings</span>
      </v-card-title>

      <v-card-text>
        <p class="text-caption tw-mb-1">Save Location</p>
        <v-text-field variant="outlined" density="compact" v-model="setting.saveLocation" :rules="[validation.required]" hint="Make sure it's a valid location" single-line></v-text-field>
      </v-card-text>

      <v-card-text class="-tw-mt-5 tw-flex tw-justify-between">
        <div class="tw-basis-2/4">
          <p class="">Workers</p>
          <p class="text-caption tw-opacity-80">Number of downloadable chunks concurrently</p>
        </div>
        <div class="tw-my-auto -tw-mb-3 tw-basis-1/4">
          <v-text-field variant="outlined" type="number" v-model="setting.concurrency" density="comfortable" :rules="[validation.required]" hint="Number of downloadable chunks concurrently" single-line></v-text-field>
        </div>
      </v-card-text>

      <v-card-text class="-tw-mt-5 tw-flex tw-justify-between">
        <div class="tw-basis-2/4">
          <p class="">Maximum tries</p>
          <p class="text-caption tw-opacity-80">Number of retry when download failed</p>
        </div>
        <div class="tw-my-auto -tw-mb-3 tw-basis-1/4">
          <v-text-field variant="outlined" type="number" v-model="setting.maxtries" density="comfortable" :rules="[validation.required]" hint="Number of downloadable chunks concurrently" single-line></v-text-field>
        </div>
      </v-card-text>
      
      <v-card-actions class="tw-mb-2">
        <v-spacer></v-spacer>
        <v-btn variant="text" @click="dialog = false">
          Close
        </v-btn>
        <v-btn color="primary" variant="text" @click="dialog = false">
          Save
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>