<script setup lang="ts">
import DownloadDialog from './DownloadDialog.vue';
import { useTheme } from 'vuetify/lib/framework.mjs';
import { ref } from 'vue';
import { useDrawerStore } from '../store/drawer';

const theme = useTheme()
const themeIcon = ref(theme.global.current.value.dark ? 'mdi-white-balance-sunny' : 'mdi-weather-night')
const themeTooltip = ref(theme.global.current.value.dark ? 'Light Mode' : 'Dark Mode')
function toggleTheme() {
  theme.global.name.value = theme.global.current.value.dark ? 'light' : 'customDark'
  themeIcon.value = theme.global.current.value.dark ? 'mdi-white-balance-sunny' : 'mdi-weather-night'
  themeTooltip.value = theme.global.current.value.dark ? 'Light Mode' : 'Dark Mode'
}

let loading = ref(false)
let loaded = ref(false)
let input = ref<HTMLInputElement>()
function search() {
  loading.value = true

  setTimeout(() => {
    loading.value = false
    loaded.value = true
  }, 2000)

  input.value?.blur()
}
const store = useDrawerStore()

</script>

<template>
  <nav class="tw-mx-5 tw-my-5 tw-flex tw-justify-between tw-gap-8 tw-items-center">
    <p class="text-button">GOWN</p>

    <v-text-field :loading="loading" density="comfortable" variant="outlined" color="primary" label="Search" append-inner-icon="mdi-magnify" single-line hide-details @click:append-inner="search" v-on:keyup.enter="search" ref="input"/>

    <div class="tw-flex tw-gap-5">
      <v-tooltip :text="themeTooltip" location="bottom">
        <template v-slot:activator="{ props }">
          <v-btn v-bind="props" density="comfortable" :icon="themeIcon" @click="toggleTheme" variant="flat"/>
        </template>
      </v-tooltip>

      <v-divider vertical/>

      <v-tooltip text="Resume All" location="bottom">
        <template v-slot:activator="{ props }">
          <v-btn v-bind="props" density="comfortable" icon="mdi-play" variant="flat"/>
        </template>
      </v-tooltip>

      <v-tooltip text="Pause All" location="bottom">
        <template v-slot:activator="{ props }">
          <v-btn v-bind="props" density="comfortable" icon="mdi-pause" variant="flat"/>
        </template>
      </v-tooltip>

      <v-tooltip text="Stop All" location="bottom">
        <template v-slot:activator="{ props }">
          <v-btn v-bind="props" density="comfortable" icon="mdi-stop" variant="flat"/>
        </template>
      </v-tooltip>

      <v-divider vertical/>
      
      <v-tooltip text="Add into Queue" location="bottom">
        <template v-slot:activator="{ props }">
          <v-btn v-bind="props" density="comfortable" icon="mdi-tray-plus" variant="flat"/>
        </template>
      </v-tooltip>

      <v-tooltip text="New" location="bottom">
        <template v-slot:activator="{ props }">
          <v-btn v-bind:="props" variant="flat" density="comfortable" icon>
            <v-icon icon="mdi-plus"></v-icon>
            <download-dialog/>
          </v-btn>
        </template>
      </v-tooltip>

      <v-tooltip text="Menu" location="bottom">
        <template v-slot:activator="{ props }">
          <v-btn class="xl:tw-hidden" v-bind="props" density="comfortable" icon="mdi-dots-vertical" variant="flat" @click="store.openDrawer()"/>
        </template>
      </v-tooltip>
    </div>
  </nav>
</template>